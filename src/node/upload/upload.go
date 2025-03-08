package upload

import (
	"context"
	"crypto/sha256"
	"dfs/schema/upload"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type uploadServer struct {
	upload.UnimplementedUploadServiceServer
}

func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	return hex.EncodeToString(hash[:])
}

var nodeID uint32
var ip string
var port string

func notifyMaster(filename string) {
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to master: %v\n", err)
		return
	}
	defer conn.Close()

	client := upload.NewUploadServiceClient(conn)

	req := &upload.NotifyMasterRequest{
		NodeId: nodeID,
		FileInfo: &upload.FileInfo{
			FileName: filename,
		},
	}

	// Send the request
	_, err = client.NotifyMaster(context.Background(), req)
	if err != nil {
		fmt.Printf("Failed to notify master: %v\n", err)
	} else {
		fmt.Println("Successfully notified master.")
	}
}
func (s *uploadServer) UploadFile(stream upload.UploadService_UploadFileServer) error {
	var file *os.File
	var filename string
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("File %s upload complete.\n", filename)
			notifyMaster(filename)
			return stream.SendAndClose(&upload.UploadFileResponse{})
		}
		if err != nil {
			fmt.Printf("Error receiving stream: %v\n", err)
			return err
		}

		if data := req.GetFileInfo(); data != nil {
			originalFilename := data.FileName
			hashedFilename := hashFilename(originalFilename) + filepath.Ext(originalFilename)
			fmt.Printf("Starting new upload: %s\n", hashedFilename)

			file, err = os.Create(filename)
			if err != nil {
				fmt.Printf("Failed to create file: %v\n", err)
				return err
			}
			continue
		}

		if data := req.GetChunks(); data != nil {
			if file == nil {
				fmt.Println("Error: FileInfo not received before chunks.")
				return fmt.Errorf("file info missing before chunks")
			}

			_, err := file.Write(data)
			if err != nil {
				fmt.Printf("Failed to write chunk: %v\n", err)
				file.Close()
				return err
			}
		}
	}
}
func StartUploadServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	upload.RegisterUploadServiceServer(s, &uploadServer{})
	fmt.Printf("Upload Server is running on port: %s\n", port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
