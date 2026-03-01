package kinds

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

type IterationBallotAssign struct {
	Preballots   *kinds.BallotAssign
	Preendorsements *kinds.BallotAssign
}

var FaultAttainedBallotOriginatingUndesiredIteration = errors.New(
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
type AltitudeBallotAssign struct {
	successionUUID           string
	altitude            int64
	itemAssign            *kinds.AssessorAssign
	additionsActivated bool

	mtx               sync.Mutex
	iteration             int32                  //
	iterationBallotCollections     map[int32]IterationBallotAssign //
	nodeOvertakeCycles map[p2p.ID][]int32     //
}

func FreshAltitudeBallotAssign(successionUUID string, altitude int64, itemAssign *kinds.AssessorAssign) *AltitudeBallotAssign {
	hvs := &AltitudeBallotAssign{
		successionUUID:           successionUUID,
		additionsActivated: false,
	}
	hvs.Restore(altitude, itemAssign)
	return hvs
}

func FreshExpandedAltitudeBallotAssign(successionUUID string, altitude int64, itemAssign *kinds.AssessorAssign) *AltitudeBallotAssign {
	hvs := &AltitudeBallotAssign{
		successionUUID:           successionUUID,
		additionsActivated: true,
	}
	hvs.Restore(altitude, itemAssign)
	return hvs
}

func (hvs *AltitudeBallotAssign) Restore(altitude int64, itemAssign *kinds.AssessorAssign) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()

	hvs.altitude = altitude
	hvs.itemAssign = itemAssign
	hvs.iterationBallotCollections = make(map[int32]IterationBallotAssign)
	hvs.nodeOvertakeCycles = make(map[p2p.ID][]int32)

	hvs.appendIteration(0)
	hvs.iteration = 0
}

func (hvs *AltitudeBallotAssign) Altitude() int64 {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.altitude
}

func (hvs *AltitudeBallotAssign) Iteration() int32 {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.iteration
}

//
func (hvs *AltitudeBallotAssign) AssignIteration(iteration int32) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	freshIteration := strongarithmetic.SecureSubtractIntThirtyTwo(hvs.iteration, 1)
	if hvs.iteration != 0 && (iteration < freshIteration) {
		panic("REDACTED")
	}
	for r := freshIteration; r <= iteration; r++ {
		if _, ok := hvs.iterationBallotCollections[r]; ok {
			continue //
		}
		hvs.appendIteration(r)
	}
	hvs.iteration = iteration
}

func (hvs *AltitudeBallotAssign) appendIteration(iteration int32) {
	if _, ok := hvs.iterationBallotCollections[iteration]; ok {
		panic("REDACTED")
	}
	//
	preballots := kinds.FreshBallotAssign(hvs.successionUUID, hvs.altitude, iteration, commitchema.PreballotKind, hvs.itemAssign)
	var preendorsements *kinds.BallotAssign
	if hvs.additionsActivated {
		preendorsements = kinds.FreshExpandedBallotAssign(hvs.successionUUID, hvs.altitude, iteration, commitchema.PreendorseKind, hvs.itemAssign)
	} else {
		preendorsements = kinds.FreshBallotAssign(hvs.successionUUID, hvs.altitude, iteration, commitchema.PreendorseKind, hvs.itemAssign)
	}
	hvs.iterationBallotCollections[iteration] = IterationBallotAssign{
		Preballots:   preballots,
		Preendorsements: preendorsements,
	}
}

//
//
func (hvs *AltitudeBallotAssign) AppendBallot(ballot *kinds.Ballot, nodeUUID p2p.ID, addnActivated bool) (appended bool, err error) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	if hvs.additionsActivated != addnActivated {
		panic(fmt.Errorf("REDACTED", hvs.additionsActivated, addnActivated))
	}
	if !kinds.EqualsBallotKindSound(ballot.Kind) {
		return
	}
	ballotAssign := hvs.obtainBallotAssign(ballot.Iteration, ballot.Kind)
	if ballotAssign == nil {
		if rendezvous := hvs.nodeOvertakeCycles[nodeUUID]; len(rendezvous) < 2 {
			hvs.appendIteration(ballot.Iteration)
			ballotAssign = hvs.obtainBallotAssign(ballot.Iteration, ballot.Kind)
			hvs.nodeOvertakeCycles[nodeUUID] = append(rendezvous, ballot.Iteration)
		} else {
			//
			err = FaultAttainedBallotOriginatingUndesiredIteration
			return
		}
	}
	appended, err = ballotAssign.AppendBallot(ballot)
	return
}

func (hvs *AltitudeBallotAssign) Preballots(iteration int32) *kinds.BallotAssign {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.obtainBallotAssign(iteration, commitchema.PreballotKind)
}

