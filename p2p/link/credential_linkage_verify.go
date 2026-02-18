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

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/async"
	cometos "github.com/valkyrieworks/utils/os"
	engineseed "github.com/valkyrieworks/utils/random"
)

//
//
var modify = flag.Bool("REDACTED", false, "REDACTED")

type objectdepotLink struct {
	*io.PipeScanner
	*io.PipeRecorder
}

func (drw objectdepotLink) End() (err error) {
	err2 := drw.PipeRecorder.CloseWithError(io.EOF)
	fault1 := drw.PipeScanner.Close()
	if err2 != nil {
		return err
	}
	return fault1
}

type privateKeyWithNullPublicKey struct {
	orig vault.PrivateKey
}

func (pk privateKeyWithNullPublicKey) Octets() []byte                   { return pk.orig.Octets() }
func (pk privateKeyWithNullPublicKey) Attest(msg []byte) ([]byte, error) { return pk.orig.Attest(msg) }
func (pk privateKeyWithNullPublicKey) PublicKey() vault.PublicKey           { return nil }
func (pk privateKeyWithNullPublicKey) Matches(pk2 vault.PrivateKey) bool  { return pk.orig.Matches(pk2) }
func (pk privateKeyWithNullPublicKey) Kind() string                    { return "REDACTED" }

func VerifyCredentialLinkageGreeting(t *testing.T) {
	fooSecurityLink, barSecurityLink := createTokenLinkCouple(t)
	if err := fooSecurityLink.End(); err != nil {
		t.Error(err)
	}
	if err := barSecurityLink.End(); err != nil {
		t.Error(err)
	}
}

func VerifyParallelRecord(t *testing.T) {
	fooSecurityLink, barSecurityLink := createTokenLinkCouple(t)
	fooRecordContent := engineseed.Str(dataMaximumVolume)

	//
	//
	//
	n := 100
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go recordMany(t, wg, fooSecurityLink, fooRecordContent, n)
	go recordMany(t, wg, fooSecurityLink, fooRecordContent, n)

	//
	readMany(t, wg, barSecurityLink, n*2)
	wg.Wait()

	if err := fooSecurityLink.End(); err != nil {
		t.Error(err)
	}
}

func VerifyParallelRead(t *testing.T) {
	fooSecurityLink, barSecurityLink := createTokenLinkCouple(t)
	fooRecordContent := engineseed.Str(dataMaximumVolume)
	n := 100

	//
	//
	//
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go readMany(t, wg, fooSecurityLink, n/2)
	go readMany(t, wg, fooSecurityLink, n/2)

	//
	recordMany(t, wg, barSecurityLink, fooRecordContent, n)
	wg.Wait()

	if err := fooSecurityLink.End(); err != nil {
		t.Error(err)
	}
}

func VerifyCredentialLinkageReadRecord(t *testing.T) {
	fooLink, barLink := createObjectDepotLinkCouple()
	fooPersists, barPersists := []string{}, []string{}
	fooFetches, barFetches := []string{}, []string{}

	//
	for i := 0; i < 100; i++ {
		fooPersists = append(fooPersists, engineseed.Str((engineseed.Int()%(dataMaximumVolume*5))+1))
		barPersists = append(barPersists, engineseed.Str((engineseed.Int()%(dataMaximumVolume*5))+1))
	}

	//
	generateMemberExecutor := func(id string, memberLink objectdepotLink, memberPersists []string, memberFetches *[]string) async.Task {
		return func(_ int) (any, bool, error) {
			//
			memberInternalKey := ed25519.GeneratePrivateKey()
			memberCredentialLink, err := CreateTokenLinkage(memberLink, memberInternalKey)
			if err != nil {
				t.Errorf("REDACTED", err)
				return nil, true, err
			}
			//
			trs, ok := async.Concurrent(
				func(_ int) (any, bool, error) {
					//
					for _, memberRecord := range memberPersists {
						n, err := memberCredentialLink.Record([]byte(memberRecord))
						if err != nil {
							t.Errorf("REDACTED", err)
							return nil, true, err
						}
						if n != len(memberRecord) {
							err = fmt.Errorf("REDACTED", len(memberRecord), n)
							t.Error(err)
							return nil, true, err
						}
					}
					if err := memberLink.PipeRecorder.Close(); err != nil {
						t.Error(err)
						return nil, true, err
					}
					return nil, false, nil
				},
				func(_ int) (any, bool, error) {
					//
					readFrame := make([]byte, dataMaximumVolume)
					for {
						n, err := memberCredentialLink.Scan(readFrame)
						if err == io.EOF {
							if err := memberLink.PipeScanner.Close(); err != nil {
								t.Error(err)
								return nil, true, err
							}
							return nil, false, nil
						} else if err != nil {
							t.Errorf("REDACTED", err)
							return nil, true, err
						}
						*memberFetches = append(*memberFetches, string(readFrame[:n]))
					}
				},
			)
			assert.True(t, ok, "REDACTED")

			//
			if trs.InitialFault() != nil {
				return nil, true, trs.InitialFault()
			}

			//
			return nil, false, nil
		}
	}

	//
	trs, ok := async.Concurrent(
		generateMemberExecutor("REDACTED", fooLink, fooPersists, &fooFetches),
		generateMemberExecutor("REDACTED", barLink, barPersists, &barFetches),
	)
	require.Nil(t, trs.InitialFault())
	require.True(t, ok, "REDACTED")

	//
	//
	contrastPersistsFetches := func(persists []string, fetches []string) {
		for {
			//
			reader := "REDACTED"
			record := persists[0]
			readNumber := 0
			for _, readSegment := range fetches {
				reader += readSegment
				readNumber++
				if len(record) <= len(reader) {
					break
				}
				if len(record) <= dataMaximumVolume {
					break //
				}
			}
			//
			if record != reader {
				t.Errorf("REDACTED", record, reader)
			}
			//
			persists = persists[1:]
			fetches = fetches[readNumber:]
			if len(persists) == 0 {
				break
			}
		}
	}

	contrastPersistsFetches(fooPersists, barFetches)
	contrastPersistsFetches(barPersists, fooFetches)
}

