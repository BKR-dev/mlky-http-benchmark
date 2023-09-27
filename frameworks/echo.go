package frameworks

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var (
	portEcho = "3443"
)

func healthCheckEndpointEcho(c echo.Context) error {
	startTime := time.Now()
	responseBody := map[string]string{
		"message": "just some JSON",
	}
	responseJson, _ := json.Marshal(responseBody)
	totalTime := time.Since(startTime)
	fmt.Printf("the route took %s long for the standard lib", totalTime)
	return c.String(http.StatusOK, string(responseJson))
}

func StartEchoServer() {
	// echo code goes here
	fmt.Printf("Echo Server Listens on Port %s with provided endpoint /healthCheck\n", portEcho)
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.GET("/healthCheck", healthCheckEndpointEcho)
	e.Start(":" + portEcho)
}
