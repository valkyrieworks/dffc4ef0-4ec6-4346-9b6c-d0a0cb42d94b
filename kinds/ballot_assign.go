package kinds

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

const (
	//
	//
	//
	MaximumBallotsTally = 10000
)

//
//
//
//
type Nodeid string

/**
a
.

s
d
.

.
s
k

,
f
t
.

e
.
.
.

`
.

s
.

r
,
.

.
*/
type BallotAssign struct {
	successionUUID           string
	altitude            int64
	iteration             int32
	notatedSignalKind     commitchema.AttestedSignalKind
	itemAssign            *AssessorAssign
	additionsActivated bool

	mtx           commitchronize.Exclusion
	ballotsDigitSeries *digits.DigitSeries
	ballots         []*Ballot                //
	sum           int64                  //
	major23         *LedgerUUID               //
	ballotsViaLedger  map[string]*ledgerBallots //
	nodeMajor23s    map[Nodeid]LedgerUUID      //
}

//
//
func FreshBallotAssign(successionUUID string, altitude int64, iteration int32,
	notatedSignalKind commitchema.AttestedSignalKind, itemAssign *AssessorAssign,
) *BallotAssign {
	if altitude == 0 {
		panic("REDACTED")
	}
	return &BallotAssign{
		successionUUID:       successionUUID,
		altitude:        altitude,
		iteration:         iteration,
		notatedSignalKind: notatedSignalKind,
		itemAssign:        itemAssign,
		ballotsDigitSeries: digits.FreshDigitCollection(itemAssign.Extent()),
		ballots:         make([]*Ballot, itemAssign.Extent()),
		sum:           0,
		major23:         nil,
		ballotsViaLedger:  make(map[string]*ledgerBallots, itemAssign.Extent()),
		nodeMajor23s:    make(map[Nodeid]LedgerUUID),
	}
}

//
//
//
func FreshExpandedBallotAssign(successionUUID string, altitude int64, iteration int32,
	notatedSignalKind commitchema.AttestedSignalKind, itemAssign *AssessorAssign,
) *BallotAssign {
	vs := FreshBallotAssign(successionUUID, altitude, iteration, notatedSignalKind, itemAssign)
	vs.additionsActivated = true
	return vs
}

func (ballotAssign *BallotAssign) SuccessionUUID() string {
	return ballotAssign.successionUUID
}

//
func (ballotAssign *BallotAssign) ObtainAltitude() int64 {
	if ballotAssign == nil {
		return 0
	}
	return ballotAssign.altitude
}

//
func (ballotAssign *BallotAssign) ObtainIteration() int32 {
	if ballotAssign == nil {
		return -1
	}
	return ballotAssign.iteration
}

//
func (ballotAssign *BallotAssign) Kind() byte {
	if ballotAssign == nil {
		return 0x00
	}
	return byte(ballotAssign.notatedSignalKind)
}

//
func (ballotAssign *BallotAssign) Extent() int {
	if ballotAssign == nil {
		return 0
	}
	return ballotAssign.itemAssign.Extent()
}

