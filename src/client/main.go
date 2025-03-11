package main

import (
	"bufio"
	"dfs/client/download"
	"dfs/client/upload"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	clientId := uint32(1)
	fmt.Print("Enter filename: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	filename = strings.TrimSpace(filename)
	fmt.Printf("Client Started\n")
	upload.MasterRequestUpload("localhost", "5052", filename, clientId)

	time.Sleep(3 * time.Second)
	filename = fmt.Sprintf("%d_%s", clientId, filename)
	info, _ := download.RequestDownloadInfo(filename, "localhost", "5052")
	for _, ipport := range info {
		fmt.Printf("Ip: %s, Port: %s\n", ipport.Ip, ipport.Port)
	}
}
