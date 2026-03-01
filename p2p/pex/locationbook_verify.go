package pex

import (
	"encoding/hex"
	"fmt"
	"math"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

//

func VerifyLocationRegisterSelectLocator(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	//
	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	assert.Zero(t, register.Extent())

	location := register.SelectLocator(50)
	assert.Nil(t, location, "REDACTED")

	arbitraryLocations := arbitraryNetworkLocatorCouples(t, 1)
	locationOrigin := arbitraryLocations[0]
	err := register.AppendLocator(locationOrigin.location, locationOrigin.src)
	require.NoError(t, err)

	//
	location = register.SelectLocator(0)
	assert.NotNil(t, location, "REDACTED")
	location = register.SelectLocator(50)
	assert.NotNil(t, location, "REDACTED")
	location = register.SelectLocator(100)
	assert.NotNil(t, location, "REDACTED")

	//
	register.LabelValid(locationOrigin.location.ID)
	location = register.SelectLocator(0)
	assert.NotNil(t, location, "REDACTED")
	location = register.SelectLocator(50)
	assert.NotNil(t, location, "REDACTED")

	//
	location = register.SelectLocator(100)
	assert.Nil(t, location, "REDACTED")
}

func VerifyLocationRegisterPersistFetch(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	//
	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	register.Persist()

	register = FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	err := register.Initiate()
	require.NoError(t, err)

	assert.True(t, register.Blank())

	//
	arbitraryLocations := arbitraryNetworkLocatorCouples(t, 100)

	for _, locationOrigin := range arbitraryLocations {
		err := register.AppendLocator(locationOrigin.location, locationOrigin.src)
		require.NoError(t, err)
	}

	assert.Equal(t, 100, register.Extent())
	register.Persist()

	register = FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	err = register.Initiate()
	require.NoError(t, err)

	assert.Equal(t, 100, register.Extent())
}

func VerifyLocationRegisterSearch(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	arbitraryLocations := arbitraryNetworkLocatorCouples(t, 100)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	for _, locationOrigin := range arbitraryLocations {
		location := locationOrigin.location
		src := locationOrigin.src
		err := register.AppendLocator(location, src)
		require.NoError(t, err)

		ka := register.OwnsLocation(location)
		assert.True(t, ka, "REDACTED", location)
	}
}

func VerifyLocationRegisterElevateTowardAged(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	arbitraryLocations := arbitraryNetworkLocatorCouples(t, 100)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	for _, locationOrigin := range arbitraryLocations {
		err := register.AppendLocator(locationOrigin.location, locationOrigin.src)
		require.NoError(t, err)
	}

	//
	for _, locationOrigin := range arbitraryLocations {
		register.LabelEffort(locationOrigin.location)
	}

	//
	for i, locationOrigin := range arbitraryLocations {
		if i%2 == 0 {
			register.LabelValid(locationOrigin.location.ID)
		}
	}

	//

	preference := register.FetchPreference()
	t.Logf("REDACTED", preference)

	if len(preference) > register.Extent() {
		t.Errorf("REDACTED")
	}

	preference = register.FetchPreferenceUsingTendency(30)
	t.Logf("REDACTED", preference)

	if len(preference) > register.Extent() {
		t.Errorf("REDACTED")
	}

	assert.Equal(t, register.Extent(), 100, "REDACTED")
}

func VerifyLocationRegisterOverseesReplicas(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())

	arbitraryLocations := arbitraryNetworkLocatorCouples(t, 100)

	distinctOrigin := arbitraryIDXIpv4address(t)
	for _, locationOrigin := range arbitraryLocations {
		err := register.AppendLocator(locationOrigin.location, locationOrigin.src)
		require.NoError(t, err)
		err = register.AppendLocator(locationOrigin.location, locationOrigin.src) //
		require.NoError(t, err)
		err = register.AppendLocator(locationOrigin.location, distinctOrigin) //
		require.NoError(t, err)
	}

	assert.Equal(t, 100, register.Extent())
}

type networkLocatorDuo struct {
	location *p2p.NetworkLocator
	src  *p2p.NetworkLocator
}

func arbitraryNetworkLocatorCouples(t *testing.T, n int) []networkLocatorDuo {
	arbitraryLocations := make([]networkLocatorDuo, n)
	for i := 0; i < n; i++ {
		arbitraryLocations[i] = networkLocatorDuo{location: arbitraryIDXIpv4address(t), src: arbitraryIDXIpv4address(t)}
	}
	return arbitraryLocations
}

