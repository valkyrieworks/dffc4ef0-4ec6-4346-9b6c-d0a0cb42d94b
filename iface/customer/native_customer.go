package abciend

import (
	"context"

	kinds "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
)

//
//
//
//
type nativeCustomer struct {
	daemon.RootDaemon

	mtx *engineconnect.Lock
	kinds.Software
	Callback
}

var _ Customer = (*nativeCustomer)(nil)

//
//
//
//
func NewNativeCustomer(mtx *engineconnect.Lock, app kinds.Software) Customer {
	if mtx == nil {
		mtx = new(engineconnect.Lock)
	}
	cli := &nativeCustomer{
		mtx:         mtx,
		Software: app,
	}
	cli.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", cli)
	return cli
}

func (app *nativeCustomer) CollectionReplyCallback(cb Callback) {
	app.mtx.Lock()
	app.Callback = cb
	app.mtx.Unlock()
}

func (app *nativeCustomer) InspectTransferAsync(ctx context.Context, req *kinds.QueryInspectTransfer) (*RequestOutput, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	res, err := app.Software.InspectTransfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return app.callback(
		kinds.ToQueryInspectTransfer(req),
		kinds.ToReplyInspectTransfer(res),
	), nil
}

func (app *nativeCustomer) callback(req *kinds.Query, res *kinds.Reply) *RequestOutput {
	app.Callback(req, res)
	rr := newNativeRequestOutput(req, res)
	rr.callbackExecuted = true
	return rr
}

func newNativeRequestOutput(req *kinds.Query, res *kinds.Reply) *RequestOutput {
	requestOutput := NewRequestOutput(req)
	requestOutput.Reply = res
	return requestOutput
}

//

func (app *nativeCustomer) Fault() error {
	return nil
}

func (app *nativeCustomer) Purge(context.Context) error {
	return nil
}

func (app *nativeCustomer) Replicate(_ context.Context, msg string) (*kinds.ReplyReverberate, error) {
	return &kinds.ReplyReverberate{Signal: msg}, nil
}

func (app *nativeCustomer) Details(ctx context.Context, req *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.Details(ctx, req)
}

func (app *nativeCustomer) InspectTransfer(ctx context.Context, req *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.InspectTransfer(ctx, req)
}

func (app *nativeCustomer) EmbedTransfer(ctx context.Context, req *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error) {
	//
	return app.Software.EmbedTransfer(ctx, req)
}

func (app *nativeCustomer) HarvestTrans(ctx context.Context, req *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	//
	return app.Software.HarvestTrans(ctx, req)
}

func (app *nativeCustomer) Inquire(ctx context.Context, req *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.Inquire(ctx, req)
}

func (app *nativeCustomer) Endorse(ctx context.Context, req *kinds.QueryEndorse) (*kinds.ReplyEndorse, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.Endorse(ctx, req)
}

func (app *nativeCustomer) InitSeries(ctx context.Context, req *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.InitSeries(ctx, req)
}

func (app *nativeCustomer) CatalogMirrors(ctx context.Context, req *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.CatalogMirrors(ctx, req)
}

func (app *nativeCustomer) ProposalMirror(ctx context.Context, req *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.ProposalMirror(ctx, req)
}

func (app *nativeCustomer) ImportMirrorSegment(ctx context.Context,
	req *kinds.QueryImportMirrorSegment,
) (*kinds.ReplyImportMirrorSegment, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.ImportMirrorSegment(ctx, req)
}

func (app *nativeCustomer) ExecuteMirrorSegment(ctx context.Context,
	req *kinds.QueryExecuteMirrorSegment,
) (*kinds.ReplyExecuteMirrorSegment, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.ExecuteMirrorSegment(ctx, req)
}

func (app *nativeCustomer) ArrangeNomination(ctx context.Context, req *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.ArrangeNomination(ctx, req)
}

func (app *nativeCustomer) HandleNomination(ctx context.Context, req *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.HandleNomination(ctx, req)
}

func (app *nativeCustomer) ExpandBallot(ctx context.Context, req *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.ExpandBallot(ctx, req)
}

func (app *nativeCustomer) ValidateBallotAddition(ctx context.Context, req *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.ValidateBallotAddition(ctx, req)
}

func (app *nativeCustomer) CompleteLedger(ctx context.Context, req *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Software.CompleteLedger(ctx, req)
}
