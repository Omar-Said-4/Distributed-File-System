package download

import (
	"context"
	"dfs/schema/download"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func cleanFilename(filename string) string {
	re := regexp.MustCompile(`^[^_]+_\d+_`)
	return re.ReplaceAllString(filename, "")
}
func RequestDownloadInfo(filename, ip, port string) error {
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to server at %s:%s - error: %v", ip, port, err)
	}
	defer conn.Close()
	client := download.NewDownloadServiceClient(conn)
	resp, err := client.RequestDownloadInfo(context.Background(), &download.MasterDownloadRequest{
		FileName: filename,
	})
	if err != nil {
		return fmt.Errorf("failed to request download info for file %s: %v", filename, err)
	}
	nodes := resp.IpPorts
	filesize := resp.FileSize
	n_nodes := len(nodes)
	chunksize := uint64(math.Ceil(float64(filesize) / float64(n_nodes)))
	fmt.Printf("To be Downloaded Filesize: %d, n_nodes: %d, chunksize: %d\n", filesize, n_nodes, chunksize)
	chunkData := make([][]byte, n_nodes)
	mu := sync.Mutex{}
	var wg sync.WaitGroup
	for i, node := range nodes {
		startByte := uint64(i) * chunksize
		endByte := uint64(i+1) * chunksize
		if i == n_nodes-1 {
			endByte = filesize
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := requestChunk(filename, node.Ip, node.Port, startByte, endByte)
			if err != nil {
				fmt.Printf("failed to request chunk from %s:%s: %v\n", node.Ip, node.Port, err)
				return
			}
			mu.Lock()
			chunkData[i] = data
			mu.Unlock()
			fmt.Printf("Downloaded chunk %d from %s:%s\n", i, node.Ip, node.Port)
		}()
	}
	wg.Wait()
	cleaned_filename := cleanFilename(filename)
	out_path := fmt.Sprintf("../downloads/%s", cleaned_filename)
	file, err := os.Create(out_path)
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		return err
	}
	defer file.Close()
	for _, data := range chunkData {
		_, err := file.Write(data)
		if err != nil {
			fmt.Printf("Failed to write to file: %v\n", err)
			return err
		}
	}
	fmt.Printf("Successfully downloaded file %s\n", filename)

	return nil

}

func requestChunk(filename, ip, port string, startByte uint64, endByte uint64) ([]byte, error) {
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server at %s:%s - error: %v", ip, port, err)
	}
	defer conn.Close()
	client := download.NewDownloadServiceClient(conn)
	req := &download.ChunkDownloadRequest{
		FileName:  filename,
		StartByte: startByte,
		EndByte:   endByte,
	}
	stream, err := client.DownloadChunk(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to request a chunk of file %s: %v", filename, err)
	}
	var data []byte
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to download chunk of file %s: %v", filename, err)
		}
		data = append(data, resp.Chunk...)
	}
	return data, nil
}
