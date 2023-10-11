//go:build httprouter

package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

var (
	portHttpRouter = "9000"
)

func init() {
	StartHttprouterServer()
}

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func StartHttprouterServer() {
	fmt.Printf("Httprouter started at port %s, with provided endpoint /healthcheck\n", portHttpRouter)
	hr := httprouter.New()
	hr.GET("/healthcheck", healthCheck)
	log.Fatal(http.ListenAndServe(":"+portHttpRouter, hr))
}
