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
//
//

package protoio

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/cosmos/gogoproto/proto"
)

//
//
//
func NewSeparatedRecorder(w io.Writer) RecordTerminator {
	return &variableintRecorder{w, nil, nil}
}

type variableintRecorder struct {
	w      io.Writer
	sizeImage []byte
	buffer []byte
}

func (w *variableintRecorder) RecordMessage(msg proto.Message) (int, error) {
	if m, ok := msg.(serializer); ok {
		n, ok := fetchVolume(m)
		if ok {
			if n+binary.MaxVarintLen64 >= len(w.buffer) {
				w.buffer = make([]byte, n+binary.MaxVarintLen64)
			}
			sizeOffset := binary.PutUvarint(w.buffer, uint64(n))
			_, err := m.SerializeTo(w.buffer[sizeOffset:])
			if err != nil {
				return 0, err
			}
			_, err = w.w.Write(w.buffer[:sizeOffset+n])
			return sizeOffset + n, err
		}
	}

	//
	if w.sizeImage == nil {
		w.sizeImage = make([]byte, binary.MaxVarintLen64)
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return 0, err
	}
	extent := uint64(len(data))
	n := binary.PutUvarint(w.sizeImage, extent)
	_, err = w.w.Write(w.sizeImage[:n])
	if err != nil {
		return 0, err
	}
	_, err = w.w.Write(data)
	return len(data) + n, err
}

func (w *variableintRecorder) End() error {
	if terminator, ok := w.w.(io.Closer); ok {
		return terminator.Close()
	}
	return nil
}

func SerializeSeparated(msg proto.Message) ([]byte, error) {
	var buf bytes.Buffer
	_, err := NewSeparatedRecorder(&buf).RecordMessage(msg)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
