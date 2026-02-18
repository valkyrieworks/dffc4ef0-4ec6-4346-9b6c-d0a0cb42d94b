package host

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"

	"github.com/valkyrieworks/iface/kinds"
	cmttrace "github.com/valkyrieworks/utils/log"
	cometnet "github.com/valkyrieworks/utils/net"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
)

//
//
//
//
//
type SocketHost struct {
	daemon.RootDaemon
	isTracerCollection bool

	schema    string
	address     string
	observer net.Listener

	linksMutex   engineconnect.Lock
	links      map[int]net.Conn
	followingLinkUID int

	applicationMutex engineconnect.Lock
	app    kinds.Software
}

const replyBufferVolume = 1000

//
func NewSocketHost(schemaAddress string, app kinds.Software) daemon.Daemon {
	schema, address := cometnet.ProtocolAndLocation(schemaAddress)
	s := &SocketHost{
		schema:    schema,
		address:     address,
		observer: nil,
		app:      app,
		links:    make(map[int]net.Conn),
	}
	s.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", s)
	return s
}

func (s *SocketHost) AssignTracer(l cmttrace.Tracer) {
	s.RootDaemon.AssignTracer(l)
	s.isTracerCollection = true
}

func (s *SocketHost) OnBegin() error {
	ln, err := net.Listen(s.schema, s.address)
	if err != nil {
		return err
	}

	s.observer = ln
	go s.allowLinkagesProcedure()

	return nil
}

func (s *SocketHost) OnHalt() {
	if err := s.observer.Close(); err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	s.linksMutex.Lock()
	defer s.linksMutex.Unlock()
	for id, link := range s.links {
		delete(s.links, id)
		if err := link.Close(); err != nil {
			s.Tracer.Fault("REDACTED", "REDACTED", id, "REDACTED", link, "REDACTED", err)
		}
	}
}

func (s *SocketHost) appendLink(link net.Conn) int {
	s.linksMutex.Lock()
	defer s.linksMutex.Unlock()

	linkUID := s.followingLinkUID
	s.followingLinkUID++
	s.links[linkUID] = link

	return linkUID
}

//
func (s *SocketHost) removeLink(linkUID int) error {
	s.linksMutex.Lock()
	defer s.linksMutex.Unlock()

	link, ok := s.links[linkUID]
	if !ok {
		return ErrLinkageDoesNotOccur{LinkUID: linkUID}
	}

	delete(s.links, linkUID)
	return link.Close()
}

func (s *SocketHost) allowLinkagesProcedure() {
	for {
		//
		s.Tracer.Details("REDACTED")
		link, err := s.observer.Accept()
		if err != nil {
			if !s.IsActive() {
				return //
			}
			s.Tracer.Fault("REDACTED", "REDACTED", err)
			continue
		}

		s.Tracer.Details("REDACTED")

		linkUID := s.appendLink(link)

		endLink := make(chan error, 2)                            //
		replies := make(chan *kinds.Reply, replyBufferVolume) //

		//
		go s.processQueries(endLink, link, replies)
		//
		go s.processReplies(endLink, link, replies)

		//
		go s.waitForEnd(endLink, linkUID)
	}
}

