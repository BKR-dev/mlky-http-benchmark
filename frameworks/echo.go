package frameworks

import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func StartEchoServer() {
	// echo code goes here
	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/healthCheck", healthCheckEndpoint)
	e.Logger.Fatal(e.Start(":8443"))	
}





func healthCheckEndpoint(c echo.Context) error{
	startTime := time.Now()
	responseBody := map[string]string {
		"message":"just some JSON",
	}
	responseJson, _ := json.Marshal(responseBody)

	totalTime := time.Since(startTime)
	fmt.Printf("the route took %s long for the standard lib", totalTime)
	return c.String(http.StatusOK, string(responseJson))
}
