package lookup

import (
	"fmt"
	"sync"
)

type File struct {
	file_name   string
	node_id     uint32
	file_path   string
	replica_id1 uint32
	file_path1  string
	replica_id2 uint32
	file_path2  string
	n_replicas  uint32
	file_size   uint64
}

type FileLookup struct {
	mutex sync.RWMutex
	table map[string]*File
}

func AddFileTable() *FileLookup {
	return &FileLookup{
		table: make(map[string]*File),
	}
}

func (table *FileLookup) AddFile(file_name string, node_id uint32, filepath string, file_size uint64) {
	file := &File{
		file_name:  file_name,
		node_id:    node_id,
		file_path:  filepath,
		file_size:  file_size,
		n_replicas: 1,
	}
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[file_name] = file
}

func (table *FileLookup) RemoveFile(file_name string) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	delete(table.table, file_name)
}
func (table *FileLookup) GetNumberOfReplicas(file_name string) uint32 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[file_name].n_replicas
}
func (table *FileLookup) GetFileLocation(file_name string) (uint32, uint32, uint32) {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	if table.table[file_name].n_replicas == 1 {
		return table.table[file_name].node_id, table.table[file_name].node_id, table.table[file_name].node_id
	}
	if table.table[file_name].n_replicas == 2 {
		return table.table[file_name].node_id, table.table[file_name].replica_id1, table.table[file_name].node_id
	}
	return table.table[file_name].node_id, table.table[file_name].replica_id1, table.table[file_name].replica_id2
}
func (table *FileLookup) AddReplica(file_name string, node_id uint32, filepath string) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[file_name].n_replicas++
	if table.table[file_name].n_replicas == 2 {
		table.table[file_name].replica_id1 = node_id
		table.table[file_name].file_path1 = filepath
	}
	if table.table[file_name].n_replicas == 3 {
		table.table[file_name].replica_id2 = node_id
		table.table[file_name].file_path2 = filepath
	}
	fmt.Printf("FileLookup: %v\n", table.table[file_name])
}
func (table *FileLookup) RemoveReplica1(file_name string, node_id uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[file_name].n_replicas--
	if table.table[file_name].n_replicas == 3 {
		table.table[file_name].replica_id1 = table.table[file_name].replica_id2
		table.table[file_name].replica_id2 = node_id
	} else if table.table[file_name].n_replicas == 2 {
		table.table[file_name].replica_id1 = node_id
	}
}
func (table *FileLookup) RemoveReplica2(file_name string, node_id uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[file_name].n_replicas--
	table.table[file_name].replica_id2 = node_id
}
func (table *FileLookup) RemoveMainNode(file_name string, node_id uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[file_name].n_replicas--
	if table.table[file_name].n_replicas == 3 {
		table.table[file_name].node_id = table.table[file_name].replica_id1
		table.table[file_name].replica_id1 = table.table[file_name].replica_id2
	} else if table.table[file_name].n_replicas == 2 {
		table.table[file_name].node_id = table.table[file_name].replica_id1
	}
}

func (table *FileLookup) SetFileSize(file_name string, file_size uint64) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[file_name].file_size = file_size
}

func (table *FileLookup) GetFileSize(file_name string) uint64 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[file_name].file_size
}
