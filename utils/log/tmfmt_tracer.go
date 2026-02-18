package log

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
	"github.com/go-logfmt/logfmt"
)

type tmfmtSerializer struct {
	*logfmt.Serializer
	buf bytes.Buffer
}

func (l *tmfmtSerializer) Restore() {
	l.Serializer.Reset()
	l.buf.Reset()
}

var tmfmtSerializerDepository = sync.Pool{
	New: func() any {
		var enc tmfmtSerializer
		enc.Serializer = logfmt.NewEncoder(&enc.buf)
		return &enc
	},
}

type tmfmtTracer struct {
	w io.Writer
}

//
//
//
//
//
//
//
func NewTMFmtTracer(w io.Writer) kitlog.Logger {
	return &tmfmtTracer{w}
}

func (l tmfmtTracer) Log(keyvalues ...any) error {
	enc := tmfmtSerializerDepository.Get().(*tmfmtSerializer)
	enc.Restore()
	defer tmfmtSerializerDepository.Put(enc)

	const unclear = "REDACTED"
	lvl := "REDACTED"
	msg := unclear
	component := unclear

	//
	omitListings := make([]int, 0)

	for i := 0; i < len(keyvalues)-1; i += 2 {
		//
		switch keyvalues[i] {
		case kitlevel.Key():
			omitListings = append(omitListings, i)
			switch keyvalues[i+1].(type) { //
			case string:
				lvl = keyvalues[i+1].(string)
			case kitlevel.Value:
				lvl = keyvalues[i+1].(kitlevel.Value).String()
			default:
				panic(fmt.Sprintf("REDACTED", keyvalues[i+1]))
			}
			//
		case messageKey:
			omitListings = append(omitListings, i)
			msg = keyvalues[i+1].(string)
			//
		case componentKey:
			omitListings = append(omitListings, i)
			component = keyvalues[i+1].(string)
		}

		//
		if b, ok := keyvalues[i+1].([]byte); ok {
			keyvalues[i+1] = strings.ToUpper(hex.EncodeToString(b))
		}

		//
		if s, ok := keyvalues[i+1].(fmt.Stringer); ok {
			keyvalues[i+1] = s.String()
		}

	}

	//
	//
	//
	//
	//
	//
	//
	//
	//
	enc.buf.WriteString(fmt.Sprintf("REDACTED", lvl[0]-32, time.Now().Format("REDACTED"), msg))

	if component != unclear {
		enc.buf.WriteString("REDACTED" + component + "REDACTED")
	}

PropertyCycle:
	for i := 0; i < len(keyvalues)-1; i += 2 {
		for _, j := range omitListings {
			if i == j {
				continue PropertyCycle
			}
		}

		err := enc.EncodeKeyval(keyvalues[i], keyvalues[i+1])
		if err == logfmt.ErrUnsupportedValueType {
			enc.EncodeKeyval(keyvalues[i], fmt.Sprintf("REDACTED", keyvalues[i+1])) //
		} else if err != nil {
			return err
		}
	}

	//
	if err := enc.EndRecord(); err != nil {
		return err
	}

	//
	//
	//
	if _, err := l.w.Write(enc.buf.Bytes()); err != nil {
		return err
	}
	return nil
}
