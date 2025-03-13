package replicate

import (
	"context"
	lookup "dfs/master/lookup/file"
	lookup2 "dfs/master/lookup/node"
	"dfs/schema/replicate"
	"fmt"
	"sync"

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

func getNodeIdToCopyTo(filename string) (uint32, error) {
	nodeId, r1, r2, _ := FilesTable.GetFileLocation(filename)
	ids := NodesTable.GetLeastLoadedNodes(3)
	if nodeId != ids[0] && r1 != ids[0] && r2 != ids[0] {
		return ids[0], nil
	}
	if nodeId != ids[1] && r1 != ids[1] && r2 != ids[1] {
		return ids[1], nil
	}
	if nodeId != ids[2] && r1 != ids[2] && r2 != ids[2] {
		return ids[2], nil
	}
	return 0, fmt.Errorf("failed to find node to copy to")
}

var fileLocks sync.Map

func getFileLock(filename string) *sync.Mutex {
	lock, _ := fileLocks.LoadOrStore(filename, &sync.Mutex{})
	return lock.(*sync.Mutex)
}

func NotifyClients(filename string, id uint32) {
	lock := getFileLock(filename)
	lock.Lock()
	defer lock.Unlock()
	err := FilesTable.IncrementNumberUploading(filename)
	if err != nil {
		fmt.Printf("Failed to increment number of uploading for file %s: %v\n", filename, err)
		return
	}
	nodeId, err := getNodeIdToCopyTo(filename)
	if err != nil {
		fmt.Printf("Failed to get node to copy to for file %s: %v\n", filename, err)
		FilesTable.DecrementNumberUploading(filename)
		return
	}
	src_ip, _ := NodesTable.GetNodeFileService(id)
	dest_ip, _ := NodesTable.GetNodeFileService(nodeId)
	src_port := NodesTable.GetNotifyToCopyPort(id)
	dest_port := NodesTable.GetNotifyToCopyPort(nodeId)
	NodesTable.AddUploadingFile(nodeId, filename, id, true)
	NodesTable.AddUploadingFile(id, filename, nodeId, false)
	conn1, err := grpc.NewClient(src_ip+":"+src_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server at %s:%s - Error: %v\n", src_ip, src_port, err)
		FilesTable.DecrementNumberUploading(filename)
		NodesTable.RemoveUploadingFile(nodeId, filename)
		NodesTable.RemoveUploadingFile(id, filename)
		return
	}
	defer conn1.Close()
	client := replicate.NewReplicateServiceClient(conn1)
	_, err = client.NotifyToCopy(context.Background(), &replicate.NotifyToCopyRequest{
		FileName:    filename,
		DestAddress: dest_ip,
		SrcAddress:  src_ip,
		From:        false,
	})
	if err != nil {
		fmt.Printf("Failed to notify client for file %s: %v\n", filename, err)
		FilesTable.DecrementNumberUploading(filename)
		NodesTable.RemoveUploadingFile(nodeId, filename)
		NodesTable.RemoveUploadingFile(id, filename)
		return
	}
	conn2, err := grpc.NewClient(dest_ip+":"+dest_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server at %s:%s - Error: %v\n", dest_ip, dest_port, err)
		FilesTable.DecrementNumberUploading(filename)
		NodesTable.RemoveUploadingFile(nodeId, filename)
		NodesTable.RemoveUploadingFile(id, filename)
		return
	}
	defer conn2.Close()
	_, replicationPort := NodesTable.GetNodeReplicationService(id)
	fmt.Printf("Replication port: %s\n", replicationPort)
	client = replicate.NewReplicateServiceClient(conn2)
	_, err = client.NotifyToCopy(context.Background(), &replicate.NotifyToCopyRequest{
		FileName:    filename,
		DestAddress: src_ip,
		SrcAddress:  dest_ip,
		From:        true,
		SrcPort:     replicationPort,
	})
	if err != nil {
		fmt.Printf("Failed to notify client for file %s: %v\n", filename, err)
		FilesTable.DecrementNumberUploading(filename)
		NodesTable.RemoveUploadingFile(nodeId, filename)
		NodesTable.RemoveUploadingFile(id, filename)
		return
	}
	fmt.Printf("Notified clients to copy file %s, Id %d\n", filename, nodeId)
}

func (s *replicateServer) ConfirmCopy(ctx context.Context, req *replicate.ConfirmCopyRequest) (*replicate.ConfirmCopyResponse, error) {
	filename := req.FileInfo.FileName
	file_path := req.FileInfo.FilePath
	fmt.Printf("Received confirm request  %s\n", filename)
	id := req.Id
	FilesTable.AddReplica(filename, id, file_path)
	NodesTable.IncrementNumberOfFiles(id)
	NodesTable.AddFileToNode(id, filename)
	FilesTable.DecrementNumberUploading(filename)
	otherId := NodesTable.RemoveUploadingFile(id, filename)
	NodesTable.RemoveUploadingFile(otherId, filename)
	fmt.Printf("File %s copied successfully.\n, Path: %s, Id %d\n", filename, file_path, id)
	return &replicate.ConfirmCopyResponse{}, nil
}

func StartConfirmCopyServer(Ftable *lookup.FileLookup, Ntable *lookup2.NodeLookup, port string, s *grpc.Server) {
	FilesTable = Ftable
	NodesTable = Ntable
	replicate.RegisterReplicateServiceServer(s, &replicateServer{})
	fmt.Printf("Confirm Copy Server is running on port: %s\n", port)
}
