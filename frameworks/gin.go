package frameworks

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	portGin = "8443"
)

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
	r := gin.Default()
	r.GET("/healthCheck", healthCheckEndpointGin)
	r.Run()
	fmt.Printf("Standard servers Listens on Port %s with provided endpoint /healthCheck\n", portGin)
	http.ListenAndServe(":"+portGin, r)
}
