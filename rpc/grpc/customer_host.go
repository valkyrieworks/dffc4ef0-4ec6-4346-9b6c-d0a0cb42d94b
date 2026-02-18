package coregrpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	cometnet "github.com/valkyrieworks/utils/net"
	"github.com/valkyrieworks/rpc/core"
)

//
//
//
type Settings struct {
	MaximumAccessLinks int
}

//
//
//
//
//
func BeginGRPCHost(env *core.Context, ln net.Listener) error {
	grpcHost := grpc.NewServer()
	EnrollMulticastAPIHost(grpcHost, &multicastAPI{env: env})
	return grpcHost.Serve(ln)
}

//
//
//
//
func BeginGRPCCustomer(schemaAddress string) MulticastAPICustomer {
	link, err := grpc.Dial(schemaAddress, grpc.WithInsecure(), grpc.WithContextDialer(callerFunction))
	if err != nil {
		panic(err)
	}
	return NewMulticastAPICustomer(link)
}

func callerFunction(_ context.Context, address string) (net.Conn, error) {
	return cometnet.Link(address)
}
