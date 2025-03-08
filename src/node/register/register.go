package register

import (
	"context"
	"dfs/schema/register"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getLocalIP() string {
	return "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting local IP:", err)
		return "127.0.0.1" // Fallback to localhost
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "127.0.0.1" // Default if no valid IP is found
}

func Register(port string) uint32 {
	ip := getLocalIP() // Get actual local IP
	// Establish a new connection each time
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return 1000000000 // Return a large number as ID to indicate failure
	}

	c := register.NewRegisterServiceClient(conn)

	// Send Ping
	resp, err := c.Register(context.Background(), &register.RegisterRequest{Ip: ip, FilePort: "1", ReplicationPort: "2"})
	if err != nil || !resp.Success {
		fmt.Println("Registration failed:", err)
	} else {
		fmt.Println("Registered successfully with IP:", ip)
	}

	conn.Close()
	return resp.Id
}
