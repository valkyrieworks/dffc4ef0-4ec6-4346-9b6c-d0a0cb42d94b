package gateway

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"

	abciend "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/kinds"
)

//

//
//

type ApplicationLinkAgreement interface {
	Fault() error
	InitSeries(context.Context, *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error)
	ArrangeNomination(context.Context, *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error)
	HandleNomination(context.Context, *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error)
	ExpandBallot(context.Context, *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error)
	ValidateBallotAddition(context.Context, *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error)
	CompleteLedger(context.Context, *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error)
	Endorse(context.Context) (*kinds.ReplyEndorse, error)
}

type ApplicationLinkTxpool interface {
	CollectionReplyCallback(abciend.Callback)
	Fault() error

	InspectTransfer(context.Context, *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error)
	InspectTransferAsync(context.Context, *kinds.QueryInspectTransfer) (*abciend.RequestOutput, error)
	EmbedTransfer(context.Context, *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error)
	HarvestTrans(context.Context, *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error)
	Purge(context.Context) error
}

type ApplicationLinkInquire interface {
	Fault() error

	Replicate(context.Context, string) (*kinds.ReplyReverberate, error)
	Details(context.Context, *kinds.QueryDetails) (*kinds.ReplyDetails, error)
	Inquire(context.Context, *kinds.QueryInquire) (*kinds.ReplyInquire, error)
}

type ApplicationLinkMirror interface {
	Fault() error

	CatalogMirrors(context.Context, *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error)
	ProposalMirror(context.Context, *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error)
	ImportMirrorSegment(context.Context, *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error)
	ExecuteMirrorSegment(context.Context, *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error)
}

//
//

type applicationLinkAgreement struct {
	stats *Stats
	applicationLink abciend.Customer
}

var _ ApplicationLinkAgreement = (*applicationLinkAgreement)(nil)

func NewApplicationLinkAgreement(applicationLink abciend.Customer, stats *Stats) ApplicationLinkAgreement {
	return &applicationLinkAgreement{
		stats: stats,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkAgreement) Fault() error {
	return app.applicationLink.Fault()
}

func (app *applicationLinkAgreement) InitSeries(ctx context.Context, req *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.InitSeries(ctx, req)
}

func (app *applicationLinkAgreement) ArrangeNomination(ctx context.Context,
	req *kinds.QueryArrangeNomination,
) (*kinds.ReplyArrangeNomination, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ArrangeNomination(ctx, req)
}

func (app *applicationLinkAgreement) HandleNomination(ctx context.Context, req *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.HandleNomination(ctx, req)
}

func (app *applicationLinkAgreement) ExpandBallot(ctx context.Context, req *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ExpandBallot(ctx, req)
}

func (app *applicationLinkAgreement) ValidateBallotAddition(ctx context.Context, req *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ValidateBallotAddition(ctx, req)
}

func (app *applicationLinkAgreement) CompleteLedger(ctx context.Context, req *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.CompleteLedger(ctx, req)
}

func (app *applicationLinkAgreement) Endorse(ctx context.Context) (*kinds.ReplyEndorse, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Endorse(ctx, &kinds.QueryEndorse{})
}

//
//

type applicationLinkTxpool struct {
	stats *Stats
	applicationLink abciend.Customer
}

func NewApplicationLinkTxpool(applicationLink abciend.Customer, stats *Stats) ApplicationLinkTxpool {
	return &applicationLinkTxpool{
		stats: stats,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkTxpool) CollectionReplyCallback(cb abciend.Callback) {
	app.applicationLink.CollectionReplyCallback(cb)
}

func (app *applicationLinkTxpool) Fault() error {
	return app.applicationLink.Fault()
}

func (app *applicationLinkTxpool) Purge(ctx context.Context) error {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Purge(ctx)
}

func (app *applicationLinkTxpool) InspectTransfer(ctx context.Context, req *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.InspectTransfer(ctx, req)
}

func (app *applicationLinkTxpool) InspectTransferAsync(ctx context.Context, req *kinds.QueryInspectTransfer) (*abciend.RequestOutput, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.InspectTransferAsync(ctx, req)
}

func (app *applicationLinkTxpool) EmbedTransfer(ctx context.Context, req *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.EmbedTransfer(ctx, req)
}

func (app *applicationLinkTxpool) HarvestTrans(ctx context.Context, req *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.HarvestTrans(ctx, req)
}

//
//

type applicationLinkInquire struct {
	stats *Stats
	applicationLink abciend.Customer
}

func NewApplicationLinkInquire(applicationLink abciend.Customer, stats *Stats) ApplicationLinkInquire {
	return &applicationLinkInquire{
		stats: stats,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkInquire) Fault() error {
	return app.applicationLink.Fault()
}

func (app *applicationLinkInquire) Replicate(ctx context.Context, msg string) (*kinds.ReplyReverberate, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Replicate(ctx, msg)
}

func (app *applicationLinkInquire) Details(ctx context.Context, req *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Details(ctx, req)
}

func (app *applicationLinkInquire) Inquire(ctx context.Context, req *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.Inquire(ctx, req)
}

//
//

type applicationLinkMirror struct {
	stats *Stats
	applicationLink abciend.Customer
}

func NewApplicationLinkMirror(applicationLink abciend.Customer, stats *Stats) ApplicationLinkMirror {
	return &applicationLinkMirror{
		stats: stats,
		applicationLink: applicationLink,
	}
}

func (app *applicationLinkMirror) Fault() error {
	return app.applicationLink.Fault()
}

func (app *applicationLinkMirror) CatalogMirrors(ctx context.Context, req *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.CatalogMirrors(ctx, req)
}

func (app *applicationLinkMirror) ProposalMirror(ctx context.Context, req *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ProposalMirror(ctx, req)
}

func (app *applicationLinkMirror) ImportMirrorSegment(ctx context.Context, req *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ImportMirrorSegment(ctx, req)
}

func (app *applicationLinkMirror) ExecuteMirrorSegment(ctx context.Context, req *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error) {
	defer appendTimeSpecimen(app.stats.ProcedureCadenceMoments.With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))()
	return app.applicationLink.ExecuteMirrorSegment(ctx, req)
}

//
//
//
//
func appendTimeSpecimen(m metrics.Histogram) func() {
	begin := time.Now()
	return func() { m.Observe(time.Since(begin).Seconds()) }
}
