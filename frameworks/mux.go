package frameworks

import (
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/healthCheck", healthCheckEndpointMux)
	http.ListenAndServe(":"+portMux, nil)
	fmt.Printf("Standard servers Listens on Port %s with provided endpoint /healthCheck\n", portMux)
}