func VerifyDeduceCredentialsAndDisputeValidated(t *testing.T) {
	validatedRoutepath := filepath.Join("REDACTED", t.Name()+"REDACTED")
	if *modify {
		t.Logf("REDACTED", validatedRoutepath)
		data := instantiateValidatedVerifyArrays(t)
		err := cometos.RecordEntry(validatedRoutepath, []byte(data), 0o644)
		require.NoError(t, err)
	}
	f, err := os.Open(validatedRoutepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	analyzer := bufio.NewScanner(f)
	for analyzer.Scan() {
		row := analyzer.Text()
		options := strings.Split(row, "REDACTED")
		randomCredentialArray, err := hex.DecodeString(options[0])
		require.Nil(t, err)
		randomCredential := new([32]byte)
		copy((*randomCredential)[:], randomCredentialArray)
		locationIsMinimum, err := strconv.ParseBool(options[1])
		require.Nil(t, err)
		anticipatedReceiveCredential, err := hex.DecodeString(options[2])
		require.Nil(t, err)
		anticipatedTransmitCredential, err := hex.DecodeString(options[3])
		require.Nil(t, err)

		receiveCredential, transmitCredential := deduceCredentials(randomCredential, locationIsMinimum)
		require.Equal(t, anticipatedReceiveCredential, (*receiveCredential)[:], "REDACTED")
		require.Equal(t, anticipatedTransmitCredential, (*transmitCredential)[:], "REDACTED")
	}
}

func VerifyNullPublickey(t *testing.T) {
	fooLink, barLink := createObjectDepotLinkCouple()
	defer fooLink.End()
	defer barLink.End()
	fooPrivateKey := ed25519.GeneratePrivateKey()
	barPrivateKey := privateKeyWithNullPublicKey{ed25519.GeneratePrivateKey()}

	go CreateTokenLinkage(fooLink, fooPrivateKey) //

	_, err := CreateTokenLinkage(barLink, barPrivateKey)
	require.Error(t, err)
	assert.Equal(t, "REDACTED", err.Error())
}

func recordMany(t *testing.T, wg *sync.WaitGroup, link io.Writer, txt string, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		_, err := link.Write([]byte(txt))
		if err != nil {
			t.Errorf("REDACTED", err)
			return
		}
	}
}

func readMany(t *testing.T, wg *sync.WaitGroup, link io.Reader, n int) {
	readFrame := make([]byte, dataMaximumVolume)
	for i := 0; i < n; i++ {
		_, err := link.Read(readFrame)
		assert.NoError(t, err)
	}
	wg.Done()
}

//
//
//
func instantiateValidatedVerifyArrays(t *testing.T) string {
	t.Helper()

	data := "REDACTED"
	for i := 0; i < 32; i++ {
		randomCredentialArray := engineseed.Octets(32)
		randomCredential := new([32]byte)
		copy((*randomCredential)[:], randomCredentialArray)
		data += hex.EncodeToString((*randomCredential)[:]) + "REDACTED"
		locationIsMinimum := engineseed.Bool()
		data += strconv.FormatBool(locationIsMinimum) + "REDACTED"
		receiveCredential, transmitCredential := deduceCredentials(randomCredential, locationIsMinimum)
		data += hex.EncodeToString((*receiveCredential)[:]) + "REDACTED"
		data += hex.EncodeToString((*transmitCredential)[:]) + "REDACTED"
	}
	return data
}

