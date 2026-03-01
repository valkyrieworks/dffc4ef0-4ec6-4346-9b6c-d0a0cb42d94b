package link

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/asyncronous"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

//
//
var revise = flag.Bool("REDACTED", false, "REDACTED")

type statedepotLink struct {
	*io.ConduitFetcher
	*io.ConduitPersistor
}

func (drw statedepotLink) Shutdown() (err error) {
	fault2 := drw.ConduitPersistor.CloseWithError(io.EOF)
	faultone := drw.ConduitFetcher.Close()
	if fault2 != nil {
		return err
	}
	return faultone
}

type privateTokenUsingVoidPublicToken struct {
	initial security.PrivateToken
}

func (pk privateTokenUsingVoidPublicToken) Octets() []byte                   { return pk.initial.Octets() }
func (pk privateTokenUsingVoidPublicToken) Attest(msg []byte) ([]byte, error) { return pk.initial.Attest(msg) }
func (pk privateTokenUsingVoidPublicToken) PublicToken() security.PublicToken           { return nil }
func (pk privateTokenUsingVoidPublicToken) Matches(pk2 security.PrivateToken) bool  { return pk.initial.Matches(pk2) }
func (pk privateTokenUsingVoidPublicToken) Kind() string                    { return "REDACTED" }

func VerifyCredentialLinkageNegotiation(t *testing.T) {
	sampleSecondLink, graphSecondLink := createCredentialLinkDuo(t)
	if err := sampleSecondLink.Shutdown(); err != nil {
		t.Error(err)
	}
	if err := graphSecondLink.Shutdown(); err != nil {
		t.Error(err)
	}
}

func VerifyParallelPersist(t *testing.T) {
	sampleSecondLink, graphSecondLink := createCredentialLinkDuo(t)
	samplePersistString := commitrand.Str(dataMaximumExtent)

	//
	//
	//
	n := 100
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go persistAbundant(t, wg, sampleSecondLink, samplePersistString, n)
	go persistAbundant(t, wg, sampleSecondLink, samplePersistString, n)

	//
	fetchAbundant(t, wg, graphSecondLink, n*2)
	wg.Wait()

	if err := sampleSecondLink.Shutdown(); err != nil {
		t.Error(err)
	}
}

func VerifyParallelFetch(t *testing.T) {
	sampleSecondLink, graphSecondLink := createCredentialLinkDuo(t)
	samplePersistString := commitrand.Str(dataMaximumExtent)
	n := 100

	//
	//
	//
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go fetchAbundant(t, wg, sampleSecondLink, n/2)
	go fetchAbundant(t, wg, sampleSecondLink, n/2)

	//
	persistAbundant(t, wg, graphSecondLink, samplePersistString, n)
	wg.Wait()

	if err := sampleSecondLink.Shutdown(); err != nil {
		t.Error(err)
	}
}

func VerifyCredentialLinkageFetchPersist(t *testing.T) {
	sampleLink, graphLink := createTokvalDepotLinkDuo()
	samplePersists, dividerPersists := []string{}, []string{}
	sampleFetches, dividerFetches := []string{}, []string{}

	//
	for i := 0; i < 100; i++ {
		samplePersists = append(samplePersists, commitrand.Str((commitrand.Int()%(dataMaximumExtent*5))+1))
		dividerPersists = append(dividerPersists, commitrand.Str((commitrand.Int()%(dataMaximumExtent*5))+1))
	}

	//
	producePeerExecutor := func(id string, peerLink statedepotLink, peerPersists []string, peerFetches *[]string) asyncronous.Activity {
		return func(_ int) (any, bool, error) {
			//
			peerPrivateToken := edwards25519.ProducePrivateToken()
			peerCredentialLink, err := CreateCredentialLinkage(peerLink, peerPrivateToken)
			if err != nil {
				t.Errorf("REDACTED", err)
				return nil, true, err
			}
			//
			trs, ok := asyncronous.Concurrent(
				func(_ int) (any, bool, error) {
					//
					for _, peerPersist := range peerPersists {
						n, err := peerCredentialLink.Record([]byte(peerPersist))
						if err != nil {
							t.Errorf("REDACTED", err)
							return nil, true, err
						}
						if n != len(peerPersist) {
							err = fmt.Errorf("REDACTED", len(peerPersist), n)
							t.Error(err)
							return nil, true, err
						}
					}
					if err := peerLink.ConduitPersistor.Close(); err != nil {
						t.Error(err)
						return nil, true, err
					}
					return nil, false, nil
				},
				func(_ int) (any, bool, error) {
					//
					fetchReserve := make([]byte, dataMaximumExtent)
					for {
						n, err := peerCredentialLink.Obtain(fetchReserve)
						if err == io.EOF {
							if err := peerLink.ConduitFetcher.Close(); err != nil {
								t.Error(err)
								return nil, true, err
							}
							return nil, false, nil
						} else if err != nil {
							t.Errorf("REDACTED", err)
							return nil, true, err
						}
						*peerFetches = append(*peerFetches, string(fetchReserve[:n]))
					}
				},
			)
			assert.True(t, ok, "REDACTED")

			//
			if trs.InitialFailure() != nil {
				return nil, true, trs.InitialFailure()
			}

			//
			return nil, false, nil
		}
	}

	//
	trs, ok := asyncronous.Concurrent(
		producePeerExecutor("REDACTED", sampleLink, samplePersists, &sampleFetches),
		producePeerExecutor("REDACTED", graphLink, dividerPersists, &dividerFetches),
	)
	require.Nil(t, trs.InitialFailure())
	require.True(t, ok, "REDACTED")

	//
	//
	contrastPersistsFetches := func(persists []string, fetches []string) {
		for {
			//
			fetch := "REDACTED"
			persist := persists[0]
			fetchTotal := 0
			for _, fetchSegment := range fetches {
				fetch += fetchSegment
				fetchTotal++
				if len(persist) <= len(fetch) {
					break
				}
				if len(persist) <= dataMaximumExtent {
					break //
				}
			}
			//
			if persist != fetch {
				t.Errorf("REDACTED", persist, fetch)
			}
			//
			persists = persists[1:]
			fetches = fetches[fetchTotal:]
			if len(persists) == 0 {
				break
			}
		}
	}

	contrastPersistsFetches(samplePersists, dividerFetches)
	contrastPersistsFetches(dividerPersists, sampleFetches)
}

