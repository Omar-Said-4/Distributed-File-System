package heartbeat

import (
	"context"
	lookup "dfs/master/lookup/node"
	"dfs/schema/heartbeat"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var NodesTable *lookup.NodeLookup

type heartbeatServer struct {
	heartbeat.UnimplementedHeartbeatServiceServer
}

func (s *heartbeatServer) Ping(ctx context.Context, req *heartbeat.HeartbeatPing) (*heartbeat.HeartbeatPong, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get peer info from context")
	}
	node_id := req.NodeId
	ip := p.Addr.String()
	fmt.Printf("Ping from NodeID: %d, IP: %s\n", node_id, ip)
	NodesTable.UpdateNodePingTime(node_id)
	// revive the node if it was dead
	if !NodesTable.GetNodeAlive(node_id) {
		NodesTable.SetNodeAlive(node_id)
	}
	return &heartbeat.HeartbeatPong{}, nil
}

func StartHeartbeatServer(table *lookup.NodeLookup, port string) {
	NodesTable = table
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	heartbeat.RegisterHeartbeatServiceServer(s, &heartbeatServer{})
	fmt.Printf("HeartBeat Server is running on port: %s\n", port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
