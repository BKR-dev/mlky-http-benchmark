//go:build fiber

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/arl/statsviz"
	"net/http"
)

var (
	portFiber = "7443"
)

func init() {
	mux := http.NewServeMux()
	statsviz.Register(mux)

	go func() {
		fmt.Println("statsviz fiber server listening on http://localhost:8082/debug/statsviz/")
		fmt.Println(http.ListenAndServe("localhost:8082", mux))
	}()
	StartFiberServer()
}

func healthCheckEndpointFiber(c *fiber.Ctx) error {
	responseBody := map[string]string{
		"message": "just some JSON",
	}
	responseJson, _ := json.Marshal(responseBody)

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
