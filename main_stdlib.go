//go:build stdlib

package main

import (
	"encoding/json"
	"fmt"
	"github.com/arl/statsviz"
	"net/http"
)

var (
	portStd = "5228"
)

func init() {
	mux := http.NewServeMux()
	statsviz.Register(mux)

	go func() {
		fmt.Println("statsviz server standard lib listening on http://localhost:8086/debug/statsviz/")
		fmt.Println(http.ListenAndServe("localhost:8086", mux))
	}()
	StartStandardServer()
}

func healthCheckEndpointStdLib(w http.ResponseWriter, r *http.Request) {
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

func StartStandardServer() {
	fmt.Printf("Standard servers Listens on Port %s with provided endpoint /healthCheck\n", portStd)
	http.HandleFunc("/healthCheck", healthCheckEndpointStdLib)
	http.ListenAndServe(":"+portStd, nil)
}
