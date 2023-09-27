package frameworks

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	portGin = "9904"
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
	fmt.Printf("the route took %s long for gin\n", totalTime)
}

func StartGinServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/healthcheckgin", healthCheckEndpointGin)
	r.Run(":" + portGin)
	fmt.Printf("Gin Listens on Port %s with provided endpoint /healthcheckgin\n", portGin)
}