func VerifyDeduceCredentialsAlsoQueryPrime(t *testing.T) {
	primePath := filepath.Join("REDACTED", t.Name()+"REDACTED")
	if *revise {
		t.Logf("REDACTED", primePath)
		data := generatePrimeVerifyArrays(t)
		err := strongos.RecordRecord(primePath, []byte(data), 0o644)
		require.NoError(t, err)
	}
	f, err := os.Open(primePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	analyzer := bufio.NewScanner(f)
	for analyzer.Scan() {
		row := analyzer.Text()
		parameters := strings.Split(row, "REDACTED")
		arbitraryCredentialArray, err := hex.DecodeString(parameters[0])
		require.Nil(t, err)
		arbitraryCredential := new([32]byte)
		copy((*arbitraryCredential)[:], arbitraryCredentialArray)
		positionEqualsMinimal, err := strconv.ParseBool(parameters[1])
		require.Nil(t, err)
		anticipatedObtainCredential, err := hex.DecodeString(parameters[2])
		require.Nil(t, err)
		anticipatedTransmitCredential, err := hex.DecodeString(parameters[3])
		require.Nil(t, err)

		obtainCredential, transmitCredential := deduceCredentials(arbitraryCredential, positionEqualsMinimal)
		require.Equal(t, anticipatedObtainCredential, (*obtainCredential)[:], "REDACTED")
		require.Equal(t, anticipatedTransmitCredential, (*transmitCredential)[:], "REDACTED")
	}
}

func VerifyVoidPublickey(t *testing.T) {
	sampleLink, graphLink := createTokvalDepotLinkDuo()
	defer sampleLink.Shutdown()
	defer graphLink.Shutdown()
	samplePrivateToken := edwards25519.ProducePrivateToken()
	graphPrivateToken := privateTokenUsingVoidPublicToken{edwards25519.ProducePrivateToken()}

	go CreateCredentialLinkage(sampleLink, samplePrivateToken) //

	_, err := CreateCredentialLinkage(graphLink, graphPrivateToken)
	require.Error(t, err)
	assert.Equal(t, "REDACTED", err.Error())
}

func persistAbundant(t *testing.T, wg *sync.WaitGroup, link io.Writer, txt string, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		_, err := link.Write([]byte(txt))
		if err != nil {
			t.Errorf("REDACTED", err)
			return
		}
	}
}

func fetchAbundant(t *testing.T, wg *sync.WaitGroup, link io.Reader, n int) {
	fetchReserve := make([]byte, dataMaximumExtent)
	for i := 0; i < n; i++ {
		_, err := link.Read(fetchReserve)
		assert.NoError(t, err)
	}
	wg.Done()
}

//
//
//
func generatePrimeVerifyArrays(t *testing.T) string {
	t.Helper()

	data := "REDACTED"
	for i := 0; i < 32; i++ {
		arbitraryCredentialArray := commitrand.Octets(32)
		arbitraryCredential := new([32]byte)
		copy((*arbitraryCredential)[:], arbitraryCredentialArray)
		data += hex.EncodeToString((*arbitraryCredential)[:]) + "REDACTED"
		positionEqualsMinimal := commitrand.Flag()
		data += strconv.FormatBool(positionEqualsMinimal) + "REDACTED"
		obtainCredential, transmitCredential := deduceCredentials(arbitraryCredential, positionEqualsMinimal)
		data += hex.EncodeToString((*obtainCredential)[:]) + "REDACTED"
		data += hex.EncodeToString((*transmitCredential)[:]) + "REDACTED"
	}
	return data
}

