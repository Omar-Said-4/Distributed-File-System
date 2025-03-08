package main

import (
	"dfs/master/heartbeat"
	lookup2 "dfs/master/lookup/file"
	lookup "dfs/master/lookup/node"
	"dfs/master/register"
	"dfs/master/upload"
	"fmt"
)

func main() {
	NodesTable := lookup.AddNodesTable()
	FilesTable := lookup2.AddFileTable()
	if NodesTable == nil { // Defensive check
		fmt.Println("Failed to initialize NodesTable")
		return
	}
	go register.StartRegisterServer(NodesTable, "5052")
	go heartbeat.StartHeartbeatServer(NodesTable, "5051")
	go heartbeat.IsIdle(NodesTable)
	go upload.StartMasterRequestUploadServer(NodesTable, "5050")
	go upload.StartNotifyMasterServer(FilesTable, "5055")
	select {}
}
