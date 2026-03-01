package kinds

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

const (
	//
	//
	//
	//
	//
	//
	//
	MaximumSumBallotingPotency = int64(math.MaxInt64) / 8

	//
	//
	//
	UrgencyFrameworkExtentElement = 2
)

//
//
var FaultSumBallotingPotencyOverrun = fmt.Errorf("REDACTED",
	MaximumSumBallotingPotency)

//
var FaultNominatorNegationInsideValues = errors.New("REDACTED")

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
//
//
type AssessorAssign struct {
	//
	Assessors []*Assessor `json:"assessors"`
	Nominator   *Assessor   `json:"nominator"`

	//
	sumBallotingPotency int64
	//
	everyTokensPossessIdenticalKind bool
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
func FreshAssessorAssign(validz []*Assessor) *AssessorAssign {
	values := &AssessorAssign{
		everyTokensPossessIdenticalKind: true,
	}
	err := values.reviseUsingModifyAssign(validz, false)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if len(validz) > 0 {
		values.AdvanceNominatorUrgency(1)
	}
	return values
}

func (values *AssessorAssign) CertifyFundamental() error {
	if values.EqualsVoidEitherBlank() {
		return errors.New("REDACTED")
	}

	for idx, val := range values.Assessors {
		if err := val.CertifyFundamental(); err != nil {
			return fmt.Errorf("REDACTED", idx, err)
		}
	}

	if err := values.Nominator.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	for _, val := range values.Assessors {
		if bytes.Equal(val.Location, values.Nominator.Location) {
			return nil
		}
	}

	return FaultNominatorNegationInsideValues
}

//
func (values *AssessorAssign) EqualsVoidEitherBlank() bool {
	return values == nil || len(values.Assessors) == 0
}

//
//
func (values *AssessorAssign) DuplicateAdvanceNominatorUrgency(multiples int32) *AssessorAssign {
	cp := values.Duplicate()
	cp.AdvanceNominatorUrgency(multiples)
	return cp
}

//
//
//
func (values *AssessorAssign) AdvanceNominatorUrgency(multiples int32) {
	if values.EqualsVoidEitherBlank() {
		panic("REDACTED")
	}
	if multiples <= 0 {
		panic("REDACTED")
	}

	//
	//
	//
	varianceMaximum := UrgencyFrameworkExtentElement * values.SumBallotingPotency()
	values.RecalibrateUrgencies(varianceMaximum)
	values.relocateViaMedianNominatorUrgency()

	var nominator *Assessor
	//
	for i := int32(0); i < multiples; i++ {
		nominator = values.advanceNominatorUrgency()
	}

	values.Nominator = nominator
}

//
//
//
func (values *AssessorAssign) RecalibrateUrgencies(varianceMaximum int64) {
	if values.EqualsVoidEitherBlank() {
		panic("REDACTED")
	}
	//
	//
	//
	if varianceMaximum <= 0 {
		return
	}

	//
	//
	//
	variance := calculateMaximumMinimumUrgencyVariance(values)
	proportion := (variance + varianceMaximum - 1) / varianceMaximum
	if variance > varianceMaximum {
		for _, val := range values.Assessors {
			val.NominatorUrgency /= proportion
		}
	}
}

func (values *AssessorAssign) advanceNominatorUrgency() *Assessor {
	for _, val := range values.Assessors {
		//
		freshUrgency := secureAppendRestrict(val.NominatorUrgency, val.BallotingPotency)
		val.NominatorUrgency = freshUrgency
	}
	//
	utmost := values.obtainItemUsingUtmostUrgency()
	//
	utmost.NominatorUrgency = secureUnderRestrict(utmost.NominatorUrgency, values.SumBallotingPotency())

	return utmost
}

//
func (values *AssessorAssign) calculateMedianNominatorUrgency() int64 {
	n := int64(len(values.Assessors))
	sum := big.NewInt(0)
	for _, val := range values.Assessors {
		sum.Add(sum, big.NewInt(val.NominatorUrgency))
	}
	avg := sum.Div(sum, big.NewInt(n))
	if avg.IsInt64() {
		return avg.Int64()
	}

	//
	panic(fmt.Sprintf("REDACTED", avg))
}

//
//
func calculateMaximumMinimumUrgencyVariance(values *AssessorAssign) int64 {
	if values.EqualsVoidEitherBlank() {
		panic("REDACTED")
	}
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)
	for _, v := range values.Assessors {
		if v.NominatorUrgency < min {
			min = v.NominatorUrgency
		}
		if v.NominatorUrgency > max {
			max = v.NominatorUrgency
		}
	}
	variance := max - min
	if variance < 0 {
		return -1 * variance
	}
	return variance
}

