package lookup

import (
	"fmt"
	"sync"
	"time"
)

type DataNode struct {
	NodeId          uint32
	Ip              string
	FilePort        string
	alive           bool
	ReplicationPort string
	lastPing        time.Time
}
type NodeLookup struct {
	mutex   sync.RWMutex
	table   map[uint32]*DataNode
	n_nodes uint32
}

func AddNodesTable() *NodeLookup {
	return &NodeLookup{
		table: make(map[uint32]*DataNode),
	}
}
func (table *NodeLookup) AddDataNode(nodeId uint32, ip string, filePort string, replicationPort string) {
	node := &DataNode{
		NodeId:          nodeId,
		Ip:              ip,
		FilePort:        filePort,
		alive:           true,
		ReplicationPort: replicationPort,
		lastPing:        time.Now(),
	}
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId] = node
	table.n_nodes++
	fmt.Printf("Node %d added to lookup table\n", nodeId)
}
func (table *NodeLookup) RemoveDataNode(nodeId uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	delete(table.table, nodeId)
	table.n_nodes--
}
func (table *NodeLookup) GetNodeFileService(nodeId uint32) (string, string) {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[nodeId].Ip, table.table[nodeId].FilePort
}
func (table *NodeLookup) GetNodeReplicationService(nodeId uint32) (string, string) {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[nodeId].Ip, table.table[nodeId].ReplicationPort
}
func (table *NodeLookup) GetNodeAlive(nodeId uint32) bool {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[nodeId].alive
}
func (table *NodeLookup) SetNodeAlive(nodeId uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId].alive = true
}
func (table *NodeLookup) SetNodeDead(nodeId uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId].alive = false
	fmt.Print("Node ", nodeId, " is dead\n")
}

func (table *NodeLookup) GetNodeCount() uint32 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.n_nodes
}

func (table *NodeLookup) UpdateNodePingTime(nodeId uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId].lastPing = time.Now()
}
func (table *NodeLookup) CheckNodeIdle(nodeId uint32) bool {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	// max idle time is 5 seconds
	return time.Since(table.table[nodeId].lastPing) > 5*time.Second
}
