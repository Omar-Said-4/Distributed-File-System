package download

import (
	"context"
	lookup "dfs/master/lookup/file"
	lookup2 "dfs/master/lookup/node"
	"dfs/schema/download"
	"fmt"

	"google.golang.org/grpc"
)

type downloadServer struct {
	download.UnimplementedDownloadServiceServer
}

var FilesTable *lookup.FileLookup
var NodesTable *lookup2.NodeLookup

func (s *downloadServer) RequestDownloadInfo(ctx context.Context, req *download.MasterDownloadRequest) (*download.MasterDownloadResponse, error) {
	if FilesTable == nil || NodesTable == nil {
		return nil, fmt.Errorf("lookup tables not initialized")
	}

	filename := req.GetFileName()
	// fmt.Print(filename)
	nodeId, r1, r2, err := FilesTable.GetFileLocation(filename)
	if err != nil {
		fmt.Print("After getting file location")
		return nil, err
	}
	fmt.Printf("Nodes for file %s: %d, %d, %d\n", filename, nodeId, r1, r2)
	node1_ip, node1_port := NodesTable.GetNodeFileService(nodeId)
	node2_ip, node2_port := NodesTable.GetNodeFileService(r1)
	node3_ip, node3_port := NodesTable.GetNodeFileService(r2)
	if node1_ip == "" && node1_port == "" && node2_ip == "" && node2_port == "" && node3_ip == "" && node3_port == "" {
		fmt.Print("No Copies of the file are found :(")
		return nil, fmt.Errorf("No Copies of the file are found :(")
	}
	return &download.MasterDownloadResponse{
		IpPorts: []*download.IPPort{
			{Ip: node1_ip, Port: node1_port},
			{Ip: node2_ip, Port: node2_port},
			{Ip: node3_ip, Port: node3_port},
		},
		FileSize: FilesTable.GetFileSize(filename),
	}, nil
}

func StartRequestDownloadInfoServer(table *lookup.FileLookup, table2 *lookup2.NodeLookup, port string, s *grpc.Server) {
	FilesTable = table
	NodesTable = table2
	download.RegisterDownloadServiceServer(s, &downloadServer{})
	fmt.Printf("Download Server is running on port: %s\n", port)
}
