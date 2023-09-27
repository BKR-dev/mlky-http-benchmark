package frameworks

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

var (
	portChi = "5445"
)

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
	fmt.Printf("the route took %s long for the chi\n", totalTime)
}

func StartChiServer() {
	r := chi.NewRouter()
	r.Get("/healthcheckchi", healthCheckEndpointChi)
	http.ListenAndServe(":"+portChi, nil)
	fmt.Printf("Chi Listens on Port %s with provided endpoint /healthcheckchi\n", portChi)
}
