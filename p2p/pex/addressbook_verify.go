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

	"github.com/valkyrieworks/utils/log"
	cometmath "github.com/valkyrieworks/utils/math"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
)

//

func VerifyAddressRegistrySelectLocation(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	//
	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	assert.Zero(t, registry.Volume())

	address := registry.SelectLocation(50)
	assert.Nil(t, address, "REDACTED")

	randomLocations := randomNetLocationCouples(t, 1)
	addressOrigin := randomLocations[0]
	err := registry.AppendLocation(addressOrigin.address, addressOrigin.src)
	require.NoError(t, err)

	//
	address = registry.SelectLocation(0)
	assert.NotNil(t, address, "REDACTED")
	address = registry.SelectLocation(50)
	assert.NotNil(t, address, "REDACTED")
	address = registry.SelectLocation(100)
	assert.NotNil(t, address, "REDACTED")

	//
	registry.StampValid(addressOrigin.address.ID)
	address = registry.SelectLocation(0)
	assert.NotNil(t, address, "REDACTED")
	address = registry.SelectLocation(50)
	assert.NotNil(t, address, "REDACTED")

	//
	address = registry.SelectLocation(100)
	assert.Nil(t, address, "REDACTED")
}

func VerifyAddressRegistryPersistImport(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	//
	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	registry.Persist()

	registry = NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	err := registry.Begin()
	require.NoError(t, err)

	assert.True(t, registry.Empty())

	//
	randomLocations := randomNetLocationCouples(t, 100)

	for _, addressOrigin := range randomLocations {
		err := registry.AppendLocation(addressOrigin.address, addressOrigin.src)
		require.NoError(t, err)
	}

	assert.Equal(t, 100, registry.Volume())
	registry.Persist()

	registry = NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	err = registry.Begin()
	require.NoError(t, err)

	assert.Equal(t, 100, registry.Volume())
}

func VerifyAddressRegistrySearch(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	randomLocations := randomNetLocationCouples(t, 100)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	for _, addressOrigin := range randomLocations {
		address := addressOrigin.address
		src := addressOrigin.src
		err := registry.AppendLocation(address, src)
		require.NoError(t, err)

		ka := registry.HasLocation(address)
		assert.True(t, ka, "REDACTED", address)
	}
}

func VerifyAddressRegistryElevateToAged(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	randomLocations := randomNetLocationCouples(t, 100)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	for _, addressOrigin := range randomLocations {
		err := registry.AppendLocation(addressOrigin.address, addressOrigin.src)
		require.NoError(t, err)
	}

	//
	for _, addressOrigin := range randomLocations {
		registry.StampEndeavor(addressOrigin.address)
	}

	//
	for i, addressOrigin := range randomLocations {
		if i%2 == 0 {
			registry.StampValid(addressOrigin.address.ID)
		}
	}

	//

	preference := registry.FetchPreference()
	t.Logf("REDACTED", preference)

	if len(preference) > registry.Volume() {
		t.Errorf("REDACTED")
	}

	preference = registry.FetchPreferenceWithTendency(30)
	t.Logf("REDACTED", preference)

	if len(preference) > registry.Volume() {
		t.Errorf("REDACTED")
	}

	assert.Equal(t, registry.Volume(), 100, "REDACTED")
}

func VerifyAddressRegistryManagersReplicates(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())

	randomLocations := randomNetLocationCouples(t, 100)

	distinctOrigin := randomIDXPv4address(t)
	for _, addressOrigin := range randomLocations {
		err := registry.AppendLocation(addressOrigin.address, addressOrigin.src)
		require.NoError(t, err)
		err = registry.AppendLocation(addressOrigin.address, addressOrigin.src) //
		require.NoError(t, err)
		err = registry.AppendLocation(addressOrigin.address, distinctOrigin) //
		require.NoError(t, err)
	}

	assert.Equal(t, 100, registry.Volume())
}

type netLocationCouple struct {
	address *p2p.NetLocation
	src  *p2p.NetLocation
}

func randomNetLocationCouples(t *testing.T, n int) []netLocationCouple {
	randomLocations := make([]netLocationCouple, n)
	for i := 0; i < n; i++ {
		randomLocations[i] = netLocationCouple{address: randomIDXPv4address(t), src: randomIDXPv4address(t)}
	}
	return randomLocations
}

