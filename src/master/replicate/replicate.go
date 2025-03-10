package replicate

import (
	"context"
	lookup "dfs/master/lookup/file"
	lookup2 "dfs/master/lookup/node"
	"dfs/schema/replicate"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var FilesTable *lookup.FileLookup
var NodesTable *lookup2.NodeLookup

func CheckNeedToCopy(filename string) bool {
	return FilesTable.GetNumberOfReplicas(filename) < 3
}

type replicateServer struct {
	replicate.UnimplementedReplicateServiceServer
}

func getNodeIdToCopyTo(filename string) uint32 {
	nodeId, r1, r2 := FilesTable.GetFileLocation(filename)
	ids := NodesTable.GetLeastLoadedNodes(3)
	if nodeId != ids[0] || r1 != ids[0] || r2 != ids[0] {
		return ids[0]
	}
	if nodeId != ids[1] || r1 != ids[1] || r2 != ids[1] {
		return ids[1]
	}
	return ids[2]
}

func NotifyClients(filename string, id uint32) {
	nodeId := getNodeIdToCopyTo(filename)
	src_ip, _ := NodesTable.GetNodeFileService(id)
	dest_ip, _ := NodesTable.GetNodeFileService(nodeId)
	src_port := NodesTable.GetNotifyToCopyPort(id)
	dest_port := NodesTable.GetNotifyToCopyPort(nodeId)
	conn, err := grpc.NewClient(src_ip+":"+src_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server at %s:%s - Error: %v\n", src_ip, src_port, err)
		return
	}
	defer conn.Close()
	client := replicate.NewReplicateServiceClient(conn)
	_, err = client.NotifyToCopy(context.Background(), &replicate.NotifyToCopyRequest{
		FileName:    filename,
		DestAddress: dest_ip,
		SrcAddress:  src_ip,
		From:        false,
	})
	if err != nil {
		fmt.Printf("Failed to notify client for file %s: %v\n", filename, err)
	}
	conn, err = grpc.NewClient(dest_ip+":"+dest_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server at %s:%s - Error: %v\n", dest_ip, dest_port, err)
		return
	}
	defer conn.Close()
	_, replicationPort := NodesTable.GetNodeReplicationService(id)
	fmt.Printf("Replication port: %s\n", replicationPort)
	client = replicate.NewReplicateServiceClient(conn)
	_, err = client.NotifyToCopy(context.Background(), &replicate.NotifyToCopyRequest{
		FileName:    filename,
		DestAddress: src_ip,
		SrcAddress:  dest_ip,
		From:        true,
		SrcPort:     replicationPort,
	})
	if err != nil {
		fmt.Printf("Failed to notify client for file %s: %v\n", filename, err)
	}
	// TODO: to be changed to the confirm, get file path from replica
	FilesTable.AddReplica(filename, nodeId, "")
	NodesTable.IncrementNumberOfFiles(nodeId)
}

func (s *replicateServer) ConfirmCopy(ctx context.Context, req *replicate.ConfirmCopyRequest) (*replicate.ConfirmCopyResponse, error) {
	filename := req.FileInfo.FileName
	file_path := req.FileInfo.FilePath
	id := req.Id
	FilesTable.AddReplica(filename, id, file_path)
	NodesTable.IncrementNumberOfFiles(id)
	fmt.Printf("File %s copied successfully.\n, Path: %s\n", filename, file_path)
	return &replicate.ConfirmCopyResponse{}, nil
}

func StartConfirmCopyServer(Ftable *lookup.FileLookup, Ntable *lookup2.NodeLookup, port string, s *grpc.Server) {
	FilesTable = Ftable
	NodesTable = Ntable
	replicate.RegisterReplicateServiceServer(s, &replicateServer{})
	fmt.Printf("ConfirmCopy Server is running on port: %s\n", port)
}
