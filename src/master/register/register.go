package register

import (
	"context"
	lookup "dfs/master/lookup/node"
	"dfs/schema/register"
	"fmt"

	"google.golang.org/grpc"
)

var NodesTable *lookup.NodeLookup

type registerServer struct {
	register.UnimplementedRegisterServiceServer
}

func (s *registerServer) Register(ctx context.Context, req *register.RegisterRequest) (*register.RegisterResponse, error) {
	file_port := req.FilePort
	replication_port := req.ReplicationPort
	ncopyPort := req.NotifyToCopyPort
	ip := req.Ip
	old_id := req.OldId
	var node_id uint32
	if old_id == -1 {
		node_id = NodesTable.GetNodeCount()
		fmt.Printf("Registering new node with ID: %d\n", node_id)
		NodesTable.AddDataNode(node_id, ip, file_port, replication_port, ncopyPort)
	} else {
		node_id = uint32(old_id)
		fmt.Printf("Updating existing node with ID: %d\n", node_id)
		NodesTable.EditDataNode(node_id, ip, file_port, replication_port, ncopyPort)
	}
	return &register.RegisterResponse{Id: node_id, Success: true}, nil
}

func StartRegisterServer(table *lookup.NodeLookup, port string, s *grpc.Server) {
	NodesTable = table

	register.RegisterRegisterServiceServer(s, &registerServer{})
	fmt.Printf("Register Server is running on port: %s\n", port)

}
