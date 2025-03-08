package heartbeat

import (
	"context"
	"dfs/schema/heartbeat"
	"fmt"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func PingServer(ip string, port string, id uint32) {
	for {
		// Establish a new connection each time
		conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println("did not connect:", err)
			time.Sleep(1 * time.Second) // Retry after 1 second
			continue
		}

		c := heartbeat.NewHeartbeatServiceClient(conn)

		// Send Ping
		_, err = c.Ping(context.Background(), &heartbeat.HeartbeatPing{NodeId: id})
		if err != nil {
			fmt.Println("Ping failed:", err)
		} else {
			fmt.Println("Ping sent successfully")
		}

		conn.Close()                // Close the connection before sleeping
		time.Sleep(1 * time.Second) // Wait for 1 second before sending again
	}
}
