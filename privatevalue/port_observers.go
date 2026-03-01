package privatevalue

import (
	"net"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	peer2peerlink "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

const (
	fallbackDeadlineEmbraceMoments = 3
)

//
//
type deadlineFailure interface {
	Deadline() bool
}

//
//

//
type TcpsocketObserverSelection func(*TcpsocketObserver)

//
//
func TcpsocketObserverDeadlineEmbrace(deadline time.Duration) TcpsocketObserverSelection {
	return func(tl *TcpsocketObserver) { tl.deadlineEmbrace = deadline }
}

//
//
func TcpsocketObserverDeadlineFetchPersist(deadline time.Duration) TcpsocketObserverSelection {
	return func(tl *TcpsocketObserver) { tl.deadlineRetrievePersist = deadline }
}

//
var _ net.Listener = (*TcpsocketObserver)(nil)

//
//
type TcpsocketObserver struct {
	*net.TcpsocketObserver

	credentialLinkToken edwards25519.PrivateToken

	deadlineEmbrace    time.Duration
	deadlineRetrievePersist time.Duration
}

//
//
func FreshTcpsocketObserver(ln net.Listener, credentialLinkToken edwards25519.PrivateToken) *TcpsocketObserver {
	return &TcpsocketObserver{
		TcpsocketObserver:      ln.(*net.TCPListener),
		credentialLinkToken:    credentialLinkToken,
		deadlineEmbrace:    time.Second * fallbackDeadlineEmbraceMoments,
		deadlineRetrievePersist: time.Second * fallbackDeadlineRetrievePersistMoments,
	}
}

//
func (ln *TcpsocketObserver) Embrace() (net.Conn, error) {
	limit := time.Now().Add(ln.deadlineEmbrace)
	err := ln.SetDeadline(limit)
	if err != nil {
		return nil, err
	}

	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}

	//
	deadlineLink := freshDeadlineLink(tc, ln.deadlineRetrievePersist)
	credentialLink, err := peer2peerlink.CreateCredentialLinkage(deadlineLink, ln.credentialLinkToken)
	if err != nil {
		_ = deadlineLink.Close()
		return nil, err
	}

	return credentialLink, nil
}

//
//

//
var _ net.Listener = (*PosixObserver)(nil)

type PosixObserverSelection func(*PosixObserver)

//
//
func PosixObserverDeadlineEmbrace(deadline time.Duration) PosixObserverSelection {
	return func(ul *PosixObserver) { ul.deadlineEmbrace = deadline }
}

//
//
func PosixObserverDeadlineFetchPersist(deadline time.Duration) PosixObserverSelection {
	return func(ul *PosixObserver) { ul.deadlineRetrievePersist = deadline }
}

//
//
type PosixObserver struct {
	*net.PosixObserver

	deadlineEmbrace    time.Duration
	deadlineRetrievePersist time.Duration
}

//
//
func FreshPosixObserver(ln net.Listener) *PosixObserver {
	return &PosixObserver{
		PosixObserver:     ln.(*net.UnixListener),
		deadlineEmbrace:    time.Second * fallbackDeadlineEmbraceMoments,
		deadlineRetrievePersist: time.Second * fallbackDeadlineRetrievePersistMoments,
	}
}

//
func (ln *PosixObserver) Embrace() (net.Conn, error) {
	limit := time.Now().Add(ln.deadlineEmbrace)
	err := ln.SetDeadline(limit)
	if err != nil {
		return nil, err
	}

	tc, err := ln.AcceptUnix()
	if err != nil {
		return nil, err
	}

	//
	link := freshDeadlineLink(tc, ln.deadlineRetrievePersist)

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
func freshDeadlineLink(link net.Conn, deadline time.Duration) *deadlineLink {
	return &deadlineLink{
		link,
		deadline,
	}
}

//
func (c deadlineLink) Obtain(b []byte) (n int, err error) {
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
