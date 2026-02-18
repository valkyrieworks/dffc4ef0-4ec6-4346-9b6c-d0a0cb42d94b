package kinds

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	cometjson "github.com/valkyrieworks/utils/json"
	cometmath "github.com/valkyrieworks/utils/math"
	"github.com/valkyrieworks/p2p"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

type EpochBallotCollection struct {
	Preballots   *kinds.BallotCollection
	Preendorsements *kinds.BallotCollection
}

var ErrAcquiredBallotFromUndesiredEpoch = errors.New(
	"REDACTED",
)

/**
.

n
.

,
r
,
e
.
.
.
*/
type LevelBallotCollection struct {
	ledgerUID           string
	level            int64
	valueCollection            *kinds.RatifierAssign
	pluginsActivated bool

	mtx               sync.Mutex
	duration             int32                  //
	epochBallotCollections     map[int32]EpochBallotCollection //
	nodeOvertakeIterations map[p2p.ID][]int32     //
}

func NewLevelBallotCollection(ledgerUID string, level int64, valueCollection *kinds.RatifierAssign) *LevelBallotCollection {
	hvs := &LevelBallotCollection{
		ledgerUID:           ledgerUID,
		pluginsActivated: false,
	}
	hvs.Restore(level, valueCollection)
	return hvs
}

func NewExpandedLevelBallotCollection(ledgerUID string, level int64, valueCollection *kinds.RatifierAssign) *LevelBallotCollection {
	hvs := &LevelBallotCollection{
		ledgerUID:           ledgerUID,
		pluginsActivated: true,
	}
	hvs.Restore(level, valueCollection)
	return hvs
}

func (hvs *LevelBallotCollection) Restore(level int64, valueCollection *kinds.RatifierAssign) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()

	hvs.level = level
	hvs.valueCollection = valueCollection
	hvs.epochBallotCollections = make(map[int32]EpochBallotCollection)
	hvs.nodeOvertakeIterations = make(map[p2p.ID][]int32)

	hvs.appendEpoch(0)
	hvs.duration = 0
}

func (hvs *LevelBallotCollection) Level() int64 {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.level
}

func (hvs *LevelBallotCollection) Cycle() int32 {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.duration
}

//
func (hvs *LevelBallotCollection) CollectionEpoch(duration int32) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	newEpoch := cometmath.SecureSubtractInt32(hvs.duration, 1)
	if hvs.duration != 0 && (duration < newEpoch) {
		panic("REDACTED")
	}
	for r := newEpoch; r <= duration; r++ {
		if _, ok := hvs.epochBallotCollections[r]; ok {
			continue //
		}
		hvs.appendEpoch(r)
	}
	hvs.duration = duration
}

func (hvs *LevelBallotCollection) appendEpoch(duration int32) {
	if _, ok := hvs.epochBallotCollections[duration]; ok {
		panic("REDACTED")
	}
	//
	preballots := kinds.NewBallotCollection(hvs.ledgerUID, hvs.level, duration, engineproto.PreballotKind, hvs.valueCollection)
	var preendorsements *kinds.BallotCollection
	if hvs.pluginsActivated {
		preendorsements = kinds.NewExpandedBallotCollection(hvs.ledgerUID, hvs.level, duration, engineproto.PreendorseKind, hvs.valueCollection)
	} else {
		preendorsements = kinds.NewBallotCollection(hvs.ledgerUID, hvs.level, duration, engineproto.PreendorseKind, hvs.valueCollection)
	}
	hvs.epochBallotCollections[duration] = EpochBallotCollection{
		Preballots:   preballots,
		Preendorsements: preendorsements,
	}
}

//
//
func (hvs *LevelBallotCollection) AppendBallot(ballot *kinds.Ballot, nodeUID p2p.ID, extensionActivated bool) (appended bool, err error) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	if hvs.pluginsActivated != extensionActivated {
		panic(fmt.Errorf("REDACTED", hvs.pluginsActivated, extensionActivated))
	}
	if !kinds.IsBallotKindSound(ballot.Kind) {
		return
	}
	ballotCollection := hvs.fetchBallotCollection(ballot.Cycle, ballot.Kind)
	if ballotCollection == nil {
		if rndz := hvs.nodeOvertakeIterations[nodeUID]; len(rndz) < 2 {
			hvs.appendEpoch(ballot.Cycle)
			ballotCollection = hvs.fetchBallotCollection(ballot.Cycle, ballot.Kind)
			hvs.nodeOvertakeIterations[nodeUID] = append(rndz, ballot.Cycle)
		} else {
			//
			err = ErrAcquiredBallotFromUndesiredEpoch
			return
		}
	}
	appended, err = ballotCollection.AppendBallot(ballot)
	return
}

