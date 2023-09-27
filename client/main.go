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

}

func startAllServers() {
	var wg sync.WaitGroup
	wg.Add(frameworks.StartChiServer())
	wg.Add(frameworks.StartFiberServer())
	defer wg.Done()
	wg.Wait()

}
