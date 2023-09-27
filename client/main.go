package main

import (
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
	frameworks.StartEchoServer()
	frameworks.StartStandardServer()
	frameworks.StartGinServer()
}
