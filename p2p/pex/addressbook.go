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

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/utils/log"
	cometmath "github.com/valkyrieworks/utils/math"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
)

const (
	segmentKindNew = 0x01
	segmentKindAged = 0x02
)

//
//
//
//
type AddressLedger interface {
	daemon.Daemon

	//
	AppendOurLocation(*p2p.NetLocation)
	//
	OurLocation(*p2p.NetLocation) bool

	AppendInternalIDXDatastore([]string)

	//
	AppendLocation(address *p2p.NetLocation, src *p2p.NetLocation) error
	DeleteLocation(*p2p.NetLocation)

	//
	HasLocation(*p2p.NetLocation) bool

	//
	RequireAdditionalLocations() bool
	//
	//
	Empty() bool

	//
	SelectLocation(tendencyTowardNewLocations int) *p2p.NetLocation

	//
	StampValid(p2p.ID)
	StampEndeavor(*p2p.NetLocation)
	StampFlawed(*p2p.NetLocation, time.Duration) //
	//
	RestoreFlawedNodes()

	IsValid(*p2p.NetLocation) bool
	IsProhibited(*p2p.NetLocation) bool

	//
	FetchPreference() []*p2p.NetLocation
	//
	FetchPreferenceWithTendency(tendencyTowardNewLocations int) []*p2p.NetLocation

	Volume() int

	//
	Persist()
}

var _ AddressLedger = (*addressLedger)(nil)

//
//
type addressLedger struct {
	daemon.RootDaemon

	//
	mtx        engineconnect.Lock
	random       *engineseed.Random
	ourLocations   map[string]struct{}
	internalIDXDatastore map[p2p.ID]struct{}
	addressSearch map[p2p.ID]*recognizedLocation //
	flawedNodes   map[p2p.ID]*recognizedLocation //
	segmentsAged []map[string]*recognizedLocation
	segmentsNew []map[string]*recognizedLocation
	nAged       int
	nNew       int

	//
	entryRoute          string
	key               string //
	forwardingPrecise bool
	digester            hash.Hash64

	wg sync.WaitGroup
}

func shouldNewDigester() hash.Hash64 {
	key := vault.CRandomOctets(highwayhash.Size)
	digester, err := highwayhash.New64(key)
	if err != nil {
		panic(err)
	}
	return digester
}

//
//
func NewAddressRegistry(entryRoute string, forwardingPrecise bool) AddressLedger {
	am := &addressLedger{
		random:              engineseed.NewRandom(),
		ourLocations:          make(map[string]struct{}),
		internalIDXDatastore:        make(map[p2p.ID]struct{}),
		addressSearch:        make(map[p2p.ID]*recognizedLocation),
		flawedNodes:          make(map[p2p.ID]*recognizedLocation),
		entryRoute:          entryRoute,
		forwardingPrecise: forwardingPrecise,
	}
	am.init()
	am.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", am)
	return am
}

//
//
func (a *addressLedger) init() {
	a.key = vault.CRandomHex(24) //
	//
	a.segmentsNew = make([]map[string]*recognizedLocation, newSegmentNumber)
	for i := range a.segmentsNew {
		a.segmentsNew[i] = make(map[string]*recognizedLocation)
	}
	//
	a.segmentsAged = make([]map[string]*recognizedLocation, agedSegmentNumber)
	for i := range a.segmentsAged {
		a.segmentsAged[i] = make(map[string]*recognizedLocation)
	}
	a.digester = shouldNewDigester()
}

//
func (a *addressLedger) OnBegin() error {
	if err := a.RootDaemon.OnBegin(); err != nil {
		return err
	}
	a.importFromEntry(a.entryRoute)

	//
	//
	a.wg.Add(1)
	go a.persistProcedure()

	return nil
}

//
func (a *addressLedger) OnHalt() {
	a.RootDaemon.OnHalt()
}

func (a *addressLedger) Wait() {
	a.wg.Wait()
}

