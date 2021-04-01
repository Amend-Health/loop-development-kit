package ldk

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/client"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
	"google.golang.org/grpc"
)

// LoopPluginMap is the map of plugins we can dispense.
var LoopPluginMap = map[string]plugin.Plugin{
	"loop": &LoopPlugin{},
}

// LoopPlugin is a structure used to define the parameters of the plugin communication
type LoopPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl   Loop
	logger *utils.Logger
	loopServer *LoopServer
}

// GRPCServer is used to register the controller plugin with the GRPC server
func (p *LoopPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterLoopServer(s, &LoopServer{
		Impl:   p.Impl,
		Broker: broker,
		Logger: p.logger,
	})

	return nil
}

// GRPCClient is used to generate controller clients that can be used by the host
func (p *LoopPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &client.LoopClient{broker, proto.NewLoopClient(c), nil}, nil
}