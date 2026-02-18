package kinds

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/valkyrieworks/utils/bits"
	cometjson "github.com/valkyrieworks/utils/json"
	engineconnect "github.com/valkyrieworks/utils/align"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
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
type P2pid string

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
type BallotCollection struct {
	ledgerUID           string
	level            int64
	epoch             int32
	attestedMessageKind     engineproto.AttestedMessageKind
	valueCollection            *RatifierAssign
	pluginsActivated bool

	mtx           engineconnect.Lock
	ballotsBitList *bits.BitList
	ballots         []*Ballot                //
	sum           int64                  //
	maj23         *LedgerUID               //
	ballotsByLedger  map[string]*ledgerBallots //
	nodeMaj23s    map[P2pid]LedgerUID      //
}

//
//
func NewBallotCollection(ledgerUID string, level int64, epoch int32,
	attestedMessageKind engineproto.AttestedMessageKind, valueCollection *RatifierAssign,
) *BallotCollection {
	if level == 0 {
		panic("REDACTED")
	}
	return &BallotCollection{
		ledgerUID:       ledgerUID,
		level:        level,
		epoch:         epoch,
		attestedMessageKind: attestedMessageKind,
		valueCollection:        valueCollection,
		ballotsBitList: bits.NewBitList(valueCollection.Volume()),
		ballots:         make([]*Ballot, valueCollection.Volume()),
		sum:           0,
		maj23:         nil,
		ballotsByLedger:  make(map[string]*ledgerBallots, valueCollection.Volume()),
		nodeMaj23s:    make(map[P2pid]LedgerUID),
	}
}

//
//
//
func NewExpandedBallotCollection(ledgerUID string, level int64, epoch int32,
	attestedMessageKind engineproto.AttestedMessageKind, valueCollection *RatifierAssign,
) *BallotCollection {
	vs := NewBallotCollection(ledgerUID, level, epoch, attestedMessageKind, valueCollection)
	vs.pluginsActivated = true
	return vs
}

func (ballotCollection *BallotCollection) LedgerUID() string {
	return ballotCollection.ledgerUID
}

//
func (ballotCollection *BallotCollection) FetchLevel() int64 {
	if ballotCollection == nil {
		return 0
	}
	return ballotCollection.level
}

//
func (ballotCollection *BallotCollection) FetchDuration() int32 {
	if ballotCollection == nil {
		return -1
	}
	return ballotCollection.epoch
}

//
func (ballotCollection *BallotCollection) Kind() byte {
	if ballotCollection == nil {
		return 0x00
	}
	return byte(ballotCollection.attestedMessageKind)
}

