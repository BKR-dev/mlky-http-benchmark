//go:build stdlib

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	portStd = "5443"
)

func init() {
	StartStandardServer()
}

func healthCheckEndpointStdLib(w http.ResponseWriter, r *http.Request) {
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

func StartStandardServer() {
	fmt.Printf("Standard servers Listens on Port %s with provided endpoint /healthCheck\n", portStd)
	http.HandleFunc("/healthCheck", healthCheckEndpointStdLib)
	http.ListenAndServe(":"+portStd, nil)
}
