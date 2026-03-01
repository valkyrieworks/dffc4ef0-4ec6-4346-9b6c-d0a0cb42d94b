package abcicustomer

import (
	"context"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
//
//
//
type regionalCustomer struct {
	facility.FoundationFacility

	mtx *commitchronize.Exclusion
	kinds.Platform
	Clbk
}

var _ Customer = (*regionalCustomer)(nil)

//
//
//
//
func FreshRegionalCustomer(mtx *commitchronize.Exclusion, app kinds.Platform) Customer {
	if mtx == nil {
		mtx = new(commitchronize.Exclusion)
	}
	cli := &regionalCustomer{
		mtx:         mtx,
		Platform: app,
	}
	cli.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", cli)
	return cli
}

func (app *regionalCustomer) AssignReplyClbk(cb Clbk) {
	app.mtx.Lock()
	app.Clbk = cb
	app.mtx.Unlock()
}

func (app *regionalCustomer) InspectTransferAsyncronous(ctx context.Context, req *kinds.SolicitInspectTransfer) (*RequestResult, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	res, err := app.Platform.InspectTransfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return app.clbk(
		kinds.TowardSolicitInspectTransfer(req),
		kinds.TowardReplyInspectTransfer(res),
	), nil
}

func (app *regionalCustomer) clbk(req *kinds.Solicit, res *kinds.Reply) *RequestResult {
	app.Clbk(req, res)
	rr := freshRegionalRequestResult(req, res)
	rr.clbkExecuted = true
	return rr
}

func freshRegionalRequestResult(req *kinds.Solicit, res *kinds.Reply) *RequestResult {
	requestResult := FreshRequestResult(req)
	requestResult.Reply = res
	return requestResult
}

//

func (app *regionalCustomer) Failure() error {
	return nil
}

func (app *regionalCustomer) Purge(context.Context) error {
	return nil
}

func (app *regionalCustomer) Reverberate(_ context.Context, msg string) (*kinds.ReplyReverberate, error) {
	return &kinds.ReplyReverberate{Signal: msg}, nil
}

func (app *regionalCustomer) Details(ctx context.Context, req *kinds.SolicitDetails) (*kinds.ReplyDetails, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.Details(ctx, req)
}

func (app *regionalCustomer) InspectTransfer(ctx context.Context, req *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.InspectTransfer(ctx, req)
}

func (app *regionalCustomer) AppendTransfer(ctx context.Context, req *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error) {
	//
	return app.Platform.AppendTransfer(ctx, req)
}

func (app *regionalCustomer) HarvestTrans(ctx context.Context, req *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	//
	return app.Platform.HarvestTrans(ctx, req)
}

func (app *regionalCustomer) Inquire(ctx context.Context, req *kinds.SolicitInquire) (*kinds.ReplyInquire, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.Inquire(ctx, req)
}

func (app *regionalCustomer) Endorse(ctx context.Context, req *kinds.SolicitEndorse) (*kinds.ReplyEndorse, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.Endorse(ctx, req)
}

func (app *regionalCustomer) InitializeSuccession(ctx context.Context, req *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.InitializeSuccession(ctx, req)
}

func (app *regionalCustomer) CollectionImages(ctx context.Context, req *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.CollectionImages(ctx, req)
}

func (app *regionalCustomer) ExtendImage(ctx context.Context, req *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.ExtendImage(ctx, req)
}

func (app *regionalCustomer) FetchImageSegment(ctx context.Context,
	req *kinds.SolicitFetchImageSegment,
) (*kinds.ReplyFetchImageSegment, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.FetchImageSegment(ctx, req)
}

func (app *regionalCustomer) ExecuteImageSegment(ctx context.Context,
	req *kinds.SolicitExecuteImageSegment,
) (*kinds.ReplyExecuteImageSegment, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.ExecuteImageSegment(ctx, req)
}

func (app *regionalCustomer) ArrangeNomination(ctx context.Context, req *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.ArrangeNomination(ctx, req)
}

func (app *regionalCustomer) HandleNomination(ctx context.Context, req *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.HandleNomination(ctx, req)
}

func (app *regionalCustomer) BroadenBallot(ctx context.Context, req *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.BroadenBallot(ctx, req)
}

func (app *regionalCustomer) ValidateBallotAddition(ctx context.Context, req *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.ValidateBallotAddition(ctx, req)
}

func (app *regionalCustomer) CulminateLedger(ctx context.Context, req *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
	app.mtx.Lock()
	defer app.mtx.Unlock()

	return app.Platform.CulminateLedger(ctx, req)
}