func (hvs *AltitudeBallotAssign) Preendorsements(iteration int32) *kinds.BallotAssign {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return hvs.obtainBallotAssign(iteration, commitchema.PreendorseKind)
}

//
//
func (hvs *AltitudeBallotAssign) PolicyDetails() (policyIteration int32, policyLedgerUUID kinds.LedgerUUID) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	for r := hvs.iteration; r >= 0; r-- {
		rvs := hvs.obtainBallotAssign(r, commitchema.PreballotKind)
		policyLedgerUUID, ok := rvs.CoupleTrinityPreponderance()
		if ok {
			return r, policyLedgerUUID
		}
	}
	return -1, kinds.LedgerUUID{}
}

func (hvs *AltitudeBallotAssign) obtainBallotAssign(iteration int32, ballotKind commitchema.AttestedSignalKind) *kinds.BallotAssign {
	rvs, ok := hvs.iterationBallotCollections[iteration]
	if !ok {
		return nil
	}
	switch ballotKind {
	case commitchema.PreballotKind:
		return rvs.Preballots
	case commitchema.PreendorseKind:
		return rvs.Preendorsements
	default:
		panic(fmt.Sprintf("REDACTED", ballotKind))
	}
}

//
//
//
//
func (hvs *AltitudeBallotAssign) AssignNodeMajor23(
	iteration int32,
	ballotKind commitchema.AttestedSignalKind,
	nodeUUID p2p.ID,
	ledgerUUID kinds.LedgerUUID,
) error {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	if !kinds.EqualsBallotKindSound(ballotKind) {
		return fmt.Errorf("REDACTED", ballotKind)
	}
	ballotAssign := hvs.obtainBallotAssign(iteration, ballotKind)
	if ballotAssign == nil {
		return nil //
	}
	return ballotAssign.AssignNodeMajor23(kinds.Nodeid(nodeUUID), ledgerUUID)
}

//
//

func (hvs *AltitudeBallotAssign) Text() string {
	return hvs.TextFormatted("REDACTED")
}

func (hvs *AltitudeBallotAssign) TextFormatted(format string) string {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	versusTexts := make([]string, 0, (len(hvs.iterationBallotCollections)+1)*2)
	//
	for iteration := int32(0); iteration <= hvs.iteration; iteration++ {
		ballotAssignText := hvs.iterationBallotCollections[iteration].Preballots.TextBrief()
		versusTexts = append(versusTexts, ballotAssignText)
		ballotAssignText = hvs.iterationBallotCollections[iteration].Preendorsements.TextBrief()
		versusTexts = append(versusTexts, ballotAssignText)
	}
	//
	for iteration, iterationBallotAssign := range hvs.iterationBallotCollections {
		if iteration <= hvs.iteration {
			continue
		}
		ballotAssignText := iterationBallotAssign.Preballots.TextBrief()
		versusTexts = append(versusTexts, ballotAssignText)
		ballotAssignText = iterationBallotAssign.Preendorsements.TextBrief()
		versusTexts = append(versusTexts, ballotAssignText)
	}
	return fmt.Sprintf(`REDACTEDv
REDACTEDv
REDACTED`,
		hvs.altitude, hvs.iteration,
		format, strings.Join(versusTexts, "REDACTED"+format+"REDACTED"),
		format)
}

func (hvs *AltitudeBallotAssign) SerializeJSN() ([]byte, error) {
	hvs.mtx.Lock()
	defer hvs.mtx.Unlock()
	return strongmindjson.Serialize(hvs.towardEveryIterationBallots())
}

func (hvs *AltitudeBallotAssign) towardEveryIterationBallots() []iterationBallots {
	sumCycles := hvs.iteration + 1
	everyBallots := make([]iterationBallots, sumCycles)
	//
	for iteration := int32(0); iteration < sumCycles; iteration++ {
		everyBallots[iteration] = iterationBallots{
			Iteration:              iteration,
			Preballots:           hvs.iterationBallotCollections[iteration].Preballots.BallotTexts(),
			PreballotsDigitSeries:   hvs.iterationBallotCollections[iteration].Preballots.DigitSeriesText(),
			Preendorsements:         hvs.iterationBallotCollections[iteration].Preendorsements.BallotTexts(),
			PreendorsementsDigitSeries: hvs.iterationBallotCollections[iteration].Preendorsements.DigitSeriesText(),
		}
	}
	//
	return everyBallots
}

type iterationBallots struct {
	Iteration              int32    `json:"iteration"`
	Preballots           []string `json:"preballots"`
	PreballotsDigitSeries   string   `json:"preballots_digit_series"`
	Preendorsements         []string `json:"preendorsements"`
	PreendorsementsDigitSeries string   `json:"preendorsements_digit_series"`
}
