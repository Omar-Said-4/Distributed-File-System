package heartbeat

import (
	lookup2 "dfs/master/lookup/file"
	"dfs/master/replicate"
	"fmt"
	"time"
)

var table *lookup2.FileLookup

func Init(table2 *lookup2.FileLookup) {
	table = table2
}

func getSourceMachine(filename string) (uint32, string) {
	n, _, _, err := table.GetFileLocation(filename)
	if err != nil {
		fmt.Printf("Failed to get file location for file %s: %v\n", filename, err)
		return 0, ""
	}
	fp1, _, _ := table.GetFilePaths(filename)
	return n, fp1
}

func FilesCheck() {
	for {
		files := table.GetFileNames()
		for _, file := range files {
			total_num := table.GetNumberUploading(file) + table.GetNumberOfReplicas(file)
			i := total_num
			for i < 3 {
				nodeId, _ := getSourceMachine(file)
				fmt.Printf("File Check Replicating file %s \n", file)
				go replicate.NotifyClients(file, nodeId)
				i++
			}
		}
		time.Sleep(10 * time.Second)
	}
}