func arbitraryIDXIpv4address(t *testing.T) *p2p.NetworkLocator {
	for {
		ip := fmt.Sprintf("REDACTED",
			commitrand.Integern(254)+1,
			commitrand.Integern(255),
			commitrand.Integern(255),
			commitrand.Integern(255),
		)
		channel := commitrand.Integern(65535-1) + 1
		id := p2p.ID(hex.EncodeToString(commitrand.Octets(p2p.UUIDOctetMagnitude)))
		uuidLocation := p2p.UUIDLocationText(id, fmt.Sprintf("REDACTED", ip, channel))
		location, err := p2p.FreshNetworkLocatorText(uuidLocation)
		assert.Nil(t, err, "REDACTED")
		if location.Directable() {
			return location
		}
	}
}

func VerifyLocationRegisterDiscardLocator(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())

	location := arbitraryIDXIpv4address(t)
	err := register.AppendLocator(location, location)
	require.NoError(t, err)
	assert.Equal(t, 1, register.Extent())

	register.DiscardLocator(location)
	assert.Equal(t, 0, register.Extent())

	unCurrentLocation := arbitraryIDXIpv4address(t)
	register.DiscardLocator(unCurrentLocation)
	assert.Equal(t, 0, register.Extent())
}

func VerifyLocationRegisterFetchPreferenceUsingSingleFlaggedValid(t *testing.T) {
	//
	register, filename := generateLocationRegisterUsingModuleAgedAlsoNTHFreshLocations(t, 1, 9)
	defer eraseTransientRecord(filename)

	locations := register.FetchPreferenceUsingTendency(tendencyTowardPreferFreshNodes)
	assert.NotNil(t, locations)
	affirmModuleAgedAlsoNTHFreshLocationsInsidePreference(t, 1, 9, locations, register)
}

func VerifyLocationRegisterFetchPreferenceUsingSingleNegationFlaggedValid(t *testing.T) {
	//
	register, filename := generateLocationRegisterUsingModuleAgedAlsoNTHFreshLocations(t, 9, 1)
	defer eraseTransientRecord(filename)

	locations := register.FetchPreferenceUsingTendency(tendencyTowardPreferFreshNodes)
	assert.NotNil(t, locations)
	affirmModuleAgedAlsoNTHFreshLocationsInsidePreference(t, 9, 1, locations, register)
}

func VerifyLocationRegisterFetchPreferenceYieldsVoidWheneverLocationRegisterEqualsBlank(t *testing.T) {
	register, filename := generateLocationRegisterUsingModuleAgedAlsoNTHFreshLocations(t, 0, 0)
	defer eraseTransientRecord(filename)

	locations := register.FetchPreferenceUsingTendency(tendencyTowardPreferFreshNodes)
	assert.Nil(t, locations)
}

func VerifyLocationRegisterFetchPreference(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())

	//
	assert.Empty(t, register.FetchPreference())

	//
	location := arbitraryIDXIpv4address(t)
	err := register.AppendLocator(location, location)
	require.NoError(t, err)

	assert.Equal(t, 1, len(register.FetchPreference()))
	assert.Equal(t, location, register.FetchPreference()[0])

	//
	arbitraryLocations := arbitraryNetworkLocatorCouples(t, 100)
	for _, locationOrigin := range arbitraryLocations {
		err := register.AppendLocator(locationOrigin.location, locationOrigin.src)
		require.NoError(t, err)
	}

	//
	locations := make(map[string]*p2p.NetworkLocator)
	preference := register.FetchPreference()
	for _, location := range preference {
		if dup, ok := locations[location.Text()]; ok {
			t.Fatalf("REDACTED", preference, dup)
		}
		locations[location.Text()] = location
	}

	if len(preference) > register.Extent() {
		t.Errorf("REDACTED", preference)
	}
}

