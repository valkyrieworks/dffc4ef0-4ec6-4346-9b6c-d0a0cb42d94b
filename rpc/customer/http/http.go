package http

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/valkyrieworks/utils/octets"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	rpccustomer "github.com/valkyrieworks/rpc/customer"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	jsonrpccustomer "github.com/valkyrieworks/rpc/jsonrpc/customer"
	"github.com/valkyrieworks/kinds"
)

/**
r
.

.
s
.

.
s
o
d
.

o
e
.

:

)
{
r
}

s
)
{
r
}
)

)
{
r
}

t
*/
type HTTP struct {
	external string
	rpc    *jsonrpccustomer.Customer

	*rootRPCCustomer
	*WSEvents
}

//
//
//
//
//
//
//
//
//
//
type GroupHTTP struct {
	rpcGroup *jsonrpccustomer.QueryGroup
	*rootRPCCustomer
}

//
//
//
type rpcCustomer interface {
	rpccustomer.IfaceCustomer
	rpccustomer.LogbookCustomer
	rpccustomer.FabricCustomer
	rpccustomer.AttestCustomer
	rpccustomer.StateCustomer
}

//
//
type rootRPCCustomer struct {
	invoker jsonrpccustomer.Invoker
}

var (
	_ rpcCustomer = (*HTTP)(nil)
	_ rpcCustomer = (*GroupHTTP)(nil)
	_ rpcCustomer = (*rootRPCCustomer)(nil)
)

//
//

//
//
//
func New(external, wsTerminus string) (*HTTP, error) {
	httpCustomer, err := jsonrpccustomer.StandardHTTPCustomer(external)
	if err != nil {
		return nil, err
	}
	return NewWithCustomer(external, wsTerminus, httpCustomer)
}

//
func NewWithDeadline(external, wsTerminus string, deadline uint) (*HTTP, error) {
	httpCustomer, err := jsonrpccustomer.StandardHTTPCustomer(external)
	if err != nil {
		return nil, err
	}
	httpCustomer.Timeout = time.Duration(deadline) * time.Second
	return NewWithCustomer(external, wsTerminus, httpCustomer)
}

//
//
func NewWithCustomer(external, wsTerminus string, customer *http.Client) (*HTTP, error) {
	if customer == nil {
		panic("REDACTED")
	}

	rc, err := jsonrpccustomer.NewWithHTTPCustomer(external, customer)
	if err != nil {
		return nil, err
	}

	wsEvents, err := newWSEvents(external, wsTerminus)
	if err != nil {
		return nil, err
	}

	httpCustomer := &HTTP{
		rpc:           rc,
		external:        external,
		rootRPCCustomer: &rootRPCCustomer{invoker: rc},
		WSEvents:      wsEvents,
	}

	return httpCustomer, nil
}

var _ rpccustomer.Customer = (*HTTP)(nil)

//
func (c *HTTP) AssignTracer(l log.Tracer) {
	c.WSEvents.AssignTracer(l)
}

//
func (c *HTTP) External() string {
	return c.external
}

//
func (c *HTTP) NewGroup() *GroupHTTP {
	rpcGroup := c.rpc.NewQueryGroup()
	return &GroupHTTP{
		rpcGroup: rpcGroup,
		rootRPCCustomer: &rootRPCCustomer{
			invoker: rpcGroup,
		},
	}
}

//
//

//
//
//
//
func (b *GroupHTTP) Transmit(ctx context.Context) ([]any, error) {
	return b.rpcGroup.Transmit(ctx)
}

//
//
func (b *GroupHTTP) Flush() int {
	return b.rpcGroup.Flush()
}

//
func (b *GroupHTTP) Number() int {
	return b.rpcGroup.Number()
}

//
//

func (c *rootRPCCustomer) Status(ctx context.Context) (*ctypes.OutcomeState, error) {
	outcome := new(ctypes.OutcomeState)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *rootRPCCustomer) IfaceDetails(ctx context.Context) (*ctypes.OutcomeIfaceDetails, error) {
	outcome := new(ctypes.OutcomeIfaceDetails)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *rootRPCCustomer) IfaceInquire(
	ctx context.Context,
	route string,
	data octets.HexOctets,
) (*ctypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireWithSettings(ctx, route, data, rpccustomer.StandardIfaceInquireSettings)
}

