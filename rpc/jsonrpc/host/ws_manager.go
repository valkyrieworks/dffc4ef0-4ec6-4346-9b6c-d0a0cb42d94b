package host

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

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//

const (
	standardWSRecordChannelAbility = 100
	standardWSRecordWait         = 10 * time.Second
	standardWSReadWait          = 30 * time.Second
	standardWSPingDuration        = (standardWSReadWait * 9) / 10
)

//
//
//
type WebchannelAdministrator struct {
	websocket.Converter

	functionIndex       map[string]*RPCFunction
	tracer        log.Tracer
	wsLinkSettings []func(*wsLinkage)
}

//
//
func NewWebchannelOverseer(
	functionIndex map[string]*RPCFunction,
	wsLinkSettings ...func(*wsLinkage),
) *WebchannelAdministrator {
	return &WebchannelAdministrator{
		functionIndex: functionIndex,
		Converter: websocket.Upgrader{
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
		tracer:        log.NewNoopTracer(),
		wsLinkSettings: wsLinkSettings,
	}
}

//
func (wm *WebchannelAdministrator) AssignTracer(l log.Tracer) {
	wm.tracer = l
}

//
//
func (wm *WebchannelAdministrator) WebchannelManager(w http.ResponseWriter, r *http.Request) {
	wsLink, err := wm.Upgrade(w, r, nil)
	if err != nil {
		//
		wm.tracer.Fault("REDACTED", "REDACTED", err)
		return
	}
	defer func() {
		if err := wsLink.Close(); err != nil {
			wm.tracer.Fault("REDACTED", "REDACTED", err)
		}
	}()

	//
	con := newWSLinkage(wsLink, wm.functionIndex, wm.wsLinkSettings...)
	con.AssignTracer(wm.tracer.With("REDACTED", wsLink.RemoteAddr()))
	wm.tracer.Details("REDACTED", "REDACTED", con.distantAddress)
	err = con.Begin() //
	if err != nil {
		wm.tracer.Fault("REDACTED", "REDACTED", err)
		return
	}
	if err := con.Halt(); err != nil {
		wm.tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//

//
//
//
//
type wsLinkage struct {
	daemon.RootDaemon

	distantAddress string
	rootLink   *websocket.Conn
	//
	recordChannel chan kinds.RPCAnswer

	//
	//
	readProcedureExit chan struct{}

	functionIndex map[string]*RPCFunction

	//
	recordChannelAbility int

	//
	recordWait time.Duration

	//
	readWait time.Duration

	//
	pingDuration time.Duration

	//
	readCeiling int64

	//
	onDetach func(distantAddress string)

	ctx    context.Context
	revoke context.CancelFunc
}

//
//
//
//
//
//
func newWSLinkage(
	rootLink *websocket.Conn,
	functionIndex map[string]*RPCFunction,
	options ...func(*wsLinkage),
) *wsLinkage {
	wsc := &wsLinkage{
		distantAddress:        rootLink.RemoteAddr().String(),
		rootLink:          rootLink,
		functionIndex:           functionIndex,
		recordWait:         standardWSRecordWait,
		recordChannelAbility: standardWSRecordChannelAbility,
		readWait:          standardWSReadWait,
		pingDuration:        standardWSPingDuration,
		readProcedureExit:   make(chan struct{}),
	}
	for _, setting := range options {
		setting(wsc)
	}
	wsc.rootLink.SetReadLimit(wsc.readCeiling)
	wsc.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", wsc)
	return wsc
}

//
//
func OnDetach(onDetach func(distantAddress string)) func(*wsLinkage) {
	return func(wsc *wsLinkage) {
		wsc.onDetach = onDetach
	}
}

//
//
func RecordWait(recordWait time.Duration) func(*wsLinkage) {
	return func(wsc *wsLinkage) {
		wsc.recordWait = recordWait
	}
}

//
//
func RecordChannelAbility(cap int) func(*wsLinkage) {
	return func(wsc *wsLinkage) {
		wsc.recordChannelAbility = cap
	}
}

//
//
func ReadWait(readWait time.Duration) func(*wsLinkage) {
	return func(wsc *wsLinkage) {
		wsc.readWait = readWait
	}
}

//
//
func PingDuration(pingDuration time.Duration) func(*wsLinkage) {
	return func(wsc *wsLinkage) {
		wsc.pingDuration = pingDuration
	}
}

//
//
func ScanCeiling(readCeiling int64) func(*wsLinkage) {
	return func(wsc *wsLinkage) {
		wsc.readCeiling = readCeiling
	}
}

//
//
func (wsc *wsLinkage) OnBegin() error {
	wsc.recordChannel = make(chan kinds.RPCAnswer, wsc.recordChannelAbility)

	//
	go wsc.readProcedure()
	//
	wsc.recordProcedure()

	return nil
}

//
//
func (wsc *wsLinkage) OnHalt() {
	if wsc.onDetach != nil {
		wsc.onDetach(wsc.distantAddress)
	}

	if wsc.ctx != nil {
		wsc.revoke()
	}
}

//
//
func (wsc *wsLinkage) FetchDistantAddress() string {
	return wsc.distantAddress
}

//
//
//
func (wsc *wsLinkage) RecordRPCReply(ctx context.Context, reply kinds.RPCAnswer) error {
	select {
	case <-wsc.Exit():
		return errors.New("REDACTED")
	case <-ctx.Done():
		return ctx.Err()
	case wsc.recordChannel <- reply:
		return nil
	}
}

//
//
//
func (wsc *wsLinkage) AttemptRecordRPCReply(reply kinds.RPCAnswer) bool {
	select {
	case <-wsc.Exit():
		return false
	case wsc.recordChannel <- reply:
		return true
	default:
		return false
	}
}

//
//
func (wsc *wsLinkage) Context() context.Context {
	if wsc.ctx != nil {
		return wsc.ctx
	}
	wsc.ctx, wsc.revoke = context.WithCancel(context.Background())
	return wsc.ctx
}

//
func (wsc *wsLinkage) readProcedure() {
	//
	recordCtx := context.Background()

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("REDACTED", r)
			}
			wsc.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", string(debug.Stack()))
			if err := wsc.RecordRPCReply(recordCtx, kinds.RPCIntrinsicFault(kinds.JsonrpcIntegerUID(-1), err)); err != nil {
				wsc.Tracer.Fault("REDACTED", "REDACTED", err)
			}
			go wsc.readProcedure()
		}
	}()

	wsc.rootLink.SetPongHandler(func(m string) error {
		return wsc.rootLink.SetReadDeadline(time.Now().Add(wsc.readWait))
	})

	for {
		select {
		case <-wsc.Exit():
			return
		default:
			//
			if err := wsc.rootLink.SetReadDeadline(time.Now().Add(wsc.readWait)); err != nil {
				wsc.Tracer.Fault("REDACTED", "REDACTED", err)
			}

			_, r, err := wsc.rootLink.NextReader()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					wsc.Tracer.Details("REDACTED")
				} else {
					wsc.Tracer.Fault("REDACTED", "REDACTED", err)
				}
				if err := wsc.Halt(); err != nil {
					wsc.Tracer.Fault("REDACTED", "REDACTED", err)
				}
				close(wsc.readProcedureExit)
				return
			}

			dec := json.NewDecoder(r)
			var query kinds.RPCQuery
			err = dec.Decode(&query)
			if err != nil {
				if err := wsc.RecordRPCReply(recordCtx,
					kinds.RPCAnalyzeFault(fmt.Errorf("REDACTED", err))); err != nil {
					wsc.Tracer.Fault("REDACTED", "REDACTED", err)
				}
				continue
			}

			//
			//
			if query.ID == nil {
				wsc.Tracer.Diagnose(
					"REDACTED",
					"REDACTED", query,
				)
				continue
			}

			//
			rpcFunction := wsc.functionIndex[query.Procedure]
			if rpcFunction == nil {
				if err := wsc.RecordRPCReply(recordCtx, kinds.RPCProcedureNegateLocatedFault(query.ID)); err != nil {
					wsc.Tracer.Fault("REDACTED", "REDACTED", err)
				}
				continue
			}

			ctx := &kinds.Context{JSONRequest: &query, WSLink: wsc}
			args := []reflect.Value{reflect.ValueOf(ctx)}
			if len(query.Options) > 0 {
				fnArgs, err := jsonOptionsToArgs(rpcFunction, query.Options)
				if err != nil {
					if err := wsc.RecordRPCReply(recordCtx,
						kinds.RPCIntrinsicFault(query.ID, fmt.Errorf("REDACTED", err)),
					); err != nil {
						wsc.Tracer.Fault("REDACTED", "REDACTED", err)
					}
					continue
				}
				args = append(args, fnArgs...)
			}

			yields := rpcFunction.f.Call(args)

			//
			wsc.Tracer.Details("REDACTED", "REDACTED", query.Procedure)

			outcome, err := unmirrorOutcome(yields)
			if err != nil {
				if err := wsc.RecordRPCReply(recordCtx, kinds.RPCIntrinsicFault(query.ID, err)); err != nil {
					wsc.Tracer.Fault("REDACTED", "REDACTED", err)
				}
				continue
			}

			if err := wsc.RecordRPCReply(recordCtx, kinds.NewRPCSuccessReply(query.ID, outcome)); err != nil {
				wsc.Tracer.Fault("REDACTED", "REDACTED", err)
			}
		}
	}
}

