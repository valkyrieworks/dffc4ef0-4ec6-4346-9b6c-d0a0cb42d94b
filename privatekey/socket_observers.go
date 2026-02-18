package privatekey

import (
	"net"
	"time"

	"github.com/valkyrieworks/vault/ed25519"
	p2plink "github.com/valkyrieworks/p2p/link"
)

const (
	standardDeadlineAllowMoments = 3
)

//
//
type deadlineFault interface {
	Deadline() bool
}

//
//

//
type TCPObserverSetting func(*TCPObserver)

//
//
func TCPObserverDeadlineAllow(deadline time.Duration) TCPObserverSetting {
	return func(tl *TCPObserver) { tl.deadlineAllow = deadline }
}

//
//
func TCPObserverDeadlineScanRecord(deadline time.Duration) TCPObserverSetting {
	return func(tl *TCPObserver) { tl.deadlineScanRecord = deadline }
}

//
var _ net.Listener = (*TCPObserver)(nil)

//
//
type TCPObserver struct {
	*net.TCPObserver

	tokenLinkKey ed25519.PrivateKey

	deadlineAllow    time.Duration
	deadlineScanRecord time.Duration
}

//
//
func NewTCPObserver(ln net.Listener, tokenLinkKey ed25519.PrivateKey) *TCPObserver {
	return &TCPObserver{
		TCPObserver:      ln.(*net.TCPListener),
		tokenLinkKey:    tokenLinkKey,
		deadlineAllow:    time.Second * standardDeadlineAllowMoments,
		deadlineScanRecord: time.Second * standardDeadlineScanRecordMoments,
	}
}

//
func (ln *TCPObserver) Allow() (net.Conn, error) {
	limit := time.Now().Add(ln.deadlineAllow)
	err := ln.SetDeadline(limit)
	if err != nil {
		return nil, err
	}

	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}

	//
	deadlineLink := newDeadlineLink(tc, ln.deadlineScanRecord)
	tokenLink, err := p2plink.CreateTokenLinkage(deadlineLink, ln.tokenLinkKey)
	if err != nil {
		_ = deadlineLink.Close()
		return nil, err
	}

	return tokenLink, nil
}

//
//

//
var _ net.Listener = (*UnixObserver)(nil)

type UnixObserverSetting func(*UnixObserver)

//
//
func UnixObserverDeadlineAllow(deadline time.Duration) UnixObserverSetting {
	return func(ul *UnixObserver) { ul.deadlineAllow = deadline }
}

//
//
func UnixObserverDeadlineScanRecord(deadline time.Duration) UnixObserverSetting {
	return func(ul *UnixObserver) { ul.deadlineScanRecord = deadline }
}

//
//
type UnixObserver struct {
	*net.UnixObserver

	deadlineAllow    time.Duration
	deadlineScanRecord time.Duration
}

//
//
func NewUnixObserver(ln net.Listener) *UnixObserver {
	return &UnixObserver{
		UnixObserver:     ln.(*net.UnixListener),
		deadlineAllow:    time.Second * standardDeadlineAllowMoments,
		deadlineScanRecord: time.Second * standardDeadlineScanRecordMoments,
	}
}

//
func (ln *UnixObserver) Allow() (net.Conn, error) {
	limit := time.Now().Add(ln.deadlineAllow)
	err := ln.SetDeadline(limit)
	if err != nil {
		return nil, err
	}

	tc, err := ln.AcceptUnix()
	if err != nil {
		return nil, err
	}

	//
	link := newDeadlineLink(tc, ln.deadlineScanRecord)

	//
	//

	return link, nil
}

//
//

//
var _ net.Conn = (*deadlineLink)(nil)

//
type deadlineLink struct {
	net.Link
	deadline time.Duration
}

//
func newDeadlineLink(link net.Conn, deadline time.Duration) *deadlineLink {
	return &deadlineLink{
		link,
		deadline,
	}
}

//
func (c deadlineLink) Scan(b []byte) (n int, err error) {
	//
	limit := time.Now().Add(c.deadline)
	err = c.SetReadDeadline(limit)
	if err != nil {
		return
	}

	return c.Link.Read(b)
}

//
func (c deadlineLink) Record(b []byte) (n int, err error) {
	//
	limit := time.Now().Add(c.deadline)
	err = c.SetWriteDeadline(limit)
	if err != nil {
		return
	}

	return c.Link.Write(b)
}
