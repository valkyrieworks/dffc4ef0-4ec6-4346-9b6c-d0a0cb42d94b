//
//
//

package pex

import (
	"encoding/binary"
	"fmt"
	"hash"
	"math"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/minio/highwayhash"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

const (
	segmentKindFresh = 0x01
	segmentKindAged = 0x02
)

//
//
//
//
type LocationRegister interface {
	facility.Facility

	//
	AppendMineLocator(*p2p.NetworkLocator)
	//
	MineLocator(*p2p.NetworkLocator) bool

	AppendSecludedIDXDstore([]string)

	//
	AppendLocator(location *p2p.NetworkLocator, src *p2p.NetworkLocator) error
	DiscardLocator(*p2p.NetworkLocator)

	//
	OwnsLocation(*p2p.NetworkLocator) bool

	//
	RequireExtraLocations() bool
	//
	//
	Blank() bool

	//
	SelectLocator(tendencyTowardFreshLocations int) *p2p.NetworkLocator

	//
	LabelValid(p2p.ID)
	LabelEffort(*p2p.NetworkLocator)
	LabelFlawed(*p2p.NetworkLocator, time.Duration) //
	//
	RestoreFlawedNodes()

	EqualsValid(*p2p.NetworkLocator) bool
	EqualsProhibited(*p2p.NetworkLocator) bool

	//
	FetchPreference() []*p2p.NetworkLocator
	//
	FetchPreferenceUsingTendency(tendencyTowardFreshLocations int) []*p2p.NetworkLocator

	Extent() int

	//
	Persist()
}

var _ LocationRegister = (*locationRegister)(nil)

//
//
type locationRegister struct {
	facility.FoundationFacility

	//
	mtx        commitchronize.Exclusion
	arbitrary       *commitrand.Arbitrary
	mineLocations   map[string]struct{}
	secludedIDXDstore map[p2p.ID]struct{}
	locationSearch map[p2p.ID]*recognizedLocator //
	flawedNodes   map[p2p.ID]*recognizedLocator //
	segmentsAged []map[string]*recognizedLocator
	segmentsFresh []map[string]*recognizedLocator
	nthAged       int
	nthFresh       int

	//
	recordRoute          string
	key               string //
	reachabilityStringent bool
	digester            hash.Hash64

	wg sync.WaitGroup
}

func shouldFreshDigester() hash.Hash64 {
	key := security.CHARArbitraryOctets(highwayhash.Size)
	digester, err := highwayhash.New64(key)
	if err != nil {
		panic(err)
	}
	return digester
}

//
//
func FreshLocationRegister(recordRoute string, reachabilityStringent bool) LocationRegister {
	am := &locationRegister{
		arbitrary:              commitrand.FreshArbitrary(),
		mineLocations:          make(map[string]struct{}),
		secludedIDXDstore:        make(map[p2p.ID]struct{}),
		locationSearch:        make(map[p2p.ID]*recognizedLocator),
		flawedNodes:          make(map[p2p.ID]*recognizedLocator),
		recordRoute:          recordRoute,
		reachabilityStringent: reachabilityStringent,
	}
	am.initialize()
	am.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", am)
	return am
}

//
//
func (a *locationRegister) initialize() {
	a.key = security.CHARArbitraryHexadecimal(24) //
	//
	a.segmentsFresh = make([]map[string]*recognizedLocator, freshSegmentTotal)
	for i := range a.segmentsFresh {
		a.segmentsFresh[i] = make(map[string]*recognizedLocator)
	}
	//
	a.segmentsAged = make([]map[string]*recognizedLocator, agedSegmentTotal)
	for i := range a.segmentsAged {
		a.segmentsAged[i] = make(map[string]*recognizedLocator)
	}
	a.digester = shouldFreshDigester()
}

//
func (a *locationRegister) UponInitiate() error {
	if err := a.FoundationFacility.UponInitiate(); err != nil {
		return err
	}
	a.fetchOriginatingRecord(a.recordRoute)

	//
	//
	a.wg.Add(1)
	go a.persistProcedure()

	return nil
}

//
func (a *locationRegister) UponHalt() {
	a.FoundationFacility.UponHalt()
}

func (a *locationRegister) Pause() {
	a.wg.Wait()
}

func (a *locationRegister) RecordRoute() string {
	return a.recordRoute
}

//

//
func (a *locationRegister) AppendMineLocator(location *p2p.NetworkLocator) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.Tracer.Details("REDACTED", "REDACTED", location)
	a.mineLocations[location.Text()] = struct{}{}
}

