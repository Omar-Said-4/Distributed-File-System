package main

import (
	"dfs/node/heartbeat"
	"dfs/node/register"
	"dfs/node/upload"
	"fmt"
	"sync"
)

var id uint32

func main() {
	var wg sync.WaitGroup
	id = register.Register("5052")
	fmt.Printf("Node registered with id: %d\n", id)

	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		heartbeat.PingServer("localhost", "5051", id)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		upload.StartUploadServer("4000", "localhost", "5055")
	}()

	wg.Wait()
}