func (c *rootRPCCustomer) IfaceInquireWithSettings(
	ctx context.Context,
	route string,
	data octets.HexOctets,
	opts rpccustomer.IfaceInquireSettings,
) (*ctypes.OutcomeIfaceInquire, error) {
	outcome := new(ctypes.OutcomeIfaceInquire)
	_, err := c.invoker.Invoke(ctx, "REDACTED",
		map[string]any{"REDACTED": route, "REDACTED": data, "REDACTED": opts.Level, "REDACTED": opts.Demonstrate},
		outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *rootRPCCustomer) MulticastTransferEndorse(
	ctx context.Context,
	tx kinds.Tx,
) (*ctypes.OutcomeMulticastTransferEndorse, error) {
	outcome := new(ctypes.OutcomeMulticastTransferEndorse)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{"REDACTED": tx}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) MulticastTransferAsync(
	ctx context.Context,
	tx kinds.Tx,
) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.multicastTransfer(ctx, "REDACTED", tx)
}

func (c *rootRPCCustomer) MulticastTransferAlign(
	ctx context.Context,
	tx kinds.Tx,
) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.multicastTransfer(ctx, "REDACTED", tx)
}

func (c *rootRPCCustomer) multicastTransfer(
	ctx context.Context,
	path string,
	tx kinds.Tx,
) (*ctypes.OutcomeMulticastTransfer, error) {
	outcome := new(ctypes.OutcomeMulticastTransfer)
	_, err := c.invoker.Invoke(ctx, path, map[string]any{"REDACTED": tx}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) UnattestedTrans(
	ctx context.Context,
	ceiling *int,
) (*ctypes.OutcomeUnattestedTrans, error) {
	outcome := new(ctypes.OutcomeUnattestedTrans)
	options := make(map[string]any)
	if ceiling != nil {
		options["REDACTED"] = ceiling
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) CountUnattestedTrans(ctx context.Context) (*ctypes.OutcomeUnattestedTrans, error) {
	outcome := new(ctypes.OutcomeUnattestedTrans)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) InspectTransfer(ctx context.Context, tx kinds.Tx) (*ctypes.OutcomeInspectTransfer, error) {
	outcome := new(ctypes.OutcomeInspectTransfer)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{"REDACTED": tx}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) NetDetails(ctx context.Context) (*ctypes.OutcomeNetDetails, error) {
	outcome := new(ctypes.OutcomeNetDetails)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) ExportAgreementStatus(ctx context.Context) (*ctypes.OutcomeExportAgreementStatus, error) {
	outcome := new(ctypes.OutcomeExportAgreementStatus)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) AgreementStatus(ctx context.Context) (*ctypes.OutcomeAgreementStatus, error) {
	outcome := new(ctypes.OutcomeAgreementStatus)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) AgreementOptions(
	ctx context.Context,
	level *int64,
) (*ctypes.OutcomeAgreementOptions, error) {
	outcome := new(ctypes.OutcomeAgreementOptions)
	options := make(map[string]any)
	if level != nil {
		options["REDACTED"] = level
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) Vitality(ctx context.Context) (*ctypes.OutcomeVitality, error) {
	outcome := new(ctypes.OutcomeVitality)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) LedgerchainDetails(
	ctx context.Context,
	minimumLevel,
	maximumLevel int64,
) (*ctypes.OutcomeLedgerchainDetails, error) {
	outcome := new(ctypes.OutcomeLedgerchainDetails)
	_, err := c.invoker.Invoke(ctx, "REDACTED",
		map[string]any{"REDACTED": minimumLevel, "REDACTED": maximumLevel},
		outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) Origin(ctx context.Context) (*ctypes.OutcomeOrigin, error) {
	outcome := new(ctypes.OutcomeOrigin)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) OriginSegmented(ctx context.Context, id uint) (*ctypes.OutcomeOriginSegment, error) {
	outcome := new(ctypes.OutcomeOriginSegment)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{"REDACTED": id}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) Ledger(ctx context.Context, level *int64) (*ctypes.OutcomeLedger, error) {
	outcome := new(ctypes.OutcomeLedger)
	options := make(map[string]any)
	if level != nil {
		options["REDACTED"] = level
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) LedgerByDigest(ctx context.Context, digest []byte) (*ctypes.OutcomeLedger, error) {
	outcome := new(ctypes.OutcomeLedger)
	options := map[string]any{
		"REDACTED": digest,
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) LedgerOutcomes(
	ctx context.Context,
	level *int64,
) (*ctypes.OutcomeLedgerOutcomes, error) {
	outcome := new(ctypes.OutcomeLedgerOutcomes)
	options := make(map[string]any)
	if level != nil {
		options["REDACTED"] = level
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) Heading(ctx context.Context, level *int64) (*ctypes.OutcomeHeading, error) {
	outcome := new(ctypes.OutcomeHeading)
	options := make(map[string]any)
	if level != nil {
		options["REDACTED"] = level
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) HeadingByDigest(ctx context.Context, digest octets.HexOctets) (*ctypes.OutcomeHeading, error) {
	outcome := new(ctypes.OutcomeHeading)
	options := map[string]any{
		"REDACTED": digest,
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) Endorse(ctx context.Context, level *int64) (*ctypes.OutcomeEndorse, error) {
	outcome := new(ctypes.OutcomeEndorse)
	options := make(map[string]any)
	if level != nil {
		options["REDACTED"] = level
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) Tx(ctx context.Context, digest []byte, demonstrate bool) (*ctypes.OutcomeTransfer, error) {
	outcome := new(ctypes.OutcomeTransfer)
	options := map[string]any{
		"REDACTED":  digest,
		"REDACTED": demonstrate,
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) TransferScan(
	ctx context.Context,
	inquire string,
	demonstrate bool,
	screen,
	eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeTransferScan, error) {
	outcome := new(ctypes.OutcomeTransferScan)
	options := map[string]any{
		"REDACTED":    inquire,
		"REDACTED":    demonstrate,
		"REDACTED": arrangeBy,
	}

	if screen != nil {
		options["REDACTED"] = screen
	}
	if eachScreen != nil {
		options["REDACTED"] = eachScreen
	}

	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *rootRPCCustomer) LedgerScan(
	ctx context.Context,
	inquire string,
	screen, eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeLedgerScan, error) {
	outcome := new(ctypes.OutcomeLedgerScan)
	options := map[string]any{
		"REDACTED":    inquire,
		"REDACTED": arrangeBy,
	}

	if screen != nil {
		options["REDACTED"] = screen
	}
	if eachScreen != nil {
		options["REDACTED"] = eachScreen
	}

	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *rootRPCCustomer) Ratifiers(
	ctx context.Context,
	level *int64,
	screen,
	eachScreen *int,
) (*ctypes.OutcomeRatifiers, error) {
	outcome := new(ctypes.OutcomeRatifiers)
	options := make(map[string]any)
	if screen != nil {
		options["REDACTED"] = screen
	}
	if eachScreen != nil {
		options["REDACTED"] = eachScreen
	}
	if level != nil {
		options["REDACTED"] = level
	}
	_, err := c.invoker.Invoke(ctx, "REDACTED", options, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *rootRPCCustomer) MulticastProof(
	ctx context.Context,
	ev kinds.Proof,
) (*ctypes.OutcomeMulticastProof, error) {
	outcome := new(ctypes.OutcomeMulticastProof)
	_, err := c.invoker.Invoke(ctx, "REDACTED", map[string]any{"REDACTED": ev}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

//
//

var errNegateActive = errors.New("REDACTED")

//
type WSEvents struct {
	daemon.RootDaemon
	external   string
	gateway string
	ws       *jsonrpccustomer.WSCustomer

	mtx           engineconnect.ReadwriteLock
	registrations map[string]chan ctypes.OutcomeEvent //
}

func newWSEvents(external, gateway string) (*WSEvents, error) {
	w := &WSEvents{
		gateway:      gateway,
		external:        external,
		registrations: make(map[string]chan ctypes.OutcomeEvent),
	}
	w.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", w)

	var err error
	w.ws, err = jsonrpccustomer.NewWS(w.external, w.gateway, jsonrpccustomer.OnReestablish(func() {
		//
		w.reworkRegistrationsAfter(0 * time.Second)
	}))
	if err != nil {
		return nil, err
	}
	w.ws.AssignTracer(w.Tracer)

	return w, nil
}

//
func (w *WSEvents) OnBegin() error {
	if err := w.ws.Begin(); err != nil {
		return err
	}

	go w.eventObserver()

	return nil
}

//
func (w *WSEvents) OnHalt() {
	if err := w.ws.Halt(); err != nil {
		w.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//
//
//
//
//
//
//
func (w *WSEvents) Enrol(ctx context.Context, _, inquire string,
	outVolume ...int,
) (out <-chan ctypes.OutcomeEvent, err error) {
	if !w.IsActive() {
		return nil, errNegateActive
	}

	if err := w.ws.Enrol(ctx, inquire); err != nil {
		return nil, err
	}

	outCeiling := 1
	if len(outVolume) > 0 {
		outCeiling = outVolume[0]
	}

	outchan := make(chan ctypes.OutcomeEvent, outCeiling)
	w.mtx.Lock()
	//
	//
	w.registrations[inquire] = outchan
	w.mtx.Unlock()

	return outchan, nil
}

//
//
//
//
func (w *WSEvents) Deenroll(ctx context.Context, _, inquire string) error {
	if !w.IsActive() {
		return errNegateActive
	}

	if err := w.ws.Deenroll(ctx, inquire); err != nil {
		return err
	}

	w.mtx.Lock()
	_, ok := w.registrations[inquire]
	if ok {
		delete(w.registrations, inquire)
	}
	w.mtx.Unlock()

	return nil
}

//
//
//
//
func (w *WSEvents) DeenrollAll(ctx context.Context, _ string) error {
	if !w.IsActive() {
		return errNegateActive
	}

	if err := w.ws.DeenrollAll(ctx); err != nil {
		return err
	}

	w.mtx.Lock()
	w.registrations = make(map[string]chan ctypes.OutcomeEvent)
	w.mtx.Unlock()

	return nil
}

//
//
func (w *WSEvents) reworkRegistrationsAfter(d time.Duration) {
	time.Sleep(d)

	w.mtx.RLock()
	defer w.mtx.RUnlock()
	for q := range w.registrations {
		err := w.ws.Enrol(context.Background(), q)
		if err != nil {
			w.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
}

func isErrYetRegistered(err error) bool {
	return strings.Contains(err.Error(), cometbroadcast.ErrYetActivated.Error())
}

func (w *WSEvents) eventObserver() {
	for {
		select {
		case reply, ok := <-w.ws.RepliesChan:
			if !ok {
				return
			}

			if reply.Fault != nil {
				w.Tracer.Fault("REDACTED", "REDACTED", reply.Fault.Fault())
				//
				//
				//
				//
				if !isErrYetRegistered(reply.Fault) {
					//
					//
					w.reworkRegistrationsAfter(1 * time.Second)
				}
				continue
			}

			outcome := new(ctypes.OutcomeEvent)
			err := cometjson.Unserialize(reply.Outcome, outcome)
			if err != nil {
				w.Tracer.Fault("REDACTED", "REDACTED", err)
				continue
			}

			w.mtx.RLock()
			if out, ok := w.registrations[outcome.Inquire]; ok {
				if cap(out) == 0 {
					out <- *outcome
				} else {
					select {
					case out <- *outcome:
					default:
						w.Tracer.Fault("REDACTED", "REDACTED", outcome, "REDACTED", outcome.Inquire)
					}
				}
			}
			w.mtx.RUnlock()
		case <-w.Exit():
			return
		}
	}
}
