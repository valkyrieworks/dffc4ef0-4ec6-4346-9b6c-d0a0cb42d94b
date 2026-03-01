package node

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

type GRPSDaemon struct {
	facility.FoundationFacility

	schema    string
	location     string
	observer net.Listener
	node   *grpc.Server

	app kinds.Platform
}

//
func FreshGRPSDaemon(schemaLocation string, app kinds.Platform) facility.Facility {
	schema, location := strongmindnet.SchemeAlsoLocation(schemaLocation)
	s := &GRPSDaemon{
		schema:    schema,
		location:     location,
		observer: nil,
		app:      app,
	}
	s.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", s)
	return s
}

//
func (s *GRPSDaemon) UponInitiate() error {
	ln, err := net.Listen(s.schema, s.location)
	if err != nil {
		return err
	}

	s.observer = ln
	s.node = grpc.NewServer()
	kinds.EnrollIfaceDaemon(s.node, &gRemotePlatform{s.app})

	s.Tracer.Details("REDACTED", "REDACTED", s.schema, "REDACTED", s.location)
	go func() {
		if err := s.node.Serve(s.observer); err != nil {
			s.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}()
	return nil
}

//
func (s *GRPSDaemon) UponHalt() {
	s.node.Stop()
}

//

//
type gRemotePlatform struct {
	kinds.Platform
}

func (app *gRemotePlatform) Reverberate(_ context.Context, req *kinds.SolicitReverberate) (*kinds.ReplyReverberate, error) {
	return &kinds.ReplyReverberate{Signal: req.Signal}, nil
}

func (app *gRemotePlatform) Purge(context.Context, *kinds.SolicitPurge) (*kinds.ReplyPurge, error) {
	return &kinds.ReplyPurge{}, nil
}
