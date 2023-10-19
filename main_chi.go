//go:build chi

package main

import (
	"encoding/json"
	"fmt"
	"github.com/arl/statsviz"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

var (
	portChi = "3000"
)

func init() {
	mux := http.NewServeMux()
	statsviz.Register(mux)

	go func() {
		fmt.Println("statsviz server chi listening on http://localhost:8080/debug/statsviz/")
		fmt.Println(http.ListenAndServe("localhost:8080", mux))
	}()
	startChiServer()
}

func healthCheckEndpointChi(w http.ResponseWriter, r *http.Request) {
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

func startChiServer() {
	r := chi.NewRouter()
	r.Get("/healthCheck", healthCheckEndpointChi)
	fmt.Printf("Standard servers Listens on Port %s with provided endpoint /healthCheck\n", portChi)
	http.ListenAndServe(":"+portChi, r)
}
