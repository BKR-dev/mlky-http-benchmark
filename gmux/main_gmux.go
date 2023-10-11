//go:build gorillamux

package gorillamux

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	portMux = "5443"
)

func healthCheckEndpointMux(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	responseBody := map[string]string{
		"message": "just some JSON",
	}
	responseJson, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "failed to marshal JSON", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)

	totalTime := time.Since(startTime)
	fmt.Printf("the route took %s long for the standard lib", totalTime)
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
	fmt.Printf("Standard servers Listens on Port %s with provided endpoint /healthCheck\n", portMux)
	srv.ListenAndServe()
}
