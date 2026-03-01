package node

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"
	"time"

	"github.com/gorilla/websocket"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//

const (
	fallbackSocketRecordChnVolume = 100
	fallbackSocketRecordPause         = 10 * time.Second
	fallbackSocketFetchPause          = 30 * time.Second
	fallbackSocketPingSpan        = (fallbackSocketFetchPause * 9) / 10
)

//
//
//
type WebterminalAdministrator struct {
	websocket.Enhancer

	methodIndex       map[string]*RemoteMethod
	tracer        log.Tracer
	socketLinkChoices []func(*socketLinkage)
}

//
//
func FreshWebterminalAdministrator(
	methodIndex map[string]*RemoteMethod,
	socketLinkChoices ...func(*socketLinkage),
) *WebterminalAdministrator {
	return &WebterminalAdministrator{
		methodIndex: methodIndex,
		Enhancer: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				//
				//
				//
				//
				//
				//
				//
				//
				//
				return true
			},
		},
		tracer:        log.FreshNooperationTracer(),
		socketLinkChoices: socketLinkChoices,
	}
}

//
func (wm *WebterminalAdministrator) AssignTracer(l log.Tracer) {
	wm.tracer = l
}

//
//
func (wm *WebterminalAdministrator) WebterminalProcessor(w http.ResponseWriter, r *http.Request) {
	socketLink, err := wm.Upgrade(w, r, nil)
	if err != nil {
		//
		wm.tracer.Failure("REDACTED", "REDACTED", err)
		return
	}
	defer func() {
		if err := socketLink.Close(); err != nil {
			wm.tracer.Failure("REDACTED", "REDACTED", err)
		}
	}()

	//
	con := freshSocketLinkage(socketLink, wm.methodIndex, wm.socketLinkChoices...)
	con.AssignTracer(wm.tracer.Using("REDACTED", socketLink.RemoteAddr()))
	wm.tracer.Details("REDACTED", "REDACTED", con.distantLocation)
	err = con.Initiate() //
	if err != nil {
		wm.tracer.Failure("REDACTED", "REDACTED", err)
		return
	}
	if err := con.Halt(); err != nil {
		wm.tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//

//
//
//
//
type socketLinkage struct {
	facility.FoundationFacility

	distantLocation string
	foundationLink   *websocket.Conn
	//
	recordChn chan kinds.RemoteReply

	//
	//
	fetchProcedureExit chan struct{}

	methodIndex map[string]*RemoteMethod

	//
	recordChnVolume int

	//
	recordPause time.Duration

	//
	fetchPause time.Duration

	//
	pingSpan time.Duration

	//
	fetchThreshold int64

	//
	uponDetach func(distantLocation string)

	ctx    context.Context
	abort context.CancelFunc
}

//
//
//
//
//
//
func freshSocketLinkage(
	foundationLink *websocket.Conn,
	methodIndex map[string]*RemoteMethod,
	choices ...func(*socketLinkage),
) *socketLinkage {
	wsc := &socketLinkage{
		distantLocation:        foundationLink.RemoteAddr().String(),
		foundationLink:          foundationLink,
		methodIndex:           methodIndex,
		recordPause:         fallbackSocketRecordPause,
		recordChnVolume: fallbackSocketRecordChnVolume,
		fetchPause:          fallbackSocketFetchPause,
		pingSpan:        fallbackSocketPingSpan,
		fetchProcedureExit:   make(chan struct{}),
	}
	for _, selection := range choices {
		selection(wsc)
	}
	wsc.foundationLink.SetReadLimit(wsc.fetchThreshold)
	wsc.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", wsc)
	return wsc
}

//
//
func UponDetach(uponDetach func(distantLocation string)) func(*socketLinkage) {
	return func(wsc *socketLinkage) {
		wsc.uponDetach = uponDetach
	}
}

//
//
func RecordPause(recordPause time.Duration) func(*socketLinkage) {
	return func(wsc *socketLinkage) {
		wsc.recordPause = recordPause
	}
}

//
//
func PersistChnVolume(cap int) func(*socketLinkage) {
	return func(wsc *socketLinkage) {
		wsc.recordChnVolume = cap
	}
}

//
//
func FetchPause(fetchPause time.Duration) func(*socketLinkage) {
	return func(wsc *socketLinkage) {
		wsc.fetchPause = fetchPause
	}
}

//
//
func PingSpan(pingSpan time.Duration) func(*socketLinkage) {
	return func(wsc *socketLinkage) {
		wsc.pingSpan = pingSpan
	}
}

//
//
func RetrieveThreshold(fetchThreshold int64) func(*socketLinkage) {
	return func(wsc *socketLinkage) {
		wsc.fetchThreshold = fetchThreshold
	}
}

//
//
func (wsc *socketLinkage) UponInitiate() error {
	wsc.recordChn = make(chan kinds.RemoteReply, wsc.recordChnVolume)

	//
	go wsc.fetchProcedure()
	//
	wsc.recordProcedure()

	return nil
}

//
//
func (wsc *socketLinkage) UponHalt() {
	if wsc.uponDetach != nil {
		wsc.uponDetach(wsc.distantLocation)
	}

	if wsc.ctx != nil {
		wsc.abort()
	}
}

//
//
func (wsc *socketLinkage) ObtainDistantLocation() string {
	return wsc.distantLocation
}

//
//
//
func (wsc *socketLinkage) PersistRemoteReply(ctx context.Context, reply kinds.RemoteReply) error {
	select {
	case <-wsc.Exit():
		return errors.New("REDACTED")
	case <-ctx.Done():
		return ctx.Err()
	case wsc.recordChn <- reply:
		return nil
	}
}

//
//
//
func (wsc *socketLinkage) AttemptPersistRemoteReply(reply kinds.RemoteReply) bool {
	select {
	case <-wsc.Exit():
		return false
	case wsc.recordChn <- reply:
		return true
	default:
		return false
	}
}

//
//
func (wsc *socketLinkage) Env() context.Context {
	if wsc.ctx != nil {
		return wsc.ctx
	}
	wsc.ctx, wsc.abort = context.WithCancel(context.Background())
	return wsc.ctx
}

//
func (wsc *socketLinkage) fetchProcedure() {
	//
	persistContext := context.Background()

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("REDACTED", r)
			}
			wsc.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", string(debug.Stack()))
			if err := wsc.PersistRemoteReply(persistContext, kinds.RemoteIntrinsicFailure(kinds.JsonifaceIntegerUUID(-1), err)); err != nil {
				wsc.Tracer.Failure("REDACTED", "REDACTED", err)
			}
			go wsc.fetchProcedure()
		}
	}()

	wsc.foundationLink.SetPongHandler(func(m string) error {
		return wsc.foundationLink.SetReadDeadline(time.Now().Add(wsc.fetchPause))
	})

	for {
		select {
		case <-wsc.Exit():
			return
		default:
			//
			if err := wsc.foundationLink.SetReadDeadline(time.Now().Add(wsc.fetchPause)); err != nil {
				wsc.Tracer.Failure("REDACTED", "REDACTED", err)
			}

			_, r, err := wsc.foundationLink.NextReader()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					wsc.Tracer.Details("REDACTED")
				} else {
					wsc.Tracer.Failure("REDACTED", "REDACTED", err)
				}
				if err := wsc.Halt(); err != nil {
					wsc.Tracer.Failure("REDACTED", "REDACTED", err)
				}
				close(wsc.fetchProcedureExit)
				return
			}

			dec := json.NewDecoder(r)
			var solicit kinds.RemoteSolicit
			err = dec.Decode(&solicit)
			if err != nil {
				if err := wsc.PersistRemoteReply(persistContext,
					kinds.RemoteAnalyzeFailure(fmt.Errorf("REDACTED", err))); err != nil {
					wsc.Tracer.Failure("REDACTED", "REDACTED", err)
				}
				continue
			}

			//
			//
			if solicit.ID == nil {
				wsc.Tracer.Diagnose(
					"REDACTED",
					"REDACTED", solicit,
				)
				continue
			}

			//
			remoteMethod := wsc.methodIndex[solicit.Procedure]
			if remoteMethod == nil {
				if err := wsc.PersistRemoteReply(persistContext, kinds.RemoteProcedureNegationDetectedFailure(solicit.ID)); err != nil {
					wsc.Tracer.Failure("REDACTED", "REDACTED", err)
				}
				continue
			}

			ctx := &kinds.Env{JSNRequest: &solicit, SocketLink: wsc}
			arguments := []reflect.Value{reflect.ValueOf(ctx)}
			if len(solicit.Parameters) > 0 {
				procArguments, err := jsnParametersTowardArguments(remoteMethod, solicit.Parameters)
				if err != nil {
					if err := wsc.PersistRemoteReply(persistContext,
						kinds.RemoteIntrinsicFailure(solicit.ID, fmt.Errorf("REDACTED", err)),
					); err != nil {
						wsc.Tracer.Failure("REDACTED", "REDACTED", err)
					}
					continue
				}
				arguments = append(arguments, procArguments...)
			}

			yields := remoteMethod.f.Call(arguments)

			//
			wsc.Tracer.Details("REDACTED", "REDACTED", solicit.Procedure)

			outcome, err := unmirrorOutcome(yields)
			if err != nil {
				if err := wsc.PersistRemoteReply(persistContext, kinds.RemoteIntrinsicFailure(solicit.ID, err)); err != nil {
					wsc.Tracer.Failure("REDACTED", "REDACTED", err)
				}
				continue
			}

			if err := wsc.PersistRemoteReply(persistContext, kinds.FreshRemoteTriumphReply(solicit.ID, outcome)); err != nil {
				wsc.Tracer.Failure("REDACTED", "REDACTED", err)
			}
		}
	}
}

