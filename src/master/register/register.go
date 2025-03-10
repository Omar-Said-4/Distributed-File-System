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
	node_id := NodesTable.GetNodeCount()
	NodesTable.AddDataNode(node_id, ip, file_port, replication_port, ncopyPort)
	return &register.RegisterResponse{Id: node_id, Success: true}, nil
}

func StartRegisterServer(table *lookup.NodeLookup, port string, s *grpc.Server) {
	NodesTable = table

	register.RegisterRegisterServiceServer(s, &registerServer{})
	fmt.Printf("Register Server is running on port: %s\n", port)

}
