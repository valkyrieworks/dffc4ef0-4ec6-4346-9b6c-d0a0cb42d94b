package kinds

import (
	"fmt"

	iface "github.com/valkyrieworks/iface/kinds"
	cometjson "github.com/valkyrieworks/utils/json"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	cmtinquire "github.com/valkyrieworks/utils/broadcast/inquire"
)

//
const (
	//
	//
	//
	//
	//
	EventNewLedger            = "REDACTED"
	EventNewLedgerHeading      = "REDACTED"
	EventNewLedgerEvents      = "REDACTED"
	EventNewProof         = "REDACTED"
	EventTransfer                  = "REDACTED"
	EventRatifierCollectionRefreshes = "REDACTED"

	//
	//
	//
	EventFinishedNomination   = "REDACTED"
	EventSecure               = "REDACTED"
	EventNewEpoch           = "REDACTED"
	EventNewDurationPhase       = "REDACTED"
	EventPolka              = "REDACTED"
	EventResecure             = "REDACTED"
	EventDeadlineNominate     = "REDACTED"
	EventDeadlineWait        = "REDACTED"
	EventRelease             = "REDACTED"
	EventSoundLedger         = "REDACTED"
	EventBallot               = "REDACTED"
	EventNewAgreementOptions = "REDACTED"
)

//

//
type TMEventData interface {
	//
}

func init() {
	cometjson.EnrollKind(EventDataNewLedger{}, "REDACTED")
	cometjson.EnrollKind(EventDataNewLedgerHeading{}, "REDACTED")
	cometjson.EnrollKind(EventDataNewLedgerEvents{}, "REDACTED")
	cometjson.EnrollKind(EventDataNewProof{}, "REDACTED")
	cometjson.EnrollKind(EventDataTransfer{}, "REDACTED")
	cometjson.EnrollKind(EventDataDurationStatus{}, "REDACTED")
	cometjson.EnrollKind(EventDataNewEpoch{}, "REDACTED")
	cometjson.EnrollKind(EventDataFinishedNomination{}, "REDACTED")
	cometjson.EnrollKind(EventDataBallot{}, "REDACTED")
	cometjson.EnrollKind(EventDataRatifierCollectionRefreshes{}, "REDACTED")
	cometjson.EnrollKind(EventDataString("REDACTED"), "REDACTED")
}

//
//

type EventDataNewLedger struct {
	Ledger               *Ledger                     `json:"ledger"`
	LedgerUID             LedgerUID                    `json:"ledger_uid"`
	OutcomeCompleteLedger iface.ReplyCompleteLedger `json:"outcome_complete_ledger"`
}

type EventDataNewLedgerHeading struct {
	Heading Heading `json:"heading"`
}

type EventDataNewLedgerEvents struct {
	Level int64        `json:"level"`
	Events []iface.Event `json:"events"`
	CountTrans int64        `json:"count_trans,string"` //
}

type EventDataNewProof struct {
	Level   int64    `json:"level"`
	Proof Proof `json:"proof"`
}

//
type EventDataTransfer struct {
	iface.TransOutcome
}

//
type EventDataDurationStatus struct {
	Level int64  `json:"level"`
	Cycle  int32  `json:"epoch"`
	Phase   string `json:"phase"`
}

type RatifierDetails struct {
	Location Location `json:"location"`
	Ordinal   int32   `json:"ordinal"`
}

type EventDataNewEpoch struct {
	Level int64  `json:"level"`
	Cycle  int32  `json:"epoch"`
	Phase   string `json:"phase"`

	Recommender RatifierDetails `json:"recommender"`
}

type EventDataFinishedNomination struct {
	Level int64  `json:"level"`
	Cycle  int32  `json:"epoch"`
	Phase   string `json:"phase"`

	LedgerUID LedgerUID `json:"ledger_uid"`
}

type EventDataBallot struct {
	Ballot *Ballot
}

type EventDataString string

type EventDataRatifierCollectionRefreshes struct {
	RatifierRefreshes []*Ratifier `json:"ratifier_refreshes"`
}

//

const (
	//
	EventKindKey = "REDACTED"

	//
	//
	TransferDigestKey = "REDACTED"

	//
	//
	TransferLevelKey = "REDACTED"

	//
	LedgerLevelKey = "REDACTED"
)

var (
	EventInquireFinishedNomination    = InquireForEvent(EventFinishedNomination)
	EventInquireSecure                = InquireForEvent(EventSecure)
	EventInquireNewLedger            = InquireForEvent(EventNewLedger)
	EventInquireNewLedgerHeading      = InquireForEvent(EventNewLedgerHeading)
	EventInquireNewLedgerEvents      = InquireForEvent(EventNewLedgerEvents)
	EventInquireNewProof         = InquireForEvent(EventNewProof)
	EventInquireNewEpoch            = InquireForEvent(EventNewEpoch)
	EventInquireNewEpochPhase        = InquireForEvent(EventNewDurationPhase)
	EventInquirePolka               = InquireForEvent(EventPolka)
	EventInquireResecure              = InquireForEvent(EventResecure)
	EventInquireDeadlineNominate      = InquireForEvent(EventDeadlineNominate)
	EventInquireDeadlineWait         = InquireForEvent(EventDeadlineWait)
	EventInquireTransfer                  = InquireForEvent(EventTransfer)
	EventInquireRelease              = InquireForEvent(EventRelease)
	EventInquireRatifierCollectionRefreshes = InquireForEvent(EventRatifierCollectionRefreshes)
	EventInquireSoundLedger          = InquireForEvent(EventSoundLedger)
	EventInquireBallot                = InquireForEvent(EventBallot)
)

func EventInquireTransferFor(tx Tx) cometbroadcast.Inquire {
	return cmtinquire.ShouldBuild(fmt.Sprintf("REDACTED", EventKindKey, EventTransfer, TransferDigestKey, tx.Digest()))
}

func InquireForEvent(eventKind string) cometbroadcast.Inquire {
	return cmtinquire.ShouldBuild(fmt.Sprintf("REDACTED", EventKindKey, eventKind))
}

//
type LedgerEventBroadcaster interface {
	BroadcastEventNewLedger(ledger EventDataNewLedger) error
	BroadcastEventNewLedgerHeading(heading EventDataNewLedgerHeading) error
	BroadcastEventNewLedgerEvents(events EventDataNewLedgerEvents) error
	BroadcastEventNewProof(proof EventDataNewProof) error
	BroadcastEventTransfer(EventDataTransfer) error
	BroadcastEventRatifierCollectionRefreshes(EventDataRatifierCollectionRefreshes) error
}

type TransferEventBroadcaster interface {
	BroadcastEventTransfer(EventDataTransfer) error
}