//
func createObjectDepotLinkCouple() (fooLink, barLink objectdepotLink) {
	barScanner, fooRecorder := io.Pipe()
	fooScanner, barRecorder := io.Pipe()
	return objectdepotLink{fooScanner, fooRecorder}, objectdepotLink{barScanner, barRecorder}
}

func createTokenLinkCouple(tb testing.TB) (fooSecurityLink, barSecurityLink *TokenLinkage) {
	var (
		fooLink, barLink = createObjectDepotLinkCouple()
		fooPrivateKey        = ed25519.GeneratePrivateKey()
		fooPublicKey        = fooPrivateKey.PublicKey()
		barPrivateKey        = ed25519.GeneratePrivateKey()
		barPublicKey        = barPrivateKey.PublicKey()
	)

	//
	trs, ok := async.Concurrent(
		func(_ int) (val any, cancel bool, err error) {
			fooSecurityLink, err = CreateTokenLinkage(fooLink, fooPrivateKey)
			if err != nil {
				tb.Errorf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := fooSecurityLink.DistantPublicKey()
			if !distantPublicOctets.Matches(barPublicKey) {
				err = fmt.Errorf("REDACTED",
					barPublicKey, fooSecurityLink.DistantPublicKey())
				tb.Error(err)
				return nil, true, err
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			barSecurityLink, err = CreateTokenLinkage(barLink, barPrivateKey)
			if barSecurityLink == nil {
				tb.Errorf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := barSecurityLink.DistantPublicKey()
			if !distantPublicOctets.Matches(fooPublicKey) {
				err = fmt.Errorf("REDACTED",
					fooPublicKey, barSecurityLink.DistantPublicKey())
				tb.Error(err)
				return nil, true, err
			}
			return nil, false, nil
		},
	)

	require.Nil(tb, trs.InitialFault())
	require.True(tb, ok, "REDACTED")

	return fooSecurityLink, barSecurityLink
}

//

func CriterionRecordCredentialLinkage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	fooSecurityLink, barSecurityLink := createTokenLinkCouple(b)
	arbitraryMessageExtents := []int{
		dataMaximumVolume / 10,
		dataMaximumVolume / 3,
		dataMaximumVolume / 2,
		dataMaximumVolume,
		dataMaximumVolume * 3 / 2,
		dataMaximumVolume * 2,
		dataMaximumVolume * 7 / 2,
	}
	fooRecordOctets := make([][]byte, 0, len(arbitraryMessageExtents))
	for _, volume := range arbitraryMessageExtents {
		fooRecordOctets = append(fooRecordOctets, engineseed.Octets(volume))
	}
	//
	go func() {
		readFrame := make([]byte, dataMaximumVolume)
		for {
			_, err := barSecurityLink.Scan(readFrame)
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
		idx := engineseed.Intn(len(fooRecordOctets))
		_, err := fooSecurityLink.Record(fooRecordOctets[idx])
		if err != nil {
			b.Errorf("REDACTED", err)
			return
		}
	}
	b.StopTimer()

	if err := fooSecurityLink.End(); err != nil {
		b.Error(err)
	}
	//
}

func CriterionReadCredentialLinkage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	fooSecurityLink, barSecurityLink := createTokenLinkCouple(b)
	arbitraryMessageExtents := []int{
		dataMaximumVolume / 10,
		dataMaximumVolume / 3,
		dataMaximumVolume / 2,
		dataMaximumVolume,
		dataMaximumVolume * 3 / 2,
		dataMaximumVolume * 2,
		dataMaximumVolume * 7 / 2,
	}
	fooRecordOctets := make([][]byte, 0, len(arbitraryMessageExtents))
	for _, volume := range arbitraryMessageExtents {
		fooRecordOctets = append(fooRecordOctets, engineseed.Octets(volume))
	}
	go func() {
		for i := 0; i < b.N; i++ {
			idx := engineseed.Intn(len(fooRecordOctets))
			_, err := fooSecurityLink.Record(fooRecordOctets[idx])
			if err != nil {
				b.Errorf("REDACTED", err, i, b.N)
				return
			}
		}
	}()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		readFrame := make([]byte, dataMaximumVolume)
		_, err := barSecurityLink.Scan(readFrame)

		if err == io.EOF {
			return
		} else if err != nil {
			b.Fatalf("REDACTED", err)
		}
	}
	b.StopTimer()
}
