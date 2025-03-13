package heartbeat

import (
	lookup2 "dfs/master/lookup/file"
	lookup "dfs/master/lookup/node"
	"fmt"
	"time"
)

func IsIdle(table *lookup.NodeLookup, table2 *lookup2.FileLookup) {
	fmt.Printf("Checking for idle nodes\n")
	for {
		n_nodes := table.GetNodeCount()
		for i := uint32(0); i < n_nodes; i++ {
			if table.CheckNodeIdle(i) {
				table.SetNodeDead(i)
				fmt.Println("Node ", i, " is idle")
				files := table.GetNodeFiles(i)
				for _, file := range files {
					table.RemoveFileFromNode(i, file)
					table.DecrementNumberOfFiles(i)
					n, r1, r2, _ := table2.GetFileLocation(file)
					if n == i {
						table2.RemoveMainNode(file, i)
					} else if r1 == i {
						table2.RemoveReplica1(file, i)
					} else if r2 == i {
						table2.RemoveReplica2(file, i)
					}
				}
				uploads := table.GetNodeUploadingFiles(i)
				for _, upload := range uploads {
					otherId := table.RemoveUploadingFile(i, upload.Filename)
					table.RemoveUploadingFile(otherId, upload.Filename)
					table2.DecrementNumberUploading(upload.Filename)
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
}
