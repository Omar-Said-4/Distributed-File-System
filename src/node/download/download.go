package download

import (
	"dfs/schema/download"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc"
)

type DownloadServer struct {
	download.UnimplementedDownloadServiceServer
}

const chunkSize = 4096 // 4KB

func (s *DownloadServer) DownloadChunk(req *download.ChunkDownloadRequest, stream download.DownloadService_DownloadChunkServer) error {
	filePath := fmt.Sprintf("../uploads/%s", req.FileName)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", req.FileName, err)
	}
	defer file.Close()
	_, err = file.Seek(int64(req.StartByte), io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek to start byte %d: %v", req.StartByte, err)
	}
	buf := make([]byte, chunkSize)
	var totalBytes uint64 = req.EndByte - req.StartByte
	fmt.Printf("Sending %d bytes of file %s\n", totalBytes, req.FileName)
	var bytesSent uint64 = 0
	for bytesSent < totalBytes {
		remaining := totalBytes - bytesSent
		if remaining < chunkSize {
			buf = buf[:remaining]
		}
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read file %s: %v", req.FileName, err)
		}
		err = stream.Send(&download.ChunkDownloadResponse{Chunk: buf[:n]})
		if err != nil {
			return fmt.Errorf("failed to send chunk of file %s: %v", req.FileName, err)
		}
		bytesSent += uint64(n)
	}
	fmt.Printf("Sent %d bytes of file %s\n", bytesSent, req.FileName)
	return nil
}
func StartDownloadServer(port string, s *grpc.Server) {
	download.RegisterDownloadServiceServer(s, &DownloadServer{})
	fmt.Printf("Download Server is running on port: %s\n", port)
}