func randomIDXPv4address(t *testing.T) *p2p.NetLocation {
	for {
		ip := fmt.Sprintf("REDACTED",
			engineseed.Intn(254)+1,
			engineseed.Intn(255),
			engineseed.Intn(255),
			engineseed.Intn(255),
		)
		port := engineseed.Intn(65535-1) + 1
		id := p2p.ID(hex.EncodeToString(engineseed.Octets(p2p.UIDOctetExtent)))
		uidAddress := p2p.UIDLocationString(id, fmt.Sprintf("REDACTED", ip, port))
		address, err := p2p.NewNetLocationString(uidAddress)
		assert.Nil(t, err, "REDACTED")
		if address.Forwardable() {
			return address
		}
	}
}

func VerifyAddressRegistryDeleteLocation(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())

	address := randomIDXPv4address(t)
	err := registry.AppendLocation(address, address)
	require.NoError(t, err)
	assert.Equal(t, 1, registry.Volume())

	registry.DeleteLocation(address)
	assert.Equal(t, 0, registry.Volume())

	notPresentAddress := randomIDXPv4address(t)
	registry.DeleteLocation(notPresentAddress)
	assert.Equal(t, 0, registry.Volume())
}

func VerifyAddressRegistryFetchPreferenceWithOneLabeledValid(t *testing.T) {
	//
	registry, filename := instantiateAddressRegistryWithMAgedAndNNewLocations(t, 1, 9)
	defer eraseTemporaryEntry(filename)

	locations := registry.FetchPreferenceWithTendency(tendencyToChooseNewNodes)
	assert.NotNil(t, locations)
	affirmMAgedAndNNewLocationsInPreference(t, 1, 9, locations, registry)
}

func VerifyAddressRegistryFetchPreferenceWithOneNegateLabeledValid(t *testing.T) {
	//
	registry, filename := instantiateAddressRegistryWithMAgedAndNNewLocations(t, 9, 1)
	defer eraseTemporaryEntry(filename)

	locations := registry.FetchPreferenceWithTendency(tendencyToChooseNewNodes)
	assert.NotNil(t, locations)
	affirmMAgedAndNNewLocationsInPreference(t, 9, 1, locations, registry)
}

func VerifyAddressRegistryFetchPreferenceYieldsNullWhenAddressRegistryIsEmpty(t *testing.T) {
	registry, filename := instantiateAddressRegistryWithMAgedAndNNewLocations(t, 0, 0)
	defer eraseTemporaryEntry(filename)

	locations := registry.FetchPreferenceWithTendency(tendencyToChooseNewNodes)
	assert.Nil(t, locations)
}

func VerifyAddressRegistryFetchPreference(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())

	//
	assert.Empty(t, registry.FetchPreference())

	//
	address := randomIDXPv4address(t)
	err := registry.AppendLocation(address, address)
	require.NoError(t, err)

	assert.Equal(t, 1, len(registry.FetchPreference()))
	assert.Equal(t, address, registry.FetchPreference()[0])

	//
	randomLocations := randomNetLocationCouples(t, 100)
	for _, addressOrigin := range randomLocations {
		err := registry.AppendLocation(addressOrigin.address, addressOrigin.src)
		require.NoError(t, err)
	}

	//
	locations := make(map[string]*p2p.NetLocation)
	preference := registry.FetchPreference()
	for _, address := range preference {
		if dup, ok := locations[address.String()]; ok {
			t.Fatalf("REDACTED", preference, dup)
		}
		locations[address.String()] = address
	}

	if len(preference) > registry.Volume() {
		t.Errorf("REDACTED", preference)
	}
}

