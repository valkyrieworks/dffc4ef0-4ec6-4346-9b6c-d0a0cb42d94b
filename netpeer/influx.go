package netpeer

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
const ProtocolUIDPrefix = "REDACTED"

//
const DeadlineInflux = 10 * time.Second

//
//
func ProtocolUID(conduitUID byte) protocol.ID {
	return protocol.ID(
		fmt.Sprintf("REDACTED", ProtocolUIDPrefix, conduitUID),
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
		heading  = uint64toUvarint(uint64(len(data)))
		shipment = append(heading, data...)
	)

	octetsInscribed, err := s.Write(shipment)
	if err != nil {
		err = errors.Wrapf(err, "REDACTED", octetsInscribed, len(shipment))
	}

	return octetsInscribed, err
}

//
//
//
func InfluxRecordEnd(s network.Stream, data []byte) (err error) {
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

	if err := endInflux(s); err != nil {
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

	scanner := bufio.NewReader(s)

	shipmentVolume, err := binary.ReadUvarint(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	shipment, err := fetchPrecisely(scanner, shipmentVolume)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

//
//
func InfluxFetchEnd(s network.Stream) (shipment []byte, err error) {
	defer func() {
		if err != nil {
			//
			_ = s.Reset()
		}
	}()

	shipment, err = InfluxFetch(s)
	if err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	if err := endInflux(s); err != nil {
		return nil, errors.Wrap(err, "REDACTED")
	}

	return shipment, nil
}

//
func fetchPrecisely(r io.Reader, volume uint64) ([]byte, error) {
	var (
		out       = make([]byte, volume)
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
		case eof && octetsFetch == volume:
			//
			return out, nil
		case eof && octetsFetch != volume:
			//
			return nil, errors.Wrapf(err, "REDACTED", octetsFetch, volume)
		case err != nil:
			//
			return nil, errors.Wrapf(err, "REDACTED", octetsFetch, volume)
		case octetsFetch < volume:
			//
			continue
		case octetsFetch > volume:
			//
			return nil, errors.Errorf("REDACTED", octetsFetch, volume)
		default:
			//
			return out, nil
		}
	}
}

func endInflux(s network.Stream) error {
	err := s.Close()
	switch {
	case isErrAborted(err):
		//
	case err != nil:
		return errors.Wrap(err, "REDACTED")
	}

	return nil
}

//
func isErrAborted(err error) bool {
	if err == nil {
		return false
	}

	const template = "REDACTED"

	return strings.Contains(err.Error(), template)
}

func uint64toUvarint(len uint64) []byte {
	out := make([]byte, binary.MaxVarintLen64)
	octetsInscribed := binary.PutUvarint(out, len)

	return out[:octetsInscribed]
}
