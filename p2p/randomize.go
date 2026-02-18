package p2p

import (
	"net"
	"time"

	"github.com/valkyrieworks/settings"
	engineseed "github.com/valkyrieworks/utils/random"
	engineconnect "github.com/valkyrieworks/utils/align"
)

//
//
type RandomizedLinkage struct {
	link net.Conn

	mtx    engineconnect.Lock
	begin  <-chan time.Time
	enabled bool

	settings *settings.RandomizeLinkSettings
}

//
func RandomizeLink(link net.Conn) net.Conn {
	return RandomizeLinkFromSettings(link, settings.StandardRandomizeLinkSettings())
}

//
//
func RandomizeLinkFromSettings(link net.Conn, settings *settings.RandomizeLinkSettings) net.Conn {
	return &RandomizedLinkage{
		link:   link,
		begin:  make(<-chan time.Time),
		enabled: true,
		settings: settings,
	}
}

//
//
func RandomizeLinkAfter(link net.Conn, d time.Duration) net.Conn {
	return RandomizeLinkAfterFromSettings(link, d, settings.StandardRandomizeLinkSettings())
}

//
//
func RandomizeLinkAfterFromSettings(
	link net.Conn,
	d time.Duration,
	settings *settings.RandomizeLinkSettings,
) net.Conn {
	return &RandomizedLinkage{
		link:   link,
		begin:  time.After(d),
		enabled: false,
		settings: settings,
	}
}

//
func (fc *RandomizedLinkage) Settings() *settings.RandomizeLinkSettings {
	return fc.settings
}

//
func (fc *RandomizedLinkage) Scan(data []byte) (n int, err error) {
	if fc.randomize() {
		return 0, nil
	}
	return fc.link.Read(data)
}

//
func (fc *RandomizedLinkage) Record(data []byte) (n int, err error) {
	if fc.randomize() {
		return 0, nil
	}
	return fc.link.Write(data)
}

//
func (fc *RandomizedLinkage) End() error { return fc.link.Close() }

//
func (fc *RandomizedLinkage) NativeAddress() net.Addr { return fc.link.LocalAddr() }

//
func (fc *RandomizedLinkage) DistantAddress() net.Addr { return fc.link.RemoteAddr() }

//
func (fc *RandomizedLinkage) CollectionLimit(t time.Time) error { return fc.link.SetDeadline(t) }

//
func (fc *RandomizedLinkage) CollectionReadLimit(t time.Time) error {
	return fc.link.SetReadDeadline(t)
}

//
func (fc *RandomizedLinkage) CollectionRecordLimit(t time.Time) error {
	return fc.link.SetWriteDeadline(t)
}

func (fc *RandomizedLinkage) arbitraryPeriod() time.Duration {
	maximumDeferralMillis := int(fc.settings.MaximumDeferral.Nanoseconds() / 1000)
	return time.Millisecond * time.Duration(engineseed.Int()%maximumDeferralMillis) //
}

//
//
func (fc *RandomizedLinkage) randomize() bool {
	if !fc.mustRandomize() {
		return false
	}

	switch fc.settings.Style {
	case settings.RandomizeStyleDiscard:
		//
		r := engineseed.Float64()
		switch {
		case r <= fc.settings.LikelihoodDiscardReadwrite:
			return true
		case r < fc.settings.LikelihoodDiscardReadwrite+fc.settings.LikelihoodDiscardLink:
			//
			//
			fc.End()
			return true
		case r < fc.settings.LikelihoodDiscardReadwrite+fc.settings.LikelihoodDiscardLink+fc.settings.LikelihoodPause:
			time.Sleep(fc.arbitraryPeriod())
		}
	case settings.RandomizeStyleDeferral:
		//
		time.Sleep(fc.arbitraryPeriod())
	}
	return false
}

func (fc *RandomizedLinkage) mustRandomize() bool {
	if fc.enabled {
		return true
	}

	fc.mtx.Lock()
	defer fc.mtx.Unlock()

	select {
	case <-fc.begin:
		fc.enabled = true
		return true
	default:
		return false
	}
}
