package main

import (
	"fmt"
	"frameworks"
)

// /////////////////////////////////////
//
// here goes the client code to query
// the servers rom the frameworks module
// i think HUGE amounts of queries, like
// 100 million or something LETS GOOOO
//
// /////////////////////////////////////
func QueryServers() {
}

func main() {
	allServers := map[string]func(){
		"echo":       frameworks.StartEchoServer,
		"httprouter": frameworks.StartHttprouterServer,
		"gin":        frameworks.StartGinServer,
		"fiber":      frameworks.StartFiberServer,
		"chi":        frameworks.StartChiServer,
		"gorilla":    frameworks.StartGorillaMuxServer,
		"standard":   frameworks.StartStandardServer,
	}

	for name, server := range allServers {
		fmt.Printf("starting %s", name)
		go server()
	}
	fmt.Println("All servers started")
}

func printName(firstName, lastName string) {
	fmt.Println("Hello " + firstName + lastName)
}
