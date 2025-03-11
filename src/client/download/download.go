package download

import (
	"context"
	"dfs/schema/download"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RequestDownloadInfo(filename, ip, port string) ([]*download.IPPort, error) {
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server at %s:%s - error: %v", ip, port, err)
	}
	defer conn.Close()
	client := download.NewDownloadServiceClient(conn)
	resp, err := client.RequestDownloadInfo(context.Background(), &download.MasterDownloadRequest{
		FileName: filename,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to request download info for file %s: %v", filename, err)
	}
	return resp.IpPorts, nil

}