//
func (wsc *wsLinkage) recordProcedure() {
	pingTimer := time.NewTicker(wsc.pingDuration)
	defer pingTimer.Stop()

	//
	pongs := make(chan string, 1)
	wsc.rootLink.SetPingHandler(func(m string) error {
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
		case <-wsc.readProcedureExit: //
			return
		case m := <-pongs:
			err := wsc.recordSignalWithExpiration(websocket.PongMessage, []byte(m))
			if err != nil {
				wsc.Tracer.Details("REDACTED", "REDACTED", err)
			}
		case <-pingTimer.C:
			err := wsc.recordSignalWithExpiration(websocket.PingMessage, []byte{})
			if err != nil {
				wsc.Tracer.Fault("REDACTED", "REDACTED", err)
				return
			}
		case msg := <-wsc.recordChannel:
			//
			//
			//
			jsonOctets, err := json.Marshal(msg)
			if err != nil {
				wsc.Tracer.Fault("REDACTED", "REDACTED", err)
				continue
			}
			if err = wsc.recordSignalWithExpiration(websocket.TextMessage, jsonOctets); err != nil {
				wsc.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", msg)
				return
			}
		}
	}
}

//
//
//
func (wsc *wsLinkage) recordSignalWithExpiration(messageKind int, msg []byte) error {
	if err := wsc.rootLink.SetWriteDeadline(time.Now().Add(wsc.recordWait)); err != nil {
		return err
	}
	return wsc.rootLink.WriteMessage(messageKind, msg)
}
