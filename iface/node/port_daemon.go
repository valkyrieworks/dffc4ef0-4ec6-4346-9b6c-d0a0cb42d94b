package node

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	endorsementlog "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
//
//
//
//
type PortDaemon struct {
	facility.FoundationFacility
	equalsTracerAssign bool

	schema    string
	location     string
	observer net.Listener

	linksMutex   commitchronize.Exclusion
	links      map[int]net.Conn
	followingLinkUUID int

	applicationMutex commitchronize.Exclusion
	app    kinds.Platform
}

const replyReserveExtent = 1000

//
func FreshPortDaemon(schemaLocation string, app kinds.Platform) facility.Facility {
	schema, location := strongmindnet.SchemeAlsoLocation(schemaLocation)
	s := &PortDaemon{
		schema:    schema,
		location:     location,
		observer: nil,
		app:      app,
		links:    make(map[int]net.Conn),
	}
	s.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", s)
	return s
}

func (s *PortDaemon) AssignTracer(l endorsementlog.Tracer) {
	s.FoundationFacility.AssignTracer(l)
	s.equalsTracerAssign = true
}

func (s *PortDaemon) UponInitiate() error {
	ln, err := net.Listen(s.schema, s.location)
	if err != nil {
		return err
	}

	s.observer = ln
	go s.embraceLinkagesProcedure()

	return nil
}

func (s *PortDaemon) UponHalt() {
	if err := s.observer.Close(); err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	s.linksMutex.Lock()
	defer s.linksMutex.Unlock()
	for id, link := range s.links {
		delete(s.links, id)
		if err := link.Close(); err != nil {
			s.Tracer.Failure("REDACTED", "REDACTED", id, "REDACTED", link, "REDACTED", err)
		}
	}
}

func (s *PortDaemon) appendLink(link net.Conn) int {
	s.linksMutex.Lock()
	defer s.linksMutex.Unlock()

	linkUUID := s.followingLinkUUID
	s.followingLinkUUID++
	s.links[linkUUID] = link

	return linkUUID
}

//
func (s *PortDaemon) delLink(linkUUID int) error {
	s.linksMutex.Lock()
	defer s.linksMutex.Unlock()

	link, ok := s.links[linkUUID]
	if !ok {
		return FaultLinkageExecutesNegationPrevail{LinkUUID: linkUUID}
	}

	delete(s.links, linkUUID)
	return link.Close()
}

func (s *PortDaemon) embraceLinkagesProcedure() {
	for {
		//
		s.Tracer.Details("REDACTED")
		link, err := s.observer.Accept()
		if err != nil {
			if !s.EqualsActive() {
				return //
			}
			s.Tracer.Failure("REDACTED", "REDACTED", err)
			continue
		}

		s.Tracer.Details("REDACTED")

		linkUUID := s.appendLink(link)

		shutdownLink := make(chan error, 2)                            //
		replies := make(chan *kinds.Reply, replyReserveExtent) //

		//
		go s.processSolicits(shutdownLink, link, replies)
		//
		go s.processReplies(shutdownLink, link, replies)

		//
		go s.pauseForeachShutdown(shutdownLink, linkUUID)
	}
}

