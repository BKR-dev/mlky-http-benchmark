//go:build gin

package main

import (
	"encoding/json"
	"fmt"
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	portGin = "8443"
)

func init() {
	mux := http.NewServeMux()
	statsviz.Register(mux)

	go func() {
		log.Println(http.ListenAndServe("localhost:8084", mux))
	}()
	StartGinServer()
}

func healthCheckEndpointGin(c *gin.Context) {
	startTime := time.Now()
	responseBody := map[string]string{
		"message": "just some JSON",
	}
	responseJson, err := json.Marshal(responseBody)
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to marshal JSON"})
	}
	c.JSON(200, gin.H{"message": responseJson})
	totalTime := time.Since(startTime)
	fmt.Printf("the route took %s long for the standard lib", totalTime)
}

func StartGinServer() {
	fmt.Printf("Gin Servers Listens on Port %s with provided endpoint /healthCheck\n", portGin)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/healthCheck", healthCheckEndpointGin)
	r.Run()
	http.ListenAndServe(":"+portGin, r)
}