func (a *addressLedger) EntryRoute() string {
	return a.entryRoute
}

//

//
func (a *addressLedger) AppendOurLocation(address *p2p.NetLocation) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.Tracer.Details("REDACTED", "REDACTED", address)
	a.ourLocations[address.String()] = struct{}{}
}

//
func (a *addressLedger) OurLocation(address *p2p.NetLocation) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	_, ok := a.ourLocations[address.String()]
	return ok
}

func (a *addressLedger) AppendInternalIDXDatastore(ids []string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	for _, id := range ids {
		a.internalIDXDatastore[p2p.ID(id)] = struct{}{}
	}
}

//
//
//
//
func (a *addressLedger) AppendLocation(address *p2p.NetLocation, src *p2p.NetLocation) error {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	return a.appendLocation(address, src)
}

//
func (a *addressLedger) DeleteLocation(address *p2p.NetLocation) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.deleteLocation(address)
}

//
//
func (a *addressLedger) IsValid(address *p2p.NetLocation) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	return a.addressSearch[address.ID].isAged()
}

//
func (a *addressLedger) IsProhibited(address *p2p.NetLocation) bool {
	a.mtx.Lock()
	_, ok := a.flawedNodes[address.ID]
	a.mtx.Unlock()

	return ok
}

//
func (a *addressLedger) HasLocation(address *p2p.NetLocation) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	ka := a.addressSearch[address.ID]
	return ka != nil
}

//
func (a *addressLedger) RequireAdditionalLocations() bool {
	return a.Volume() < requireLocationLimit
}

//
//
func (a *addressLedger) Empty() bool {
	return a.Volume() == 0
}

//
//
//
//
//
//
func (a *addressLedger) SelectLocation(tendencyTowardNewLocations int) *p2p.NetLocation {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	registryVolume := a.volume()
	if registryVolume <= 0 {
		if registryVolume < 0 {
			panic(fmt.Sprintf("REDACTED", a.nNew+a.nAged, a.nNew, a.nAged))
		}
		return nil
	}
	if tendencyTowardNewLocations > 100 {
		tendencyTowardNewLocations = 100
	}
	if tendencyTowardNewLocations < 0 {
		tendencyTowardNewLocations = 0
	}

	//
	agedRelation := math.Sqrt(float64(a.nAged)) * (100.0 - float64(tendencyTowardNewLocations))
	newRelation := math.Sqrt(float64(a.nNew)) * float64(tendencyTowardNewLocations)

	//
	var segment map[string]*recognizedLocation
	selectFromAgedSegment := (newRelation+agedRelation)*a.random.Float64() < agedRelation
	if (selectFromAgedSegment && a.nAged == 0) ||
		(!selectFromAgedSegment && a.nNew == 0) {
		return nil
	}
	//
	for len(segment) == 0 {
		if selectFromAgedSegment {
			segment = a.segmentsAged[a.random.Intn(len(a.segmentsAged))]
		} else {
			segment = a.segmentsNew[a.random.Intn(len(a.segmentsNew))]
		}
	}
	//
	randomOrdinal := a.random.Intn(len(segment))
	for _, ka := range segment {
		if randomOrdinal == 0 {
			return ka.Address
		}
		randomOrdinal--
	}
	return nil
}

//
//
func (a *addressLedger) StampValid(id p2p.ID) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	ka := a.addressSearch[id]
	if ka == nil {
		return
	}
	ka.stampValid()
	if ka.isNew() {
		if err := a.shiftToAged(ka); err != nil {
			a.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
}

//
func (a *addressLedger) StampEndeavor(address *p2p.NetLocation) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	ka := a.addressSearch[address.ID]
	if ka == nil {
		return
	}
	ka.stampEndeavor()
}

//
//
func (a *addressLedger) StampFlawed(address *p2p.NetLocation, prohibitTime time.Duration) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	if a.appendFlawedNode(address, prohibitTime) {
		a.deleteLocation(address)
	}
}