//
//
//
//
//
//
//
//
//
//
//
func (ballotAssign *BallotAssign) AppendBallot(ballot *Ballot) (appended bool, err error) {
	if ballotAssign == nil {
		panic("REDACTED")
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()

	return ballotAssign.appendBallot(ballot)
}

//
func (ballotAssign *BallotAssign) appendBallot(ballot *Ballot) (appended bool, err error) {
	if ballot == nil {
		return false, FaultBallotVoid
	}
	itemOrdinal := ballot.AssessorOrdinal
	itemLocation := ballot.AssessorLocation
	ledgerToken := ballot.LedgerUUID.Key()

	//
	if itemOrdinal < 0 {
		return false, fmt.Errorf("REDACTED", FaultBallotUnfitAssessorPosition)
	} else if len(itemLocation) == 0 {
		return false, fmt.Errorf("REDACTED", FaultBallotUnfitAssessorLocator)
	}

	//
	if (ballot.Altitude != ballotAssign.altitude) ||
		(ballot.Iteration != ballotAssign.iteration) ||
		(ballot.Kind != ballotAssign.notatedSignalKind) {
		return false, fmt.Errorf("REDACTED",
			ballotAssign.altitude, ballotAssign.iteration, ballotAssign.notatedSignalKind,
			ballot.Altitude, ballot.Iteration, ballot.Kind, FaultBallotUnforeseenPhase)
	}

	//
	searchLocation, val := ballotAssign.itemAssign.ObtainViaOrdinal(itemOrdinal)
	if val == nil {
		return false, fmt.Errorf(
			"REDACTED",
			itemOrdinal, ballotAssign.itemAssign.Extent(), FaultBallotUnfitAssessorPosition)
	}

	//
	if !bytes.Equal(itemLocation, searchLocation) {
		return false, fmt.Errorf(
			"REDACTED"+
				"REDACTED",
			itemLocation, searchLocation, itemOrdinal, FaultBallotUnfitAssessorLocator)
	}

	//
	if current, ok := ballotAssign.obtainBallot(itemOrdinal, ledgerToken, &ballot.LedgerUUID); ok {
		if bytes.Equal(current.Notation, ballot.Notation) {
			return false, nil //
		}
		return false, fmt.Errorf("REDACTED", current, ballot, FaultBallotUnCertainNotation)
	}

	//
	if ballotAssign.additionsActivated {
		if err := ballot.ValidateBallotAlsoAddition(ballotAssign.successionUUID, val.PublicToken); err != nil {
			return false, fmt.Errorf("REDACTED", ballotAssign.successionUUID, val.PublicToken, err)
		}
	} else {
		if err := ballot.Validate(ballotAssign.successionUUID, val.PublicToken); err != nil {
			return false, fmt.Errorf("REDACTED", ballotAssign.successionUUID, val.PublicToken, err)
		}
		if len(ballot.AdditionNotation) > 0 || len(ballot.Addition) > 0 {
			return false, fmt.Errorf("REDACTED",
				len(ballot.Addition),
				len(ballot.AdditionNotation),
			)
		}
	}

	//
	appended, discordant := ballotAssign.appendAttestedBallot(ballot, ledgerToken, val.BallotingPotency)
	if discordant != nil {
		return appended, FreshDiscordantBallotFailure(discordant, ballot)
	}
	if !appended {
		panic("REDACTED")
	}
	return appended, nil
}

//
func (ballotAssign *BallotAssign) obtainBallot(itemOrdinal int32, ledgerToken string, ledgerUUID *LedgerUUID) (ballot *Ballot, ok bool) {
	if current := ballotAssign.ballots[itemOrdinal]; current != nil && ledgerUUID.Matches(current.LedgerUUID) {
		return current, true
	}
	if current := ballotAssign.ballotsViaLedger[ledgerToken].obtainViaPosition(itemOrdinal); current != nil {
		return current, true
	}
	return nil, false
}

//
//
func (ballotAssign *BallotAssign) appendAttestedBallot(
	ballot *Ballot,
	ledgerToken string,
	ballotingPotency int64,
) (appended bool, discordant *Ballot) {
	itemOrdinal := ballot.AssessorOrdinal

	//
	if current := ballotAssign.ballots[itemOrdinal]; current != nil {
		if current.LedgerUUID.Matches(ballot.LedgerUUID) {
			panic("REDACTED")
		}
		discordant = current
		//
		if ballotAssign.major23 != nil && ballotAssign.major23.Matches(ballot.LedgerUUID) {
			ballotAssign.ballots[itemOrdinal] = ballot
			ballotAssign.ballotsDigitSeries.AssignOrdinal(int(itemOrdinal), true)
		}
		//
	} else {
		//
		ballotAssign.ballots[itemOrdinal] = ballot
		ballotAssign.ballotsDigitSeries.AssignOrdinal(int(itemOrdinal), true)
		ballotAssign.sum += ballotingPotency
	}

	ballotsViaLedger, ok := ballotAssign.ballotsViaLedger[ledgerToken]
	if ok {
		if discordant != nil && !ballotsViaLedger.nodeMajor23 {
			//
			return false, discordant
		}
		//
	} else {
		//
		if discordant != nil {
			//
			//
			return false, discordant
		}
		//
		//
		ballotsViaLedger = freshLedgerBallots(false, ballotAssign.itemAssign.Extent())
		ballotAssign.ballotsViaLedger[ledgerToken] = ballotsViaLedger
		//
	}

	//
	sourceTotal := ballotsViaLedger.sum
	assembly := ballotAssign.itemAssign.SumBallotingPotency()*2/3 + 1

	//
	ballotsViaLedger.appendAttestedBallot(ballot, ballotingPotency)

	//
	if sourceTotal < assembly && assembly <= ballotsViaLedger.sum {
		//
		if ballotAssign.major23 == nil {
			major23ledgerUUID := ballot.LedgerUUID
			ballotAssign.major23 = &major23ledgerUUID
			//
			for i, ballot := range ballotsViaLedger.ballots {
				if ballot != nil {
					ballotAssign.ballots[i] = ballot
				}
			}
		}
	}

	return true, discordant
}

//
//
//
//
//
func (ballotAssign *BallotAssign) AssignNodeMajor23(nodeUUID Nodeid, ledgerUUID LedgerUUID) error {
	if ballotAssign == nil {
		panic("REDACTED")
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()

	ledgerToken := ledgerUUID.Key()
	//
	if current, ok := ballotAssign.nodeMajor23s[nodeUUID]; ok {
		if current.Matches(ledgerUUID) {
			return nil //
		}
		return fmt.Errorf("REDACTED",
			nodeUUID, ledgerUUID, current)
	}
	ballotAssign.nodeMajor23s[nodeUUID] = ledgerUUID

	//
	ballotsViaLedger, ok := ballotAssign.ballotsViaLedger[ledgerToken]
	if ok {
		if ballotsViaLedger.nodeMajor23 {
			return nil //
		}
		ballotsViaLedger.nodeMajor23 = true
		//
	} else {
		ballotsViaLedger = freshLedgerBallots(true, ballotAssign.itemAssign.Extent())
		ballotAssign.ballotsViaLedger[ledgerToken] = ballotsViaLedger
		//
	}
	return nil
}

//
func (ballotAssign *BallotAssign) DigitSeries() *digits.DigitSeries {
	if ballotAssign == nil {
		return nil
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.ballotsDigitSeries.Duplicate()
}

func (ballotAssign *BallotAssign) DigitCollectionViaLedgerUUID(ledgerUUID LedgerUUID) *digits.DigitSeries {
	if ballotAssign == nil {
		return nil
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	ballotsViaLedger, ok := ballotAssign.ballotsViaLedger[ledgerUUID.Key()]
	if ok {
		return ballotsViaLedger.digitSeries.Duplicate()
	}
	return nil
}

//
//
func (ballotAssign *BallotAssign) ObtainViaOrdinal(itemOrdinal int32) *Ballot {
	if ballotAssign == nil {
		return nil
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.ballots[itemOrdinal]
}

//
func (ballotAssign *BallotAssign) Catalog() []Ballot {
	if ballotAssign == nil || ballotAssign.ballots == nil {
		return nil
	}
	ballots := make([]Ballot, 0, len(ballotAssign.ballots))
	for i := range ballotAssign.ballots {
		if ballotAssign.ballots[i] != nil {
			ballots = append(ballots, *ballotAssign.ballots[i])
		}
	}
	return ballots
}

func (ballotAssign *BallotAssign) ObtainViaLocation(location []byte) *Ballot {
	if ballotAssign == nil {
		return nil
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	itemOrdinal, val := ballotAssign.itemAssign.ObtainViaLocationAlterable(location)
	if val == nil {
		panic("REDACTED")
	}
	return ballotAssign.ballots[itemOrdinal]
}

func (ballotAssign *BallotAssign) OwnsCoupleTrinityPreponderance() bool {
	if ballotAssign == nil {
		return false
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.major23 != nil
}

//
func (ballotAssign *BallotAssign) EqualsEndorse() bool {
	if ballotAssign == nil {
		return false
	}
	if ballotAssign.notatedSignalKind != commitchema.PreendorseKind {
		return false
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.major23 != nil
}

func (ballotAssign *BallotAssign) OwnsCoupleTrinitySome() bool {
	if ballotAssign == nil {
		return false
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.sum > ballotAssign.itemAssign.SumBallotingPotency()*2/3
}

func (ballotAssign *BallotAssign) OwnsEvery() bool {
	if ballotAssign == nil {
		return false
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.sum == ballotAssign.itemAssign.SumBallotingPotency()
}

//
//
func (ballotAssign *BallotAssign) CoupleTrinityPreponderance() (ledgerUUID LedgerUUID, ok bool) {
	if ballotAssign == nil {
		return LedgerUUID{}, false
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	if ballotAssign.major23 != nil {
		return *ballotAssign.major23, true
	}
	return LedgerUUID{}, false
}

//
//

const voidBallotAssignText = "REDACTED"

//
//
//
func (ballotAssign *BallotAssign) Text() string {
	if ballotAssign == nil {
		return voidBallotAssignText
	}
	return ballotAssign.TextFormatted("REDACTED")
}

//
//
//
//
//
//
//
//
func (ballotAssign *BallotAssign) TextFormatted(format string) string {
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()

	ballotTexts := make([]string, len(ballotAssign.ballots))
	for i, ballot := range ballotAssign.ballots {
		if ballot == nil {
			ballotTexts[i] = voidBallotTxt
		} else {
			ballotTexts[i] = ballot.Text()
		}
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		format, ballotAssign.altitude, ballotAssign.iteration, ballotAssign.notatedSignalKind,
		format, strings.Join(ballotTexts, "REDACTED"+format+"REDACTED"),
		format, ballotAssign.ballotsDigitSeries,
		format, ballotAssign.nodeMajor23s,
		format)
}

//
//
func (ballotAssign *BallotAssign) SerializeJSN() ([]byte, error) {
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return strongmindjson.Serialize(BallotAssignJSN{
		ballotAssign.ballotTexts(),
		ballotAssign.digitSeriesText(),
		ballotAssign.nodeMajor23s,
	})
}

//
//
//
type BallotAssignJSN struct {
	Ballots         []string          `json:"ballots"`
	BallotsDigitSeries string            `json:"ballots_digit_series"`
	NodeMajor23s    map[Nodeid]LedgerUUID `json:"node_major_twenty3s"`
}

//
//
//
func (ballotAssign *BallotAssign) DigitSeriesText() string {
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.digitSeriesText()
}

func (ballotAssign *BallotAssign) digitSeriesText() string {
	byteANText := ballotAssign.ballotsDigitSeries.Text()
	balloted, sum, divisionBalloted := ballotAssign.totalSumDivision()
	return fmt.Sprintf("REDACTED", byteANText, balloted, sum, divisionBalloted)
}

//
func (ballotAssign *BallotAssign) BallotTexts() []string {
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	return ballotAssign.ballotTexts()
}

func (ballotAssign *BallotAssign) ballotTexts() []string {
	ballotTexts := make([]string, len(ballotAssign.ballots))
	for i, ballot := range ballotAssign.ballots {
		if ballot == nil {
			ballotTexts[i] = voidBallotTxt
		} else {
			ballotTexts[i] = ballot.Text()
		}
	}
	return ballotTexts
}

//
//
//
//
//
//
//
//
//
func (ballotAssign *BallotAssign) TextBrief() string {
	if ballotAssign == nil {
		return voidBallotAssignText
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	_, _, division := ballotAssign.totalSumDivision()
	return fmt.Sprintf("REDACTED",
		ballotAssign.altitude, ballotAssign.iteration, ballotAssign.notatedSignalKind, ballotAssign.major23, division, ballotAssign.ballotsDigitSeries, ballotAssign.nodeMajor23s)
}

//
//
func (ballotAssign *BallotAssign) RecordText() string {
	if ballotAssign == nil {
		return voidBallotAssignText
	}
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()
	balloted, sum, division := ballotAssign.totalSumDivision()

	return fmt.Sprintf("REDACTED", balloted, sum, division)
}

//
func (ballotAssign *BallotAssign) totalSumDivision() (int64, int64, float64) {
	balloted, sum := ballotAssign.sum, ballotAssign.itemAssign.SumBallotingPotency()
	divisionBalloted := float64(balloted) / float64(sum)
	return balloted, sum, divisionBalloted
}

//
//

//
//
//
//
//
func (ballotAssign *BallotAssign) CreateExpandedEndorse(ap IfaceParameters) *ExpandedEndorse {
	ballotAssign.mtx.Lock()
	defer ballotAssign.mtx.Unlock()

	if ballotAssign.notatedSignalKind != commitchema.PreendorseKind {
		panic("REDACTED")
	}

	//
	if ballotAssign.major23 == nil {
		panic("REDACTED")
	}

	//
	signatures := make([]ExpandedEndorseSignature, len(ballotAssign.ballots))
	for i, v := range ballotAssign.ballots {
		sig := v.ExpandedEndorseSignature()
		//
		if sig.LedgerUUIDMarker == LedgerUUIDMarkerEndorse && !v.LedgerUUID.Matches(*ballotAssign.major23) {
			sig = FreshExpandedEndorseSignatureMissing()
		}

		signatures[i] = sig
	}

	ec := &ExpandedEndorse{
		Altitude:             ballotAssign.ObtainAltitude(),
		Iteration:              ballotAssign.ObtainIteration(),
		LedgerUUID:            *ballotAssign.major23,
		ExpandedNotations: signatures,
	}
	if err := ec.AssureAdditions(ap.BallotAdditionsActivated(ec.Altitude)); err != nil {
		panic(fmt.Errorf("REDACTED",
			ec.Altitude, err))
	}
	return ec
}

//

/**
k
.
)
)
*/
type ledgerBallots struct {
	nodeMajor23 bool           //
	digitSeries  *digits.DigitSeries //
	ballots     []*Ballot        //
	sum       int64          //
}

func freshLedgerBallots(nodeMajor23 bool, countAssessors int) *ledgerBallots {
	return &ledgerBallots{
		nodeMajor23: nodeMajor23,
		digitSeries:  digits.FreshDigitCollection(countAssessors),
		ballots:     make([]*Ballot, countAssessors),
		sum:       0,
	}
}

func (vs *ledgerBallots) appendAttestedBallot(ballot *Ballot, ballotingPotency int64) {
	itemOrdinal := ballot.AssessorOrdinal
	if current := vs.ballots[itemOrdinal]; current == nil {
		vs.digitSeries.AssignOrdinal(int(itemOrdinal), true)
		vs.ballots[itemOrdinal] = ballot
		vs.sum += ballotingPotency
	}
}

func (vs *ledgerBallots) obtainViaPosition(ordinal int32) *Ballot {
	if vs == nil {
		return nil
	}
	return vs.ballots[ordinal]
}

//

//
type BallotAssignFetcher interface {
	ObtainAltitude() int64
	ObtainIteration() int32
	Kind() byte
	Extent() int
	DigitSeries() *digits.DigitSeries
	ObtainViaOrdinal(int32) *Ballot
	EqualsEndorse() bool
}
