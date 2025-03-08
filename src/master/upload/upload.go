package upload

import (
	"context"
	lookup "dfs/master/lookup/file"
	lookup2 "dfs/master/lookup/node"
	"dfs/schema/upload"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var FilesTable *lookup.FileLookup
var NodesTable *lookup2.NodeLookup

type uploadServer struct {
	upload.UnimplementedUploadServiceServer
}

func (s *uploadServer) NotifyMaster(ctx context.Context, req *upload.NotifyMasterRequest) (*upload.NotifyMasterResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get peer info from context")
	}

	nodeID := req.NodeId
	filename := req.FileInfo.FileName
	ip := p.Addr.String()

	fmt.Printf("NotifyMaster from NodeID: %d, IP: %s, Filename: %s\n", nodeID, ip, filename)

	FilesTable.AddFile(filename, nodeID)
	fmt.Printf("Added file %s to FilesTable\n", filename)

	return &upload.NotifyMasterResponse{}, nil
}

func StartNotifyMasterServer(table *lookup.FileLookup, port string) {
	FilesTable = table
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	upload.RegisterUploadServiceServer(s, &uploadServer{})
	fmt.Printf("NotifyMaster Server is running on port: %s\n", port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}

func (s *uploadServer) MasterRequestUpload(ctx context.Context, req *upload.MasterUploadRequest) (*upload.MasterUploadResponse, error) {
	// !TODO: Edit this to be chosen based on some criteria from NodesTable
	node_ip := "localhost"
	port := "4000"
	_, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get peer info from context")
	}
	return &upload.MasterUploadResponse{
		NodeIp:   node_ip,
		NodePort: port,
	}, nil
}

func StartMasterRequestUploadServer(table *lookup2.NodeLookup, port string) {
	NodesTable = table
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	upload.RegisterUploadServiceServer(s, &uploadServer{})
	fmt.Printf("MasterRequestUpload Server is running on port: %s\n", port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
