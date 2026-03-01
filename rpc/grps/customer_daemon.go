package coregrpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base"
)

//
//
//
type Settings struct {
	MaximumInitiateLinks int
}

//
//
//
//
//
func InitiateGRPSDaemon(env *base.Context, ln net.Listener) error {
	grpsDaemon := grpc.NewServer()
	EnrollMulticastAPIDaemon(grpsDaemon, &multicastAPI{env: env})
	return grpsDaemon.Serve(ln)
}

//
//
//
//
func InitiateGRPSCustomer(schemaLocation string) MulticastAPICustomer {
	link, err := grpc.Dial(schemaLocation, grpc.WithInsecure(), grpc.WithContextDialer(callerMethod))
	if err != nil {
		panic(err)
	}
	return FreshMulticastAPICustomer(link)
}

func callerMethod(_ context.Context, location string) (net.Conn, error) {
	return strongmindnet.Relate(location)
}
