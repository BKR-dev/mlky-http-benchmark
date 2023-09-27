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
	for i := 1; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			frameworks.StartEchoServer()
		}()
	}

	wg.Wait()
}
