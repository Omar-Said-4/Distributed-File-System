package upload

import (
	"context"
	"dfs/schema/upload"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type uploadServer struct {
	upload.UnimplementedUploadServiceServer
}

var nodeID uint32
var masterIp string
var masterPort string

func getFilePath(filename string) (string, error) {
	absPath, err := filepath.Abs(filepath.Join("../uploads", filename))
	if err != nil {
		return "", err
	}
	return absPath, nil
}
func notifyMaster(filename string, filepath string) {
	conn, err := grpc.NewClient(masterIp+":"+masterPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
			FilePath: filepath,
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
	var file_path string
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			notifyMaster(filename, file_path)
			fmt.Printf("File %s upload complete.\n", filename)
			return stream.SendAndClose(&upload.UploadFileResponse{})
		}
		if err != nil {
			fmt.Printf("Error receiving stream: %v\n", err)
			return err
		}

		if data := req.GetFileInfo(); data != nil {
			filename = data.FileName
			fmt.Printf("Starting new upload: %s\n", filename)

			file, err = os.Create(filepath.Join("../uploads", filename))

			if err != nil {
				fmt.Printf("Failed to create file: %v\n", err)
				return err
			}
			file_path, err = getFilePath(filename)
			if err != nil {
				fmt.Printf("File %s upload complete.\nError getting file path: %v\n", filename, err)
				return err
			}
			fmt.Printf("File path: %s\n", file_path)
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
func StartUploadServer(port string, mIp string, mPort string, id uint32, s *grpc.Server) {
	masterIp = mIp
	masterPort = mPort
	nodeID = id
	upload.RegisterUploadServiceServer(s, &uploadServer{})
	fmt.Printf("Upload Server is running on port: %s\n", port)

}
