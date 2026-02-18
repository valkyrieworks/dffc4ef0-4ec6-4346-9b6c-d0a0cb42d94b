//

package verifies

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/async"
	sc "github.com/valkyrieworks/p2p/link"
)

func RandomizeP2PTokenLinkage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		randomize(data)
	})
}

func randomize(data []byte) {
	if len(data) == 0 {
		return
	}

	fooLink, barLink := createTokenLinkCouple()

	//
	//
	go func() {
		//
		dataToRecord := make([]byte, len(data))
		copy(dataToRecord, data)

		n, err := fooLink.Record(dataToRecord)
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
		m, err := barLink.Scan(buf)
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

//
func createObjectDepotLinkCouple() (fooLink, barLink objectdepotLink) {
	barScanner, fooRecorder := io.Pipe()
	fooScanner, barRecorder := io.Pipe()
	return objectdepotLink{fooScanner, fooRecorder}, objectdepotLink{barScanner, barRecorder}
}

func createTokenLinkCouple() (fooSecurityLink, barSecurityLink *sc.TokenLinkage) {
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
			fooSecurityLink, err = sc.CreateTokenLinkage(fooLink, fooPrivateKey)
			if err != nil {
				log.Printf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := fooSecurityLink.DistantPublicKey()
			if !distantPublicOctets.Matches(barPublicKey) {
				err = fmt.Errorf("REDACTED",
					barPublicKey, fooSecurityLink.DistantPublicKey())
				log.Print(err)
				return nil, true, err
			}
			return nil, false, nil
		},
		func(_ int) (val any, cancel bool, err error) {
			barSecurityLink, err = sc.CreateTokenLinkage(barLink, barPrivateKey)
			if barSecurityLink == nil {
				log.Printf("REDACTED", err)
				return nil, true, err
			}
			distantPublicOctets := barSecurityLink.DistantPublicKey()
			if !distantPublicOctets.Matches(fooPublicKey) {
				err = fmt.Errorf("REDACTED",
					fooPublicKey, barSecurityLink.DistantPublicKey())
				log.Print(err)
				return nil, true, err
			}
			return nil, false, nil
		},
	)

	if trs.InitialFault() != nil {
		log.Fatalf("REDACTED", trs.InitialFault())
	}
	if !ok {
		log.Fatal("REDACTED")
	}

	return fooSecurityLink, barSecurityLink
}
