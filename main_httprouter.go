//go:build httprouter

package main

import (
	"encoding/json"
	"fmt"
	"github.com/arl/statsviz"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var (
	portHttpRouter = "9000"
)

func init() {
	mux := http.NewServeMux()
	statsviz.Register(mux)

	go func() {
		fmt.Println("statsviz server httprouter listening on http://localhost:8085/debug/statsviz/")
		fmt.Println(http.ListenAndServe("localhost:8085", mux))
	}()
	StartHttprouterServer()
}

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func StartHttprouterServer() {
	fmt.Printf("Httprouter started at port %s, with provided endpoint /healthCheck\n", portHttpRouter)
	hr := httprouter.New()
	hr.GET("/healthcheck", healthCheck)
	log.Fatal(http.ListenAndServe(":"+portHttpRouter, hr))
}
