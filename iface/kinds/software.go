package kinds

import "context"

//

//
//
type Software interface {
	//
	Details(context.Context, *QueryDetails) (*ReplyDetails, error)    //
	Inquire(context.Context, *QueryInquire) (*ReplyInquire, error) //

	//
	InspectTransfer(context.Context, *QueryInspectTransfer) (*ReplyInspectTransfer, error)    //
	EmbedTransfer(context.Context, *QueryEmbedTransfer) (*ReplyEmbedTransfer, error) //
	HarvestTrans(context.Context, *QueryHarvestTrans) (*ReplyHarvestTrans, error)    //

	//
	InitSeries(context.Context, *QueryInitSeries) (*ReplyInitSeries, error) //
	ArrangeNomination(context.Context, *QueryArrangeNomination) (*ReplyArrangeNomination, error)
	HandleNomination(context.Context, *QueryHandleNomination) (*ReplyHandleNomination, error)
	//
	CompleteLedger(context.Context, *QueryCompleteLedger) (*ReplyCompleteLedger, error)
	//
	ExpandBallot(context.Context, *QueryExpandBallot) (*ReplyExpandBallot, error)
	//
	ValidateBallotAddition(context.Context, *QueryValidateBallotAddition) (*ReplyValidateBallotAddition, error)
	//
	Endorse(context.Context, *QueryEndorse) (*ReplyEndorse, error)

	//
	CatalogMirrors(context.Context, *QueryCatalogMirrors) (*ReplyCatalogMirrors, error)                //
	ProposalMirror(context.Context, *QueryProposalMirror) (*ReplyProposalMirror, error)                //
	ImportMirrorSegment(context.Context, *QueryImportMirrorSegment) (*ReplyImportMirrorSegment, error)    //
	ExecuteMirrorSegment(context.Context, *QueryExecuteMirrorSegment) (*ReplyExecuteMirrorSegment, error) //
}

//
//

var _ Software = (*RootSoftware)(nil)

type RootSoftware struct{}

func NewRootSoftware() *RootSoftware {
	return &RootSoftware{}
}

func (RootSoftware) Details(context.Context, *QueryDetails) (*ReplyDetails, error) {
	return &ReplyDetails{}, nil
}

func (RootSoftware) InspectTransfer(context.Context, *QueryInspectTransfer) (*ReplyInspectTransfer, error) {
	return &ReplyInspectTransfer{Code: CodeKindSuccess}, nil
}

func (RootSoftware) EmbedTransfer(context.Context, *QueryEmbedTransfer) (*ReplyEmbedTransfer, error) {
	return &ReplyEmbedTransfer{Code: CodeKindSuccess}, nil
}

func (RootSoftware) HarvestTrans(context.Context, *QueryHarvestTrans) (*ReplyHarvestTrans, error) {
	return &ReplyHarvestTrans{}, nil
}

func (RootSoftware) Endorse(context.Context, *QueryEndorse) (*ReplyEndorse, error) {
	return &ReplyEndorse{}, nil
}

func (RootSoftware) Inquire(context.Context, *QueryInquire) (*ReplyInquire, error) {
	return &ReplyInquire{Code: CodeKindSuccess}, nil
}

func (RootSoftware) InitSeries(context.Context, *QueryInitSeries) (*ReplyInitSeries, error) {
	return &ReplyInitSeries{}, nil
}

func (RootSoftware) CatalogMirrors(context.Context, *QueryCatalogMirrors) (*ReplyCatalogMirrors, error) {
	return &ReplyCatalogMirrors{}, nil
}

func (RootSoftware) ProposalMirror(context.Context, *QueryProposalMirror) (*ReplyProposalMirror, error) {
	return &ReplyProposalMirror{}, nil
}

func (RootSoftware) ImportMirrorSegment(context.Context, *QueryImportMirrorSegment) (*ReplyImportMirrorSegment, error) {
	return &ReplyImportMirrorSegment{}, nil
}

func (RootSoftware) ExecuteMirrorSegment(context.Context, *QueryExecuteMirrorSegment) (*ReplyExecuteMirrorSegment, error) {
	return &ReplyExecuteMirrorSegment{}, nil
}

func (RootSoftware) ArrangeNomination(_ context.Context, req *QueryArrangeNomination) (*ReplyArrangeNomination, error) {
	txs := make([][]byte, 0, len(req.Txs))
	var sumOctets int64
	for _, tx := range req.Txs {
		sumOctets += int64(len(tx))
		if sumOctets > req.MaximumTransferOctets {
			break
		}
		txs = append(txs, tx)
	}
	return &ReplyArrangeNomination{Txs: txs}, nil
}

func (RootSoftware) HandleNomination(context.Context, *QueryHandleNomination) (*ReplyHandleNomination, error) {
	return &ReplyHandleNomination{Status: Responseprocessnomination_ALLOW}, nil
}

func (RootSoftware) ExpandBallot(context.Context, *QueryExpandBallot) (*ReplyExpandBallot, error) {
	return &ReplyExpandBallot{}, nil
}

func (RootSoftware) ValidateBallotAddition(context.Context, *QueryValidateBallotAddition) (*ReplyValidateBallotAddition, error) {
	return &ReplyValidateBallotAddition{
		Status: Responseverifyballotextension_ALLOW,
	}, nil
}

func (RootSoftware) CompleteLedger(_ context.Context, req *QueryCompleteLedger) (*ReplyCompleteLedger, error) {
	txs := make([]*InvokeTransferOutcome, len(req.Txs))
	for i := range req.Txs {
		txs[i] = &InvokeTransferOutcome{Code: CodeKindSuccess}
	}
	return &ReplyCompleteLedger{
		TransOutcomes: txs,
	}, nil
}
