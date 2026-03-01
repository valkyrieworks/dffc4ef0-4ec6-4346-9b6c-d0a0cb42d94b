//

package verifies

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/asyncronous"
	sc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

func RandomizePeer2peerCredentialLinkage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		randomize(data)
	})
}

func randomize(data []byte) {
	if len(data) == 0 {
		return
	}

	sampleLink, graphLink := createCredentialLinkDuo()

	//
	//
	go func() {
		//
		dataTowardPersist := make([]byte, len(data))
		copy(dataTowardPersist, data)

		n, err := sampleLink.Record(dataTowardPersist)
		if err != nil {
			panic(err)
		}
		if n < len(data) {
			panic(fmt.Sprintf("REDACTED", len(data), n))
		}
	}()

	dataFetch := make([]byte, len(data))
	sumFetch := 0
	for sumFetch < len(data) {
		buf := make([]byte, len(data)-sumFetch)
		m, err := graphLink.Obtain(buf)
		if err != nil {
			panic(err)
		}
		copy(dataFetch[sumFetch:], buf[:m])
		sumFetch += m
	}

	if !bytes.Equal(data, dataFetch) {
		panic("REDACTED")
	}
}

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

//
func createTokvalDepotLinkDuo() (sampleLink, graphLink statedepotLink) {
	graphFetcher, samplePersistor := io.Pipe()
	sampleFetcher, graphPersistor := io.Pipe()
	return statedepotLink{sampleFetcher, samplePersistor}, statedepotLink{graphFetcher, graphPersistor}
}

func createCredentialLinkDuo() (sampleSecondLink, graphSecondLink *sc.CredentialLinkage) {
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
			sampleSecondLink, err = sc.CreateCredentialLinkage(sampleLink, samplePrivateToken)
			if err != nil {
				log.Printf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := sampleSecondLink.DistantPublicToken()
			if !distantPublicOctets.Matches(graphPublicToken) {
				err = fmt.Errorf("REDACTED",
					graphPublicToken, sampleSecondLink.DistantPublicToken())
				log.Print(err)
				return nil, true, err
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			graphSecondLink, err = sc.CreateCredentialLinkage(graphLink, graphPrivateToken)
			if graphSecondLink == nil {
				log.Printf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := graphSecondLink.DistantPublicToken()
			if !distantPublicOctets.Matches(samplePublicToken) {
				err = fmt.Errorf("REDACTED",
					samplePublicToken, graphSecondLink.DistantPublicToken())
				log.Print(err)
				return nil, true, err
			}
			return nil, false, nil
		},
	)

	if trs.InitialFailure() != nil {
		log.Fatalf("REDACTED", trs.InitialFailure())
	}
	if !ok {
		log.Fatal("REDACTED")
	}

	return sampleSecondLink, graphSecondLink
}