func VerifyAddressRegistryFetchPreferenceWithTendency(t *testing.T) {
	const tendencyTowardNewLocations = 30

	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())

	//
	preference := registry.FetchPreferenceWithTendency(tendencyTowardNewLocations)
	assert.Empty(t, preference)

	//
	address := randomIDXPv4address(t)
	err := registry.AppendLocation(address, address)
	require.NoError(t, err)

	preference = registry.FetchPreferenceWithTendency(tendencyTowardNewLocations)
	assert.Equal(t, 1, len(preference))
	assert.Equal(t, address, preference[0])

	//
	randomLocations := randomNetLocationCouples(t, 100)
	for _, addressOrigin := range randomLocations {
		err := registry.AppendLocation(addressOrigin.address, addressOrigin.src)
		require.NoError(t, err)
	}

	//
	locations := make(map[string]*p2p.NetLocation)
	preference = registry.FetchPreferenceWithTendency(tendencyTowardNewLocations)
	for _, address := range preference {
		if dup, ok := locations[address.String()]; ok {
			t.Fatalf("REDACTED", preference, dup)
		}
		locations[address.String()] = address
	}

	if len(preference) > registry.Volume() {
		t.Fatalf("REDACTED", preference)
	}

	//
	randomLocationsSize := len(randomLocations)
	for i, addressOrigin := range randomLocations {
		if int((float64(i)/float64(randomLocationsSize))*100) >= 20 {
			registry.StampValid(addressOrigin.address.ID)
		}
	}

	preference = registry.FetchPreferenceWithTendency(tendencyTowardNewLocations)

	//
	sound := 0
	for _, address := range preference {
		if registry.IsValid(address) {
			sound++
		}
	}

	got, anticipated := int((float64(sound)/float64(len(preference)))*100), 100-tendencyTowardNewLocations

	//
	deficiency := int(math.Round(float64(100) / float64(len(preference))))
	if got > anticipated+deficiency {
		t.Fatalf(
			"REDACTED",
			got,
			anticipated,
			sound,
			len(preference),
		)
	}
	if got < anticipated-deficiency {
		t.Fatalf(
			"REDACTED",
			got,
			anticipated,
			sound,
			len(preference),
		)
	}
}

func VerifyAddressRegistryHasLocation(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	address := randomIDXPv4address(t)
	err := registry.AppendLocation(address, address)
	require.NoError(t, err)

	assert.True(t, registry.HasLocation(address))

	registry.DeleteLocation(address)

	assert.False(t, registry.HasLocation(address))
}

func verifyInstantiateInternalLocations(t *testing.T, countLocations int) ([]*p2p.NetLocation, []string) {
	locations := make([]*p2p.NetLocation, countLocations)
	for i := 0; i < countLocations; i++ {
		locations[i] = randomIDXPv4address(t)
	}

	internal := make([]string, countLocations)
	for i, address := range locations {
		internal[i] = string(address.ID)
	}
	return locations, internal
}

func VerifyProhibitFlawedNodes(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())

	address := randomIDXPv4address(t)
	_ = registry.AppendLocation(address, address)

	registry.StampFlawed(address, 1*time.Second)
	//
	assert.False(t, registry.HasLocation(address))
	assert.True(t, registry.IsProhibited(address))

	err := registry.AppendLocation(address, address)
	//
	require.Error(t, err)

	time.Sleep(1 * time.Second)
	registry.RestoreFlawedNodes()
	//
	assert.EqualValues(t, 1, registry.Volume())
	assert.True(t, registry.HasLocation(address))
	assert.False(t, registry.IsValid(address))
}

func VerifyAddressRegistryEmpty(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	//
	require.True(t, registry.Empty())
	//
	registry.AppendOurLocation(randomIDXPv4address(t))
	require.True(t, registry.Empty())
	//
	_, internalIdentifiers := verifyInstantiateInternalLocations(t, 5)
	registry.AppendInternalIDXDatastore(internalIdentifiers)
	require.True(t, registry.Empty())

	//
	err := registry.AppendLocation(randomIDXPv4address(t), randomIDXPv4address(t))
	require.NoError(t, err)
	require.False(t, registry.Empty())
}

func VerifyInternalNodes(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())

	locations, internal := verifyInstantiateInternalLocations(t, 10)
	registry.AppendInternalIDXDatastore(internal)

	//
	for _, address := range locations {
		err := registry.AppendLocation(address, address)
		if assert.Error(t, err) {
			_, ok := err.(ErrAddressRegistryInternal)
			assert.True(t, ok)
		}
	}

	//
	err := registry.AppendLocation(randomIDXPv4address(t), locations[0])
	if assert.Error(t, err) {
		_, ok := err.(ErrAddressRegistryInternalOrigin)
		assert.True(t, ok)
	}
}