func VerifyLocationRegisterFetchPreferenceUsingTendency(t *testing.T) {
	const tendencyTowardFreshLocations = 30

	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())

	//
	preference := register.FetchPreferenceUsingTendency(tendencyTowardFreshLocations)
	assert.Empty(t, preference)

	//
	location := arbitraryIDXIpv4address(t)
	err := register.AppendLocator(location, location)
	require.NoError(t, err)

	preference = register.FetchPreferenceUsingTendency(tendencyTowardFreshLocations)
	assert.Equal(t, 1, len(preference))
	assert.Equal(t, location, preference[0])

	//
	arbitraryLocations := arbitraryNetworkLocatorCouples(t, 100)
	for _, locationOrigin := range arbitraryLocations {
		err := register.AppendLocator(locationOrigin.location, locationOrigin.src)
		require.NoError(t, err)
	}

	//
	locations := make(map[string]*p2p.NetworkLocator)
	preference = register.FetchPreferenceUsingTendency(tendencyTowardFreshLocations)
	for _, location := range preference {
		if dup, ok := locations[location.Text()]; ok {
			t.Fatalf("REDACTED", preference, dup)
		}
		locations[location.Text()] = location
	}

	if len(preference) > register.Extent() {
		t.Fatalf("REDACTED", preference)
	}

	//
	arbitraryLocationsLength := len(arbitraryLocations)
	for i, locationOrigin := range arbitraryLocations {
		if int((float64(i)/float64(arbitraryLocationsLength))*100) >= 20 {
			register.LabelValid(locationOrigin.location.ID)
		}
	}

	preference = register.FetchPreferenceUsingTendency(tendencyTowardFreshLocations)

	//
	valid := 0
	for _, location := range preference {
		if register.EqualsValid(location) {
			valid++
		}
	}

	got, anticipated := int((float64(valid)/float64(len(preference)))*100), 100-tendencyTowardFreshLocations

	//
	leeway := int(math.Round(float64(100) / float64(len(preference))))
	if got > anticipated+leeway {
		t.Fatalf(
			"REDACTED",
			got,
			anticipated,
			valid,
			len(preference),
		)
	}
	if got < anticipated-leeway {
		t.Fatalf(
			"REDACTED",
			got,
			anticipated,
			valid,
			len(preference),
		)
	}
}

func VerifyLocationRegisterOwnsLocator(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	location := arbitraryIDXIpv4address(t)
	err := register.AppendLocator(location, location)
	require.NoError(t, err)

	assert.True(t, register.OwnsLocation(location))

	register.DiscardLocator(location)

	assert.False(t, register.OwnsLocation(location))
}

func verifyGenerateSecludedLocations(t *testing.T, countLocations int) ([]*p2p.NetworkLocator, []string) {
	locations := make([]*p2p.NetworkLocator, countLocations)
	for i := 0; i < countLocations; i++ {
		locations[i] = arbitraryIDXIpv4address(t)
	}

	secluded := make([]string, countLocations)
	for i, location := range locations {
		secluded[i] = string(location.ID)
	}
	return locations, secluded
}

func VerifyProhibitFlawedNodes(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())

	location := arbitraryIDXIpv4address(t)
	_ = register.AppendLocator(location, location)

	register.LabelFlawed(location, 1*time.Second)
	//
	assert.False(t, register.OwnsLocation(location))
	assert.True(t, register.EqualsProhibited(location))

	err := register.AppendLocator(location, location)
	//
	require.Error(t, err)

	time.Sleep(1 * time.Second)
	register.RestoreFlawedNodes()
	//
	assert.EqualValues(t, 1, register.Extent())
	assert.True(t, register.OwnsLocation(location))
	assert.False(t, register.EqualsValid(location))
}

func VerifyLocationRegisterBlank(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	//
	require.True(t, register.Blank())
	//
	register.AppendMineLocator(arbitraryIDXIpv4address(t))
	require.True(t, register.Blank())
	//
	_, secludedIndexes := verifyGenerateSecludedLocations(t, 5)
	register.AppendSecludedIDXDstore(secludedIndexes)
	require.True(t, register.Blank())

	//
	err := register.AppendLocator(arbitraryIDXIpv4address(t), arbitraryIDXIpv4address(t))
	require.NoError(t, err)
	require.False(t, register.Blank())
}

func VerifySecludedNodes(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())

	locations, secluded := verifyGenerateSecludedLocations(t, 10)
	register.AppendSecludedIDXDstore(secluded)

	//
	for _, location := range locations {
		err := register.AppendLocator(location, location)
		if assert.Error(t, err) {
			_, ok := err.(FaultLocationRegisterSecluded)
			assert.True(t, ok)
		}
	}

	//
	err := register.AppendLocator(arbitraryIDXIpv4address(t), locations[0])
	if assert.Error(t, err) {
		_, ok := err.(FaultLocationRegisterSecludedOrigin)
		assert.True(t, ok)
	}
}