func (s *PortDaemon) pauseForeachShutdown(shutdownLink chan error, linkUUID int) {
	err := <-shutdownLink
	switch {
	case err == io.EOF:
		s.Tracer.Failure("REDACTED")
	case err != nil:
		s.Tracer.Failure("REDACTED", "REDACTED", err)
	default:
		//
		s.Tracer.Failure("REDACTED")
	}

	//
	if err := s.delLink(linkUUID); err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//
func (s *PortDaemon) processSolicits(shutdownLink chan error, link io.Reader, replies chan<- *kinds.Reply) {
	bufferFetcher := bufio.NewReader(link)

	defer func() {
		//
		//
		//
		r := recover()
		if r != nil {
			const extent = 64 << 10
			buf := make([]byte, extent)
			buf = buf[:runtime.Stack(buf, false)]
			err := fmt.Errorf("REDACTED", r, buf)
			if !s.equalsTracerAssign {
				fmt.Fprintln(os.Stderr, err)
			}
			shutdownLink <- err
			s.applicationMutex.Unlock()
		}
	}()

	for {

		req := &kinds.Solicit{}
		err := kinds.FetchArtifact(bufferFetcher, req)
		if err != nil {
			if err == io.EOF {
				shutdownLink <- err
			} else {
				shutdownLink <- fmt.Errorf("REDACTED", err)
			}
			return
		}
		s.applicationMutex.Lock()
		reply, err := s.processSolicit(context.TODO(), req)
		if err != nil {
			//
			//
			//
			replies <- kinds.TowardReplyExemption(err.Error())
		} else {
			replies <- reply
		}
		s.applicationMutex.Unlock()
	}
}

//
func (s *PortDaemon) processSolicit(ctx context.Context, req *kinds.Solicit) (*kinds.Reply, error) {
	switch r := req.Datum.(type) {
	case *kinds.Solicit_Reverberate:
		return kinds.TowardReplyReverberate(r.Reverberate.Signal), nil
	case *kinds.Solicit_Purge:
		return kinds.TowardReplyPurge(), nil
	case *kinds.Solicit_Details:
		res, err := s.app.Details(ctx, r.Details)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyDetails(res), nil
	case *kinds.Solicit_Inspecttrans:
		res, err := s.app.InspectTransfer(ctx, r.InspectTransfer)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyInspectTransfer(res), nil
	case *kinds.Solicit_Endorse:
		res, err := s.app.Endorse(ctx, r.Endorse)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyEndorse(res), nil
	case *kinds.Solicit_Inquire:
		res, err := s.app.Inquire(ctx, r.Inquire)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyInquire(res), nil
	case *kinds.Solicit_Initiatechain:
		res, err := s.app.InitializeSuccession(ctx, r.InitializeSuccession)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyInitializeSuccession(res), nil
	case *kinds.Solicit_Finalizeledger:
		res, err := s.app.CulminateLedger(ctx, r.CulminateLedger)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyCulminateLedger(res), nil
	case *kinds.Solicit_Catalogimages:
		res, err := s.app.CollectionImages(ctx, r.CollectionImages)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyCatalogImages(res), nil
	case *kinds.Solicit_Extendimage:
		res, err := s.app.ExtendImage(ctx, r.ExtendImage)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyExtendImage(res), nil
	case *kinds.Solicit_Prepareitem:
		res, err := s.app.ArrangeNomination(ctx, r.ArrangeNomination)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyArrangeNomination(res), nil
	case *kinds.Solicit_Executeitem:
		res, err := s.app.HandleNomination(ctx, r.HandleNomination)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyHandleNomination(res), nil
	case *kinds.Solicit_Loadimagefragment:
		res, err := s.app.FetchImageSegment(ctx, r.FetchImageSegment)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyFetchImageSegment(res), nil
	case *kinds.Solicit_Executeimagefragment:
		res, err := s.app.ExecuteImageSegment(ctx, r.ExecuteImageSegment)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyExecuteImageSegment(res), nil
	case *kinds.Solicit_Extendballot:
		res, err := s.app.BroadenBallot(ctx, r.BroadenBallot)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyBroadenBallot(res), nil
	case *kinds.Solicit_Verifyballotaddition:
		res, err := s.app.ValidateBallotAddition(ctx, r.ValidateBallotAddition)
		if err != nil {
			return nil, err
		}
		return kinds.TowardReplyValidateBallotAddition(res), nil
	default:
		return nil, FaultUnfamiliarSolicit{Solicit: *req}
	}
}

//
func (s *PortDaemon) processReplies(shutdownLink chan error, link io.Writer, replies <-chan *kinds.Reply) {
	bufferPersistor := bufio.NewWriter(link)
	for {
		res := <-replies
		err := kinds.PersistArtifact(res, bufferPersistor)
		if err != nil {
			shutdownLink <- fmt.Errorf("REDACTED", err)
			return
		}
		if _, ok := res.Datum.(*kinds.Reply_Purge); ok {
			err = bufferPersistor.Flush()
			if err != nil {
				shutdownLink <- fmt.Errorf("REDACTED", err)
				return
			}
		}

		//
		//
		//
		if e, ok := res.Datum.(*kinds.Reply_Exemption); ok {
			shutdownLink <- errors.New(e.Exemption.Failure)
		}
	}
}
