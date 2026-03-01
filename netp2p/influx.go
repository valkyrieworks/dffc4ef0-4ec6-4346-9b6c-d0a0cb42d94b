package netp2p

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/pkg/errors"
)

//
const SchemeUUIDHeading = "REDACTED"

//
const DeadlineInflux = 10 * time.Second

//
//
const MaximumInfluxExtent = 4 * (1 << 20)

//
//
func SchemeUUID(conduitUUID byte) protocol.ID {
	return protocol.ID(
		fmt.Sprintf("REDACTED", SchemeUUIDHeading, conduitUUID),
	)
}

//
//
//
func InfluxRecord(s network.Stream, data []byte) (int, error) {
	switch {
	case len(data) == 0:
		//
		return 0, nil
	case s.Conn().IsClosed():
		return 0, fmt.Errorf("REDACTED")
	}

	//
	var (
		heading  = uinttobyteVaruint(uint64(len(data)))
		content = append(heading, data...)
	)

	octetsRecorded, err := s.Write(content)
	if err != nil {
		err = errors.Wrapf(err, "REDACTED", octetsRecorded, len(content))
	}

	return octetsRecorded, err
}

//
//
//
func InfluxRecordShutdown(s network.Stream, data []byte) (err error) {
	defer func() {
		if err != nil {
			//
			_ = s.Reset()
		}
	}()

	//

	if _, err := InfluxRecord(s, data); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	if err := shutdownInflux(s); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	return nil
}

//
//
func InfluxFetch(s network.Stream) ([]byte, error) {
	if s.Conn().IsClosed() {
		return nil, fmt.Errorf("REDACTED")
	}

	fetcher := bufio.NewReader(s)

	//
	workloadExtent, err := binary.ReadUvarint(fetcher)
	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	if workloadExtent > MaximumInfluxExtent {
		return nil, errors.Errorf("REDACTED", workloadExtent, MaximumInfluxExtent)
	}

	content, err := fetchPrecisely(fetcher, workloadExtent)
	if err != nil {
		return nil, err
	}

	return content, nil
}

//
//
func InfluxFetchShutdown(s network.Stream) (content []byte, err error) {
	defer func() {
		if err != nil {
			//
			_ = s.Reset()
		}
	}()

	content, err = InfluxFetch(s)
	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	if err := shutdownInflux(s); err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	return content, nil
}

//
func fetchPrecisely(r io.Reader, extent uint64) ([]byte, error) {
	var (
		out       = make([]byte, extent)
		octetsFetch uint64
		n         int
		err       error
		eof       bool
	)

	for {
		n, err = r.Read(out[octetsFetch:])
		eof = errors.Is(err, io.EOF)

		octetsFetch += uint64(n)

		switch {
		case eof && octetsFetch == extent:
			//
			return out, nil
		case eof && octetsFetch != extent:
			//
			return nil, errors.Wrapf(err, "REDACTED", octetsFetch, extent)
		case err != nil:
			//
			return nil, errors.Wrapf(err, "REDACTED", octetsFetch, extent)
		case octetsFetch < extent:
			//
			continue
		case octetsFetch > extent:
			//
			return nil, errors.Errorf("REDACTED", octetsFetch, extent)
		default:
			//
			return out, nil
		}
	}
}

func shutdownInflux(s network.Stream) error {
	err := s.Close()
	switch {
	case equalsFaultAborted(err):
		//
	case err != nil:
		return errors.Wrap(err, "REDACTED")
	}

	return nil
}

//
func equalsFaultAborted(err error) bool {
	if err == nil {
		return false
	}

	const template = "REDACTED"

	return strings.Contains(err.Error(), template)
}

func uinttobyteVaruint(len uint64) []byte {
	out := make([]byte, binary.MaxVarintLen64)
	octetsRecorded := binary.PutUvarint(out, len)

	return out[:octetsRecorded]
}
