package frameworks

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var (
	portEcho = "9902"
)

func healthCheckEndpointEcho(c echo.Context) error {
	startTime := time.Now()
	responseBody := map[string]string{
		"message": "just some JSON",
	}
	responseJson, _ := json.Marshal(responseBody)

	totalTime := time.Since(startTime)
	fmt.Printf("the route took %s long for echo\n", totalTime)
	return c.String(http.StatusOK, string(responseJson))
}

func StartEchoServer() {
	// echo code goes here
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.GET("/healthcheckec", healthCheckEndpointEcho)
	if err := e.Start(":" + portEcho); err != nil {
		fmt.Printf("error starting echo server: %v", err)
	}
	fmt.Printf("Echo Listens on Port %s with provided endpoint /healthcheckec\n", portEcho)
}
