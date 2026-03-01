package kinds

import "context"

//

//
//
type Platform interface {
	//
	Details(context.Context, *SolicitDetails) (*ReplyDetails, error)    //
	Inquire(context.Context, *SolicitInquire) (*ReplyInquire, error) //

	//
	InspectTransfer(context.Context, *SolicitInspectTransfer) (*ReplyInspectTransfer, error)    //
	AppendTransfer(context.Context, *SolicitAppendTransfer) (*ReplyAppendTransfer, error) //
	HarvestTrans(context.Context, *SolicitHarvestTrans) (*ReplyHarvestTrans, error)    //

	//
	InitializeSuccession(context.Context, *SolicitInitializeSuccession) (*ReplyInitializeSuccession, error) //
	ArrangeNomination(context.Context, *SolicitArrangeNomination) (*ReplyArrangeNomination, error)
	HandleNomination(context.Context, *SolicitHandleNomination) (*ReplyHandleNomination, error)
	//
	CulminateLedger(context.Context, *SolicitCulminateLedger) (*ReplyCulminateLedger, error)
	//
	BroadenBallot(context.Context, *SolicitBroadenBallot) (*ReplyBroadenBallot, error)
	//
	ValidateBallotAddition(context.Context, *SolicitValidateBallotAddition) (*ReplyValidateBallotAddition, error)
	//
	Endorse(context.Context, *SolicitEndorse) (*ReplyEndorse, error)

	//
	CollectionImages(context.Context, *SolicitCollectionImages) (*ReplyCatalogImages, error)                //
	ExtendImage(context.Context, *SolicitExtendImage) (*ReplyExtendImage, error)                //
	FetchImageSegment(context.Context, *SolicitFetchImageSegment) (*ReplyFetchImageSegment, error)    //
	ExecuteImageSegment(context.Context, *SolicitExecuteImageSegment) (*ReplyExecuteImageSegment, error) //
}

//
//

var _ Platform = (*FoundationPlatform)(nil)

type FoundationPlatform struct{}

func FreshFoundationPlatform() *FoundationPlatform {
	return &FoundationPlatform{}
}

func (FoundationPlatform) Details(context.Context, *SolicitDetails) (*ReplyDetails, error) {
	return &ReplyDetails{}, nil
}

func (FoundationPlatform) InspectTransfer(context.Context, *SolicitInspectTransfer) (*ReplyInspectTransfer, error) {
	return &ReplyInspectTransfer{Cipher: CipherKindOKAY}, nil
}

func (FoundationPlatform) AppendTransfer(context.Context, *SolicitAppendTransfer) (*ReplyAppendTransfer, error) {
	return &ReplyAppendTransfer{Cipher: CipherKindOKAY}, nil
}

func (FoundationPlatform) HarvestTrans(context.Context, *SolicitHarvestTrans) (*ReplyHarvestTrans, error) {
	return &ReplyHarvestTrans{}, nil
}

func (FoundationPlatform) Endorse(context.Context, *SolicitEndorse) (*ReplyEndorse, error) {
	return &ReplyEndorse{}, nil
}

func (FoundationPlatform) Inquire(context.Context, *SolicitInquire) (*ReplyInquire, error) {
	return &ReplyInquire{Cipher: CipherKindOKAY}, nil
}

func (FoundationPlatform) InitializeSuccession(context.Context, *SolicitInitializeSuccession) (*ReplyInitializeSuccession, error) {
	return &ReplyInitializeSuccession{}, nil
}

func (FoundationPlatform) CollectionImages(context.Context, *SolicitCollectionImages) (*ReplyCatalogImages, error) {
	return &ReplyCatalogImages{}, nil
}

func (FoundationPlatform) ExtendImage(context.Context, *SolicitExtendImage) (*ReplyExtendImage, error) {
	return &ReplyExtendImage{}, nil
}

func (FoundationPlatform) FetchImageSegment(context.Context, *SolicitFetchImageSegment) (*ReplyFetchImageSegment, error) {
	return &ReplyFetchImageSegment{}, nil
}

func (FoundationPlatform) ExecuteImageSegment(context.Context, *SolicitExecuteImageSegment) (*ReplyExecuteImageSegment, error) {
	return &ReplyExecuteImageSegment{}, nil
}

func (FoundationPlatform) ArrangeNomination(_ context.Context, req *SolicitArrangeNomination) (*ReplyArrangeNomination, error) {
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

func (FoundationPlatform) HandleNomination(context.Context, *SolicitHandleNomination) (*ReplyHandleNomination, error) {
	return &ReplyHandleNomination{Condition: Responseexecuteitem_EMBRACE}, nil
}

func (FoundationPlatform) BroadenBallot(context.Context, *SolicitBroadenBallot) (*ReplyBroadenBallot, error) {
	return &ReplyBroadenBallot{}, nil
}

func (FoundationPlatform) ValidateBallotAddition(context.Context, *SolicitValidateBallotAddition) (*ReplyValidateBallotAddition, error) {
	return &ReplyValidateBallotAddition{
		Condition: Responsecertifyballotaddition_EMBRACE,
	}, nil
}

func (FoundationPlatform) CulminateLedger(_ context.Context, req *SolicitCulminateLedger) (*ReplyCulminateLedger, error) {
	txs := make([]*InvokeTransferOutcome, len(req.Txs))
	for i := range req.Txs {
		txs[i] = &InvokeTransferOutcome{Cipher: CipherKindOKAY}
	}
	return &ReplyCulminateLedger{
		TransferOutcomes: txs,
	}, nil
}
