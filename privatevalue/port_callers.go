package privatevalue

import (
	"errors"
	"net"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	peer2peerlink "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

//
var (
	FaultCallReissueMaximum = errors.New("REDACTED")
)

//
type PortCaller func() (net.Conn, error)

//
//
func CallStreamProc(location string, deadlineRetrievePersist time.Duration, privateToken security.PrivateToken) PortCaller {
	return func() (net.Conn, error) {
		link, err := strongmindnet.Relate(location)
		if err == nil {
			limit := time.Now().Add(deadlineRetrievePersist)
			err = link.SetDeadline(limit)
		}
		if err == nil {
			link, err = peer2peerlink.CreateCredentialLinkage(link, privateToken)
		}
		return link, err
	}
}

//
func CallPosixProc(location string) PortCaller {
	return func() (net.Conn, error) {
		posixLocation := &net.UnixAddr{Name: location, Net: "REDACTED"}
		return net.DialUnix("REDACTED", nil, posixLocation)
	}
}