func verifyAddressRegistryLocationPreference(t *testing.T, registryVolume int) {
	//
	for nRegistryAged := 0; nRegistryAged <= registryVolume; nRegistryAged++ {
		nRegistryNew := registryVolume - nRegistryAged
		dbgStr := fmt.Sprintf("REDACTED", registryVolume, nRegistryNew, nRegistryAged)

		//
		registry, filename := instantiateAddressRegistryWithMAgedAndNNewLocations(t, nRegistryAged, nRegistryNew)
		defer eraseTemporaryEntry(filename)
		locations := registry.FetchPreferenceWithTendency(tendencyToChooseNewNodes)
		assert.NotNil(t, locations, "REDACTED", dbgStr)
		nLocations := len(locations)
		assert.NotZero(t, nLocations, "REDACTED", dbgStr)

		//
		for _, address := range locations {
			if address == nil {
				t.Fatalf("REDACTED", dbgStr, locations)
			}
		}

		//
		nAged, nNew := numberAgedAndNewLocationsInPreference(locations, registry)

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
			k      = fractionOfCount(tendencyToChooseNewNodes, nLocations)
			expirationNew = cometmath.MinimumInteger(nNew, cometmath.MaximumInteger(k, nLocations-nRegistryAged))
			expirationAged = cometmath.MinimumInteger(nAged, nLocations-expirationNew)
		)

		//
		if nNew != expirationNew {
			t.Fatalf("REDACTED", dbgStr, expirationNew, nNew)
		}
		if nAged != expirationAged {
			t.Fatalf("REDACTED", dbgStr, expirationAged, nAged)
		}

		//
		//
		seqSizes, seqKinds, err := examinePreferenceArrangement(registry, locations)
		assert.NoError(t, err, "REDACTED", dbgStr)

		//
		//
		//
		var expirationSeqSizes []int
		var expirationSeqKinds []int

		switch {
		case expirationAged == 0: //
			expirationSeqSizes = []int{nLocations}
			expirationSeqKinds = []int{1}
		case expirationNew == 0: //
			expirationSeqSizes = []int{nLocations}
			expirationSeqKinds = []int{2}
		case nLocations-expirationNew-expirationAged == 0: //
			expirationSeqSizes = []int{expirationNew, expirationAged}
			expirationSeqKinds = []int{1, 2}
		}

		assert.Equal(t, expirationSeqSizes, seqSizes,
			"REDACTED",
			dbgStr, expirationSeqSizes, seqSizes)
		assert.Equal(t, expirationSeqKinds, seqKinds,
			"REDACTED",
			dbgStr, expirationSeqKinds, seqKinds)
	}
}

func VerifyVariedAddressRegistryLocationPreference(t *testing.T) {
	//
	const N = 32
	for registryVolume := 1; registryVolume < N; registryVolume++ {
		verifyAddressRegistryLocationPreference(t, registryVolume)
	}

	//
	spans := [...][]int{{33, 100}, {100, 175}}
	registryExtents := make([]int, 0, len(spans))
	for _, r := range spans {
		registryExtents = append(registryExtents, engineseed.Intn(r[1]-r[0])+r[0])
	}
	t.Logf("REDACTED", registryExtents)
	for _, registryVolume := range registryExtents {
		verifyAddressRegistryLocationPreference(t, registryVolume)
	}
}