func verifyLocationRegisterLocatorPreference(t *testing.T, registerExtent int) {
	//
	for nthRegisterAged := 0; nthRegisterAged <= registerExtent; nthRegisterAged++ {
		nthRegisterFresh := registerExtent - nthRegisterAged
		traceTxt := fmt.Sprintf("REDACTED", registerExtent, nthRegisterFresh, nthRegisterAged)

		//
		register, filename := generateLocationRegisterUsingModuleAgedAlsoNTHFreshLocations(t, nthRegisterAged, nthRegisterFresh)
		defer eraseTransientRecord(filename)
		locations := register.FetchPreferenceUsingTendency(tendencyTowardPreferFreshNodes)
		assert.NotNil(t, locations, "REDACTED", traceTxt)
		nthLocations := len(locations)
		assert.NotZero(t, nthLocations, "REDACTED", traceTxt)

		//
		for _, location := range locations {
			if location == nil {
				t.Fatalf("REDACTED", traceTxt, locations)
			}
		}

		//
		nthAged, nthFresh := totalAgedAlsoFreshLocationsInsidePreference(locations, register)

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
		var (
			k      = fractionBelongingCount(tendencyTowardPreferFreshNodes, nthLocations)
			expirationFresh = strongarithmetic.MinimumInteger(nthFresh, strongarithmetic.MaximumInteger(k, nthLocations-nthRegisterAged))
			expirationAged = strongarithmetic.MinimumInteger(nthAged, nthLocations-expirationFresh)
		)

		//
		if nthFresh != expirationFresh {
			t.Fatalf("REDACTED", traceTxt, expirationFresh, nthFresh)
		}
		if nthAged != expirationAged {
			t.Fatalf("REDACTED", traceTxt, expirationAged, nthAged)
		}

		//
		//
		orderMagnitudes, orderKinds, err := scrutinizePreferenceSchema(register, locations)
		assert.NoError(t, err, "REDACTED", traceTxt)

		//
		//
		//
		var expirationOrderMagnitudes []int
		var expirationOrderKinds []int

		switch {
		case expirationAged == 0: //
			expirationOrderMagnitudes = []int{nthLocations}
			expirationOrderKinds = []int{1}
		case expirationFresh == 0: //
			expirationOrderMagnitudes = []int{nthLocations}
			expirationOrderKinds = []int{2}
		case nthLocations-expirationFresh-expirationAged == 0: //
			expirationOrderMagnitudes = []int{expirationFresh, expirationAged}
			expirationOrderKinds = []int{1, 2}
		}

		assert.Equal(t, expirationOrderMagnitudes, orderMagnitudes,
			"REDACTED",
			traceTxt, expirationOrderMagnitudes, orderMagnitudes)
		assert.Equal(t, expirationOrderKinds, orderKinds,
			"REDACTED",
			traceTxt, expirationOrderKinds, orderKinds)
	}
}

func VerifyVariousLocationRegisterLocatorPreference(t *testing.T) {
	//
	const N = 32
	for registerExtent := 1; registerExtent < N; registerExtent++ {
		verifyLocationRegisterLocatorPreference(t, registerExtent)
	}

	//
	extents := [...][]int{{33, 100}, {100, 175}}
	registerExtents := make([]int, 0, len(extents))
	for _, r := range extents {
		registerExtents = append(registerExtents, commitrand.Integern(r[1]-r[0])+r[0])
	}
	t.Logf("REDACTED", registerExtents)
	for _, registerExtent := range registerExtents {
		verifyLocationRegisterLocatorPreference(t, registerExtent)
	}
}

func VerifyLocationRegisterAppendExecutesNegationSupersedeAgedINET(t *testing.T) {
	filename := generateTransientRecordAlias("REDACTED")
	defer eraseTransientRecord(filename)

	//
	//
	//
	nodeUUID := "REDACTED"
	nodeActualINET := "REDACTED"
	nodeSupplantEffortINET := "REDACTED"
	OriginLocation := "REDACTED"

	//
	//
	//
	countSupplantEndeavors := 10

	nodeActualLocation, err := p2p.FreshNetworkLocatorText(nodeUUID + "REDACTED" + nodeActualINET)
	require.Nil(t, err)

	nodeSupplantEffortLocation, err := p2p.FreshNetworkLocatorText(nodeUUID + "REDACTED" + nodeSupplantEffortINET)
	require.Nil(t, err)

	src, err := p2p.FreshNetworkLocatorText(OriginLocation)
	require.Nil(t, err)

	register := FreshLocationRegister(filename, true)
	register.AssignTracer(log.VerifyingTracer())
	err = register.AppendLocator(nodeActualLocation, src)
	require.Nil(t, err)
	register.LabelEffort(nodeActualLocation)
	register.LabelValid(nodeActualLocation.ID)

	//
	err = register.AppendLocator(nodeActualLocation, src)
	require.Nil(t, err)

	//
	//
	for i := 0; i < countSupplantEndeavors; i++ {
		err = register.AppendLocator(nodeSupplantEffortLocation, src)
		require.Nil(t, err)
	}
	//
	//
	//
	//
	for i := 0; i < countSupplantEndeavors; i++ {
		preference := register.FetchPreference()
		for _, location := range preference {
			require.Equal(t, location.IP, nodeActualLocation.IP)
		}
	}
}

