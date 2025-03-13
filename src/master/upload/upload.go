package upload

import (
	"context"
	lookup "dfs/master/lookup/file"
	lookup2 "dfs/master/lookup/node"
	"dfs/master/replicate"
	"dfs/schema/upload"
	"fmt"

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
	filepath := req.FileInfo.FilePath
	filesize := req.FileInfo.FileSize
	ip := p.Addr.String()

	fmt.Printf("NotifyMaster from NodeID: %d, IP: %s, Filename: %s\n", nodeID, ip, filename)
	fmt.Printf("Added file %s to FilesTable\n", filename)

	FilesTable.AddFile(filename, nodeID, filepath, filesize)
	NodesTable.IncrementNumberOfFiles(nodeID)
	NodesTable.AddFileToNode(nodeID, filename)
	// replicate the file to 2 nodes
	replicate.NotifyClients(filename, nodeID)
	replicate.NotifyClients(filename, nodeID)
	fmt.Printf("Notified clients to Copy for file %s\n", filename)

	return &upload.NotifyMasterResponse{}, nil
}

func StartNotifyMasterServer(table *lookup.FileLookup, port string, s *grpc.Server) {
	FilesTable = table

	upload.RegisterUploadServiceServer(s, &uploadServer{})
	fmt.Printf("NotifyMaster Server is running on port: %s\n", port)

}

func (s *uploadServer) MasterRequestUpload(ctx context.Context, req *upload.MasterUploadRequest) (*upload.MasterUploadResponse, error) {
	// !TODO: Edit this to be chosen based on some criteria from NodesTable
	node := NodesTable.GetLeastLoadedNode()
	node_ip, port := NodesTable.GetNodeFileService(node)
	_, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get peer info from context")
	}
	return &upload.MasterUploadResponse{
		NodeIp:   node_ip,
		NodePort: port,
	}, nil
}

func StartMasterRequestUploadServer(table *lookup2.NodeLookup, port string, s *grpc.Server) {
	NodesTable = table

	// upload.RegisterUploadServiceServer(s, &uploadServer{})
	fmt.Printf("MasterRequestUpload Server is running on port: %s\n", port)

}
