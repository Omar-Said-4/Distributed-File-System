package heartbeat

import (
	"context"
	lookup "dfs/master/lookup/node"
	"dfs/schema/heartbeat"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var NodesTable *lookup.NodeLookup

type heartbeatServer struct {
	heartbeat.UnimplementedHeartbeatServiceServer
}

func (s *heartbeatServer) Ping(ctx context.Context, req *heartbeat.HeartbeatPing) (*heartbeat.HeartbeatPong, error) {
	_, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get peer info from context")
	}
	node_id := req.NodeId
	// ip := p.Addr.String()
	// fmt.Printf("Ping from NodeID: %d, IP: %s\n", node_id, ip)
	NodesTable.UpdateNodePingTime(node_id)
	// revive the node if it was dead
	if !NodesTable.GetNodeAlive(node_id) {
		NodesTable.SetNodeAlive(node_id)
		fmt.Printf("Node %d is alive again\n", node_id)
	}
	return &heartbeat.HeartbeatPong{}, nil
}

func StartHeartbeatServer(table *lookup.NodeLookup, port string, s *grpc.Server) {
	NodesTable = table

	heartbeat.RegisterHeartbeatServiceServer(s, &heartbeatServer{})
	fmt.Printf("HeartBeat Server is running on port: %s\n", port)

}
