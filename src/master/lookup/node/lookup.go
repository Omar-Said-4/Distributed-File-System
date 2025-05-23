package lookup

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type uploading struct {
	Filename  string
	othernode uint32
	src       bool
}
type DataNode struct {
	NodeId           uint32
	Ip               string
	FilePort         string
	alive            bool
	ReplicationPort  string
	lastPing         time.Time
	n_files          uint32
	files            []string
	uploading        []uploading
	n_uploading      uint32
	NotifyToCopyPort string
}

type NodeHeap []*DataNode

func (h NodeHeap) Len() int { return len(h) }
func (h NodeHeap) Less(i, j int) bool {
	// fmt.Printf("Node Id %d Sum of files is %d\n", h[i].NodeId, int(h[i].n_files)+int(h[i].n_uploading))
	// fmt.Printf("Node Id %d Sum of files is %d\n", h[j].NodeId, int(h[j].n_files)+int(h[j].n_uploading))

	return int(h[i].n_files)+int(h[i].n_uploading) < int(h[j].n_files)+int(h[j].n_uploading)
}
func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*DataNode))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type NodeLookup struct {
	mutex   sync.RWMutex
	table   map[uint32]*DataNode
	heap    NodeHeap
	n_nodes uint32
}

func AddNodesTable() *NodeLookup {
	nl := &NodeLookup{
		table: make(map[uint32]*DataNode),
		heap:  make(NodeHeap, 0),
	}
	heap.Init(&nl.heap)
	return nl
}
func (table *NodeLookup) AddDataNode(nodeId uint32, ip string, filePort string, replicationPort string, ncopyport string) {
	node := &DataNode{
		NodeId:           nodeId,
		Ip:               ip,
		FilePort:         filePort,
		alive:            true,
		ReplicationPort:  replicationPort,
		lastPing:         time.Now(),
		NotifyToCopyPort: ncopyport,
	}
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId] = node
	heap.Push(&table.heap, node)
	table.n_nodes++
	fmt.Printf("Node %d added to lookup table\n", nodeId)
}
func (table *NodeLookup) EditDataNode(nodeId uint32, ip string, filePort string, replicationPort string, ncopyport string) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	if _, ok := table.table[nodeId]; ok {
		table.table[nodeId].Ip = ip
		table.table[nodeId].FilePort = filePort
		table.table[nodeId].ReplicationPort = replicationPort
		table.table[nodeId].NotifyToCopyPort = ncopyport
		fmt.Printf("Node %d edited in lookup table, new IP %s, new Port %s\n", nodeId, ip, filePort)
	}
}
func (table *NodeLookup) RemoveDataNode(nodeId uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	delete(table.table, nodeId)

	table.n_nodes--
	for i, n := range table.heap {
		if n.NodeId == nodeId {
			heap.Remove(&table.heap, i)
			break
		}
	}
}
func (table *NodeLookup) GetNodeFileService(nodeId uint32) (string, string) {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	if len(table.heap) == 0 {
		fmt.Printf("No nodes registered in the system\n")
		return "", ""
	}
	_, ok := table.table[nodeId]
	if !ok {
		fmt.Printf("Node with ID %d not found in NodeLookup\n", nodeId)
		return "", ""
	}
	return table.table[nodeId].Ip, table.table[nodeId].FilePort
}
func (table *NodeLookup) GetNodeReplicationService(nodeId uint32) (string, string) {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	if len(table.heap) == 0 {
		fmt.Printf("No nodes registered in the \n")
		return "", ""
	}
	_, ok := table.table[nodeId]
	if !ok {
		fmt.Printf("Node with ID %d not found in NodeLookup\n", nodeId)
		return "", ""
	}
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
func (table *NodeLookup) GetNumberOfFiles(nodeId uint32) uint32 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[nodeId].n_files
}
func (table *NodeLookup) IncrementNumberOfFiles(nodeId uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId].n_files++
	for i, n := range table.heap {
		if n.NodeId == nodeId {
			heap.Fix(&table.heap, i)
			break
		}
	}
}

func (table *NodeLookup) DecrementNumberOfFiles(nodeId uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId].n_files--
	for i, n := range table.heap {
		if n.NodeId == nodeId {
			heap.Fix(&table.heap, i)
			break
		}
	}
}
func (table *NodeLookup) GetLeastLoadedNode() uint32 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()

	if len(table.heap) == 0 {
		return 0
	}
	tempHeap := make(NodeHeap, len(table.heap))
	copy(tempHeap, table.heap)
	heap.Init(&tempHeap)
	leastLoaded := heap.Pop(&tempHeap).(*DataNode)
	for !leastLoaded.alive {
		leastLoaded = heap.Pop(&tempHeap).(*DataNode)
	}
	return leastLoaded.NodeId
}
func (table *NodeLookup) GetLeastLoadedNodes(n int) []uint32 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()

	if len(table.heap) == 0 {
		return nil
	}
	tempHeap := make(NodeHeap, len(table.heap))
	copy(tempHeap, table.heap)
	heap.Init(&tempHeap)
	result := []uint32{}
	i := 0
	for i < n && len(tempHeap) > 0 {
		// Extract min element
		leastLoaded := heap.Pop(&tempHeap).(*DataNode)
		if leastLoaded.alive {
			result = append(result, leastLoaded.NodeId)
			i++
		} else {
			i-- // Skip dead nodes
		}
	}
	// handle case where n > number of nodes

	for len(result) < n && len(result) > 0 {
		fmt.Printf("len result: %d, len heap: %d\n", len(result), len(table.heap))
		result = append(result, result[i%len(result)])
		i++
	}
	fmt.Printf("Least loaded nodes: %v\n", result)
	return result
}
func (table *NodeLookup) GetNotifyToCopyPort(nodeId uint32) string {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[nodeId].NotifyToCopyPort
}
func (table *NodeLookup) AddFileToNode(nodeId uint32, filename string) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId].files = append(table.table[nodeId].files, filename)
}
func (table *NodeLookup) RemoveFileFromNode(nodeId uint32, filename string) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	for i, f := range table.table[nodeId].files {
		if f == filename {
			table.table[nodeId].files = append(table.table[nodeId].files[:i], table.table[nodeId].files[i+1:]...)
			break
		}
	}
}
func (table *NodeLookup) GetNodeFiles(nodeId uint32) []string {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[nodeId].files
}
func (table *NodeLookup) AddUploadingFile(nodeId uint32, filename string, othernode uint32, src bool) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[nodeId].uploading = append(table.table[nodeId].uploading, uploading{filename, othernode, src})
	if !src {
		table.table[nodeId].n_uploading++
		for i, n := range table.heap {
			if n.NodeId == nodeId {
				heap.Fix(&table.heap, i)
				break
			}
		}
	}
}
func (table *NodeLookup) RemoveUploadingFile(nodeId uint32, filename string) uint32 {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	for i, f := range table.table[nodeId].uploading {
		if f.Filename == filename {
			otherId := f.othernode
			table.table[nodeId].uploading = append(table.table[nodeId].uploading[:i], table.table[nodeId].uploading[i+1:]...)
			if !f.src {
				table.table[nodeId].n_uploading--
				for i, n := range table.heap {
					if n.NodeId == nodeId {
						heap.Fix(&table.heap, i)
						break
					}
				}
			}
			return otherId
		}
	}
	return ^uint32(0)
}
func (table *NodeLookup) GetNodeUploadingFiles(nodeId uint32) []uploading {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[nodeId].uploading
}