//
func (a *locationRegister) MineLocator(location *p2p.NetworkLocator) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	_, ok := a.mineLocations[location.Text()]
	return ok
}

func (a *locationRegister) AppendSecludedIDXDstore(ids []string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	for _, id := range ids {
		a.secludedIDXDstore[p2p.ID(id)] = struct{}{}
	}
}

//
//
//
//
func (a *locationRegister) AppendLocator(location *p2p.NetworkLocator, src *p2p.NetworkLocator) error {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	return a.appendLocator(location, src)
}

//
func (a *locationRegister) DiscardLocator(location *p2p.NetworkLocator) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.discardLocator(location)
}

//
//
func (a *locationRegister) EqualsValid(location *p2p.NetworkLocator) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	return a.locationSearch[location.ID].equalsAged()
}

//
func (a *locationRegister) EqualsProhibited(location *p2p.NetworkLocator) bool {
	a.mtx.Lock()
	_, ok := a.flawedNodes[location.ID]
	a.mtx.Unlock()

	return ok
}

//
func (a *locationRegister) OwnsLocation(location *p2p.NetworkLocator) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	ka := a.locationSearch[location.ID]
	return ka != nil
}

//
func (a *locationRegister) RequireExtraLocations() bool {
	return a.Extent() < requireLocatorLimit
}

//
//
func (a *locationRegister) Blank() bool {
	return a.Extent() == 0
}

//
//
//
//
//
//
func (a *locationRegister) SelectLocator(tendencyTowardFreshLocations int) *p2p.NetworkLocator {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	registerExtent := a.extent()
	if registerExtent <= 0 {
		if registerExtent < 0 {
			panic(fmt.Sprintf("REDACTED", a.nthFresh+a.nthAged, a.nthFresh, a.nthAged))
		}
		return nil
	}
	if tendencyTowardFreshLocations > 100 {
		tendencyTowardFreshLocations = 100
	}
	if tendencyTowardFreshLocations < 0 {
		tendencyTowardFreshLocations = 0
	}

	//
	agedRelation := math.Sqrt(float64(a.nthAged)) * (100.0 - float64(tendencyTowardFreshLocations))
	freshRelation := math.Sqrt(float64(a.nthFresh)) * float64(tendencyTowardFreshLocations)

	//
	var segment map[string]*recognizedLocator
	selectOriginatingAgedSegment := (freshRelation+agedRelation)*a.arbitrary.Float64() < agedRelation
	if (selectOriginatingAgedSegment && a.nthAged == 0) ||
		(!selectOriginatingAgedSegment && a.nthFresh == 0) {
		return nil
	}
	//
	for len(segment) == 0 {
		if selectOriginatingAgedSegment {
			segment = a.segmentsAged[a.arbitrary.Integern(len(a.segmentsAged))]
		} else {
			segment = a.segmentsFresh[a.arbitrary.Integern(len(a.segmentsFresh))]
		}
	}
	//
	arbitraryPosition := a.arbitrary.Integern(len(segment))
	for _, ka := range segment {
		if arbitraryPosition == 0 {
			return ka.Location
		}
		arbitraryPosition--
	}
	return nil
}

//
//
func (a *locationRegister) LabelValid(id p2p.ID) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	ka := a.locationSearch[id]
	if ka == nil {
		return
	}
	ka.labelValid()
	if ka.equalsFresh() {
		if err := a.shiftTowardAged(ka); err != nil {
			a.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
}

//
func (a *locationRegister) LabelEffort(location *p2p.NetworkLocator) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	ka := a.locationSearch[location.ID]
	if ka == nil {
		return
	}
	ka.labelEffort()
}

//
//
func (a *locationRegister) LabelFlawed(location *p2p.NetworkLocator, prohibitMoment time.Duration) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	if a.appendFlawedNode(location, prohibitMoment) {
		a.discardLocator(location)
	}
}

