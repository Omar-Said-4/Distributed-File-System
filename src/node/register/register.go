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
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces:", err)
		return "127.0.0.1"
	}

	for _, iface := range interfaces {
		// Skip interfaces that are down or loopback
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				ip := ipNet.IP
				// Check for valid IPv4 (non-loopback, non-link-local)
				if ip.To4() != nil && !ip.IsLoopback() && !ip.IsLinkLocalUnicast() {
					return ip.String()
				}
			}
		}
	}

	return "127.0.0.1"
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
func Register(serverIp string, port string, old_id int64) (uint32, string, string, string) {
	ip := getLocalIP() // Get actual local IP
	// Establish a new connection each time
	conn, err := grpc.NewClient(serverIp+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return 1000000000, "", "", "" // Return a large number as ID to indicate failure
	}

	c := register.NewRegisterServiceClient(conn)

	// Send Ping
	Nport := getRandomPort()
	// Rport := getRandomPort()
	// nCopyport := getRandomPort()
	fmt.Println("Registering with IP:", ip, "FilePort:", Nport, "ReplicationPort:", Nport, "NotifyToCopyPort:", Nport, "OldId:", old_id)
	resp, err := c.Register(context.Background(), &register.RegisterRequest{Ip: ip, FilePort: Nport, ReplicationPort: Nport, NotifyToCopyPort: Nport, OldId: old_id})
	if err != nil || !resp.Success {
		fmt.Println("Registration failed:", err)
	} else {
		fmt.Println("Registered successfully with IP:", ip)
		fmt.Println("Node ID:", resp.Id)
	}

	conn.Close()
	return resp.Id, Nport, Nport, Nport
}
