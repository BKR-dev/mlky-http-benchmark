package frameworks

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

var (
	portFiber = "7443"
)

func healthCheckEndpointFiber(c *fiber.Ctx) error {
	startTime := time.Now()
	responseBody := map[string]string{
		"message": "just some JSON",
	}
	responseJson, _ := json.Marshal(responseBody)

	totalTime := time.Since(startTime)
	fmt.Printf("the route took %s long for the standard lib", totalTime)
	return c.SendString(string(responseJson))
}

func StartFiberServer() {
	// echo code goes here
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Get("/healthcheck", healthCheckEndpointFiber)
	app.Listen(":" + portFiber)
	fmt.Printf("Standard servers Listens on Port %s with provided endpoint /healthCheck\n", portFiber)
}
