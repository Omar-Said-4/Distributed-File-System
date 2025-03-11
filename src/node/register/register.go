package register

import (
	"context"
	"dfs/schema/register"
	"fmt"
	"net"
	"strconv"

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
func getRandomPort() string {
	listener, err := net.Listen("tcp", ":0") // Bind to an available port
	if err != nil {
		return "" // Handle error appropriately
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	return strconv.Itoa(port)
}
func Register(port string) (uint32, string, string, string) {
	ip := getLocalIP() // Get actual local IP
	// Establish a new connection each time
	conn, err := grpc.NewClient(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return 1000000000, "", "", "" // Return a large number as ID to indicate failure
	}

	c := register.NewRegisterServiceClient(conn)

	// Send Ping
	Nport := getRandomPort()
	// Rport := getRandomPort()
	// nCopyport := getRandomPort()
	fmt.Println("Registering with IP:", ip, "FilePort:", Nport, "ReplicationPort:", Nport, "NotifyToCopyPort:", Nport)
	resp, err := c.Register(context.Background(), &register.RegisterRequest{Ip: ip, FilePort: Nport, ReplicationPort: Nport, NotifyToCopyPort: Nport})
	if err != nil || !resp.Success {
		fmt.Println("Registration failed:", err)
	} else {
		fmt.Println("Registered successfully with IP:", ip)
	}

	conn.Close()
	return resp.Id, Nport, Nport, Nport
}
