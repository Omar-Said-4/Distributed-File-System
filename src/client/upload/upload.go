package upload

import (
	"context"
	"dfs/schema/upload"
	"fmt"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const chunkSize = 4096 // 4KB
func MasterRequestUpload(ip string, port string, filename string, id string) {
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to master at %s:%s - Error: %v\n", ip, port, err)
		return
	}
	defer conn.Close()

	client := upload.NewUploadServiceClient(conn)

	req := &upload.MasterUploadRequest{}
	resp, err := client.MasterRequestUpload(context.Background(), req)
	if err != nil {
		fmt.Printf("Failed to notify master: %v\n", err)
		return
	}

	fmt.Printf("MasterRequestUpload response - Node IP: %s, Node Port: %s\n", resp.NodeIp, resp.NodePort)
	UploadFile(id, filename, resp.NodeIp, resp.NodePort)
}

func UploadFile(id string, filename string, ip string, port string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", filename, err)
		return
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %v\n", err)
		return
	}
	fileSize := fileStat.Size()

	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server at %s:%s - Error: %v\n", ip, port, err)
		return
	}
	defer conn.Close()
	c := upload.NewUploadServiceClient(conn)
	stream, err := c.UploadFile(context.Background())
	if err != nil {
		fmt.Printf("Failed to create upload stream: %v\n", err)
		return
	}
	// concatenate id and filename to create a unique filename
	fileInfo := &upload.FileInfo{FileName: fmt.Sprintf("%s_%s", id, filename)}
	err = stream.Send(&upload.UploadFileRequest{Data: &upload.UploadFileRequest_FileInfo{FileInfo: fileInfo}})
	if err != nil {
		fmt.Printf("Failed to send file info: %v\n", err)
		return
	}
	buf := make([]byte, chunkSize)
	bar := progressbar.DefaultBytes(fileSize, "Uploading")

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Failed to read file: %v\n", err)
			return
		}

		// Send chunk
		err = stream.Send(&upload.UploadFileRequest{Data: &upload.UploadFileRequest_Chunks{Chunks: buf[:n]}})
		if err != nil {
			fmt.Printf("Failed to send chunk: %v\n", err)
			return
		}
		bar.Add(n)

	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Failed to close upload stream: %v\n", err)
	} else {
		fmt.Printf("File %s uploaded successfully.\n", filename)
	}

}
