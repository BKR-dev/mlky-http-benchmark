package main

import (
	"frameworks"
	"sync"
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
	startAllServers()
}

func startAllServers() {
	var wg sync.WaitGroup
	serverSlice := []func(){
		frameworks.StartChiServer,
		frameworks.StartEchoServer,
		frameworks.StartFiberServer,
		frameworks.StartGinServer,
		frameworks.StartGorillaMuxServer,
		frameworks.StartHttprouterServer,
	}
	// counter for each server (not too nice)
	for i := 1; i < len(serverSlice); i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			serverSlice[i]()
		}()
	}
	wg.Wait()
}
