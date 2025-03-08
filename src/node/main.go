package main

import (
	"dfs/node/heartbeat"
	"dfs/node/register"
	"fmt"
	"sync"
)

var id uint32

func main() {

	id = register.Register("5052")
	fmt.Printf("Node registered with id: %d\n", id)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		heartbeat.PingServer("localhost", "5051", 0)
	}()

	wg.Wait()
}
