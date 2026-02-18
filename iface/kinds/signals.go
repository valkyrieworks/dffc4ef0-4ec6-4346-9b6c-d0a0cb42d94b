package kinds

import (
	"io"
	"math"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/utils/protoio"
)

const (
	maximumMessageVolume = math.MaxInt32 //
)

//
func RecordSignal(msg proto.Message, w io.Writer) error {
	schemaRecorder := protoio.NewSeparatedRecorder(w)
	_, err := schemaRecorder.RecordMessage(msg)
	return err
}

//
func ScanSignal(r io.Reader, msg proto.Message) error {
	_, err := protoio.NewSeparatedScanner(r, maximumMessageVolume).ScanMessage(msg)
	return err
}

//

func ToQueryReverberate(signal string) *Query {
	return &Query{
		Item: &Query_Reverberate{&QueryReverberate{Signal: signal}},
	}
}

func ToQueryPurge() *Query {
	return &Query{
		Item: &Query_Purge{&QueryPurge{}},
	}
}

func ToQueryDetails(req *QueryDetails) *Query {
	return &Query{
		Item: &Query_Details{req},
	}
}

func ToQueryInspectTransfer(req *QueryInspectTransfer) *Query {
	return &Query{
		Item: &Query_Transfercheck{req},
	}
}

func ToQueryEmbedTransfer(req *QueryEmbedTransfer) *Query {
	return &Query{
		Item: &Query_Transferinsert{req},
	}
}

func ToQueryHarvestTrans(req *QueryHarvestTrans) *Query {
	return &Query{
		Item: &Query_Reaptransfers{req},
	}
}

func ToQueryEndorse() *Query {
	return &Query{
		Item: &Query_Endorse{&QueryEndorse{}},
	}
}

func ToQueryInquire(req *QueryInquire) *Query {
	return &Query{
		Item: &Query_Inquire{req},
	}
}

func ToQueryInitSeries(req *QueryInitSeries) *Query {
	return &Query{
		Item: &Query_Initiatechain{req},
	}
}

func ToQueryCatalogMirrors(req *QueryCatalogMirrors) *Query {
	return &Query{
		Item: &Query_Catalogmirrors{req},
	}
}

func ToQueryProposalMirror(req *QueryProposalMirror) *Query {
	return &Query{
		Item: &Query_Mirrorsnapshot{req},
	}
}

func ToQueryImportMirrorSegment(req *QueryImportMirrorSegment) *Query {
	return &Query{
		Item: &Query_Loadmirrorsegment{req},
	}
}

func ToQueryExecuteMirrorSegment(req *QueryExecuteMirrorSegment) *Query {
	return &Query{
		Item: &Query_Executemirrorsegment{req},
	}
}

func ToQueryArrangeNomination(req *QueryArrangeNomination) *Query {
	return &Query{
		Item: &Query_Arrangenomination{req},
	}
}

func ToQueryHandleNomination(req *QueryHandleNomination) *Query {
	return &Query{
		Item: &Query_Processnomination{req},
	}
}

func ToQueryExpandBallot(req *QueryExpandBallot) *Query {
	return &Query{
		Item: &Query_Ballotextend{req},
	}
}

func ToQueryValidateBallotAddition(req *QueryValidateBallotAddition) *Query {
	return &Query{
		Item: &Query_Validateballotextension{req},
	}
}

func ToQueryCompleteLedger(req *QueryCompleteLedger) *Query {
	return &Query{
		Item: &Query_Terminateblock{req},
	}
}

//

func ToReplyExemption(errStr string) *Reply {
	return &Reply{
		Item: &Reply_Exemption{&ReplyExemption{Fault: errStr}},
	}
}

func ToReplyReverberate(signal string) *Reply {
	return &Reply{
		Item: &Reply_Reverberate{&ReplyReverberate{Signal: signal}},
	}
}

func ToReplyPurge() *Reply {
	return &Reply{
		Item: &Reply_Purge{&ReplyPurge{}},
	}
}

func ToReplyDetails(res *ReplyDetails) *Reply {
	return &Reply{
		Item: &Reply_Details{res},
	}
}

func ToReplyInspectTransfer(res *ReplyInspectTransfer) *Reply {
	return &Reply{
		Item: &Reply_Transfercheck{res},
	}
}

func ToReplyEndorse(res *ReplyEndorse) *Reply {
	return &Reply{
		Item: &Reply_Endorse{res},
	}
}

func ToReplyInquire(res *ReplyInquire) *Reply {
	return &Reply{
		Item: &Reply_Inquire{res},
	}
}

func ToReplyInitSeries(res *ReplyInitSeries) *Reply {
	return &Reply{
		Item: &Reply_Initiatechain{res},
	}
}

func ToReplyCatalogMirrors(res *ReplyCatalogMirrors) *Reply {
	return &Reply{
		Item: &Reply_Catalogmirrors{res},
	}
}

func ToReplyProposalMirror(res *ReplyProposalMirror) *Reply {
	return &Reply{
		Item: &Reply_Mirrorsnapshot{res},
	}
}

func ToReplyImportMirrorSegment(res *ReplyImportMirrorSegment) *Reply {
	return &Reply{
		Item: &Reply_Loadmirrorsegment{res},
	}
}

func ToReplyExecuteMirrorSegment(res *ReplyExecuteMirrorSegment) *Reply {
	return &Reply{
		Item: &Reply_Executemirrorsegment{res},
	}
}

func ToReplyArrangeNomination(res *ReplyArrangeNomination) *Reply {
	return &Reply{
		Item: &Reply_Arrangenomination{res},
	}
}

func ToReplyHandleNomination(res *ReplyHandleNomination) *Reply {
	return &Reply{
		Item: &Reply_Processnomination{res},
	}
}

func ToReplyExpandBallot(res *ReplyExpandBallot) *Reply {
	return &Reply{
		Item: &Reply_Ballotextend{res},
	}
}

func ToReplyValidateBallotAddition(res *ReplyValidateBallotAddition) *Reply {
	return &Reply{
		Item: &Reply_Validateballotextension{res},
	}
}

func ToReplyCompleteLedger(res *ReplyCompleteLedger) *Reply {
	return &Reply{
		Item: &Reply_Terminateblock{res},
	}
}
