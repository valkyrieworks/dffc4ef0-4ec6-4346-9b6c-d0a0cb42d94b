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
//
//
//
//
//
//
//
//
//

package schema_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/cosmos/gogoproto/test"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
)

func iovalidate(persistor protocolio.PersistTerminator, fetcher protocolio.FetchTerminator) error {
	variableint := make([]byte, binary.MaxVarintLen64)
	extent := 1000
	signals := make([]*test.NinOptNative, extent)
	magnitudes := make([]int, extent)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range signals {
		signals[i] = test.NewPopulatedNinOptNative(r, true)
		//
		if i == 5 {
			signals[i] = &test.NinOptNative{}
		}
		//
		if i == 999 {
			signals[i] = &test.NinOptNative{}
		}
		//
		bz, err := proto.Marshal(signals[i])
		if err != nil {
			return err
		}
		visiblemagnitude := binary.PutUvarint(variableint, uint64(len(bz)))
		n, err := persistor.PersistSignal(signals[i])
		if err != nil {
			return err
		}
		if n != len(bz)+visiblemagnitude {
			return fmt.Errorf("REDACTED", n, len(bz)+visiblemagnitude)
		}
		magnitudes[i] = n
	}
	if err := persistor.Close(); err != nil {
		return err
	}
	i := 0
	for {
		msg := &test.NinOptNative{}
		if n, err := fetcher.FetchSignal(msg); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else if n != magnitudes[i] {
			return fmt.Errorf("REDACTED", n, magnitudes[i])
		}
		if err := msg.VerboseEqual(signals[i]); err != nil {
			return err
		}
		i++
	}
	if i != extent {
		panic("REDACTED")
	}
	return fetcher.Close()
}

type reserve struct {
	*bytes.Reserve
	terminated bool
}

func (b *reserve) Shutdown() error {
	b.terminated = true
	return nil
}

func freshReserve() *reserve {
	return &reserve{bytes.NewBuffer(nil), false}
}

func VerifyVariableintTypical(t *testing.T) {
	buf := freshReserve()
	persistor := protocolio.FreshSeparatedPersistor(buf)
	fetcher := protocolio.FreshSeparatedFetcher(buf, 1024*1024)
	err := iovalidate(persistor, fetcher)
	require.NoError(t, err)
	require.True(t, buf.terminated, "REDACTED")
}

func VerifyVariableintNegativeShutdown(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	persistor := protocolio.FreshSeparatedPersistor(buf)
	fetcher := protocolio.FreshSeparatedFetcher(buf, 1024*1024)
	err := iovalidate(persistor, fetcher)
	require.NoError(t, err)
}

//
func VerifyVariableintMaximumExtent(t *testing.T) {
	buf := freshReserve()
	persistor := protocolio.FreshSeparatedPersistor(buf)
	fetcher := protocolio.FreshSeparatedFetcher(buf, 20)
	err := iovalidate(persistor, fetcher)
	require.Error(t, err)
}

func VerifyVariableintFailure(t *testing.T) {
	buf := freshReserve()
	buf.Write([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f})
	fetcher := protocolio.FreshSeparatedFetcher(buf, 1024*1024)
	msg := &test.NinOptNative{}
	n, err := fetcher.FetchSignal(msg)
	require.Error(t, err)
	require.Equal(t, 10, n)
}

func VerifyVariableintAbridged(t *testing.T) {
	buf := freshReserve()
	buf.Write([]byte{0xff, 0xff})
	fetcher := protocolio.FreshSeparatedFetcher(buf, 1024*1024)
	msg := &test.NinOptNative{}
	n, err := fetcher.FetchSignal(msg)
	require.Error(t, err)
	require.Equal(t, 2, n)
}

func VerifyBrief(t *testing.T) {
	buf := freshReserve()

	variableintArea := make([]byte, binary.MaxVarintLen64)
	variableintLength := binary.PutUvarint(variableintArea, 100)
	_, err := buf.Write(variableintArea[:variableintLength])
	require.NoError(t, err)

	bz, err := proto.Marshal(&test.NinOptNative{Field15: []byte{0x01, 0x02, 0x03}})
	require.NoError(t, err)
	buf.Write(bz)

	fetcher := protocolio.FreshSeparatedFetcher(buf, 1024*1024)
	require.NoError(t, err)
	msg := &test.NinOptNative{}
	n, err := fetcher.FetchSignal(msg)
	require.Error(t, err)
	require.Equal(t, variableintLength+len(bz), n)
}
