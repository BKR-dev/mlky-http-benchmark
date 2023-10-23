//go:build gorillamux

package main

import (
	"encoding/json"
	"fmt"
	"github.com/arl/statsviz"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	portMux = "5443"
)

func init() {
	mux := http.NewServeMux()
	statsviz.Register(mux)

	go func() {
		fmt.Println("statsviz server gorulla mux listening on http://localhost:8084/debug/statsviz/")
		fmt.Println(http.ListenAndServe("localhost:8084", mux))
	}()
	StartGorillaMuxServer()
}

func healthCheckEndpointMux(w http.ResponseWriter, r *http.Request) {
	responseBody := map[string]string{
		"message": "just some JSON",
	}
	responseJson, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "failed to marshal JSON", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func StartGorillaMuxServer() {
	r := mux.NewRouter()
	r.HandleFunc("/healthCheck", healthCheckEndpointMux)
	srv := &http.Server{
		Hanlder:      r,
		Addr:         "127.0.0.1:" + portMux,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	fmt.Printf("Gorilla mux Listens on Port %s with provided endpoint /healthCheck\n", portMux)
	srv.ListenAndServe()
}