//
//
func (a *addressLedger) RestoreFlawedNodes() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	for _, ka := range a.flawedNodes {
		if ka.isProhibited() {
			continue
		}

		segment, err := a.computeNewSegment(ka.Address, ka.Src)
		if err != nil {
			a.Tracer.Fault("REDACTED",
				"REDACTED", ka.Address, "REDACTED", err)
			continue
		}

		if err := a.appendToNewSegment(ka, segment); err != nil {
			a.Tracer.Fault("REDACTED", "REDACTED", err)
		}
		delete(a.flawedNodes, ka.ID())

		a.Tracer.Details("REDACTED", "REDACTED", ka.Address)
	}
}

//
//
//
func (a *addressLedger) FetchPreference() []*p2p.NetLocation {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	registryVolume := a.volume()
	if registryVolume <= 0 {
		if registryVolume < 0 {
			panic(fmt.Sprintf("REDACTED", a.nNew+a.nAged, a.nNew, a.nAged))
		}
		return nil
	}

	countAddresses := cometmath.MaximumInteger(
		cometmath.MinimumInteger(minimumFetchPreference, registryVolume),
		registryVolume*fetchPreferenceFraction/100)
	countAddresses = cometmath.MinimumInteger(maximumFetchPreference, countAddresses)

	//
	//
	allAddress := make([]*p2p.NetLocation, registryVolume)
	i := 0
	for _, ka := range a.addressSearch {
		allAddress[i] = ka.Address
		i++
	}

	//
	//
	for i := 0; i < countAddresses; i++ {
		//
		j := engineseed.Intn(len(allAddress)-i) + i
		allAddress[i], allAddress[j] = allAddress[j], allAddress[i]
	}

	//
	return allAddress[:countAddresses]
}

