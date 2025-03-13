package main

import (
	"bytes"

	ui "dfs/client/interface"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
			ui.UploadFile(clientId)
		case 2:
			ui.DownloadFile(clientId)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}