func (values *AssessorAssign) obtainItemUsingUtmostUrgency() *Assessor {
	var res *Assessor
	for _, val := range values.Assessors {
		res = res.ContrastNominatorUrgency(val)
	}
	return res
}

func (values *AssessorAssign) relocateViaMedianNominatorUrgency() {
	if values.EqualsVoidEitherBlank() {
		panic("REDACTED")
	}
	medianNominatorUrgency := values.calculateMedianNominatorUrgency()
	for _, val := range values.Assessors {
		val.NominatorUrgency = secureUnderRestrict(val.NominatorUrgency, medianNominatorUrgency)
	}
}

//
func assessorCatalogDuplicate(valuesCatalog []*Assessor) []*Assessor {
	if valuesCatalog == nil {
		return nil
	}
	valuesDuplicate := make([]*Assessor, len(valuesCatalog))
	for i, val := range valuesCatalog {
		valuesDuplicate[i] = val.Duplicate()
	}
	return valuesDuplicate
}

//
func (values *AssessorAssign) Duplicate() *AssessorAssign {
	return &AssessorAssign{
		Assessors:          assessorCatalogDuplicate(values.Assessors),
		Nominator:            values.Nominator,
		sumBallotingPotency:    values.sumBallotingPotency,
		everyTokensPossessIdenticalKind: values.everyTokensPossessIdenticalKind,
	}
}

//
//
func (values *AssessorAssign) OwnsLocation(location []byte) bool {
	for _, val := range values.Assessors {
		if bytes.Equal(val.Location, location) {
			return true
		}
	}
	return false
}

//
//
func (values *AssessorAssign) ObtainViaLocation(location []byte) (ordinal int32, val *Assessor) {
	i, val := values.ObtainViaLocationAlterable(location)
	if i == -1 {
		return -1, nil
	}
	return i, val.Duplicate()
}

//
//
//
//
func (values *AssessorAssign) ObtainViaLocationAlterable(location []byte) (ordinal int32, val *Assessor) {
	for idx, val := range values.Assessors {
		if bytes.Equal(val.Location, location) {
			return int32(idx), val
		}
	}
	return -1, nil
}

//
//
//
//
func (values *AssessorAssign) ObtainViaOrdinal(ordinal int32) (location []byte, val *Assessor) {
	if ordinal < 0 || int(ordinal) >= len(values.Assessors) {
		return nil, nil
	}
	val = values.Assessors[ordinal]
	return val.Location, val.Duplicate()
}

//
func (values *AssessorAssign) Extent() int {
	return len(values.Assessors)
}

//
//
func (values *AssessorAssign) reviseSumBallotingPotency() error {
	sum := int64(0)
	for _, val := range values.Assessors {
		//
		sum = secureAppendRestrict(sum, val.BallotingPotency)
		if sum > MaximumSumBallotingPotency {
			return fmt.Errorf("REDACTED", sum, MaximumSumBallotingPotency)
		}
	}

	values.sumBallotingPotency = sum
	return nil
}

//
//
func (values *AssessorAssign) SumBallotingPotencySecure() (int64, error) {
	if values.sumBallotingPotency == 0 {
		if err := values.reviseSumBallotingPotency(); err != nil {
			return 0, err
		}
	}
	return values.sumBallotingPotency, nil
}

//
//
func (values *AssessorAssign) SumBallotingPotency() int64 {
	if values.sumBallotingPotency == 0 {
		if err := values.reviseSumBallotingPotency(); err != nil {
			panic(err)
		}
	}
	return values.sumBallotingPotency
}

//
//
func (values *AssessorAssign) ObtainNominator() (nominator *Assessor) {
	if len(values.Assessors) == 0 {
		return nil
	}
	if values.Nominator == nil {
		values.Nominator = values.locateNominator()
	}
	return values.Nominator.Duplicate()
}

func (values *AssessorAssign) locateNominator() *Assessor {
	var nominator *Assessor
	for _, val := range values.Assessors {
		if nominator == nil || !bytes.Equal(val.Location, nominator.Location) {
			nominator = nominator.ContrastNominatorUrgency(val)
		}
	}
	return nominator
}

