package heartbeat

import (
	lookup "dfs/master/lookup/node"
	"fmt"
	"time"
)

func IsIdle(table *lookup.NodeLookup) {
	fmt.Printf("Checking for idle nodes\n")
	for {
		n_nodes := table.GetNodeCount()
		for i := uint32(0); i < n_nodes; i++ {
			if table.CheckNodeIdle(i) {
				table.SetNodeDead(i)
				fmt.Println("Node ", i, " is idle")
			}
		}
		time.Sleep(2 * time.Second)
	}
}
