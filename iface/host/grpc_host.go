package host

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/valkyrieworks/iface/kinds"
	cometnet "github.com/valkyrieworks/utils/net"
	"github.com/valkyrieworks/utils/daemon"
)

type GRPCHost struct {
	daemon.RootDaemon

	schema    string
	address     string
	observer net.Listener
	host   *grpc.Server

	app kinds.Software
}

//
func NewGRPCHost(schemaAddress string, app kinds.Software) daemon.Daemon {
	schema, address := cometnet.ProtocolAndLocation(schemaAddress)
	s := &GRPCHost{
		schema:    schema,
		address:     address,
		observer: nil,
		app:      app,
	}
	s.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", s)
	return s
}

//
func (s *GRPCHost) OnBegin() error {
	ln, err := net.Listen(s.schema, s.address)
	if err != nil {
		return err
	}

	s.observer = ln
	s.host = grpc.NewServer()
	kinds.EnrollIfaceHost(s.host, &gRPCSoftware{s.app})

	s.Tracer.Details("REDACTED", "REDACTED", s.schema, "REDACTED", s.address)
	go func() {
		if err := s.host.Serve(s.observer); err != nil {
			s.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}()
	return nil
}

//
func (s *GRPCHost) OnHalt() {
	s.host.Stop()
}

//

//
type gRPCSoftware struct {
	kinds.Software
}

func (app *gRPCSoftware) Replicate(_ context.Context, req *kinds.QueryReverberate) (*kinds.ReplyReverberate, error) {
	return &kinds.ReplyReverberate{Signal: req.Signal}, nil
}

func (app *gRPCSoftware) Purge(context.Context, *kinds.QueryPurge) (*kinds.ReplyPurge, error) {
	return &kinds.ReplyPurge{}, nil
}
