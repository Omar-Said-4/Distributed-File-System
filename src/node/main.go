package main

import (
	"dfs/node/heartbeat"
	"dfs/node/register"
	"dfs/node/replicate"
	"dfs/node/upload"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup
	id, Fport, RepPort, _ := register.Register("5052")
	fmt.Printf("Node registered with id: %d\n", id)
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+Fport)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		heartbeat.PingServer("localhost", "5052", id)
	}()
	upload.StartUploadServer(Fport, "localhost", "5052", id, s)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done() // Mark goroutine as finished
	// 	replicate.StartNotifytoCopyServer(NotifyToCopyPort, id)
	// }()

	replicate.StartReplicateServer(RepPort, id, s)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
	wg.Wait()
}