func (hvs *LevelBallotCollection) Preballots(duration int32) *kinds.BallotCollection {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.fetchBallotCollection(duration, engineproto.PreballotKind)
}

func (hvs *LevelBallotCollection) Preendorsements(duration int32) *kinds.BallotCollection {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.fetchBallotCollection(duration, engineproto.PreendorseKind)
}

//
//
func (hvs *LevelBallotCollection) POLDetails() (polEpoch int32, polLedgerUID kinds.LedgerUID) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	for r := hvs.duration; r >= 0; r-- {
		rvs := hvs.fetchBallotCollection(r, engineproto.PreballotKind)
		polLedgerUID, ok := rvs.DualThirdsBulk()
		if ok {
			return r, polLedgerUID
		}
	}
	return -1, kinds.LedgerUID{}
}

func (hvs *LevelBallotCollection) fetchBallotCollection(duration int32, ballotKind engineproto.AttestedMessageKind) *kinds.BallotCollection {
	rvs, ok := hvs.epochBallotCollections[duration]
	if !ok {
		return nil
	}
	switch ballotKind {
	case engineproto.PreballotKind:
		return rvs.Preballots
	case engineproto.PreendorseKind:
		return rvs.Preendorsements
	default:
		panic(fmt.Sprintf("REDACTED", ballotKind))
	}
}

//
//
//
//
func (hvs *LevelBallotCollection) AssignNodeMaj23(
	duration int32,
	ballotKind engineproto.AttestedMessageKind,
	nodeUID p2p.ID,
	ledgerUID kinds.LedgerUID,
) error {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	if !kinds.IsBallotKindSound(ballotKind) {
		return fmt.Errorf("REDACTED", ballotKind)
	}
	ballotCollection := hvs.fetchBallotCollection(duration, ballotKind)
	if ballotCollection == nil {
		return nil //
	}
	return ballotCollection.AssignNodeMaj23(kinds.P2pid(nodeUID), ledgerUID)
}

//
//

func (hvs *LevelBallotCollection) String() string {
	return hvs.StringIndented("REDACTED")
}

func (hvs *LevelBallotCollection) StringIndented(indent string) string {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	vsStrings := make([]string, 0, (len(hvs.epochBallotCollections)+1)*2)
	//
	for duration := int32(0); duration <= hvs.duration; duration++ {
		ballotCollectionString := hvs.epochBallotCollections[duration].Preballots.StringBrief()
		vsStrings = append(vsStrings, ballotCollectionString)
		ballotCollectionString = hvs.epochBallotCollections[duration].Preendorsements.StringBrief()
		vsStrings = append(vsStrings, ballotCollectionString)
	}
	//
	for duration, epochBallotCollection := range hvs.epochBallotCollections {
		if duration <= hvs.duration {
			continue
		}
		ballotCollectionString := epochBallotCollection.Preballots.StringBrief()
		vsStrings = append(vsStrings, ballotCollectionString)
		ballotCollectionString = epochBallotCollection.Preendorsements.StringBrief()
		vsStrings = append(vsStrings, ballotCollectionString)
	}
	return fmt.Sprintf(`REDACTEDv
REDACTEDv
REDACTED`,
		hvs.level, hvs.duration,
		indent, strings.Join(vsStrings, "REDACTED"+indent+"REDACTED"),
		indent)
}

func (hvs *LevelBallotCollection) SerializeJSON() ([]byte, error) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return cometjson.Serialize(hvs.toAllEpochBallots())
}

func (hvs *LevelBallotCollection) toAllEpochBallots() []epochBallots {
	sumIterations := hvs.duration + 1
	allBallots := make([]epochBallots, sumIterations)
	//
	for duration := int32(0); duration < sumIterations; duration++ {
		allBallots[duration] = epochBallots{
			Cycle:              duration,
			Preballots:           hvs.epochBallotCollections[duration].Preballots.BallotStrings(),
			PreballotsBitList:   hvs.epochBallotCollections[duration].Preballots.BitListString(),
			Preendorsements:         hvs.epochBallotCollections[duration].Preendorsements.BallotStrings(),
			PreendorsementsBitList: hvs.epochBallotCollections[duration].Preendorsements.BitListString(),
		}
	}
	//
	return allBallots
}

type epochBallots struct {
	Cycle              int32    `json:"duration"`
	Preballots           []string `json:"preballots"`
	PreballotsBitList   string   `json:"preballots_bit_list"`
	Preendorsements         []string `json:"preendorsements"`
	PreendorsementsBitList string   `json:"preendorsements_bit_list"`
}
