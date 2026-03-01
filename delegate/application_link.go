package delegate

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

//

//
//

type ApplicationLinkAgreement interface {
	Failure() error
	InitializeSuccession(context.Context, *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error)
	ArrangeNomination(context.Context, *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error)
	HandleNomination(context.Context, *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error)
	BroadenBallot(context.Context, *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error)
	ValidateBallotAddition(context.Context, *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error)
	CulminateLedger(context.Context, *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error)
	Endorse(context.Context) (*kinds.ReplyEndorse, error)
}

type ApplicationLinkTxpool interface {
	AssignReplyClbk(abcicustomer.Clbk)
	Failure() error

	InspectTransfer(context.Context, *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error)
	InspectTransferAsyncronous(context.Context, *kinds.SolicitInspectTransfer) (*abcicustomer.RequestResult, error)
	AppendTransfer(context.Context, *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error)
	HarvestTrans(context.Context, *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error)
	Purge(context.Context) error
}

type PlatformLinkInquire interface {
	Failure() error

	Reverberate(context.Context, string) (*kinds.ReplyReverberate, error)
	Details(context.Context, *kinds.SolicitDetails) (*kinds.ReplyDetails, error)
	Inquire(context.Context, *kinds.SolicitInquire) (*kinds.ReplyInquire, error)
}

type PlatformLinkImage interface {
	Failure() error

	CollectionImages(context.Context, *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error)
	ExtendImage(context.Context, *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error)
	FetchImageSegment(context.Context, *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error)
	ExecuteImageSegment(context.Context, *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error)
}

//
//

type applicationLinkAgreement struct {
	telemetry *Telemetry
	applicationLink abcicustomer.Customer
}

var _ ApplicationLinkAgreement = (*applicationLinkAgreement)(nil)

func FreshApplicationLinkAgreement(applicationLink abcicustomer.Customer, telemetry *Telemetry) ApplicationLinkAgreement {
	return &applicationLinkAgreement{
		telemetry: telemetry,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkAgreement) Failure() error {
	return app.applicationLink.Failure()
}

func (app *applicationLinkAgreement) InitializeSuccession(ctx context.Context, req *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.InitializeSuccession(ctx, req)
}

func (app *applicationLinkAgreement) ArrangeNomination(ctx context.Context,
	req *kinds.SolicitArrangeNomination,
) (*kinds.ReplyArrangeNomination, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ArrangeNomination(ctx, req)
}

func (app *applicationLinkAgreement) HandleNomination(ctx context.Context, req *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.HandleNomination(ctx, req)
}

func (app *applicationLinkAgreement) BroadenBallot(ctx context.Context, req *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.BroadenBallot(ctx, req)
}

func (app *applicationLinkAgreement) ValidateBallotAddition(ctx context.Context, req *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ValidateBallotAddition(ctx, req)
}

func (app *applicationLinkAgreement) CulminateLedger(ctx context.Context, req *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.CulminateLedger(ctx, req)
}

func (app *applicationLinkAgreement) Endorse(ctx context.Context) (*kinds.ReplyEndorse, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Endorse(ctx, &kinds.SolicitEndorse{})
}

//
//

type applicationLinkTxpool struct {
	telemetry *Telemetry
	applicationLink abcicustomer.Customer
}

func FreshApplicationLinkTxpool(applicationLink abcicustomer.Customer, telemetry *Telemetry) ApplicationLinkTxpool {
	return &applicationLinkTxpool{
		telemetry: telemetry,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkTxpool) AssignReplyClbk(cb abcicustomer.Clbk) {
	app.applicationLink.AssignReplyClbk(cb)
}

func (app *applicationLinkTxpool) Failure() error {
	return app.applicationLink.Failure()
}

func (app *applicationLinkTxpool) Purge(ctx context.Context) error {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Purge(ctx)
}

func (app *applicationLinkTxpool) InspectTransfer(ctx context.Context, req *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.InspectTransfer(ctx, req)
}

func (app *applicationLinkTxpool) InspectTransferAsyncronous(ctx context.Context, req *kinds.SolicitInspectTransfer) (*abcicustomer.RequestResult, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.InspectTransferAsyncronous(ctx, req)
}

func (app *applicationLinkTxpool) AppendTransfer(ctx context.Context, req *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.AppendTransfer(ctx, req)
}

func (app *applicationLinkTxpool) HarvestTrans(ctx context.Context, req *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.HarvestTrans(ctx, req)
}

//
//

type applicationLinkInquire struct {
	telemetry *Telemetry
	applicationLink abcicustomer.Customer
}

func FreshApplicationLinkInquire(applicationLink abcicustomer.Customer, telemetry *Telemetry) PlatformLinkInquire {
	return &applicationLinkInquire{
		telemetry: telemetry,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkInquire) Failure() error {
	return app.applicationLink.Failure()
}

func (app *applicationLinkInquire) Reverberate(ctx context.Context, msg string) (*kinds.ReplyReverberate, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Reverberate(ctx, msg)
}

func (app *applicationLinkInquire) Details(ctx context.Context, req *kinds.SolicitDetails) (*kinds.ReplyDetails, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Details(ctx, req)
}

func (app *applicationLinkInquire) Inquire(ctx context.Context, req *kinds.SolicitInquire) (*kinds.ReplyInquire, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Inquire(ctx, req)
}

//
//

type applicationLinkImage struct {
	telemetry *Telemetry
	applicationLink abcicustomer.Customer
}

func FreshApplicationLinkImage(applicationLink abcicustomer.Customer, telemetry *Telemetry) PlatformLinkImage {
	return &applicationLinkImage{
		telemetry: telemetry,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkImage) Failure() error {
	return app.applicationLink.Failure()
}

func (app *applicationLinkImage) CollectionImages(ctx context.Context, req *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.CollectionImages(ctx, req)
}

func (app *applicationLinkImage) ExtendImage(ctx context.Context, req *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ExtendImage(ctx, req)
}

func (app *applicationLinkImage) FetchImageSegment(ctx context.Context, req *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.FetchImageSegment(ctx, req)
}

func (app *applicationLinkImage) ExecuteImageSegment(ctx context.Context, req *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error) {
	defer appendMomentSpecimen(app.telemetry.ProcedureScheduleMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ExecuteImageSegment(ctx, req)
}

//
//
//
//
func appendMomentSpecimen(m metrics.Histogram) func() {
	initiate := time.Now()
	return func() { m.Observe(time.Since(initiate).Seconds()) }
}
