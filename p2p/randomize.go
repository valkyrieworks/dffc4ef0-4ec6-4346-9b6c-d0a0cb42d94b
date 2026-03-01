package p2p

import (
	"net"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
//
type TestedLinkage struct {
	link net.Conn

	mtx    commitchronize.Exclusion
	initiate  <-chan time.Time
	dynamic bool

	settings *settings.RandomizeLinkSettings
}

//
func RandomizeLink(link net.Conn) net.Conn {
	return RandomizeLinkOriginatingSettings(link, settings.FallbackRandomizeLinkSettings())
}

//
//
func RandomizeLinkOriginatingSettings(link net.Conn, settings *settings.RandomizeLinkSettings) net.Conn {
	return &TestedLinkage{
		link:   link,
		initiate:  make(<-chan time.Time),
		dynamic: true,
		settings: settings,
	}
}

//
//
func RandomizeLinkSubsequent(link net.Conn, d time.Duration) net.Conn {
	return RandomizeLinkSubsequentOriginatingSettings(link, d, settings.FallbackRandomizeLinkSettings())
}

//
//
func RandomizeLinkSubsequentOriginatingSettings(
	link net.Conn,
	d time.Duration,
	settings *settings.RandomizeLinkSettings,
) net.Conn {
	return &TestedLinkage{
		link:   link,
		initiate:  time.After(d),
		dynamic: false,
		settings: settings,
	}
}

//
func (fc *TestedLinkage) Settings() *settings.RandomizeLinkSettings {
	return fc.settings
}

//
func (fc *TestedLinkage) Obtain(data []byte) (n int, err error) {
	if fc.randomize() {
		return 0, nil
	}
	return fc.link.Read(data)
}

//
func (fc *TestedLinkage) Record(data []byte) (n int, err error) {
	if fc.randomize() {
		return 0, nil
	}
	return fc.link.Write(data)
}

//
func (fc *TestedLinkage) Shutdown() error { return fc.link.Close() }

//
func (fc *TestedLinkage) RegionalLocation() net.Addr { return fc.link.LocalAddr() }

//
func (fc *TestedLinkage) DistantLocation() net.Addr { return fc.link.RemoteAddr() }

//
func (fc *TestedLinkage) AssignExpiration(t time.Time) error { return fc.link.SetDeadline(t) }

//
func (fc *TestedLinkage) AssignFetchLimit(t time.Time) error {
	return fc.link.SetReadDeadline(t)
}

//
func (fc *TestedLinkage) AssignPersistLimit(t time.Time) error {
	return fc.link.SetWriteDeadline(t)
}

func (fc *TestedLinkage) unpredictableInterval() time.Duration {
	maximumDeferralMilli := int(fc.settings.MaximumDeferral.Nanoseconds() / 1000)
	return time.Millisecond * time.Duration(commitrand.Int()%maximumDeferralMilli) //
}

//
//
func (fc *TestedLinkage) randomize() bool {
	if !fc.mustRandomize() {
		return false
	}

	switch fc.settings.Style {
	case settings.RandomizeStyleDiscard:
		//
		r := commitrand.Float64()
		switch {
		case r <= fc.settings.LikelihoodDiscardReadwrite:
			return true
		case r < fc.settings.LikelihoodDiscardReadwrite+fc.settings.LikelihoodDiscardLink:
			//
			//
			fc.Shutdown()
			return true
		case r < fc.settings.LikelihoodDiscardReadwrite+fc.settings.LikelihoodDiscardLink+fc.settings.LikelihoodSnooze:
			time.Sleep(fc.unpredictableInterval())
		}
	case settings.RandomizeStyleDeferral:
		//
		time.Sleep(fc.unpredictableInterval())
	}
	return false
}

func (fc *TestedLinkage) mustRandomize() bool {
	if fc.dynamic {
		return true
	}

	fc.mtx.Lock()
	defer fc.mtx.Unlock()

	select {
	case <-fc.initiate:
		fc.dynamic = true
		return true
	default:
		return false
	}
}
