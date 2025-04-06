package clientinterface

import (
	"bufio"
	"dfs/client/download"
	"dfs/client/upload"
	"fmt"
	"os"
	"strings"
)

func UploadFile(id string, serverIP string, serverPort string) {
	fmt.Print("Please Enter filename to upload: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	filename = strings.TrimSpace(filename)
	upload.MasterRequestUpload(serverIP, serverPort, filename, id)
}

func DownloadFile(id string, serverIP string, serverPort string) {
	fmt.Print("Please Enter filename to download: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	filename = strings.TrimSpace(filename)
	filename = fmt.Sprintf("%s_%s", id, filename)
	err = download.RequestDownloadInfo(filename, serverIP, serverPort)
	if err != nil {
		fmt.Println(err)
	}
}
