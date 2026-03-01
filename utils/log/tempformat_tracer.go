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

type tempformatSerializer struct {
	*logfmt.Serializer
	buf bytes.Buffer
}

func (l *tempformatSerializer) Restore() {
	l.Serializer.Reset()
	l.buf.Reset()
}

var tempformatSerializerHub = sync.Pool{
	New: func() any {
		var enc tempformatSerializer
		enc.Serializer = logfmt.NewEncoder(&enc.buf)
		return &enc
	},
}

type tempformatTracer struct {
	w io.Writer
}

//
//
//
//
//
//
//
func FreshTEMPTextformatTracer(w io.Writer) kitlog.Logger {
	return &tempformatTracer{w}
}

func (l tempformatTracer) Log(tokvals ...any) error {
	enc := tempformatSerializerHub.Get().(*tempformatSerializer)
	enc.Restore()
	defer tempformatSerializerHub.Put(enc)

	const unfamiliar = "REDACTED"
	lvl := "REDACTED"
	msg := unfamiliar
	component := unfamiliar

	//
	omitPositions := make([]int, 0)

	for i := 0; i < len(tokvals)-1; i += 2 {
		//
		switch tokvals[i] {
		case kitlevel.Key():
			omitPositions = append(omitPositions, i)
			switch tokvals[i+1].(type) { //
			case string:
				lvl = tokvals[i+1].(string)
			case kitlevel.Value:
				lvl = tokvals[i+1].(kitlevel.Value).String()
			default:
				panic(fmt.Sprintf("REDACTED", tokvals[i+1]))
			}
			//
		case signalToken:
			omitPositions = append(omitPositions, i)
			msg = tokvals[i+1].(string)
			//
		case componentToken:
			omitPositions = append(omitPositions, i)
			component = tokvals[i+1].(string)
		}

		//
		if b, ok := tokvals[i+1].([]byte); ok {
			tokvals[i+1] = strings.ToUpper(hex.EncodeToString(b))
		}

		//
		if s, ok := tokvals[i+1].(fmt.Stringer); ok {
			tokvals[i+1] = s.String()
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

	if component != unfamiliar {
		enc.buf.WriteString("REDACTED" + component + "REDACTED")
	}

TokvalueCycle:
	for i := 0; i < len(tokvals)-1; i += 2 {
		for _, j := range omitPositions {
			if i == j {
				continue TokvalueCycle
			}
		}

		err := enc.EncodeKeyval(tokvals[i], tokvals[i+1])
		if err == logfmt.ErrUnsupportedValueType {
			enc.EncodeKeyval(tokvals[i], fmt.Sprintf("REDACTED", tokvals[i+1])) //
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
