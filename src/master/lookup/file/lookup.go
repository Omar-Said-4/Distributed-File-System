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
	n_uploading uint32
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
		file_name:   file_name,
		node_id:     node_id,
		file_path:   filepath,
		file_size:   file_size,
		n_replicas:  1,
		n_uploading: 0,
	}
	table.mutex.Lock()
	defer table.mutex.Unlock()
	table.table[file_name] = file
}

func (table *FileLookup) RemoveFile(file_name string) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return
	}
	delete(table.table, file_name)

}
func (table *FileLookup) GetNumberOfReplicas(file_name string) uint32 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	_, ok := table.table[file_name]
	if !ok {
		return 0
	}
	return table.table[file_name].n_replicas
}
func (table *FileLookup) GetFileLocation(file_name string) (uint32, uint32, uint32, error) {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	entry, ok := table.table[file_name]
	if !ok || entry.n_replicas == 0 {

		return 0, 0, 0, fmt.Errorf("File %s not found or has no replicas", file_name)
	}
	if table.table[file_name].n_replicas == 1 {
		return table.table[file_name].node_id, table.table[file_name].node_id, table.table[file_name].node_id, nil
	}
	if table.table[file_name].n_replicas == 2 {

		return table.table[file_name].node_id, table.table[file_name].replica_id1, table.table[file_name].node_id, nil
	}
	return table.table[file_name].node_id, table.table[file_name].replica_id1, table.table[file_name].replica_id2, nil
}
func (table *FileLookup) AddReplica(file_name string, node_id uint32, filepath string) {
	table.mutex.Lock()

	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return
	}
	table.table[file_name].n_replicas++
	if table.table[file_name].n_replicas == 2 {
		table.table[file_name].replica_id1 = node_id
		table.table[file_name].file_path1 = filepath
	}
	if table.table[file_name].n_replicas == 3 {
		table.table[file_name].replica_id2 = node_id
		table.table[file_name].file_path2 = filepath
	}
	// fmt.Printf("FileLookup: %v\n", table.table[file_name])
}
func (table *FileLookup) RemoveReplica1(file_name string, node_id uint32) {
	table.mutex.Lock()
	// fmt.Printf("Acquired RemoveReplica1 Lock\n")

	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return
	}
	table.table[file_name].n_replicas--
	if table.table[file_name].n_replicas == 2 {
		table.table[file_name].replica_id1 = table.table[file_name].replica_id2
		table.table[file_name].replica_id2 = node_id
	} else if table.table[file_name].n_replicas == 1 {
		table.table[file_name].replica_id1 = node_id
	}
	// fmt.Printf("Left RemoveReplica1 Lock\n")

}
func (table *FileLookup) RemoveReplica2(file_name string, node_id uint32) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return
	}
	table.table[file_name].n_replicas--
	table.table[file_name].replica_id2 = node_id
}
func (table *FileLookup) RemoveMainNode(file_name string, node_id uint32) {
	table.mutex.Lock()
	// fmt.Printf("Acquired RemoveMainNode Lock\n")
	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return
	}
	table.table[file_name].n_replicas--
	if table.table[file_name].n_replicas == 2 {
		table.table[file_name].node_id = table.table[file_name].replica_id1
		table.table[file_name].replica_id1 = table.table[file_name].replica_id2
	} else if table.table[file_name].n_replicas == 1 {
		table.table[file_name].node_id = table.table[file_name].replica_id1
	} else if table.table[file_name].n_replicas == 0 {
		fmt.Printf("File %s is not available anymore in the \n", file_name)
		// table.RemoveFile(file_name)
		delete(table.table, file_name)

		// fmt.Print(table)
	}
	// fmt.Printf("Left RemoveMainNode Lock\n")

}

func (table *FileLookup) SetFileSize(file_name string, file_size uint64) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return
	}
	table.table[file_name].file_size = file_size
}

func (table *FileLookup) GetFileSize(file_name string) uint64 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	_, ok := table.table[file_name]
	if !ok {
		return 0
	}
	return table.table[file_name].file_size
}

func (table *FileLookup) GetNumberUploading(file_name string) uint32 {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	_, ok := table.table[file_name]
	if !ok {
		return 0
	}
	return table.table[file_name].n_uploading
}

func (table *FileLookup) IncrementNumberUploading(file_name string) error {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return fmt.Errorf("No such file %s", file_name)
	}
	if table.table[file_name].n_uploading >= 2 {
		return fmt.Errorf("File %s is already being uploaded by 2 nodes", file_name)
	}
	table.table[file_name].n_uploading++
	return nil
}

func (table *FileLookup) DecrementNumberUploading(file_name string) {
	table.mutex.Lock()
	defer table.mutex.Unlock()
	_, ok := table.table[file_name]
	if !ok {
		return
	}

	table.table[file_name].n_uploading--
}

func (table *FileLookup) GetFilePaths(file_name string) (string, string, string) {
	table.mutex.RLock()
	defer table.mutex.RUnlock()
	return table.table[file_name].file_path, table.table[file_name].file_path1, table.table[file_name].file_path2
}

func (table *FileLookup) GetFileNames() []string {
	table.mutex.RLock()

	defer table.mutex.RUnlock()
	var result []string
	for k := range table.table {
		result = append(result, k)
	}
	return result
}
