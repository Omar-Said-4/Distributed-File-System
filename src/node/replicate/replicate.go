package replicate

import (
	"context"
	"dfs/schema/replicate"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

const chunkSize = 4096 // 4KB

var acceptedFilesToCopy = make(map[string]string)
var rwMu sync.RWMutex
var NodeId uint32
var serverIP string
var serverPort string

func getFilePath(filename string) (string, error) {
	absPath, err := filepath.Abs(filepath.Join("../uploads", filename))
	if err != nil {
		return "", err
	}
	return absPath, nil
}

// handle the case when the file is still being received
func waitUntilFileStable(filename string, checkInterval time.Duration) {
	var lastSize int64 = -1

	for {
		uploadsPath := filepath.Join("../uploads", filename)
		info, err := os.Stat(uploadsPath)
		if err != nil {
			// File does not exist yet, keep waiting
			time.Sleep(checkInterval)
			continue
		}

		currentSize := info.Size()
		if currentSize == lastSize {
			// File size hasn't changed, assume it's fully received
			break
		}

		lastSize = currentSize
		time.Sleep(checkInterval)
	}
}

type replicateServer struct {
	replicate.UnimplementedReplicateServiceServer
}

func IsAcceptedFile(filename string, destIp string) bool {
	rwMu.RLock()
	defer rwMu.RUnlock()

	ip, exists := acceptedFilesToCopy[filename]
	return exists && ip == destIp
}
func AddAcceptedFileToCopy(filename string, ip string) {
	rwMu.Lock()
	defer rwMu.Unlock()
	acceptedFilesToCopy[filename] = ip
}
func RemoveAcceptedFileToCopy(filename string) {
	rwMu.Lock()
	defer rwMu.Unlock()
	delete(acceptedFilesToCopy, filename)
}
func (s *replicateServer) CopyFile(req *replicate.CopyFileRequest, stream replicate.ReplicateService_CopyFileServer) error {

	filename := req.FileInfo.FileName
	waitUntilFileStable(filename, 1*time.Second)
	p, ok := peer.FromContext(stream.Context())
	clientIP := "unknown"
	if ok {
		host, _, err := net.SplitHostPort(p.Addr.String())
		if err == nil {
			clientIP = host
		}
	}
	if !IsAcceptedFile(filename, clientIP) {
		return fmt.Errorf("file %s is not accepted to be copied", filename)
	}
	filepath := filepath.Join("../uploads", filename)
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	fileInfoResponse := &replicate.CopyFileResponse{
		Data: &replicate.CopyFileResponse_FileInfo{
			FileInfo: &replicate.FileInfo{
				FileName: req.FileInfo.FileName,
			},
		},
	}
	if err := stream.Send(fileInfoResponse); err != nil {
		return err
	}
	buf := make([]byte, chunkSize)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Send file chunk
		chunkResponse := &replicate.CopyFileResponse{
			Data: &replicate.CopyFileResponse_Chunks{
				Chunks: buf[:n],
			},
		}
		if err := stream.Send(chunkResponse); err != nil {
			return err
		}
	}
	RemoveAcceptedFileToCopy(filename)
	return nil

}

func StartReplicateServer(sIP string, sPort string, port string, id uint32, s *grpc.Server) {
	NodeId = id
	serverIP = sIP
	serverPort = sPort
	replicate.RegisterReplicateServiceServer(s, &replicateServer{})
	fmt.Printf("Replicate Server is running on port: %s\n", port)

}

func RequestACopy(serverIP string, serverPort string, filename string, ip string, port string) {
	var file *os.File
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server at %s:%s - Error: %v\n", ip, port, err)
		return
	}
	defer conn.Close()
	client := replicate.NewReplicateServiceClient(conn)
	req := &replicate.CopyFileRequest{
		FileInfo: &replicate.FileInfo{
			FileName: filename,
		},
	}
	stream, err := client.CopyFile(context.Background(), req)
	if err != nil {
		fmt.Printf("Failed to request a copy of file %s: %v\n", filename, err)
		return
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("File %s copied successfully, nodeId = %d.\n", filename, NodeId)
			ConfirmCopy(serverIP, serverPort, filename)
			return
		}
		if err != nil {
			fmt.Printf("Error receiving stream: %v\n", err)
			return
		}
		if data := resp.GetFileInfo(); data != nil {
			fmt.Printf("Starting new copy: %s\n", data.FileName)
			file, err = os.Create(filepath.Join("../uploads", filename))
			if err != nil {
				fmt.Printf("Failed to create file: %v\n", err)
				return
			}
			defer file.Close()
			continue
		}
		if data := resp.GetChunks(); data != nil {
			_, err := file.Write(data)
			if err != nil {
				fmt.Printf("Failed to write chunk: %v\n", err)
				return
			}
		}
	}
}

func (s *replicateServer) NotifyToCopy(ctx context.Context, req *replicate.NotifyToCopyRequest) (*replicate.NotifyToCopyResponse, error) {
	from := req.From
	if !from {
		filename := req.FileName
		ip := req.DestAddress
		AddAcceptedFileToCopy(filename, ip)
		fmt.Printf("File %s is accepted to be copied to %s\n", filename, ip)
	} else {
		filename := req.FileName
		ip := req.SrcAddress
		port := req.SrcPort
		fmt.Printf("Node notified to copy %s from %s\n", filename, ip)
		RequestACopy(serverIP, serverPort, filename, ip, port)
	}
	return &replicate.NotifyToCopyResponse{}, nil
}

func StartNotifytoCopyServer(port string, id uint32) {
	NodeId = id
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	replicate.RegisterReplicateServiceServer(s, &replicateServer{})
	fmt.Printf("NotifyToCopy Server is running on port: %s\n", port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}

func ConfirmCopy(serverIP, Port, filename string) {
	file_path, _ := getFilePath(filename)
	conn, err := grpc.NewClient(serverIP+":"+Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to master at %s - Error: %v\n", serverIP+":"+serverPort, err)
		return
	}
	defer conn.Close()
	client := replicate.NewReplicateServiceClient(conn)
	_, err = client.ConfirmCopy(context.Background(), &replicate.ConfirmCopyRequest{
		FileInfo: &replicate.FileInfo{
			FileName: filename,
			FilePath: file_path,
		},
		Id: NodeId,
	})
	if err != nil {
		fmt.Printf("Failed to confirm copy for file %s: %v\n", filename, err)
	}
	fmt.Printf("File %s copied successfully, id %d.\n", filename, NodeId)
}