//
//
func (a *locationRegister) RestoreFlawedNodes() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	for _, ka := range a.flawedNodes {
		if ka.equalsProhibited() {
			continue
		}

		segment, err := a.computeFreshSegment(ka.Location, ka.Src)
		if err != nil {
			a.Tracer.Failure("REDACTED",
				"REDACTED", ka.Location, "REDACTED", err)
			continue
		}

		if err := a.appendTowardFreshSegment(ka, segment); err != nil {
			a.Tracer.Failure("REDACTED", "REDACTED", err)
		}
		delete(a.flawedNodes, ka.ID())

		a.Tracer.Details("REDACTED", "REDACTED", ka.Location)
	}
}

//
//
//
func (a *locationRegister) FetchPreference() []*p2p.NetworkLocator {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	registerExtent := a.extent()
	if registerExtent <= 0 {
		if registerExtent < 0 {
			panic(fmt.Sprintf("REDACTED", a.nthFresh+a.nthAged, a.nthFresh, a.nthAged))
		}
		return nil
	}

	countLocators := strongarithmetic.MaximumInteger(
		strongarithmetic.MinimumInteger(minimumFetchPreference, registerExtent),
		registerExtent*fetchPreferenceRatio/100)
	countLocators = strongarithmetic.MinimumInteger(maximumFetchPreference, countLocators)

	//
	//
	everyLocation := make([]*p2p.NetworkLocator, registerExtent)
	i := 0
	for _, ka := range a.locationSearch {
		everyLocation[i] = ka.Location
		i++
	}

	//
	//
	for i := 0; i < countLocators; i++ {
		//
		j := commitrand.Integern(len(everyLocation)-i) + i
		everyLocation[i], everyLocation[j] = everyLocation[j], everyLocation[i]
	}

	//
	return everyLocation[:countLocators]
}

func fractionBelongingCount(p, n int) int {
	return int(math.Round((float64(p) / float64(100)) * float64(n)))
}

//
//
//
//
//
//
//
//
func (a *locationRegister) FetchPreferenceUsingTendency(tendencyTowardFreshLocations int) []*p2p.NetworkLocator {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	registerExtent := a.extent()
	if registerExtent <= 0 {
		if registerExtent < 0 {
			panic(fmt.Sprintf("REDACTED", a.nthFresh+a.nthAged, a.nthFresh, a.nthAged))
		}
		return nil
	}

	if tendencyTowardFreshLocations > 100 {
		tendencyTowardFreshLocations = 100
	}
	if tendencyTowardFreshLocations < 0 {
		tendencyTowardFreshLocations = 0
	}

	countLocators := strongarithmetic.MaximumInteger(
		strongarithmetic.MinimumInteger(minimumFetchPreference, registerExtent),
		registerExtent*fetchPreferenceRatio/100)
	countLocators = strongarithmetic.MinimumInteger(maximumFetchPreference, countLocators)

	//
	//
	countMandatoryFreshAppend := strongarithmetic.MaximumInteger(fractionBelongingCount(tendencyTowardFreshLocations, countLocators), countLocators-a.nthAged)
	preference := a.unpredictableSelectLocators(segmentKindFresh, countMandatoryFreshAppend)
	preference = append(preference, a.unpredictableSelectLocators(segmentKindAged, countLocators-len(preference))...)
	return preference
}

//

//
func (a *locationRegister) Extent() int {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	return a.extent()
}

func (a *locationRegister) extent() int {
	return a.nthFresh + a.nthAged
}

//

//
func (a *locationRegister) Persist() {
	a.persistTowardRecord(a.recordRoute) //
}

func (a *locationRegister) persistProcedure() {
	defer a.wg.Done()

	persistRecordMetronome := time.NewTicker(exportLocatorDuration)
out:
	for {
		select {
		case <-persistRecordMetronome.C:
			a.persistTowardRecord(a.recordRoute)
		case <-a.Exit():
			break out
		}
	}
	persistRecordMetronome.Stop()
	a.persistTowardRecord(a.recordRoute)
}

//

func (a *locationRegister) fetchSegment(segmentKind byte, segmentOffset int) map[string]*recognizedLocator {
	switch segmentKind {
	case segmentKindFresh:
		return a.segmentsFresh[segmentOffset]
	case segmentKindAged:
		return a.segmentsAged[segmentOffset]
	default:
		panic("REDACTED")
	}
}

