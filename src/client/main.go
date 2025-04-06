package main

import (
	"bytes"
	"encoding/json"

	ui "dfs/client/interface"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ServerConfig struct {
	ServerIP   string `json:"serverIP"`
	ServerPort string `json:"serverPort"`
}

func getMachineID() string {

	cmd := exec.Command("wmic", "csproduct", "get", "UUID")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "unknown-machine"
	}
	lines := strings.Split(out.String(), "\n")
	if len(lines) > 1 {
		return strings.TrimSpace(lines[1]) // Extract UUID
	}
	return "unknown-machine"
}
func uniqueID() string {
	machineID := getMachineID()
	pid := os.Getpid()
	return fmt.Sprintf("%s_%d", machineID, pid)
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

	clientId := uniqueID()
	fmt.Printf("Client Started with ID %s\n", clientId)

	for {
		fmt.Println("1. Upload File")
		fmt.Println("2. Download File")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			ui.UploadFile(clientId, config.ServerIP, config.ServerPort)
		case 2:
			ui.DownloadFile(clientId, config.ServerIP, config.ServerPort)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}
