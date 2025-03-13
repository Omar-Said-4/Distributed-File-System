package main

import (
	"dfs/master/download"
	"dfs/master/heartbeat"
	lookup2 "dfs/master/lookup/file"
	lookup "dfs/master/lookup/node"
	"dfs/master/register"
	"dfs/master/replicate"
	"dfs/master/upload"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	NodesTable := lookup.AddNodesTable()
	FilesTable := lookup2.AddFileTable()
	if NodesTable == nil { // Defensive check
		fmt.Println("Failed to initialize NodesTable")
		return
	}
	s := grpc.NewServer()
	heartbeat.Init(FilesTable)
	go heartbeat.IsIdle(NodesTable, FilesTable)
	go heartbeat.FilesCheck()
	register.StartRegisterServer(NodesTable, "5052", s)
	heartbeat.StartHeartbeatServer(NodesTable, "5052", s)
	upload.StartMasterRequestUploadServer(NodesTable, "5052", s)
	upload.StartNotifyMasterServer(FilesTable, "5052", s)
	replicate.StartConfirmCopyServer(FilesTable, NodesTable, "5052", s)
	download.StartRequestDownloadInfoServer(FilesTable, NodesTable, "5052", s)
	go func() {
		lis, err := net.Listen("tcp", ":5052")
		if err != nil {
			fmt.Printf("failed to listen: %v\n", err)
			return
		}
		if err := s.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v\n", err)
		}
	}()
	select {}
}