//
//
func (a *locationRegister) appendTowardFreshSegment(ka *recognizedLocator, segmentOffset int) error {
	//
	if ka.equalsAged() {
		return faultLocationRegisterAgedLocatorFreshSegment{ka.Location, segmentOffset}
	}

	locationTxt := ka.Location.Text()
	segment := a.fetchSegment(segmentKindFresh, segmentOffset)

	//
	if _, ok := segment[locationTxt]; ok {
		return nil
	}

	//
	if len(segment) > freshSegmentExtent {
		a.Tracer.Details("REDACTED")
		a.lapseFresh(segmentOffset)
	}

	//
	segment[locationTxt] = ka
	//
	if ka.appendSegmentPointer(segmentOffset) == 1 {
		a.nthFresh++
	}

	//
	a.locationSearch[ka.ID()] = ka
	return nil
}

//
func (a *locationRegister) appendTowardAgedSegment(ka *recognizedLocator, segmentOffset int) bool {
	//
	if ka.equalsFresh() {
		a.Tracer.Failure(fmt.Sprintf("REDACTED", ka))
		return false
	}
	if len(ka.Segments) != 0 {
		a.Tracer.Failure(fmt.Sprintf("REDACTED", ka))
		return false
	}

	locationTxt := ka.Location.Text()
	segment := a.fetchSegment(segmentKindAged, segmentOffset)

	//
	if _, ok := segment[locationTxt]; ok {
		return true
	}

	//
	if len(segment) > agedSegmentExtent {
		return false
	}

	//
	segment[locationTxt] = ka
	if ka.appendSegmentPointer(segmentOffset) == 1 {
		a.nthAged++
	}

	//
	a.locationSearch[ka.ID()] = ka

	return true
}

func (a *locationRegister) discardOriginatingSegment(ka *recognizedLocator, segmentKind byte, segmentOffset int) {
	if ka.SegmentKind != segmentKind {
		a.Tracer.Failure(fmt.Sprintf("REDACTED", ka))
		return
	}
	segment := a.fetchSegment(segmentKind, segmentOffset)
	delete(segment, ka.Location.Text())
	if ka.discardSegmentPointer(segmentOffset) == 0 {
		if segmentKind == segmentKindFresh {
			a.nthFresh--
		} else {
			a.nthAged--
		}
		delete(a.locationSearch, ka.ID())
	}
}

func (a *locationRegister) discardOriginatingEverySegments(ka *recognizedLocator) {
	for _, segmentOffset := range ka.Segments {
		segment := a.fetchSegment(ka.SegmentKind, segmentOffset)
		delete(segment, ka.Location.Text())
	}
	ka.Segments = nil
	if ka.SegmentKind == segmentKindFresh {
		a.nthFresh--
	} else {
		a.nthAged--
	}
	delete(a.locationSearch, ka.ID())
}

//

func (a *locationRegister) selectSenior(segmentKind byte, segmentOffset int) *recognizedLocator {
	segment := a.fetchSegment(segmentKind, segmentOffset)
	var senior *recognizedLocator
	for _, ka := range segment {
		if senior == nil || ka.FinalEffort.Before(senior.FinalEffort) {
			senior = ka
		}
	}
	return senior
}

//
//
func (a *locationRegister) appendLocator(location, src *p2p.NetworkLocator) error {
	if location == nil || src == nil {
		return FaultLocationRegisterVoidLocation{location, src}
	}

	if err := location.Sound(); err != nil {
		return FaultLocationRegisterUnfitLocation{Location: location, LocationFault: err}
	}

	if _, ok := a.flawedNodes[location.ID]; ok {
		return FaultLocatorProhibited{location}
	}

	if _, ok := a.secludedIDXDstore[location.ID]; ok {
		return FaultLocationRegisterSecluded{location}
	}

	if _, ok := a.secludedIDXDstore[src.ID]; ok {
		return FaultLocationRegisterSecludedOrigin{src}
	}

	//
	if _, ok := a.mineLocations[location.Text()]; ok {
		return FaultLocationRegisterEgo{location}
	}

	if a.reachabilityStringent && !location.Directable() {
		return FaultLocationRegisterUnDirectable{location}
	}

	ka := a.locationSearch[location.ID]
	if ka != nil {
		//
		//
		//
		if ka.equalsAged() && ka.Location.ID == location.ID {
			return nil
		}
		//
		if len(ka.Segments) == maximumFreshSegmentsEveryLocator {
			return nil
		}
		//
		element := int32(2 * len(ka.Segments))
		if a.arbitrary.Integer31n(element) != 0 {
			return nil
		}
	} else {
		ka = freshRecognizedLocator(location, src)
	}

	segment, err := a.computeFreshSegment(location, src)
	if err != nil {
		return err
	}
	return a.appendTowardFreshSegment(ka, segment)
}