//
func (wsc *socketLinkage) recordProcedure() {
	pingMetronome := time.NewTicker(wsc.pingSpan)
	defer pingMetronome.Stop()

	//
	pongs := make(chan string, 1)
	wsc.foundationLink.SetPingHandler(func(m string) error {
		select {
		case pongs <- m:
		default:
		}
		return nil
	})

	for {
		select {
		case <-wsc.Exit():
			return
		case <-wsc.fetchProcedureExit: //
			return
		case m := <-pongs:
			err := wsc.recordArtifactUsingExpiration(websocket.PongMessage, []byte(m))
			if err != nil {
				wsc.Tracer.Details("REDACTED", "REDACTED", err)
			}
		case <-pingMetronome.C:
			err := wsc.recordArtifactUsingExpiration(websocket.PingMessage, []byte{})
			if err != nil {
				wsc.Tracer.Failure("REDACTED", "REDACTED", err)
				return
			}
		case msg := <-wsc.recordChn:
			//
			//
			//
			jsnOctets, err := json.Marshal(msg)
			if err != nil {
				wsc.Tracer.Failure("REDACTED", "REDACTED", err)
				continue
			}
			if err = wsc.recordArtifactUsingExpiration(websocket.TextMessage, jsnOctets); err != nil {
				wsc.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", msg)
				return
			}
		}
	}
}

//
//
//
func (wsc *socketLinkage) recordArtifactUsingExpiration(signalKind int, msg []byte) error {
	if err := wsc.foundationLink.SetWriteDeadline(time.Now().Add(wsc.recordPause)); err != nil {
		return err
	}
	return wsc.foundationLink.WriteMessage(signalKind, msg)
}
