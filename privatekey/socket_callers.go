package privatekey

import (
	"errors"
	"net"
	"time"

	"github.com/valkyrieworks/vault"
	cometnet "github.com/valkyrieworks/utils/net"
	p2plink "github.com/valkyrieworks/p2p/link"
)

//
var (
	ErrCallReprocessMaximum = errors.New("REDACTED")
)

//
type SocketCaller func() (net.Conn, error)

//
//
func CallTCPFn(address string, deadlineScanRecord time.Duration, privateKey vault.PrivateKey) SocketCaller {
	return func() (net.Conn, error) {
		link, err := cometnet.Link(address)
		if err == nil {
			limit := time.Now().Add(deadlineScanRecord)
			err = link.SetDeadline(limit)
		}
		if err == nil {
			link, err = p2plink.CreateTokenLinkage(link, privateKey)
		}
		return link, err
	}
}

//
func CallUnixFn(address string) SocketCaller {
	return func() (net.Conn, error) {
		unixAddress := &net.UnixAddr{Name: address, Net: "REDACTED"}
		return net.DialUnix("REDACTED", nil, unixAddress)
	}
}
