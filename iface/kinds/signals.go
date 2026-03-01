package kinds

import (
	"io"
	"math"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
)

const (
	maximumSignalExtent = math.MaxInt32 //
)

//
func PersistArtifact(msg proto.Message, w io.Writer) error {
	schemaPersistor := protocolio.FreshSeparatedPersistor(w)
	_, err := schemaPersistor.PersistSignal(msg)
	return err
}

//
func FetchArtifact(r io.Reader, msg proto.Message) error {
	_, err := protocolio.FreshSeparatedFetcher(r, maximumSignalExtent).FetchSignal(msg)
	return err
}

//

func TowardSolicitReverberate(signal string) *Solicit {
	return &Solicit{
		Datum: &Solicit_Reverberate{&SolicitReverberate{Signal: signal}},
	}
}

func TowardSolicitPurge() *Solicit {
	return &Solicit{
		Datum: &Solicit_Purge{&SolicitPurge{}},
	}
}

func TowardSolicitDetails(req *SolicitDetails) *Solicit {
	return &Solicit{
		Datum: &Solicit_Details{req},
	}
}

func TowardSolicitInspectTransfer(req *SolicitInspectTransfer) *Solicit {
	return &Solicit{
		Datum: &Solicit_Inspecttrans{req},
	}
}

func TowardSolicitAppendTransfer(req *SolicitAppendTransfer) *Solicit {
	return &Solicit{
		Datum: &Solicit_Appendtrans{req},
	}
}

func TowardSolicitHarvestTrans(req *SolicitHarvestTrans) *Solicit {
	return &Solicit{
		Datum: &Solicit_Harvesttrans{req},
	}
}

func TowardSolicitEndorse() *Solicit {
	return &Solicit{
		Datum: &Solicit_Endorse{&SolicitEndorse{}},
	}
}

func TowardSolicitInquire(req *SolicitInquire) *Solicit {
	return &Solicit{
		Datum: &Solicit_Inquire{req},
	}
}

func TowardSolicitInitializeSuccession(req *SolicitInitializeSuccession) *Solicit {
	return &Solicit{
		Datum: &Solicit_Initiatechain{req},
	}
}

func TowardSolicitCatalogImages(req *SolicitCollectionImages) *Solicit {
	return &Solicit{
		Datum: &Solicit_Catalogimages{req},
	}
}

func TowardSolicitExtendImage(req *SolicitExtendImage) *Solicit {
	return &Solicit{
		Datum: &Solicit_Extendimage{req},
	}
}

func TowardSolicitFetchImageSegment(req *SolicitFetchImageSegment) *Solicit {
	return &Solicit{
		Datum: &Solicit_Loadimagefragment{req},
	}
}

func TowardSolicitExecuteImageSegment(req *SolicitExecuteImageSegment) *Solicit {
	return &Solicit{
		Datum: &Solicit_Executeimagefragment{req},
	}
}

func TowardSolicitArrangeNomination(req *SolicitArrangeNomination) *Solicit {
	return &Solicit{
		Datum: &Solicit_Prepareitem{req},
	}
}

func TowardSolicitHandleNomination(req *SolicitHandleNomination) *Solicit {
	return &Solicit{
		Datum: &Solicit_Executeitem{req},
	}
}

func TowardSolicitBroadenBallot(req *SolicitBroadenBallot) *Solicit {
	return &Solicit{
		Datum: &Solicit_Extendballot{req},
	}
}

func TowardSolicitValidateBallotAddition(req *SolicitValidateBallotAddition) *Solicit {
	return &Solicit{
		Datum: &Solicit_Verifyballotaddition{req},
	}
}

func TowardSolicitCulminateLedger(req *SolicitCulminateLedger) *Solicit {
	return &Solicit{
		Datum: &Solicit_Finalizeledger{req},
	}
}

//

func TowardReplyExemption(faultTxt string) *Reply {
	return &Reply{
		Datum: &Reply_Exemption{&ReplyExemption{Failure: faultTxt}},
	}
}

func TowardReplyReverberate(signal string) *Reply {
	return &Reply{
		Datum: &Reply_Reverberate{&ReplyReverberate{Signal: signal}},
	}
}

func TowardReplyPurge() *Reply {
	return &Reply{
		Datum: &Reply_Purge{&ReplyPurge{}},
	}
}

func TowardReplyDetails(res *ReplyDetails) *Reply {
	return &Reply{
		Datum: &Reply_Details{res},
	}
}

func TowardReplyInspectTransfer(res *ReplyInspectTransfer) *Reply {
	return &Reply{
		Datum: &Reply_Inspecttrans{res},
	}
}

func TowardReplyEndorse(res *ReplyEndorse) *Reply {
	return &Reply{
		Datum: &Reply_Endorse{res},
	}
}

func TowardReplyInquire(res *ReplyInquire) *Reply {
	return &Reply{
		Datum: &Reply_Inquire{res},
	}
}

func TowardReplyInitializeSuccession(res *ReplyInitializeSuccession) *Reply {
	return &Reply{
		Datum: &Reply_Initiatechain{res},
	}
}

func TowardReplyCatalogImages(res *ReplyCatalogImages) *Reply {
	return &Reply{
		Datum: &Reply_Catalogimages{res},
	}
}

func TowardReplyExtendImage(res *ReplyExtendImage) *Reply {
	return &Reply{
		Datum: &Reply_Extendimage{res},
	}
}

func TowardReplyFetchImageSegment(res *ReplyFetchImageSegment) *Reply {
	return &Reply{
		Datum: &Reply_Loadimagefragment{res},
	}
}

func TowardReplyExecuteImageSegment(res *ReplyExecuteImageSegment) *Reply {
	return &Reply{
		Datum: &Reply_Executeimagefragment{res},
	}
}

func TowardReplyArrangeNomination(res *ReplyArrangeNomination) *Reply {
	return &Reply{
		Datum: &Reply_Prepareitem{res},
	}
}

func TowardReplyHandleNomination(res *ReplyHandleNomination) *Reply {
	return &Reply{
		Datum: &Reply_Executeitem{res},
	}
}

func TowardReplyBroadenBallot(res *ReplyBroadenBallot) *Reply {
	return &Reply{
		Datum: &Reply_Extendballot{res},
	}
}

func TowardReplyValidateBallotAddition(res *ReplyValidateBallotAddition) *Reply {
	return &Reply{
		Datum: &Reply_Verifyballotaddition{res},
	}
}

func TowardReplyCulminateLedger(res *ReplyCulminateLedger) *Reply {
	return &Reply{
		Datum: &Reply_Finalizeledger{res},
	}
}