//
func createTokvalDepotLinkDuo() (sampleLink, graphLink statedepotLink) {
	graphFetcher, samplePersistor := io.Pipe()
	sampleFetcher, graphPersistor := io.Pipe()
	return statedepotLink{sampleFetcher, samplePersistor}, statedepotLink{graphFetcher, graphPersistor}
}

func createCredentialLinkDuo(tb testing.TB) (sampleSecondLink, graphSecondLink *CredentialLinkage) {
	var (
		sampleLink, graphLink = createTokvalDepotLinkDuo()
		samplePrivateToken        = edwards25519.ProducePrivateToken()
		samplePublicToken        = samplePrivateToken.PublicToken()
		graphPrivateToken        = edwards25519.ProducePrivateToken()
		graphPublicToken        = graphPrivateToken.PublicToken()
	)

	//
	trs, ok := asyncronous.Concurrent(
		func(_ int) (val any, cancel bool, err error) {
			sampleSecondLink, err = CreateCredentialLinkage(sampleLink, samplePrivateToken)
			if err != nil {
				tb.Errorf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := sampleSecondLink.DistantPublicToken()
			if !distantPublicOctets.Matches(graphPublicToken) {
				err = fmt.Errorf("REDACTED",
					graphPublicToken, sampleSecondLink.DistantPublicToken())
				tb.Error(err)
				return nil, true, err
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			graphSecondLink, err = CreateCredentialLinkage(graphLink, graphPrivateToken)
			if graphSecondLink == nil {
				tb.Errorf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := graphSecondLink.DistantPublicToken()
			if !distantPublicOctets.Matches(samplePublicToken) {
				err = fmt.Errorf("REDACTED",
					samplePublicToken, graphSecondLink.DistantPublicToken())
				tb.Error(err)
				return nil, true, err
			}
			return nil, false, nil
		},
	)

	require.Nil(tb, trs.InitialFailure())
	require.True(tb, ok, "REDACTED")

	return sampleSecondLink, graphSecondLink
}

//

func AssessmentPersistCredentialLinkage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	sampleSecondLink, graphSecondLink := createCredentialLinkDuo(b)
	unpredictableSignalExtents := []int{
		dataMaximumExtent / 10,
		dataMaximumExtent / 3,
		dataMaximumExtent / 2,
		dataMaximumExtent,
		dataMaximumExtent * 3 / 2,
		dataMaximumExtent * 2,
		dataMaximumExtent * 7 / 2,
	}
	samplePersistOctets := make([][]byte, 0, len(unpredictableSignalExtents))
	for _, extent := range unpredictableSignalExtents {
		samplePersistOctets = append(samplePersistOctets, commitrand.Octets(extent))
	}
	//
	go func() {
		fetchReserve := make([]byte, dataMaximumExtent)
		for {
			_, err := graphSecondLink.Obtain(fetchReserve)
			if err == io.EOF {
				return
			} else if err != nil {
				b.Errorf("REDACTED", err)
				return
			}
		}
	}()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		idx := commitrand.Integern(len(samplePersistOctets))
		_, err := sampleSecondLink.Record(samplePersistOctets[idx])
		if err != nil {
			b.Errorf("REDACTED", err)
			return
		}
	}
	b.StopTimer()

	if err := sampleSecondLink.Shutdown(); err != nil {
		b.Error(err)
	}
	//
}

func AssessmentFetchCredentialLinkage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	sampleSecondLink, graphSecondLink := createCredentialLinkDuo(b)
	unpredictableSignalExtents := []int{
		dataMaximumExtent / 10,
		dataMaximumExtent / 3,
		dataMaximumExtent / 2,
		dataMaximumExtent,
		dataMaximumExtent * 3 / 2,
		dataMaximumExtent * 2,
		dataMaximumExtent * 7 / 2,
	}
	samplePersistOctets := make([][]byte, 0, len(unpredictableSignalExtents))
	for _, extent := range unpredictableSignalExtents {
		samplePersistOctets = append(samplePersistOctets, commitrand.Octets(extent))
	}
	go func() {
		for i := 0; i < b.N; i++ {
			idx := commitrand.Integern(len(samplePersistOctets))
			_, err := sampleSecondLink.Record(samplePersistOctets[idx])
			if err != nil {
				b.Errorf("REDACTED", err, i, b.N)
				return
			}
		}
	}()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fetchReserve := make([]byte, dataMaximumExtent)
		_, err := graphSecondLink.Obtain(fetchReserve)

		if err == io.EOF {
			return
		} else if err != nil {
			b.Fatalf("REDACTED", err)
		}
	}
	b.StopTimer()
}
