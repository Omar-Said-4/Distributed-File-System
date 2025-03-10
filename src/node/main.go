package main

import (
	"dfs/node/heartbeat"
	"dfs/node/register"
	"dfs/node/replicate"
	"dfs/node/upload"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	id, Fport, RepPort, NotifyToCopyPort := register.Register("5052")
	fmt.Printf("Node registered with id: %d\n", id)

	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		heartbeat.PingServer("localhost", "5052", id)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		upload.StartUploadServer(Fport, "localhost", "5052", id)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		replicate.StartNotifytoCopyServer(NotifyToCopyPort, id)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		replicate.StartReplicateServer(RepPort)
	}()
	wg.Wait()
}