//
func (ballotCollection *BallotCollection) Volume() int {
	if ballotCollection == nil {
		return 0
	}
	return ballotCollection.valueCollection.Volume()
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
func (ballotCollection *BallotCollection) AppendBallot(ballot *Ballot) (appended bool, err error) {
	if ballotCollection == nil {
		panic("REDACTED")
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()

	return ballotCollection.appendBallot(ballot)
}

//
func (ballotCollection *BallotCollection) appendBallot(ballot *Ballot) (appended bool, err error) {
	if ballot == nil {
		return false, ErrBallotNull
	}
	valueOrdinal := ballot.RatifierOrdinal
	valueAddress := ballot.RatifierLocation
	ledgerKey := ballot.LedgerUID.Key()

	//
	if valueOrdinal < 0 {
		return false, fmt.Errorf("REDACTED", ErrBallotCorruptRatifierOrdinal)
	} else if len(valueAddress) == 0 {
		return false, fmt.Errorf("REDACTED", ErrBallotCorruptRatifierLocation)
	}

	//
	if (ballot.Level != ballotCollection.level) ||
		(ballot.Cycle != ballotCollection.epoch) ||
		(ballot.Kind != ballotCollection.attestedMessageKind) {
		return false, fmt.Errorf("REDACTED",
			ballotCollection.level, ballotCollection.epoch, ballotCollection.attestedMessageKind,
			ballot.Level, ballot.Cycle, ballot.Kind, ErrBallotUnforeseenPhase)
	}

	//
	searchAddress, val := ballotCollection.valueCollection.FetchByOrdinal(valueOrdinal)
	if val == nil {
		return false, fmt.Errorf(
			"REDACTED",
			valueOrdinal, ballotCollection.valueCollection.Volume(), ErrBallotCorruptRatifierOrdinal)
	}

	//
	if !bytes.Equal(valueAddress, searchAddress) {
		return false, fmt.Errorf(
			"REDACTED"+
				"REDACTED",
			valueAddress, searchAddress, valueOrdinal, ErrBallotCorruptRatifierLocation)
	}

	//
	if current, ok := ballotCollection.fetchBallot(valueOrdinal, ledgerKey, &ballot.LedgerUID); ok {
		if bytes.Equal(current.Autograph, ballot.Autograph) {
			return false, nil //
		}
		return false, fmt.Errorf("REDACTED", current, ballot, ErrBallotNotCertainAutograph)
	}

	//
	if ballotCollection.pluginsActivated {
		if err := ballot.ValidateBallotAndAddition(ballotCollection.ledgerUID, val.PublicKey); err != nil {
			return false, fmt.Errorf("REDACTED", ballotCollection.ledgerUID, val.PublicKey, err)
		}
	} else {
		if err := ballot.Validate(ballotCollection.ledgerUID, val.PublicKey); err != nil {
			return false, fmt.Errorf("REDACTED", ballotCollection.ledgerUID, val.PublicKey, err)
		}
		if len(ballot.AdditionAutograph) > 0 || len(ballot.Addition) > 0 {
			return false, fmt.Errorf("REDACTED",
				len(ballot.Addition),
				len(ballot.AdditionAutograph),
			)
		}
	}

	//
	appended, clashing := ballotCollection.appendCertifiedBallot(ballot, ledgerKey, val.PollingEnergy)
	if clashing != nil {
		return appended, NewClashingBallotFault(clashing, ballot)
	}
	if !appended {
		panic("REDACTED")
	}
	return appended, nil
}

//
func (ballotCollection *BallotCollection) fetchBallot(valueOrdinal int32, ledgerKey string, ledgerUID *LedgerUID) (ballot *Ballot, ok bool) {
	if current := ballotCollection.ballots[valueOrdinal]; current != nil && ledgerUID.Matches(current.LedgerUID) {
		return current, true
	}
	if current := ballotCollection.ballotsByLedger[ledgerKey].fetchByOrdinal(valueOrdinal); current != nil {
		return current, true
	}
	return nil, false
}

//
//
func (ballotCollection *BallotCollection) appendCertifiedBallot(
	ballot *Ballot,
	ledgerKey string,
	pollingEnergy int64,
) (appended bool, clashing *Ballot) {
	valueOrdinal := ballot.RatifierOrdinal

	//
	if current := ballotCollection.ballots[valueOrdinal]; current != nil {
		if current.LedgerUID.Matches(ballot.LedgerUID) {
			panic("REDACTED")
		}
		clashing = current
		//
		if ballotCollection.maj23 != nil && ballotCollection.maj23.Matches(ballot.LedgerUID) {
			ballotCollection.ballots[valueOrdinal] = ballot
			ballotCollection.ballotsBitList.AssignOrdinal(int(valueOrdinal), true)
		}
		//
	} else {
		//
		ballotCollection.ballots[valueOrdinal] = ballot
		ballotCollection.ballotsBitList.AssignOrdinal(int(valueOrdinal), true)
		ballotCollection.sum += pollingEnergy
	}

	ballotsByLedger, ok := ballotCollection.ballotsByLedger[ledgerKey]
	if ok {
		if clashing != nil && !ballotsByLedger.nodeMaj23 {
			//
			return false, clashing
		}
		//
	} else {
		//
		if clashing != nil {
			//
			//
			return false, clashing
		}
		//
		//
		ballotsByLedger = newLedgerBallots(false, ballotCollection.valueCollection.Volume())
		ballotCollection.ballotsByLedger[ledgerKey] = ballotsByLedger
		//
	}

	//
	origTotal := ballotsByLedger.sum
	assembly := ballotCollection.valueCollection.SumPollingEnergy()*2/3 + 1

	//
	ballotsByLedger.appendCertifiedBallot(ballot, pollingEnergy)

	//
	if origTotal < assembly && assembly <= ballotsByLedger.sum {
		//
		if ballotCollection.maj23 == nil {
			maj23ledgerUID := ballot.LedgerUID
			ballotCollection.maj23 = &maj23ledgerUID
			//
			for i, ballot := range ballotsByLedger.ballots {
				if ballot != nil {
					ballotCollection.ballots[i] = ballot
				}
			}
		}
	}

	return true, clashing
}

//
//
//
//
//
func (ballotCollection *BallotCollection) AssignNodeMaj23(nodeUID P2pid, ledgerUID LedgerUID) error {
	if ballotCollection == nil {
		panic("REDACTED")
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()

	ledgerKey := ledgerUID.Key()
	//
	if current, ok := ballotCollection.nodeMaj23s[nodeUID]; ok {
		if current.Matches(ledgerUID) {
			return nil //
		}
		return fmt.Errorf("REDACTED",
			nodeUID, ledgerUID, current)
	}
	ballotCollection.nodeMaj23s[nodeUID] = ledgerUID

	//
	ballotsByLedger, ok := ballotCollection.ballotsByLedger[ledgerKey]
	if ok {
		if ballotsByLedger.nodeMaj23 {
			return nil //
		}
		ballotsByLedger.nodeMaj23 = true
		//
	} else {
		ballotsByLedger = newLedgerBallots(true, ballotCollection.valueCollection.Volume())
		ballotCollection.ballotsByLedger[ledgerKey] = ballotsByLedger
		//
	}
	return nil
}

//
func (ballotCollection *BallotCollection) BitList() *bits.BitList {
	if ballotCollection == nil {
		return nil
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.ballotsBitList.Clone()
}

func (ballotCollection *BallotCollection) BitListByLedgerUID(ledgerUID LedgerUID) *bits.BitList {
	if ballotCollection == nil {
		return nil
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	ballotsByLedger, ok := ballotCollection.ballotsByLedger[ledgerUID.Key()]
	if ok {
		return ballotsByLedger.bitList.Clone()
	}
	return nil
}

//
//
func (ballotCollection *BallotCollection) FetchByOrdinal(valueOrdinal int32) *Ballot {
	if ballotCollection == nil {
		return nil
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.ballots[valueOrdinal]
}

//
func (ballotCollection *BallotCollection) Catalog() []Ballot {
	if ballotCollection == nil || ballotCollection.ballots == nil {
		return nil
	}
	ballots := make([]Ballot, 0, len(ballotCollection.ballots))
	for i := range ballotCollection.ballots {
		if ballotCollection.ballots[i] != nil {
			ballots = append(ballots, *ballotCollection.ballots[i])
		}
	}
	return ballots
}

func (ballotCollection *BallotCollection) FetchByLocation(location []byte) *Ballot {
	if ballotCollection == nil {
		return nil
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	valueOrdinal, val := ballotCollection.valueCollection.FetchByLocationMut(location)
	if val == nil {
		panic("REDACTED")
	}
	return ballotCollection.ballots[valueOrdinal]
}

func (ballotCollection *BallotCollection) HasDualThirdsBulk() bool {
	if ballotCollection == nil {
		return false
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.maj23 != nil
}

//
func (ballotCollection *BallotCollection) IsEndorse() bool {
	if ballotCollection == nil {
		return false
	}
	if ballotCollection.attestedMessageKind != engineproto.PreendorseKind {
		return false
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.maj23 != nil
}

func (ballotCollection *BallotCollection) HasDualThirdsAny() bool {
	if ballotCollection == nil {
		return false
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.sum > ballotCollection.valueCollection.SumPollingEnergy()*2/3
}

func (ballotCollection *BallotCollection) HasAll() bool {
	if ballotCollection == nil {
		return false
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.sum == ballotCollection.valueCollection.SumPollingEnergy()
}

//
//
func (ballotCollection *BallotCollection) DualThirdsBulk() (ledgerUID LedgerUID, ok bool) {
	if ballotCollection == nil {
		return LedgerUID{}, false
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	if ballotCollection.maj23 != nil {
		return *ballotCollection.maj23, true
	}
	return LedgerUID{}, false
}

//
//

const nullBallotCollectionString = "REDACTED"

//
//
//
func (ballotCollection *BallotCollection) String() string {
	if ballotCollection == nil {
		return nullBallotCollectionString
	}
	return ballotCollection.StringIndented("REDACTED")
}

//
//
//
//
//
//
//
//
func (ballotCollection *BallotCollection) StringIndented(indent string) string {
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()

	ballotStrings := make([]string, len(ballotCollection.ballots))
	for i, ballot := range ballotCollection.ballots {
		if ballot == nil {
			ballotStrings[i] = nullBallotStr
		} else {
			ballotStrings[i] = ballot.String()
		}
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		indent, ballotCollection.level, ballotCollection.epoch, ballotCollection.attestedMessageKind,
		indent, strings.Join(ballotStrings, "REDACTED"+indent+"REDACTED"),
		indent, ballotCollection.ballotsBitList,
		indent, ballotCollection.nodeMaj23s,
		indent)
}

//
//
func (ballotCollection *BallotCollection) SerializeJSON() ([]byte, error) {
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return cometjson.Serialize(BallotCollectionJSON{
		ballotCollection.ballotStrings(),
		ballotCollection.bitListString(),
		ballotCollection.nodeMaj23s,
	})
}

//
//
//
type BallotCollectionJSON struct {
	Ballots         []string          `json:"ballots"`
	BallotsBitList string            `json:"ballots_bit_list"`
	NodeMaj23s    map[P2pid]LedgerUID `json:"node_major_23s"`
}

//
//
//
func (ballotCollection *BallotCollection) BitListString() string {
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.bitListString()
}

func (ballotCollection *BallotCollection) bitListString() string {
	byteAString := ballotCollection.ballotsBitList.String()
	polled, sum, slicePolled := ballotCollection.totalSumSlice()
	return fmt.Sprintf("REDACTED", byteAString, polled, sum, slicePolled)
}

//
func (ballotCollection *BallotCollection) BallotStrings() []string {
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	return ballotCollection.ballotStrings()
}

func (ballotCollection *BallotCollection) ballotStrings() []string {
	ballotStrings := make([]string, len(ballotCollection.ballots))
	for i, ballot := range ballotCollection.ballots {
		if ballot == nil {
			ballotStrings[i] = nullBallotStr
		} else {
			ballotStrings[i] = ballot.String()
		}
	}
	return ballotStrings
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
func (ballotCollection *BallotCollection) StringBrief() string {
	if ballotCollection == nil {
		return nullBallotCollectionString
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	_, _, slice := ballotCollection.totalSumSlice()
	return fmt.Sprintf("REDACTED",
		ballotCollection.level, ballotCollection.epoch, ballotCollection.attestedMessageKind, ballotCollection.maj23, slice, ballotCollection.ballotsBitList, ballotCollection.nodeMaj23s)
}

//
//
func (ballotCollection *BallotCollection) TraceString() string {
	if ballotCollection == nil {
		return nullBallotCollectionString
	}
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()
	polled, sum, slice := ballotCollection.totalSumSlice()

	return fmt.Sprintf("REDACTED", polled, sum, slice)
}

//
func (ballotCollection *BallotCollection) totalSumSlice() (int64, int64, float64) {
	polled, sum := ballotCollection.sum, ballotCollection.valueCollection.SumPollingEnergy()
	slicePolled := float64(polled) / float64(sum)
	return polled, sum, slicePolled
}

//
//

//
//
//
//
//
func (ballotCollection *BallotCollection) CreateExpandedEndorse(ap IfaceOptions) *ExpandedEndorse {
	ballotCollection.mtx.Lock()
	defer ballotCollection.mtx.Unlock()

	if ballotCollection.attestedMessageKind != engineproto.PreendorseKind {
		panic("REDACTED")
	}

	//
	if ballotCollection.maj23 == nil {
		panic("REDACTED")
	}

	//
	autographs := make([]ExpandedEndorseSignature, len(ballotCollection.ballots))
	for i, v := range ballotCollection.ballots {
		sig := v.ExpandedEndorseSignature()
		//
		if sig.LedgerUIDMark == LedgerUIDMarkEndorse && !v.LedgerUID.Matches(*ballotCollection.maj23) {
			sig = NewExpandedEndorseSignatureMissing()
		}

		autographs[i] = sig
	}

	ec := &ExpandedEndorse{
		Level:             ballotCollection.FetchLevel(),
		Cycle:              ballotCollection.FetchDuration(),
		LedgerUID:            *ballotCollection.maj23,
		ExpandedEndorsements: autographs,
	}
	if err := ec.AssurePlugins(ap.BallotPluginsActivated(ec.Level)); err != nil {
		panic(fmt.Errorf("REDACTED",
			ec.Level, err))
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
	nodeMaj23 bool           //
	bitList  *bits.BitList //
	ballots     []*Ballot        //
	sum       int64          //
}

func newLedgerBallots(nodeMaj23 bool, countRatifiers int) *ledgerBallots {
	return &ledgerBallots{
		nodeMaj23: nodeMaj23,
		bitList:  bits.NewBitList(countRatifiers),
		ballots:     make([]*Ballot, countRatifiers),
		sum:       0,
	}
}

func (vs *ledgerBallots) appendCertifiedBallot(ballot *Ballot, pollingEnergy int64) {
	valueOrdinal := ballot.RatifierOrdinal
	if current := vs.ballots[valueOrdinal]; current == nil {
		vs.bitList.AssignOrdinal(int(valueOrdinal), true)
		vs.ballots[valueOrdinal] = ballot
		vs.sum += pollingEnergy
	}
}

func (vs *ledgerBallots) fetchByOrdinal(ordinal int32) *Ballot {
	if vs == nil {
		return nil
	}
	return vs.ballots[ordinal]
}

//

//
type BallotAssignScanner interface {
	FetchLevel() int64
	FetchDuration() int32
	Kind() byte
	Volume() int
	BitList() *bits.BitList
	FetchByOrdinal(int32) *Ballot
	IsEndorse() bool
}