func ratioOfCount(p, n int) int {
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
func (a *addressLedger) FetchPreferenceWithTendency(tendencyTowardNewLocations int) []*p2p.NetLocation {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	registryVolume := a.volume()
	if registryVolume <= 0 {
		if registryVolume < 0 {
			panic(fmt.Sprintf("REDACTED", a.nNew+a.nAged, a.nNew, a.nAged))
		}
		return nil
	}

	if tendencyTowardNewLocations > 100 {
		tendencyTowardNewLocations = 100
	}
	if tendencyTowardNewLocations < 0 {
		tendencyTowardNewLocations = 0
	}

	countAddresses := cometmath.MaximumInteger(
		cometmath.MinimumInteger(minimumFetchPreference, registryVolume),
		registryVolume*fetchPreferenceFraction/100)
	countAddresses = cometmath.MinimumInteger(maximumFetchPreference, countAddresses)

	//
	//
	countMandatoryNewAppend := cometmath.MaximumInteger(ratioOfCount(tendencyTowardNewLocations, countAddresses), countAddresses-a.nAged)
	preference := a.arbitrarySelectAddresses(segmentKindNew, countMandatoryNewAppend)
	preference = append(preference, a.arbitrarySelectAddresses(segmentKindAged, countAddresses-len(preference))...)
	return preference
}

//

//
func (a *addressLedger) Volume() int {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	return a.volume()
}

func (a *addressLedger) volume() int {
	return a.nNew + a.nAged
}

//

//
func (a *addressLedger) Persist() {
	a.persistToEntry(a.entryRoute) //
}

func (a *addressLedger) persistProcedure() {
	defer a.wg.Done()

	persistEntryTimer := time.NewTicker(exportLocationCadence)
out:
	for {
		select {
		case <-persistEntryTimer.C:
			a.persistToEntry(a.entryRoute)
		case <-a.Exit():
			break out
		}
	}
	persistEntryTimer.Stop()
	a.persistToEntry(a.entryRoute)
}

//

func (a *addressLedger) fetchSegment(segmentKind byte, segmentIndex int) map[string]*recognizedLocation {
	switch segmentKind {
	case segmentKindNew:
		return a.segmentsNew[segmentIndex]
	case segmentKindAged:
		return a.segmentsAged[segmentIndex]
	default:
		panic("REDACTED")
	}
}

//
//
func (a *addressLedger) appendToNewSegment(ka *recognizedLocation, segmentIndex int) error {
	//
	if ka.isAged() {
		return errAddressRegistryAgedLocationNewSegment{ka.Address, segmentIndex}
	}

	addressStr := ka.Address.String()
	segment := a.fetchSegment(segmentKindNew, segmentIndex)

	//
	if _, ok := segment[addressStr]; ok {
		return nil
	}

	//
	if len(segment) > newSegmentVolume {
		a.Tracer.Details("REDACTED")
		a.deactivateNew(segmentIndex)
	}

	//
	segment[addressStr] = ka
	//
	if ka.appendSegmentReference(segmentIndex) == 1 {
		a.nNew++
	}

	//
	a.addressSearch[ka.ID()] = ka
	return nil
}

//
func (a *addressLedger) appendToAgedSegment(ka *recognizedLocation, segmentIndex int) bool {
	//
	if ka.isNew() {
		a.Tracer.Fault(fmt.Sprintf("REDACTED", ka))
		return false
	}
	if len(ka.Segments) != 0 {
		a.Tracer.Fault(fmt.Sprintf("REDACTED", ka))
		return false
	}

	addressStr := ka.Address.String()
	segment := a.fetchSegment(segmentKindAged, segmentIndex)

	//
	if _, ok := segment[addressStr]; ok {
		return true
	}

	//
	if len(segment) > agedSegmentVolume {
		return false
	}

	//
	segment[addressStr] = ka
	if ka.appendSegmentReference(segmentIndex) == 1 {
		a.nAged++
	}

	//
	a.addressSearch[ka.ID()] = ka

	return true
}

func (a *addressLedger) deleteFromSegment(ka *recognizedLocation, segmentKind byte, segmentIndex int) {
	if ka.SegmentKind != segmentKind {
		a.Tracer.Fault(fmt.Sprintf("REDACTED", ka))
		return
	}
	segment := a.fetchSegment(segmentKind, segmentIndex)
	delete(segment, ka.Address.String())
	if ka.deleteSegmentReference(segmentIndex) == 0 {
		if segmentKind == segmentKindNew {
			a.nNew--
		} else {
			a.nAged--
		}
		delete(a.addressSearch, ka.ID())
	}
}

func (a *addressLedger) deleteFromAllSegments(ka *recognizedLocation) {
	for _, segmentIndex := range ka.Segments {
		segment := a.fetchSegment(ka.SegmentKind, segmentIndex)
		delete(segment, ka.Address.String())
	}
	ka.Segments = nil
	if ka.SegmentKind == segmentKindNew {
		a.nNew--
	} else {
		a.nAged--
	}
	delete(a.addressSearch, ka.ID())
}

//

func (a *addressLedger) selectEarliest(segmentKind byte, segmentIndex int) *recognizedLocation {
	segment := a.fetchSegment(segmentKind, segmentIndex)
	var earliest *recognizedLocation
	for _, ka := range segment {
		if earliest == nil || ka.FinalEndeavor.Before(earliest.FinalEndeavor) {
			earliest = ka
		}
	}
	return earliest
}

//
//
func (a *addressLedger) appendLocation(address, src *p2p.NetLocation) error {
	if address == nil || src == nil {
		return ErrAddressRegistryNullAddress{address, src}
	}

	if err := address.Sound(); err != nil {
		return ErrAddressRegistryCorruptAddress{Address: address, AddressErr: err}
	}

	if _, ok := a.flawedNodes[address.ID]; ok {
		return ErrLocationProhibited{address}
	}

	if _, ok := a.internalIDXDatastore[address.ID]; ok {
		return ErrAddressRegistryInternal{address}
	}

	if _, ok := a.internalIDXDatastore[src.ID]; ok {
		return ErrAddressRegistryInternalOrigin{src}
	}

	//
	if _, ok := a.ourLocations[address.String()]; ok {
		return ErrAddressRegistryEgo{address}
	}

	if a.forwardingPrecise && !address.Forwardable() {
		return ErrAddressRegistryNotForwardable{address}
	}

	ka := a.addressSearch[address.ID]
	if ka != nil {
		//
		//
		//
		if ka.isAged() && ka.Address.ID == address.ID {
			return nil
		}
		//
		if len(ka.Segments) == maximumNewSegmentsEachLocation {
			return nil
		}
		//
		coefficient := int32(2 * len(ka.Segments))
		if a.random.Int31n(coefficient) != 0 {
			return nil
		}
	} else {
		ka = newRecognizedLocation(address, src)
	}

	segment, err := a.computeNewSegment(address, src)
	if err != nil {
		return err
	}
	return a.appendToNewSegment(ka, segment)
}

func (a *addressLedger) arbitrarySelectAddresses(segmentKind byte, num int) []*p2p.NetLocation {
	var segments []map[string]*recognizedLocation
	switch segmentKind {
	case segmentKindNew:
		segments = a.segmentsNew
	case segmentKindAged:
		segments = a.segmentsAged
	default:
		panic("REDACTED")
	}
	sum := 0
	for _, segment := range segments {
		sum += len(segment)
	}
	locations := make([]*recognizedLocation, 0, sum)
	for _, segment := range segments {
		for _, ka := range segment {
			locations = append(locations, ka)
		}
	}
	preference := make([]*p2p.NetLocation, 0, num)
	designatedCollection := make(map[string]bool, num)
	rand.Shuffle(sum, func(i, j int) {
		locations[i], locations[j] = locations[j], locations[i]
	})
	for _, address := range locations {
		if designatedCollection[address.Address.String()] {
			continue
		}
		designatedCollection[address.Address.String()] = true
		preference = append(preference, address.Address)
		if len(preference) >= num {
			return preference
		}
	}
	return preference
}

//
//
func (a *addressLedger) deactivateNew(segmentIndex int) {
	for addressStr, ka := range a.segmentsNew[segmentIndex] {
		//
		if ka.isFlawed() {
			a.Tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", addressStr))
			a.deleteFromSegment(ka, segmentKindNew, segmentIndex)
			return
		}
	}

	//
	earliest := a.selectEarliest(segmentKindNew, segmentIndex)
	a.deleteFromSegment(earliest, segmentKindNew, segmentIndex)
}

//
//
//
func (a *addressLedger) shiftToAged(ka *recognizedLocation) error {
	//
	if ka.isAged() {
		a.Tracer.Fault(fmt.Sprintf("REDACTED", ka))
		return nil
	}
	if len(ka.Segments) == 0 {
		a.Tracer.Fault(fmt.Sprintf("REDACTED", ka))
		return nil
	}

	//
	a.deleteFromAllSegments(ka)
	//
	ka.SegmentKind = segmentKindAged

	//
	agedSegmentIndex, err := a.computeAgedSegment(ka.Address)
	if err != nil {
		return err
	}
	appended := a.appendToAgedSegment(ka, agedSegmentIndex)
	if !appended {
		//
		earliest := a.selectEarliest(segmentKindAged, agedSegmentIndex)
		a.deleteFromSegment(earliest, segmentKindAged, agedSegmentIndex)
		newSegmentIndex, err := a.computeNewSegment(earliest.Address, earliest.Src)
		if err != nil {
			return err
		}
		if err := a.appendToNewSegment(earliest, newSegmentIndex); err != nil {
			a.Tracer.Fault("REDACTED", "REDACTED", err)
		}

		//
		appended = a.appendToAgedSegment(ka, agedSegmentIndex)
		if !appended {
			a.Tracer.Fault(fmt.Sprintf("REDACTED", ka, agedSegmentIndex))
		}
	}
	return nil
}

func (a *addressLedger) deleteLocation(address *p2p.NetLocation) {
	ka := a.addressSearch[address.ID]
	if ka == nil {
		return
	}
	a.Tracer.Details("REDACTED", "REDACTED", address)
	a.deleteFromAllSegments(ka)
}

func (a *addressLedger) appendFlawedNode(address *p2p.NetLocation, prohibitTime time.Duration) bool {
	//
	ka := a.addressSearch[address.ID]
	//
	if ka == nil {
		return false
	}

	if _, yetFlawedNode := a.flawedNodes[address.ID]; !yetFlawedNode {
		//
		ka.ban(prohibitTime)
		a.flawedNodes[address.ID] = ka
		a.Tracer.Details("REDACTED", "REDACTED", address)
	}
	return true
}

//
//

//
func (a *addressLedger) computeNewSegment(address, src *p2p.NetLocation) (int, error) {
	data1 := []byte{}
	data1 = append(data1, []byte(a.key)...)
	data1 = append(data1, []byte(a.clusterKey(address))...)
	data1 = append(data1, []byte(a.clusterKey(src))...)
	digest1, err := a.digest(data1)
	if err != nil {
		return 0, err
	}
	hash64 := binary.BigEndian.Uint64(digest1)
	hash64 %= newSegmentsEachCluster
	var hashbuf [8]byte
	binary.BigEndian.PutUint64(hashbuf[:], hash64)
	data2 := []byte{}
	data2 = append(data2, []byte(a.key)...)
	data2 = append(data2, a.clusterKey(src)...)
	data2 = append(data2, hashbuf[:]...)

	digest2, err := a.digest(data2)
	if err != nil {
		return 0, err
	}
	outcome := int(binary.BigEndian.Uint64(digest2) % newSegmentNumber)
	return outcome, nil
}

//
func (a *addressLedger) computeAgedSegment(address *p2p.NetLocation) (int, error) {
	data1 := []byte{}
	data1 = append(data1, []byte(a.key)...)
	data1 = append(data1, []byte(address.String())...)
	digest1, err := a.digest(data1)
	if err != nil {
		return 0, err
	}
	hash64 := binary.BigEndian.Uint64(digest1)
	hash64 %= agedSegmentsEachCluster
	var hashbuf [8]byte
	binary.BigEndian.PutUint64(hashbuf[:], hash64)
	data2 := []byte{}
	data2 = append(data2, []byte(a.key)...)
	data2 = append(data2, a.clusterKey(address)...)
	data2 = append(data2, hashbuf[:]...)

	digest2, err := a.digest(data2)
	if err != nil {
		return 0, err
	}
	outcome := int(binary.BigEndian.Uint64(digest2) % agedSegmentNumber)
	return outcome, nil
}

//
//
//
//
func (a *addressLedger) clusterKey(na *p2p.NetLocation) string {
	return clusterKeyFor(na, a.forwardingPrecise)
}

func clusterKeyFor(na *p2p.NetLocation, forwardingPrecise bool) string {
	if forwardingPrecise && na.Native() {
		return "REDACTED"
	}
	if forwardingPrecise && !na.Forwardable() {
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

	if na.OnionCatTor() {
		//
		return fmt.Sprintf("REDACTED", na.IP[6]&((1<<4)-1))
	}

	//
	//
	//
	bits := 32
	heNet := &net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(32, 128)}
	if heNet.Contains(na.IP) {
		bits = 36
	}
	ipv6mask := net.CIDRMask(bits, 128)
	return na.IP.Mask(ipv6mask).String()
}

func (a *addressLedger) digest(b []byte) ([]byte, error) {
	a.digester.Reset()
	a.digester.Write(b)
	return a.digester.Sum(nil), nil
}