func (a *locationRegister) unpredictableSelectLocators(segmentKind byte, num int) []*p2p.NetworkLocator {
	var segments []map[string]*recognizedLocator
	switch segmentKind {
	case segmentKindFresh:
		segments = a.segmentsFresh
	case segmentKindAged:
		segments = a.segmentsAged
	default:
		panic("REDACTED")
	}
	sum := 0
	for _, segment := range segments {
		sum += len(segment)
	}
	locators := make([]*recognizedLocator, 0, sum)
	for _, segment := range segments {
		for _, ka := range segment {
			locators = append(locators, ka)
		}
	}
	preference := make([]*p2p.NetworkLocator, 0, num)
	designatedAssign := make(map[string]bool, num)
	rand.Shuffle(sum, func(i, j int) {
		locators[i], locators[j] = locators[j], locators[i]
	})
	for _, location := range locators {
		if designatedAssign[location.Location.Text()] {
			continue
		}
		designatedAssign[location.Location.Text()] = true
		preference = append(preference, location.Location)
		if len(preference) >= num {
			return preference
		}
	}
	return preference
}

//
//
func (a *locationRegister) lapseFresh(segmentOffset int) {
	for locationTxt, ka := range a.segmentsFresh[segmentOffset] {
		//
		if ka.equalsFlawed() {
			a.Tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", locationTxt))
			a.discardOriginatingSegment(ka, segmentKindFresh, segmentOffset)
			return
		}
	}

	//
	senior := a.selectSenior(segmentKindFresh, segmentOffset)
	a.discardOriginatingSegment(senior, segmentKindFresh, segmentOffset)
}

//
//
//
func (a *locationRegister) shiftTowardAged(ka *recognizedLocator) error {
	//
	if ka.equalsAged() {
		a.Tracer.Failure(fmt.Sprintf("REDACTED", ka))
		return nil
	}
	if len(ka.Segments) == 0 {
		a.Tracer.Failure(fmt.Sprintf("REDACTED", ka))
		return nil
	}

	//
	a.discardOriginatingEverySegments(ka)
	//
	ka.SegmentKind = segmentKindAged

	//
	agedSegmentOffset, err := a.computeAgedSegment(ka.Location)
	if err != nil {
		return err
	}
	appended := a.appendTowardAgedSegment(ka, agedSegmentOffset)
	if !appended {
		//
		senior := a.selectSenior(segmentKindAged, agedSegmentOffset)
		a.discardOriginatingSegment(senior, segmentKindAged, agedSegmentOffset)
		freshSegmentOffset, err := a.computeFreshSegment(senior.Location, senior.Src)
		if err != nil {
			return err
		}
		if err := a.appendTowardFreshSegment(senior, freshSegmentOffset); err != nil {
			a.Tracer.Failure("REDACTED", "REDACTED", err)
		}

		//
		appended = a.appendTowardAgedSegment(ka, agedSegmentOffset)
		if !appended {
			a.Tracer.Failure(fmt.Sprintf("REDACTED", ka, agedSegmentOffset))
		}
	}
	return nil
}

func (a *locationRegister) discardLocator(location *p2p.NetworkLocator) {
	ka := a.locationSearch[location.ID]
	if ka == nil {
		return
	}
	a.Tracer.Details("REDACTED", "REDACTED", location)
	a.discardOriginatingEverySegments(ka)
}