func VerifyLocationRegisterCohortToken(t *testing.T) {
	//
	verifyScenarios := []struct {
		alias   string
		ip     string
		expirationToken string
	}{
		//
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},

		//
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},

		//
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},

		//
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
	}

	for i, tc := range verifyScenarios {
		nip := net.ParseIP(tc.ip)
		key := cohortTokenForeach(p2p.FreshNetworkLocatorINETChannel(nip, 26656), false)
		assert.Equal(t, tc.expirationToken, key, "REDACTED", i)
	}

	//
	verifyScenarios = []struct {
		alias   string
		ip     string
		expirationToken string
	}{
		//
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},

		//
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
	}

	for i, tc := range verifyScenarios {
		nip := net.ParseIP(tc.ip)
		key := cohortTokenForeach(p2p.FreshNetworkLocatorINETChannel(nip, 26656), true)
		assert.Equal(t, tc.expirationToken, key, "REDACTED", i)
	}
}

func affirmModuleAgedAlsoNTHFreshLocationsInsidePreference(t *testing.T, m, n int, locations []*p2p.NetworkLocator, register *locationRegister) {
	nthAged, nthFresh := totalAgedAlsoFreshLocationsInsidePreference(locations, register)
	assert.Equal(t, m, nthAged, "REDACTED")
	assert.Equal(t, n, nthFresh, "REDACTED")
}

func generateTransientRecordAlias(heading string) string {
	f, err := os.CreateTemp("REDACTED", heading)
	if err != nil {
		panic(err)
	}
	filename := f.Name()
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return filename
}

func eraseTransientRecord(filename string) {
	err := os.Remove(filename)
	if err != nil {
		panic(err)
	}
}

func generateLocationRegisterUsingModuleAgedAlsoNTHFreshLocations(t *testing.T, nthAged, nthFresh int) (register *locationRegister, filename string) {
	filename = generateTransientRecordAlias("REDACTED")

	register = FreshLocationRegister(filename, true).(*locationRegister)
	register.AssignTracer(log.VerifyingTracer())
	assert.Zero(t, register.Extent())

	arbitraryLocations := arbitraryNetworkLocatorCouples(t, nthAged)
	for _, location := range arbitraryLocations {
		err := register.AppendLocator(location.location, location.src)
		require.NoError(t, err)
		register.LabelValid(location.location.ID)
	}

	arbitraryLocations = arbitraryNetworkLocatorCouples(t, nthFresh)
	for _, location := range arbitraryLocations {
		err := register.AppendLocator(location.location, location.src)
		require.NoError(t, err)
	}

	return
}

func totalAgedAlsoFreshLocationsInsidePreference(locations []*p2p.NetworkLocator, register *locationRegister) (nthAged, nthFresh int) {
	for _, location := range locations {
		if register.EqualsValid(location) {
			nthAged++
		} else {
			nthFresh++
		}
	}
	return
}

//
//
//
//
func scrutinizePreferenceSchema(register *locationRegister, locations []*p2p.NetworkLocator) (orderMagnitudes, orderKinds []int, err error) {
	//
	var (
		previousKind      = 0
		prevailingOrderLength = 0
	)

	for _, location := range locations {
		locationKind := 0
		if register.EqualsValid(location) {
			locationKind = 2
		} else {
			locationKind = 1
		}
		if locationKind != previousKind && previousKind != 0 {
			orderMagnitudes = append(orderMagnitudes, prevailingOrderLength)
			orderKinds = append(orderKinds, previousKind)
			prevailingOrderLength = 0
		}
		prevailingOrderLength++
		previousKind = locationKind
	}

	orderMagnitudes = append(orderMagnitudes, prevailingOrderLength)
	orderKinds = append(orderKinds, previousKind)

	return
}
