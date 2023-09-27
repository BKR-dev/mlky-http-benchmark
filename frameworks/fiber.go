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
	fmt.Printf("the route took %s long for fiber\n", totalTime)
	return c.SendString(string(responseJson))
}

func StartFiberServer() {
	// echo code goes here
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Get("/healthcheckfi", healthCheckEndpointFiber)
	app.Listen(":" + portFiber)
	fmt.Printf("Fiber Listens on Port %s with provided endpoint /healthcheckfi\n", portFiber)
}
