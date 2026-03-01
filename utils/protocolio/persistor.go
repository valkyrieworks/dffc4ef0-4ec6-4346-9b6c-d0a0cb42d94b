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

package protocolio

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/cosmos/gogoproto/proto"
)

//
//
//
func FreshSeparatedPersistor(w io.Writer) PersistTerminator {
	return &variableintPersistor{w, nil, nil}
}

type variableintPersistor struct {
	w      io.Writer
	lengthArea []byte
	reserve []byte
}

func (w *variableintPersistor) PersistSignal(msg proto.Message) (int, error) {
	if m, ok := msg.(serializer); ok {
		n, ok := obtainExtent(m)
		if ok {
			if n+binary.MaxVarintLen64 >= len(w.reserve) {
				w.reserve = make([]byte, n+binary.MaxVarintLen64)
			}
			lengthDeactivated := binary.PutUvarint(w.reserve, uint64(n))
			_, err := m.SerializeToward(w.reserve[lengthDeactivated:])
			if err != nil {
				return 0, err
			}
			_, err = w.w.Write(w.reserve[:lengthDeactivated+n])
			return lengthDeactivated + n, err
		}
	}

	//
	if w.lengthArea == nil {
		w.lengthArea = make([]byte, binary.MaxVarintLen64)
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return 0, err
	}
	magnitude := uint64(len(data))
	n := binary.PutUvarint(w.lengthArea, magnitude)
	_, err = w.w.Write(w.lengthArea[:n])
	if err != nil {
		return 0, err
	}
	_, err = w.w.Write(data)
	return len(data) + n, err
}

func (w *variableintPersistor) Shutdown() error {
	if terminator, ok := w.w.(io.Closer); ok {
		return terminator.Close()
	}
	return nil
}

func SerializeSeparated(msg proto.Message) ([]byte, error) {
	var buf bytes.Buffer
	_, err := FreshSeparatedPersistor(&buf).PersistSignal(msg)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
