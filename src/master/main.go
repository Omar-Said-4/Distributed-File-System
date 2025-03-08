package main

import (
	"dfs/master/heartbeat"
	lookup "dfs/master/lookup/node"
	"dfs/master/register"
	"fmt"
)

func main() {
	NodesTable := lookup.AddNodesTable()
	if NodesTable == nil { // Defensive check
		fmt.Println("Failed to initialize NodesTable")
		return
	}
	go register.StartRegisterServer(NodesTable, "5052")
	go heartbeat.StartHeartbeatServer(NodesTable, "5051")
	go heartbeat.IsIdle(NodesTable)
	fmt.Printf("Heartbeat Server is started and running on port :5051\n")
	select {}
}