func (a *locationRegister) appendFlawedNode(location *p2p.NetworkLocator, prohibitMoment time.Duration) bool {
	//
	ka := a.locationSearch[location.ID]
	//
	if ka == nil {
		return false
	}

	if _, earlierFlawedNode := a.flawedNodes[location.ID]; !earlierFlawedNode {
		//
		ka.ban(prohibitMoment)
		a.flawedNodes[location.ID] = ka
		a.Tracer.Details("REDACTED", "REDACTED", location)
	}
	return true
}

//
//

//
func (a *locationRegister) computeFreshSegment(location, src *p2p.NetworkLocator) (int, error) {
	item1 := []byte{}
	item1 = append(item1, []byte(a.key)...)
	item1 = append(item1, []byte(a.cohortToken(location))...)
	item1 = append(item1, []byte(a.cohortToken(src))...)
	digest1, err := a.digest(item1)
	if err != nil {
		return 0, err
	}
	hash64 := binary.BigEndian.Uint64(digest1)
	hash64 %= freshSegmentsEveryCohort
	var hashreserve [8]byte
	binary.BigEndian.PutUint64(hashreserve[:], hash64)
	datum2 := []byte{}
	datum2 = append(datum2, []byte(a.key)...)
	datum2 = append(datum2, a.cohortToken(src)...)
	datum2 = append(datum2, hashreserve[:]...)

	digest2, err := a.digest(datum2)
	if err != nil {
		return 0, err
	}
	outcome := int(binary.BigEndian.Uint64(digest2) % freshSegmentTotal)
	return outcome, nil
}

//
func (a *locationRegister) computeAgedSegment(location *p2p.NetworkLocator) (int, error) {
	item1 := []byte{}
	item1 = append(item1, []byte(a.key)...)
	item1 = append(item1, []byte(location.Text())...)
	digest1, err := a.digest(item1)
	if err != nil {
		return 0, err
	}
	hash64 := binary.BigEndian.Uint64(digest1)
	hash64 %= agedSegmentsEveryCohort
	var hashreserve [8]byte
	binary.BigEndian.PutUint64(hashreserve[:], hash64)
	datum2 := []byte{}
	datum2 = append(datum2, []byte(a.key)...)
	datum2 = append(datum2, a.cohortToken(location)...)
	datum2 = append(datum2, hashreserve[:]...)

	digest2, err := a.digest(datum2)
	if err != nil {
		return 0, err
	}
	outcome := int(binary.BigEndian.Uint64(digest2) % agedSegmentTotal)
	return outcome, nil
}

//
//
//
//
func (a *locationRegister) cohortToken(na *p2p.NetworkLocator) string {
	return cohortTokenForeach(na, a.reachabilityStringent)
}

func cohortTokenForeach(na *p2p.NetworkLocator, reachabilityStringent bool) string {
	if reachabilityStringent && na.Regional() {
		return "REDACTED"
	}
	if reachabilityStringent && !na.Directable() {
		return "REDACTED"
	}

	if ipv4 := na.IP.To4(); ipv4 != nil {
		return na.IP.Mask(net.CIDRMask(16, 32)).String()
	}

	if na.Rfc6145() || na.Rfc6052() {
		//
		ip := na.IP[12:16]
		return ip.Mask(net.CIDRMask(16, 32)).String()
	}

	if na.Rfc3964() {
		ip := na.IP[2:6]
		return ip.Mask(net.CIDRMask(16, 32)).String()
	}

	if na.Rfc4380() {
		//
		//
		ip := net.IP(make([]byte, 4))
		for i, octet := range na.IP[12:16] {
			ip[i] = octet ^ 0xff
		}
		return ip.Mask(net.CIDRMask(16, 32)).String()
	}

	if na.OnionConcatenateTor() {
		//
		return fmt.Sprintf("REDACTED", na.IP[6]&((1<<4)-1))
	}

	//
	//
	//
	digits := 32
	himNetwork := &net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(32, 128)}
	if himNetwork.Contains(na.IP) {
		digits = 36
	}
	ipv6mask := net.CIDRMask(digits, 128)
	return na.IP.Mask(ipv6mask).String()
}

func (a *locationRegister) digest(b []byte) ([]byte, error) {
	a.digester.Reset()
	a.digester.Write(b)
	return a.digester.Sum(nil), nil
}