//
//
//
//
func (values *AssessorAssign) Digest() []byte {
	bzs := make([][]byte, len(values.Assessors))
	for i, val := range values.Assessors {
		bzs[i] = val.Octets()
	}
	return hashmap.DigestOriginatingOctetSegments(bzs)
}

//
//
//
func (values *AssessorAssign) NominatorUrgencyDigest() []byte {
	if len(values.Assessors) == 0 {
		return nil
	}

	buf := make([]byte, binary.MaxVarintLen64*len(values.Assessors))
	displacement := 0
	for _, val := range values.Assessors {
		n := binary.PutVarint(buf[displacement:], val.NominatorUrgency)
		displacement += n
	}
	return tenderminthash.Sum(buf[:displacement])
}

//
func (values *AssessorAssign) Traverse(fn func(ordinal int, val *Assessor) bool) {
	for i, val := range values.Assessors {
		halt := fn(i, val.Duplicate())
		if halt {
			break
		}
	}
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
func handleModifications(sourceModifications []*Assessor) (revisions, deletions []*Assessor, err error) {
	//
	modifications := assessorCatalogDuplicate(sourceModifications)
	sort.Sort(AssessorsViaLocator(modifications))

	deletions = make([]*Assessor, 0, len(modifications))
	revisions = make([]*Assessor, 0, len(modifications))
	var previousLocation Location

	//
	for _, itemRevise := range modifications {
		if bytes.Equal(itemRevise.Location, previousLocation) {
			err = fmt.Errorf("REDACTED", itemRevise, modifications)
			return nil, nil, err
		}

		switch {
		case itemRevise.BallotingPotency < 0:
			err = fmt.Errorf("REDACTED", itemRevise.BallotingPotency)
			return nil, nil, err
		case itemRevise.BallotingPotency > MaximumSumBallotingPotency:
			err = fmt.Errorf("REDACTED",
				MaximumSumBallotingPotency, itemRevise.BallotingPotency)
			return nil, nil, err
		case itemRevise.BallotingPotency == 0:
			deletions = append(deletions, itemRevise)
		default:
			revisions = append(revisions, itemRevise)
		}

		previousLocation = itemRevise.Location
	}

	return revisions, deletions, err
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
//
//
//
//
//
//
//
func validateRevisions(
	revisions []*Assessor,
	values *AssessorAssign,
	discardedPotency int64,
) (tvoterSubsequentRevisionsPriorDeletions int64, err error) {
	variation := func(revise *Assessor, values *AssessorAssign) int64 {
		_, val := values.ObtainViaLocationAlterable(revise.Location)
		if val != nil {
			return revise.BallotingPotency - val.BallotingPotency
		}
		return revise.BallotingPotency
	}

	revisionsDuplicate := assessorCatalogDuplicate(revisions)
	sort.Slice(revisionsDuplicate, func(i, j int) bool {
		return variation(revisionsDuplicate[i], values) < variation(revisionsDuplicate[j], values)
	})

	tvoterSubsequentDeletions := values.SumBallotingPotency() - discardedPotency
	for _, upd := range revisionsDuplicate {
		tvoterSubsequentDeletions += variation(upd, values)
		if tvoterSubsequentDeletions > MaximumSumBallotingPotency {
			return 0, FaultSumBallotingPotencyOverrun
		}
	}
	return tvoterSubsequentDeletions + discardedPotency, nil
}

func countFreshAssessors(revisions []*Assessor, values *AssessorAssign) int {
	countFreshAssessors := 0
	for _, itemRevise := range revisions {
		if !values.OwnsLocation(itemRevise.Location) {
			countFreshAssessors++
		}
	}
	return countFreshAssessors
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
//
func calculateFreshUrgencies(revisions []*Assessor, values *AssessorAssign, revisedSumBallotingPotency int64) {
	for _, itemRevise := range revisions {
		location := itemRevise.Location
		_, val := values.ObtainViaLocationAlterable(location)
		if val == nil {
			//
			//
			//
			//
			//
			//
			//
			//
			itemRevise.NominatorUrgency = -(revisedSumBallotingPotency + (revisedSumBallotingPotency >> 3))
		} else {
			itemRevise.NominatorUrgency = val.NominatorUrgency
		}
	}
}

//
//
//
//
func (values *AssessorAssign) executeRevisions(revisions []*Assessor) {
	current := values.Assessors
	sort.Sort(AssessorsViaLocator(current))

	integrated := make([]*Assessor, len(current)+len(revisions))
	i := 0

	for len(current) > 0 && len(revisions) > 0 {
		if bytes.Compare(current[0].Location, revisions[0].Location) < 0 { //
			integrated[i] = current[0]
			current = current[1:]
		} else {
			//
			integrated[i] = revisions[0]
			if bytes.Equal(current[0].Location, revisions[0].Location) {
				//
				current = current[1:]
			}
			revisions = revisions[1:]
		}
		i++
	}

	//
	for j := 0; j < len(current); j++ {
		integrated[i] = current[j]
		i++
	}
	//
	for j := 0; j < len(revisions); j++ {
		integrated[i] = revisions[j]
		i++
	}

	values.Assessors = integrated[:i]
}

//
//
//
func validateDeletions(removals []*Assessor, values *AssessorAssign) (ballotingPotency int64, err error) {
	discardedBallotingPotency := int64(0)
	for _, itemRevise := range removals {
		location := itemRevise.Location
		_, val := values.ObtainViaLocationAlterable(location)
		if val == nil {
			return discardedBallotingPotency, fmt.Errorf("REDACTED", location)
		}
		discardedBallotingPotency += val.BallotingPotency
	}
	if len(removals) > len(values.Assessors) {
		panic("REDACTED")
	}
	return discardedBallotingPotency, nil
}

//
//
//
//
func (values *AssessorAssign) executeDeletions(removals []*Assessor) {
	current := values.Assessors

	integrated := make([]*Assessor, len(current)-len(removals))
	i := 0

	//
	for len(removals) > 0 {
		if bytes.Equal(current[0].Location, removals[0].Location) {
			removals = removals[1:]
		} else { //
			integrated[i] = current[0]
			i++
		}
		current = current[1:]
	}

	//
	for j := 0; j < len(current); j++ {
		integrated[i] = current[j]
		i++
	}

	values.Assessors = integrated[:i]
}

//
//
//
//
//
func (values *AssessorAssign) reviseUsingModifyAssign(modifications []*Assessor, permitRemovals bool) error {
	if len(modifications) == 0 {
		return nil
	}

	//
	revisions, removals, err := handleModifications(modifications)
	if err != nil {
		return err
	}

	if !permitRemovals && len(removals) != 0 {
		return fmt.Errorf("REDACTED", removals)
	}

	//
	if countFreshAssessors(revisions, values) == 0 && len(values.Assessors) == len(removals) {
		return errors.New("REDACTED")
	}

	//
	//
	discardedBallotingPotency, err := validateDeletions(removals, values)
	if err != nil {
		return err
	}

	//
	//
	tvoterSubsequentRevisionsPriorDeletions, err := validateRevisions(revisions, values, discardedBallotingPotency)
	if err != nil {
		return err
	}

	//
	calculateFreshUrgencies(revisions, values, tvoterSubsequentRevisionsPriorDeletions)

	//
	values.executeRevisions(revisions)
	values.executeDeletions(removals)

	//
	values.inspectEveryTokensPossessIdenticalKind()

	if err = values.reviseSumBallotingPotency(); err != nil {
		panic(err)
	}

	//
	values.RecalibrateUrgencies(UrgencyFrameworkExtentElement * values.SumBallotingPotency())
	values.relocateViaMedianNominatorUrgency()

	sort.Sort(AssessorsViaBallotingPotency(values.Assessors))

	return nil
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
//
//
func (values *AssessorAssign) ReviseUsingModifyAssign(modifications []*Assessor) error {
	return values.reviseUsingModifyAssign(modifications, true)
}

//
//
func (values *AssessorAssign) ValidateEndorse(successionUUID string, ledgerUUID LedgerUUID,
	altitude int64, endorse *Endorse,
) error {
	return ValidateEndorse(successionUUID, values, ledgerUUID, altitude, endorse)
}

//

//
//
func (values *AssessorAssign) ValidateEndorseAgile(successionUUID string, ledgerUUID LedgerUUID,
	altitude int64, endorse *Endorse,
) error {
	return ValidateEndorseAgile(successionUUID, values, ledgerUUID, altitude, endorse)
}

//
//
//
//
//
//
func (values *AssessorAssign) ValidateEndorseAgileUsingStash(successionUUID string, ledgerUUID LedgerUUID,
	altitude int64, endorse *Endorse,
	attestedSigningStash *SigningStash,
) error {
	return ValidateEndorseAgileUsingStash(successionUUID, values, ledgerUUID, altitude, endorse, attestedSigningStash)
}

//
//
func (values *AssessorAssign) ValidateEndorseAgileEveryNotations(successionUUID string, ledgerUUID LedgerUUID,
	altitude int64, endorse *Endorse,
) error {
	return ValidateEndorseAgileEveryNotations(successionUUID, values, ledgerUUID, altitude, endorse)
}

//
//
//
func (values *AssessorAssign) ValidateEndorseAgileRelying(
	successionUUID string,
	endorse *Endorse,
	relianceStratum strongarithmetic.Portion,
) error {
	return ValidateEndorseAgileRelying(successionUUID, values, endorse, relianceStratum)
}

//
//
//
//
//
//
//
//
func (values *AssessorAssign) ValidateEndorseAgileRelyingUsingStash(
	successionUUID string,
	endorse *Endorse,
	relianceStratum strongarithmetic.Portion,
	attestedSigningStash *SigningStash,
) error {
	return ValidateEndorseAgileRelyingUsingStash(successionUUID, values, endorse, relianceStratum, attestedSigningStash)
}

//
//
//
func (values *AssessorAssign) ValidateEndorseAgileRelyingEveryNotations(
	successionUUID string,
	endorse *Endorse,
	relianceStratum strongarithmetic.Portion,
) error {
	return ValidateEndorseAgileRelyingEveryNotations(successionUUID, values, endorse, relianceStratum)
}

//
//
//
//
func (values *AssessorAssign) locatePrecedingNominator() *Assessor {
	var precedingNominator *Assessor
	for _, val := range values.Assessors {
		if precedingNominator == nil {
			precedingNominator = val
			continue
		}
		if precedingNominator == precedingNominator.ContrastNominatorUrgency(val) {
			precedingNominator = val
		}
	}
	return precedingNominator
}

func (values *AssessorAssign) inspectEveryTokensPossessIdenticalKind() {
	if values.Extent() == 0 {
		values.everyTokensPossessIdenticalKind = true
		return
	}

	initialTokenKind := "REDACTED"
	for _, val := range values.Assessors {
		if initialTokenKind == "REDACTED" {
			//
			if val.PublicToken == nil {
				continue
			}
			initialTokenKind = val.PublicToken.Kind()
		}
		if val.PublicToken.Kind() != initialTokenKind {
			values.everyTokensPossessIdenticalKind = false
			return
		}
	}

	values.everyTokensPossessIdenticalKind = true
}

//
//
func (values *AssessorAssign) EveryTokensPossessIdenticalKind() bool {
	return values.everyTokensPossessIdenticalKind
}

//

//
//
func EqualsFaultNegationAmpleBallotingPotencyNotated(err error) bool {
	return errors.As(err, &FaultNegationAmpleBallotingPotencyNotated{})
}

//
//
type FaultNegationAmpleBallotingPotencyNotated struct {
	Got    int64
	Required int64
}

func (e FaultNegationAmpleBallotingPotencyNotated) Failure() string {
	return fmt.Sprintf("REDACTED", e.Got, e.Required)
}

//

//
//
//
func (values *AssessorAssign) Text() string {
	return values.TextFormatted("REDACTED")
}

//
//
//
func (values *AssessorAssign) TextFormatted(format string) string {
	if values == nil {
		return "REDACTED"
	}
	var itemTexts []string
	values.Traverse(func(ordinal int, val *Assessor) bool {
		itemTexts = append(itemTexts, val.Text())
		return false
	})
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTED:
REDACTEDv
REDACTED`,
		format, values.ObtainNominator().Text(),
		format,
		format, strings.Join(itemTexts, "REDACTED"+format+"REDACTED"),
		format)
}

//

//
//
type AssessorsViaBallotingPotency []*Assessor

func (validz AssessorsViaBallotingPotency) Len() int { return len(validz) }

func (validz AssessorsViaBallotingPotency) Inferior(i, j int) bool {
	if validz[i].BallotingPotency == validz[j].BallotingPotency {
		return bytes.Compare(validz[i].Location, validz[j].Location) == -1
	}
	return validz[i].BallotingPotency > validz[j].BallotingPotency
}

func (validz AssessorsViaBallotingPotency) Exchange(i, j int) {
	validz[i], validz[j] = validz[j], validz[i]
}

//
//
type AssessorsViaLocator []*Assessor

func (validz AssessorsViaLocator) Len() int { return len(validz) }

func (validz AssessorsViaLocator) Inferior(i, j int) bool {
	return bytes.Compare(validz[i].Location, validz[j].Location) == -1
}

func (validz AssessorsViaLocator) Exchange(i, j int) {
	validz[i], validz[j] = validz[j], validz[i]
}

//
func (values *AssessorAssign) TowardSchema() (*commitchema.AssessorAssign, error) {
	if values.EqualsVoidEitherBlank() {
		return &commitchema.AssessorAssign{}, nil //
	}

	vp := new(commitchema.AssessorAssign)
	valuesSchema := make([]*commitchema.Assessor, len(values.Assessors))
	for i := 0; i < len(values.Assessors); i++ {
		validp, err := values.Assessors[i].TowardSchema()
		if err != nil {
			return nil, err
		}
		valuesSchema[i] = validp
	}
	vp.Assessors = valuesSchema

	itemNominator, err := values.Nominator.TowardSchema()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	vp.Nominator = itemNominator

	//
	//
	vp.SumBallotingPotency = 0

	return vp, nil
}

//
//
//
func AssessorAssignOriginatingSchema(vp *commitchema.AssessorAssign) (*AssessorAssign, error) {
	if vp == nil {
		return nil, errors.New("REDACTED") //
	}
	values := new(AssessorAssign)

	valuesSchema := make([]*Assessor, len(vp.Assessors))
	for i := 0; i < len(vp.Assessors); i++ {
		v, err := AssessorOriginatingSchema(vp.Assessors[i])
		if err != nil {
			return nil, err
		}
		valuesSchema[i] = v
	}
	values.Assessors = valuesSchema
	values.inspectEveryTokensPossessIdenticalKind()

	p, err := AssessorOriginatingSchema(vp.ObtainNominator())
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	values.Nominator = p

	//
	//
	//
	//
	//
	//
	if _, err := values.SumBallotingPotencySecure(); err != nil {
		return nil, err
	}

	return values, values.CertifyFundamental()
}

//
//
//
//
func AssessorAssignOriginatingCurrentAssessors(validz []*Assessor) (*AssessorAssign, error) {
	if len(validz) == 0 {
		return nil, errors.New("REDACTED")
	}
	for _, val := range validz {
		err := val.CertifyFundamental()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	values := &AssessorAssign{
		Assessors: validz,
	}
	values.inspectEveryTokensPossessIdenticalKind()
	values.Nominator = values.locatePrecedingNominator()
	if err := values.reviseSumBallotingPotency(); err != nil {
		return nil, err
	}
	sort.Sort(AssessorsViaBallotingPotency(values.Assessors))
	return values, nil
}

//

//
//
//
//
func ArbitraryAssessorAssign(countAssessors int, ballotingPotency int64) (*AssessorAssign, []PrivateAssessor) {
	var (
		validz           = make([]*Assessor, countAssessors)
		privateAssessors = make([]PrivateAssessor, countAssessors)
	)

	for i := 0; i < countAssessors; i++ {
		val, privateAssessor := ArbitraryAssessor(false, ballotingPotency)
		validz[i] = val
		privateAssessors[i] = privateAssessor
	}

	sort.Sort(PrivateAssessorsViaLocation(privateAssessors))

	return FreshAssessorAssign(validz), privateAssessors
}

//

func secureAppend(a, b int64) (int64, bool) {
	if b > 0 && a > math.MaxInt64-b {
		return -1, true
	} else if b < 0 && a < math.MinInt64-b {
		return -1, true
	}
	return a + b, false
}

func secureUnder(a, b int64) (int64, bool) {
	if b > 0 && a < math.MinInt64+b {
		return -1, true
	} else if b < 0 && a > math.MaxInt64+b {
		return -1, true
	}
	return a - b, false
}

func secureAppendRestrict(a, b int64) int64 {
	c, overrun := secureAppend(a, b)
	if overrun {
		if b < 0 {
			return math.MinInt64
		}
		return math.MaxInt64
	}
	return c
}

func secureUnderRestrict(a, b int64) int64 {
	c, overrun := secureUnder(a, b)
	if overrun {
		if b > 0 {
			return math.MinInt64
		}
		return math.MaxInt64
	}
	return c
}

func secureMultiply(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}

	absoluteBelongingBYTE := b
	if b < 0 {
		absoluteBelongingBYTE = -b
	}

	absoluteBelongingAN := a
	if a < 0 {
		absoluteBelongingAN = -a
	}

	if absoluteBelongingAN > math.MaxInt64/absoluteBelongingBYTE {
		return 0, true
	}

	return a * b, false
}