func (s *SocketHost) waitForEnd(endLink chan error, linkUID int) {
	err := <-endLink
	switch {
	case err == io.EOF:
		s.Tracer.Fault("REDACTED")
	case err != nil:
		s.Tracer.Fault("REDACTED", "REDACTED", err)
	default:
		//
		s.Tracer.Fault("REDACTED")
	}

	//
	if err := s.removeLink(linkUID); err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//
func (s *SocketHost) processQueries(endLink chan error, link io.Reader, replies chan<- *kinds.Reply) {
	bufferScanner := bufio.NewReader(link)

	defer func() {
		//
		//
		//
		r := recover()
		if r != nil {
			const volume = 64 << 10
			buf := make([]byte, volume)
			buf = buf[:runtime.Stack(buf, false)]
			err := fmt.Errorf("REDACTED", r, buf)
			if !s.isTracerCollection {
				fmt.Fprintln(os.Stderr, err)
			}
			endLink <- err
			s.applicationMutex.Unlock()
		}
	}()

	for {

		req := &kinds.Query{}
		err := kinds.FetchSignal(bufferScanner, req)
		if err != nil {
			if err == io.EOF {
				endLink <- err
			} else {
				endLink <- fmt.Errorf("REDACTED", err)
			}
			return
		}
		s.applicationMutex.Lock()
		reply, err := s.processQuery(context.TODO(), req)
		if err != nil {
			//
			//
			//
			replies <- kinds.ToReplyExemption(err.Error())
		} else {
			replies <- reply
		}
		s.applicationMutex.Unlock()
	}
}

//
func (s *SocketHost) processQuery(ctx context.Context, req *kinds.Query) (*kinds.Reply, error) {
	switch r := req.Item.(type) {
	case *kinds.Query_Reverberate:
		return kinds.ToReplyReverberate(r.Replicate.Signal), nil
	case *kinds.Query_Purge:
		return kinds.ToReplyPurge(), nil
	case *kinds.Query_Details:
		res, err := s.app.Details(ctx, r.Details)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyDetails(res), nil
	case *kinds.Query_Transfercheck:
		res, err := s.app.InspectTransfer(ctx, r.InspectTransfer)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyInspectTransfer(res), nil
	case *kinds.Query_Endorse:
		res, err := s.app.Endorse(ctx, r.Endorse)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyEndorse(res), nil
	case *kinds.Query_Inquire:
		res, err := s.app.Inquire(ctx, r.Inquire)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyInquire(res), nil
	case *kinds.Query_Initiatechain:
		res, err := s.app.InitSeries(ctx, r.InitSeries)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyInitSeries(res), nil
	case *kinds.Query_Terminateblock:
		res, err := s.app.CompleteLedger(ctx, r.CompleteLedger)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyCompleteLedger(res), nil
	case *kinds.Query_Catalogmirrors:
		res, err := s.app.CatalogMirrors(ctx, r.CatalogMirrors)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyCatalogMirrors(res), nil
	case *kinds.Query_Mirrorsnapshot:
		res, err := s.app.ProposalMirror(ctx, r.ProposalMirror)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyProposalMirror(res), nil
	case *kinds.Query_Arrangenomination:
		res, err := s.app.ArrangeNomination(ctx, r.ArrangeNomination)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyArrangeNomination(res), nil
	case *kinds.Query_Processnomination:
		res, err := s.app.HandleNomination(ctx, r.HandleNomination)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyHandleNomination(res), nil
	case *kinds.Query_Loadmirrorsegment:
		res, err := s.app.ImportMirrorSegment(ctx, r.ImportMirrorSegment)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyImportMirrorSegment(res), nil
	case *kinds.Query_Executemirrorsegment:
		res, err := s.app.ExecuteMirrorSegment(ctx, r.ExecuteMirrorSegment)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyExecuteMirrorSegment(res), nil
	case *kinds.Query_Ballotextend:
		res, err := s.app.ExpandBallot(ctx, r.ExpandBallot)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyExpandBallot(res), nil
	case *kinds.Query_Validateballotextension:
		res, err := s.app.ValidateBallotAddition(ctx, r.ValidateBallotAddition)
		if err != nil {
			return nil, err
		}
		return kinds.ToReplyValidateBallotAddition(res), nil
	default:
		return nil, ErrUnclearQuery{Query: *req}
	}
}

//
func (s *SocketHost) processReplies(endLink chan error, link io.Writer, replies <-chan *kinds.Reply) {
	bufferRecorder := bufio.NewWriter(link)
	for {
		res := <-replies
		err := kinds.RecordSignal(res, bufferRecorder)
		if err != nil {
			endLink <- fmt.Errorf("REDACTED", err)
			return
		}
		if _, ok := res.Item.(*kinds.Reply_Purge); ok {
			err = bufferRecorder.Flush()
			if err != nil {
				endLink <- fmt.Errorf("REDACTED", err)
				return
			}
		}

		//
		//
		//
		if e, ok := res.Item.(*kinds.Reply_Exemption); ok {
			endLink <- errors.New(e.Exemption.Fault)
		}
	}
}
