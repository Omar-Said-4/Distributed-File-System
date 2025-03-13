package clientinterface

import (
	"bufio"
	"dfs/client/download"
	"dfs/client/upload"
	"fmt"
	"os"
	"strings"
)

func UploadFile(id string) {
	fmt.Print("Please Enter filename to upload: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	filename = strings.TrimSpace(filename)
	upload.MasterRequestUpload("localhost", "5052", filename, id)
}

func DownloadFile(id string) {
	fmt.Print("Please Enter filename to download: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	filename = strings.TrimSpace(filename)
	filename = fmt.Sprintf("%s_%s", id, filename)
	err = download.RequestDownloadInfo(filename, "localhost", "5052")
	if err != nil {
		fmt.Println(err)
	}
}
