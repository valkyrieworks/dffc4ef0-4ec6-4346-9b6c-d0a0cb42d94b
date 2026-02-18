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

	"github.com/valkyrieworks/utils/protoio"
)

func ioverify(recorder protoio.RecordTerminator, scanner protoio.ScanTerminator) error {
	variableint := make([]byte, binary.MaxVarintLen64)
	volume := 1000
	notices := make([]*test.NinOptNative, volume)
	sizes := make([]int, volume)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range notices {
		notices[i] = test.NewPopulatedNinOptNative(r, true)
		//
		if i == 5 {
			notices[i] = &test.NinOptNative{}
		}
		//
		if i == 999 {
			notices[i] = &test.NinOptNative{}
		}
		//
		bz, err := proto.Marshal(notices[i])
		if err != nil {
			return err
		}
		viewsize := binary.PutUvarint(variableint, uint64(len(bz)))
		n, err := recorder.RecordMessage(notices[i])
		if err != nil {
			return err
		}
		if n != len(bz)+viewsize {
			return fmt.Errorf("REDACTED", n, len(bz)+viewsize)
		}
		sizes[i] = n
	}
	if err := recorder.Close(); err != nil {
		return err
	}
	i := 0
	for {
		msg := &test.NinOptNative{}
		if n, err := scanner.ScanMessage(msg); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else if n != sizes[i] {
			return fmt.Errorf("REDACTED", n, sizes[i])
		}
		if err := msg.VerboseEqual(notices[i]); err != nil {
			return err
		}
		i++
	}
	if i != volume {
		panic("REDACTED")
	}
	return scanner.Close()
}

type buffer struct {
	*bytes.Frame
	halted bool
}

func (b *buffer) End() error {
	b.halted = true
	return nil
}

func newFrame() *buffer {
	return &buffer{bytes.NewBuffer(nil), false}
}

func VerifyVariableintTypical(t *testing.T) {
	buf := newFrame()
	recorder := protoio.NewSeparatedRecorder(buf)
	scanner := protoio.NewSeparatedScanner(buf, 1024*1024)
	err := ioverify(recorder, scanner)
	require.NoError(t, err)
	require.True(t, buf.halted, "REDACTED")
}

func VerifyVariableintNoEnd(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	recorder := protoio.NewSeparatedRecorder(buf)
	scanner := protoio.NewSeparatedScanner(buf, 1024*1024)
	err := ioverify(recorder, scanner)
	require.NoError(t, err)
}

//
func VerifyVariableintMaximumVolume(t *testing.T) {
	buf := newFrame()
	recorder := protoio.NewSeparatedRecorder(buf)
	scanner := protoio.NewSeparatedScanner(buf, 20)
	err := ioverify(recorder, scanner)
	require.Error(t, err)
}

func VerifyVariableintFault(t *testing.T) {
	buf := newFrame()
	buf.Write([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f})
	scanner := protoio.NewSeparatedScanner(buf, 1024*1024)
	msg := &test.NinOptNative{}
	n, err := scanner.ScanMessage(msg)
	require.Error(t, err)
	require.Equal(t, 10, n)
}

func VerifyVariableintShortened(t *testing.T) {
	buf := newFrame()
	buf.Write([]byte{0xff, 0xff})
	scanner := protoio.NewSeparatedScanner(buf, 1024*1024)
	msg := &test.NinOptNative{}
	n, err := scanner.ScanMessage(msg)
	require.Error(t, err)
	require.Equal(t, 2, n)
}

func VerifyBrief(t *testing.T) {
	buf := newFrame()

	variableintImage := make([]byte, binary.MaxVarintLen64)
	variableintSize := binary.PutUvarint(variableintImage, 100)
	_, err := buf.Write(variableintImage[:variableintSize])
	require.NoError(t, err)

	bz, err := proto.Marshal(&test.NinOptNative{Field15: []byte{0x01, 0x02, 0x03}})
	require.NoError(t, err)
	buf.Write(bz)

	scanner := protoio.NewSeparatedScanner(buf, 1024*1024)
	require.NoError(t, err)
	msg := &test.NinOptNative{}
	n, err := scanner.ScanMessage(msg)
	require.Error(t, err)
	require.Equal(t, variableintSize+len(bz), n)
}
