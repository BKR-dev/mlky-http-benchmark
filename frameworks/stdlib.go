package frameworks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	portStd = "5443"
)

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
	fmt.Printf("healthcheck route took %s long for the standard lib\n", totalTime)
}

func StartStandardServer() {
	http.HandleFunc("/healthcheck", healthCheckEndpointStdLib)
	fmt.Printf("Standard lib Listens on Port %s with provided endpoint /healthcheck\n", portStd)
	http.ListenAndServe(":"+portStd, nil)
}