func VerifyAddressRegistryAppendDoesNegateReplaceAgedIP(t *testing.T) {
	filename := instantiateTemporaryEntryLabel("REDACTED")
	defer eraseTemporaryEntry(filename)

	//
	//
	//
	nodeUID := "REDACTED"
	nodeActualIP := "REDACTED"
	nodeSupersedeEndeavorIP := "REDACTED"
	OriginAddress := "REDACTED"

	//
	//
	//
	countSupersedeTries := 10

	nodeActualAddress, err := p2p.NewNetLocationString(nodeUID + "REDACTED" + nodeActualIP)
	require.Nil(t, err)

	nodeSupersedeEndeavorAddress, err := p2p.NewNetLocationString(nodeUID + "REDACTED" + nodeSupersedeEndeavorIP)
	require.Nil(t, err)

	src, err := p2p.NewNetLocationString(OriginAddress)
	require.Nil(t, err)

	registry := NewAddressRegistry(filename, true)
	registry.AssignTracer(log.VerifyingTracer())
	err = registry.AppendLocation(nodeActualAddress, src)
	require.Nil(t, err)
	registry.StampEndeavor(nodeActualAddress)
	registry.StampValid(nodeActualAddress.ID)

	//
	err = registry.AppendLocation(nodeActualAddress, src)
	require.Nil(t, err)

	//
	//
	for i := 0; i < countSupersedeTries; i++ {
		err = registry.AppendLocation(nodeSupersedeEndeavorAddress, src)
		require.Nil(t, err)
	}
	//
	//
	//
	//
	for i := 0; i < countSupersedeTries; i++ {
		preference := registry.FetchPreference()
		for _, address := range preference {
			require.Equal(t, address.IP, nodeActualAddress.IP)
		}
	}
}

func VerifyAddressRegistryClusterKey(t *testing.T) {
	//
	verifyScenarios := []struct {
		label   string
		ip     string
		expirationKey string
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
		key := clusterKeyFor(p2p.NewNetLocationIPPort(nip, 26656), false)
		assert.Equal(t, tc.expirationKey, key, "REDACTED", i)
	}

	//
	verifyScenarios = []struct {
		label   string
		ip     string
		expirationKey string
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
		key := clusterKeyFor(p2p.NewNetLocationIPPort(nip, 26656), true)
		assert.Equal(t, tc.expirationKey, key, "REDACTED", i)
	}
}

func affirmMAgedAndNNewLocationsInPreference(t *testing.T, m, n int, locations []*p2p.NetLocation, registry *addressLedger) {
	nAged, nNew := numberAgedAndNewLocationsInPreference(locations, registry)
	assert.Equal(t, m, nAged, "REDACTED")
	assert.Equal(t, n, nNew, "REDACTED")
}

func instantiateTemporaryEntryLabel(prefix string) string {
	f, err := os.CreateTemp("REDACTED", prefix)
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

func eraseTemporaryEntry(filename string) {
	err := os.Remove(filename)
	if err != nil {
		panic(err)
	}
}

func instantiateAddressRegistryWithMAgedAndNNewLocations(t *testing.T, nAged, nNew int) (registry *addressLedger, filename string) {
	filename = instantiateTemporaryEntryLabel("REDACTED")

	registry = NewAddressRegistry(filename, true).(*addressLedger)
	registry.AssignTracer(log.VerifyingTracer())
	assert.Zero(t, registry.Volume())

	randomLocations := randomNetLocationCouples(t, nAged)
	for _, address := range randomLocations {
		err := registry.AppendLocation(address.address, address.src)
		require.NoError(t, err)
		registry.StampValid(address.address.ID)
	}

	randomLocations = randomNetLocationCouples(t, nNew)
	for _, address := range randomLocations {
		err := registry.AppendLocation(address.address, address.src)
		require.NoError(t, err)
	}

	return
}

func numberAgedAndNewLocationsInPreference(locations []*p2p.NetLocation, registry *addressLedger) (nAged, nNew int) {
	for _, address := range locations {
		if registry.IsValid(address) {
			nAged++
		} else {
			nNew++
		}
	}
	return
}

//
//
//
//
func examinePreferenceArrangement(registry *addressLedger, locations []*p2p.NetLocation) (seqSizes, seqKinds []int, err error) {
	//
	var (
		previousKind      = 0
		ongoingSeqSize = 0
	)

	for _, address := range locations {
		addressKind := 0
		if registry.IsValid(address) {
			addressKind = 2
		} else {
			addressKind = 1
		}
		if addressKind != previousKind && previousKind != 0 {
			seqSizes = append(seqSizes, ongoingSeqSize)
			seqKinds = append(seqKinds, previousKind)
			ongoingSeqSize = 0
		}
		ongoingSeqSize++
		previousKind = addressKind
	}

	seqSizes = append(seqSizes, ongoingSeqSize)
	seqKinds = append(seqKinds, previousKind)

	return
}
