package frameworks

import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

func HealthCheckEndpoint(w http.ResponseWriter, r *http.Request){
	startTime := time.Now()
	responseBody := map[string]string {
		"message":"just some JSON",
	}
	responseJson, err := json.Marshal(responseBody)
	if err!= nil {
		http.Error(w, "failed to marshal JSON", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)

	totalTime := time.Since(startTime)
	fmt.Printf("the route took %s long for the standard lib", totalTime)
}

func StartStandardServer(){
	http.HandleFunc("/healthCheck", HealthCheckEndpoint)
	http.ListenAndServe(":9443", nil)
}
