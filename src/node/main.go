package main

import (
	"dfs/node/download"
	"dfs/node/heartbeat"
	"dfs/node/register"
	"dfs/node/replicate"
	"dfs/node/upload"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"

	"google.golang.org/grpc"
)

func DeleteAllFiles(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	for _, file := range files {
		// Skip directories
		if info, err := os.Stat(file); err == nil && info.IsDir() {
			continue
		}

		if err := os.Remove(file); err != nil {
			return fmt.Errorf("failed to remove %s: %w", file, err)
		}
	}
	return nil
}

type ServerConfig struct {
	NodeID     int    `json:"nodeID"`
	ServerIP   string `json:"serverIP"`
	ServerPort string `json:"serverPort"`
}

func main() {
	configPath := "config/config.json"
	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}
	var config ServerConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		return
	}

	var wg sync.WaitGroup
	var id uint32
	var Fport string
	var RepPort string
	fmt.Printf("Server ip %s\n", config.ServerIP)
	if config.NodeID == -1 {
		id, Fport, RepPort, _ = register.Register(config.ServerIP, config.ServerPort)
		config.NodeID = int(id)
		jsonData, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			return
		}
		err = os.WriteFile(configPath, jsonData, 0644)
		if err != nil {
			fmt.Printf("Error writing config file: %v\n", err)
			return
		}
	} else {
		id = uint32(config.NodeID)
		err := DeleteAllFiles("uploads")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("All old uploads deleted successfully")
		}
	}
	fmt.Printf("Node registered with id: %d\n", id)
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+Fport)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark goroutine as finished
		heartbeat.PingServer(config.ServerIP, config.ServerPort, id)
	}()
	upload.StartUploadServer(Fport, config.ServerIP, config.ServerPort, id, s)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done() // Mark goroutine as finished
	// 	replicate.StartNotifytoCopyServer(NotifyToCopyPort, id)
	// }()

	replicate.StartReplicateServer(config.ServerIP, config.ServerPort, RepPort, id, s)
	download.StartDownloadServer(Fport, s)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
	wg.Wait()
}
