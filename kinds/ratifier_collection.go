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

	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/vault/comethash"
	cometmath "github.com/valkyrieworks/utils/math"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

const (
	//
	//
	//
	//
	//
	//
	//
	MaximumSumPollingEnergy = int64(math.MaxInt64) / 8

	//
	//
	//
	UrgencyPeriodVolumeCoefficient = 2
)

//
//
var ErrSumPollingEnergyOverload = fmt.Errorf("REDACTED",
	MaximumSumPollingEnergy)

//
var ErrRecommenderNotInValues = errors.New("REDACTED")

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
type RatifierAssign struct {
	//
	Ratifiers []*Ratifier `json:"ratifiers"`
	Recommender   *Ratifier   `json:"recommender"`

	//
	sumPollingEnergy int64
	//
	allKeysPossessIdenticalKind bool
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
func NewRatifierCollection(valz []*Ratifier) *RatifierAssign {
	values := &RatifierAssign{
		allKeysPossessIdenticalKind: true,
	}
	err := values.modifyWithAlterCollection(valz, false)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if len(valz) > 0 {
		values.AugmentRecommenderUrgency(1)
	}
	return values
}

func (values *RatifierAssign) CertifySimple() error {
	if values.IsNullOrEmpty() {
		return errors.New("REDACTED")
	}

	for idx, val := range values.Ratifiers {
		if err := val.CertifySimple(); err != nil {
			return fmt.Errorf("REDACTED", idx, err)
		}
	}

	if err := values.Recommender.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	for _, val := range values.Ratifiers {
		if bytes.Equal(val.Location, values.Recommender.Location) {
			return nil
		}
	}

	return ErrRecommenderNotInValues
}

//
func (values *RatifierAssign) IsNullOrEmpty() bool {
	return values == nil || len(values.Ratifiers) == 0
}

//
//
func (values *RatifierAssign) CloneAugmentRecommenderUrgency(instances int32) *RatifierAssign {
	cp := values.Clone()
	cp.AugmentRecommenderUrgency(instances)
	return cp
}

//
//
//
func (values *RatifierAssign) AugmentRecommenderUrgency(instances int32) {
	if values.IsNullOrEmpty() {
		panic("REDACTED")
	}
	if instances <= 0 {
		panic("REDACTED")
	}

	//
	//
	//
	varyMaximum := UrgencyPeriodVolumeCoefficient * values.SumPollingEnergy()
	values.ReadjustUrgencies(varyMaximum)
	values.displaceByAverageRecommenderUrgency()

	var recommender *Ratifier
	//
	for i := int32(0); i < instances; i++ {
		recommender = values.augmentRecommenderUrgency()
	}

	values.Recommender = recommender
}

//
//
//
func (values *RatifierAssign) ReadjustUrgencies(varyMaximum int64) {
	if values.IsNullOrEmpty() {
		panic("REDACTED")
	}
	//
	//
	//
	if varyMaximum <= 0 {
		return
	}

	//
	//
	//
	vary := calculateMaximumMinimumUrgencyVary(values)
	proportion := (vary + varyMaximum - 1) / varyMaximum
	if vary > varyMaximum {
		for _, val := range values.Ratifiers {
			val.RecommenderUrgency /= proportion
		}
	}
}

func (values *RatifierAssign) augmentRecommenderUrgency() *Ratifier {
	for _, val := range values.Ratifiers {
		//
		newUrgency := secureAppendTruncate(val.RecommenderUrgency, val.PollingEnergy)
		val.RecommenderUrgency = newUrgency
	}
	//
	utmost := values.fetchValueWithGreatestUrgency()
	//
	utmost.RecommenderUrgency = secureSubtractTruncate(utmost.RecommenderUrgency, values.SumPollingEnergy())

	return utmost
}

//
func (values *RatifierAssign) calculateAverageRecommenderUrgency() int64 {
	n := int64(len(values.Ratifiers))
	sum := big.NewInt(0)
	for _, val := range values.Ratifiers {
		sum.Add(sum, big.NewInt(val.RecommenderUrgency))
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
func calculateMaximumMinimumUrgencyVary(values *RatifierAssign) int64 {
	if values.IsNullOrEmpty() {
		panic("REDACTED")
	}
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)
	for _, v := range values.Ratifiers {
		if v.RecommenderUrgency < min {
			min = v.RecommenderUrgency
		}
		if v.RecommenderUrgency > max {
			max = v.RecommenderUrgency
		}
	}
	vary := max - min
	if vary < 0 {
		return -1 * vary
	}
	return vary
}

func (values *RatifierAssign) fetchValueWithGreatestUrgency() *Ratifier {
	var res *Ratifier
	for _, val := range values.Ratifiers {
		res = res.ContrastRecommenderUrgency(val)
	}
	return res
}

func (values *RatifierAssign) displaceByAverageRecommenderUrgency() {
	if values.IsNullOrEmpty() {
		panic("REDACTED")
	}
	averageRecommenderUrgency := values.calculateAverageRecommenderUrgency()
	for _, val := range values.Ratifiers {
		val.RecommenderUrgency = secureSubtractTruncate(val.RecommenderUrgency, averageRecommenderUrgency)
	}
}

//
func ratifierCatalogClone(valuesCatalog []*Ratifier) []*Ratifier {
	if valuesCatalog == nil {
		return nil
	}
	valuesClone := make([]*Ratifier, len(valuesCatalog))
	for i, val := range valuesCatalog {
		valuesClone[i] = val.Clone()
	}
	return valuesClone
}

//
func (values *RatifierAssign) Clone() *RatifierAssign {
	return &RatifierAssign{
		Ratifiers:          ratifierCatalogClone(values.Ratifiers),
		Recommender:            values.Recommender,
		sumPollingEnergy:    values.sumPollingEnergy,
		allKeysPossessIdenticalKind: values.allKeysPossessIdenticalKind,
	}
}

//
//
func (values *RatifierAssign) HasLocation(location []byte) bool {
	for _, val := range values.Ratifiers {
		if bytes.Equal(val.Location, location) {
			return true
		}
	}
	return false
}

//
//
func (values *RatifierAssign) FetchByLocation(location []byte) (ordinal int32, val *Ratifier) {
	i, val := values.FetchByLocationMut(location)
	if i == -1 {
		return -1, nil
	}
	return i, val.Clone()
}

//
//
//
//
func (values *RatifierAssign) FetchByLocationMut(location []byte) (ordinal int32, val *Ratifier) {
	for idx, val := range values.Ratifiers {
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
func (values *RatifierAssign) FetchByOrdinal(ordinal int32) (location []byte, val *Ratifier) {
	if ordinal < 0 || int(ordinal) >= len(values.Ratifiers) {
		return nil, nil
	}
	val = values.Ratifiers[ordinal]
	return val.Location, val.Clone()
}

//
func (values *RatifierAssign) Volume() int {
	return len(values.Ratifiers)
}

//
//
func (values *RatifierAssign) modifySumPollingEnergy() error {
	sum := int64(0)
	for _, val := range values.Ratifiers {
		//
		sum = secureAppendTruncate(sum, val.PollingEnergy)
		if sum > MaximumSumPollingEnergy {
			return fmt.Errorf("REDACTED", sum, MaximumSumPollingEnergy)
		}
	}

	values.sumPollingEnergy = sum
	return nil
}

//
//
func (values *RatifierAssign) SumPollingEnergySecure() (int64, error) {
	if values.sumPollingEnergy == 0 {
		if err := values.modifySumPollingEnergy(); err != nil {
			return 0, err
		}
	}
	return values.sumPollingEnergy, nil
}

//
//
func (values *RatifierAssign) SumPollingEnergy() int64 {
	if values.sumPollingEnergy == 0 {
		if err := values.modifySumPollingEnergy(); err != nil {
			panic(err)
		}
	}
	return values.sumPollingEnergy
}

//
//
func (values *RatifierAssign) FetchRecommender() (recommender *Ratifier) {
	if len(values.Ratifiers) == 0 {
		return nil
	}
	if values.Recommender == nil {
		values.Recommender = values.locateRecommender()
	}
	return values.Recommender.Clone()
}

func (values *RatifierAssign) locateRecommender() *Ratifier {
	var recommender *Ratifier
	for _, val := range values.Ratifiers {
		if recommender == nil || !bytes.Equal(val.Location, recommender.Location) {
			recommender = recommender.ContrastRecommenderUrgency(val)
		}
	}
	return recommender
}

//
//
//
//
func (values *RatifierAssign) Digest() []byte {
	bzs := make([][]byte, len(values.Ratifiers))
	for i, val := range values.Ratifiers {
		bzs[i] = val.Octets()
	}
	return merkle.DigestFromOctetSegments(bzs)
}

//
//
//
func (values *RatifierAssign) RecommenderUrgencyDigest() []byte {
	if len(values.Ratifiers) == 0 {
		return nil
	}

	buf := make([]byte, binary.MaxVarintLen64*len(values.Ratifiers))
	displacement := 0
	for _, val := range values.Ratifiers {
		n := binary.PutVarint(buf[displacement:], val.RecommenderUrgency)
		displacement += n
	}
	return comethash.Sum(buf[:displacement])
}

//
func (values *RatifierAssign) Recurse(fn func(ordinal int, val *Ratifier) bool) {
	for i, val := range values.Ratifiers {
		halt := fn(i, val.Clone())
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
func handleModifications(origModifications []*Ratifier) (refreshes, deletions []*Ratifier, err error) {
	//
	modifications := ratifierCatalogClone(origModifications)
	sort.Sort(RatifiersByLocation(modifications))

	deletions = make([]*Ratifier, 0, len(modifications))
	refreshes = make([]*Ratifier, 0, len(modifications))
	var previousAddress Location

	//
	for _, valueModify := range modifications {
		if bytes.Equal(valueModify.Location, previousAddress) {
			err = fmt.Errorf("REDACTED", valueModify, modifications)
			return nil, nil, err
		}

		switch {
		case valueModify.PollingEnergy < 0:
			err = fmt.Errorf("REDACTED", valueModify.PollingEnergy)
			return nil, nil, err
		case valueModify.PollingEnergy > MaximumSumPollingEnergy:
			err = fmt.Errorf("REDACTED",
				MaximumSumPollingEnergy, valueModify.PollingEnergy)
			return nil, nil, err
		case valueModify.PollingEnergy == 0:
			deletions = append(deletions, valueModify)
		default:
			refreshes = append(refreshes, valueModify)
		}

		previousAddress = valueModify.Location
	}

	return refreshes, deletions, err
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
func validateRefreshes(
	refreshes []*Ratifier,
	values *RatifierAssign,
	deletedEnergy int64,
) (tvpAfterRefreshesPriorDeletions int64, err error) {
	variance := func(modify *Ratifier, values *RatifierAssign) int64 {
		_, val := values.FetchByLocationMut(modify.Location)
		if val != nil {
			return modify.PollingEnergy - val.PollingEnergy
		}
		return modify.PollingEnergy
	}

	refreshesClone := ratifierCatalogClone(refreshes)
	sort.Slice(refreshesClone, func(i, j int) bool {
		return variance(refreshesClone[i], values) < variance(refreshesClone[j], values)
	})

	tvpAfterDeletions := values.SumPollingEnergy() - deletedEnergy
	for _, upd := range refreshesClone {
		tvpAfterDeletions += variance(upd, values)
		if tvpAfterDeletions > MaximumSumPollingEnergy {
			return 0, ErrSumPollingEnergyOverload
		}
	}
	return tvpAfterDeletions + deletedEnergy, nil
}

func countNewRatifiers(refreshes []*Ratifier, values *RatifierAssign) int {
	countNewRatifiers := 0
	for _, valueModify := range refreshes {
		if !values.HasLocation(valueModify.Location) {
			countNewRatifiers++
		}
	}
	return countNewRatifiers
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
func calculateNewUrgencies(refreshes []*Ratifier, values *RatifierAssign, refreshedSumPollingEnergy int64) {
	for _, valueModify := range refreshes {
		location := valueModify.Location
		_, val := values.FetchByLocationMut(location)
		if val == nil {
			//
			//
			//
			//
			//
			//
			//
			//
			valueModify.RecommenderUrgency = -(refreshedSumPollingEnergy + (refreshedSumPollingEnergy >> 3))
		} else {
			valueModify.RecommenderUrgency = val.RecommenderUrgency
		}
	}
}

//
//
//
//
func (values *RatifierAssign) executeRefreshes(refreshes []*Ratifier) {
	current := values.Ratifiers
	sort.Sort(RatifiersByLocation(current))

	combined := make([]*Ratifier, len(current)+len(refreshes))
	i := 0

	for len(current) > 0 && len(refreshes) > 0 {
		if bytes.Compare(current[0].Location, refreshes[0].Location) < 0 { //
			combined[i] = current[0]
			current = current[1:]
		} else {
			//
			combined[i] = refreshes[0]
			if bytes.Equal(current[0].Location, refreshes[0].Location) {
				//
				current = current[1:]
			}
			refreshes = refreshes[1:]
		}
		i++
	}

	//
	for j := 0; j < len(current); j++ {
		combined[i] = current[j]
		i++
	}
	//
	for j := 0; j < len(refreshes); j++ {
		combined[i] = refreshes[j]
		i++
	}

	values.Ratifiers = combined[:i]
}

//
//
//
func validateDeletions(removals []*Ratifier, values *RatifierAssign) (pollingEnergy int64, err error) {
	deletedPollingEnergy := int64(0)
	for _, valueModify := range removals {
		location := valueModify.Location
		_, val := values.FetchByLocationMut(location)
		if val == nil {
			return deletedPollingEnergy, fmt.Errorf("REDACTED", location)
		}
		deletedPollingEnergy += val.PollingEnergy
	}
	if len(removals) > len(values.Ratifiers) {
		panic("REDACTED")
	}
	return deletedPollingEnergy, nil
}

//
//
//
//
func (values *RatifierAssign) executeDeletions(removals []*Ratifier) {
	current := values.Ratifiers

	combined := make([]*Ratifier, len(current)-len(removals))
	i := 0

	//
	for len(removals) > 0 {
		if bytes.Equal(current[0].Location, removals[0].Location) {
			removals = removals[1:]
		} else { //
			combined[i] = current[0]
			i++
		}
		current = current[1:]
	}

	//
	for j := 0; j < len(current); j++ {
		combined[i] = current[j]
		i++
	}

	values.Ratifiers = combined[:i]
}

//
//
//
//
//
func (values *RatifierAssign) modifyWithAlterCollection(modifications []*Ratifier, permitRemovals bool) error {
	if len(modifications) == 0 {
		return nil
	}

	//
	refreshes, removals, err := handleModifications(modifications)
	if err != nil {
		return err
	}

	if !permitRemovals && len(removals) != 0 {
		return fmt.Errorf("REDACTED", removals)
	}

	//
	if countNewRatifiers(refreshes, values) == 0 && len(values.Ratifiers) == len(removals) {
		return errors.New("REDACTED")
	}

	//
	//
	deletedPollingEnergy, err := validateDeletions(removals, values)
	if err != nil {
		return err
	}

	//
	//
	tvpAfterRefreshesPriorDeletions, err := validateRefreshes(refreshes, values, deletedPollingEnergy)
	if err != nil {
		return err
	}

	//
	calculateNewUrgencies(refreshes, values, tvpAfterRefreshesPriorDeletions)

	//
	values.executeRefreshes(refreshes)
	values.executeDeletions(removals)

	//
	values.inspectAllKeysPossessIdenticalKind()

	if err = values.modifySumPollingEnergy(); err != nil {
		panic(err)
	}

	//
	values.ReadjustUrgencies(UrgencyPeriodVolumeCoefficient * values.SumPollingEnergy())
	values.displaceByAverageRecommenderUrgency()

	sort.Sort(RatifiersByPollingEnergy(values.Ratifiers))

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
func (values *RatifierAssign) ModifyWithAlterCollection(modifications []*Ratifier) error {
	return values.modifyWithAlterCollection(modifications, true)
}

//
//
func (values *RatifierAssign) ValidateEndorse(ledgerUID string, ledgerUID LedgerUID,
	level int64, endorse *Endorse,
) error {
	return ValidateEndorse(ledgerUID, values, ledgerUID, level, endorse)
}

//

//
//
func (values *RatifierAssign) ValidateEndorseRapid(ledgerUID string, ledgerUID LedgerUID,
	level int64, endorse *Endorse,
) error {
	return ValidateEndorseRapid(ledgerUID, values, ledgerUID, level, endorse)
}

//
//
//
//
//
//
func (values *RatifierAssign) ValidateEndorseRapidWithRepository(ledgerUID string, ledgerUID LedgerUID,
	level int64, endorse *Endorse,
	certifiedAutographRepository *AutographRepository,
) error {
	return ValidateEndorseRapidWithRepository(ledgerUID, values, ledgerUID, level, endorse, certifiedAutographRepository)
}

//
//
func (values *RatifierAssign) ValidateEndorseRapidAllEndorsements(ledgerUID string, ledgerUID LedgerUID,
	level int64, endorse *Endorse,
) error {
	return ValidateEndorseRapidAllEndorsements(ledgerUID, values, ledgerUID, level, endorse)
}

//
//
//
func (values *RatifierAssign) ValidateEndorseRapidValidating(
	ledgerUID string,
	endorse *Endorse,
	validateLayer cometmath.Portion,
) error {
	return ValidateEndorseRapidValidating(ledgerUID, values, endorse, validateLayer)
}

//
//
//
//
//
//
//
//
func (values *RatifierAssign) ValidateEndorseRapidValidatingWithRepository(
	ledgerUID string,
	endorse *Endorse,
	validateLayer cometmath.Portion,
	certifiedAutographRepository *AutographRepository,
) error {
	return ValidateEndorseRapidValidatingWithRepository(ledgerUID, values, endorse, validateLayer, certifiedAutographRepository)
}

//
//
//
func (values *RatifierAssign) ValidateEndorseRapidValidatingAllEndorsements(
	ledgerUID string,
	endorse *Endorse,
	validateLayer cometmath.Portion,
) error {
	return ValidateEndorseRapidValidatingAllEndorsements(ledgerUID, values, endorse, validateLayer)
}

//
//
//
//
func (values *RatifierAssign) locatePrecedingRecommender() *Ratifier {
	var precedingRecommender *Ratifier
	for _, val := range values.Ratifiers {
		if precedingRecommender == nil {
			precedingRecommender = val
			continue
		}
		if precedingRecommender == precedingRecommender.ContrastRecommenderUrgency(val) {
			precedingRecommender = val
		}
	}
	return precedingRecommender
}

func (values *RatifierAssign) inspectAllKeysPossessIdenticalKind() {
	if values.Volume() == 0 {
		values.allKeysPossessIdenticalKind = true
		return
	}

	initialKeyKind := "REDACTED"
	for _, val := range values.Ratifiers {
		if initialKeyKind == "REDACTED" {
			//
			if val.PublicKey == nil {
				continue
			}
			initialKeyKind = val.PublicKey.Kind()
		}
		if val.PublicKey.Kind() != initialKeyKind {
			values.allKeysPossessIdenticalKind = false
			return
		}
	}

	values.allKeysPossessIdenticalKind = true
}

//
//
func (values *RatifierAssign) AllKeysPossessIdenticalKind() bool {
	return values.allKeysPossessIdenticalKind
}

//

//
//
func IsErrNotSufficientPollingEnergyAttested(err error) bool {
	return errors.As(err, &ErrNotSufficientPollingEnergyAttested{})
}

//
//
type ErrNotSufficientPollingEnergyAttested struct {
	Got    int64
	Required int64
}

func (e ErrNotSufficientPollingEnergyAttested) Fault() string {
	return fmt.Sprintf("REDACTED", e.Got, e.Required)
}

//

//
//
//
func (values *RatifierAssign) String() string {
	return values.StringIndented("REDACTED")
}

//
//
//
func (values *RatifierAssign) StringIndented(indent string) string {
	if values == nil {
		return "REDACTED"
	}
	var valueStrings []string
	values.Recurse(func(ordinal int, val *Ratifier) bool {
		valueStrings = append(valueStrings, val.String())
		return false
	})
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTED:
REDACTEDv
REDACTED`,
		indent, values.FetchRecommender().String(),
		indent,
		indent, strings.Join(valueStrings, "REDACTED"+indent+"REDACTED"),
		indent)
}

//

//
//
type RatifiersByPollingEnergy []*Ratifier

func (valz RatifiersByPollingEnergy) Len() int { return len(valz) }

func (valz RatifiersByPollingEnergy) Lower(i, j int) bool {
	if valz[i].PollingEnergy == valz[j].PollingEnergy {
		return bytes.Compare(valz[i].Location, valz[j].Location) == -1
	}
	return valz[i].PollingEnergy > valz[j].PollingEnergy
}

func (valz RatifiersByPollingEnergy) Exchange(i, j int) {
	valz[i], valz[j] = valz[j], valz[i]
}

//
//
type RatifiersByLocation []*Ratifier

func (valz RatifiersByLocation) Len() int { return len(valz) }

func (valz RatifiersByLocation) Lower(i, j int) bool {
	return bytes.Compare(valz[i].Location, valz[j].Location) == -1
}

func (valz RatifiersByLocation) Exchange(i, j int) {
	valz[i], valz[j] = valz[j], valz[i]
}

//
func (values *RatifierAssign) ToSchema() (*engineproto.RatifierAssign, error) {
	if values.IsNullOrEmpty() {
		return &engineproto.RatifierAssign{}, nil //
	}

	vp := new(engineproto.RatifierAssign)
	valuesSchema := make([]*engineproto.Ratifier, len(values.Ratifiers))
	for i := 0; i < len(values.Ratifiers); i++ {
		valp, err := values.Ratifiers[i].ToSchema()
		if err != nil {
			return nil, err
		}
		valuesSchema[i] = valp
	}
	vp.Ratifiers = valuesSchema

	valueRecommender, err := values.Recommender.ToSchema()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	vp.Recommender = valueRecommender

	//
	//
	vp.SumPollingEnergy = 0

	return vp, nil
}

//
//
//
func RatifierCollectionFromSchema(vp *engineproto.RatifierAssign) (*RatifierAssign, error) {
	if vp == nil {
		return nil, errors.New("REDACTED") //
	}
	values := new(RatifierAssign)

	valuesSchema := make([]*Ratifier, len(vp.Ratifiers))
	for i := 0; i < len(vp.Ratifiers); i++ {
		v, err := RatifierFromSchema(vp.Ratifiers[i])
		if err != nil {
			return nil, err
		}
		valuesSchema[i] = v
	}
	values.Ratifiers = valuesSchema
	values.inspectAllKeysPossessIdenticalKind()

	p, err := RatifierFromSchema(vp.FetchRecommender())
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	values.Recommender = p

	//
	//
	//
	//
	//
	//
	if _, err := values.SumPollingEnergySecure(); err != nil {
		return nil, err
	}

	return values, values.CertifySimple()
}

//
//
//
//
func RatifierCollectionFromCurrentRatifiers(valz []*Ratifier) (*RatifierAssign, error) {
	if len(valz) == 0 {
		return nil, errors.New("REDACTED")
	}
	for _, val := range valz {
		err := val.CertifySimple()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	values := &RatifierAssign{
		Ratifiers: valz,
	}
	values.inspectAllKeysPossessIdenticalKind()
	values.Recommender = values.locatePrecedingRecommender()
	if err := values.modifySumPollingEnergy(); err != nil {
		return nil, err
	}
	sort.Sort(RatifiersByPollingEnergy(values.Ratifiers))
	return values, nil
}

//

//
//
//
//
func RandomRatifierCollection(countRatifiers int, pollingEnergy int64) (*RatifierAssign, []PrivateRatifier) {
	var (
		valz           = make([]*Ratifier, countRatifiers)
		privateRatifiers = make([]PrivateRatifier, countRatifiers)
	)

	for i := 0; i < countRatifiers; i++ {
		val, privateRatifier := RandomRatifier(false, pollingEnergy)
		valz[i] = val
		privateRatifiers[i] = privateRatifier
	}

	sort.Sort(PrivateRatifiersByLocation(privateRatifiers))

	return NewRatifierCollection(valz), privateRatifiers
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

func secureSubtract(a, b int64) (int64, bool) {
	if b > 0 && a < math.MinInt64+b {
		return -1, true
	} else if b < 0 && a > math.MaxInt64+b {
		return -1, true
	}
	return a - b, false
}

func secureAppendTruncate(a, b int64) int64 {
	c, overload := secureAppend(a, b)
	if overload {
		if b < 0 {
			return math.MinInt64
		}
		return math.MaxInt64
	}
	return c
}

func secureSubtractTruncate(a, b int64) int64 {
	c, overload := secureSubtract(a, b)
	if overload {
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

	absoluteOfBYTE := b
	if b < 0 {
		absoluteOfBYTE = -b
	}

	absoluteOfA := a
	if a < 0 {
		absoluteOfA = -a
	}

	if absoluteOfA > math.MaxInt64/absoluteOfBYTE {
		return 0, true
	}

	return a * b, false
}
