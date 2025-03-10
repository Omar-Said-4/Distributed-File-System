package main

import (
	"bufio"
	"dfs/client/upload"
	"fmt"
	"os"
	"strings"
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
}
