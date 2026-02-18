package p2p

import (
	"fmt"
	"net"
)

//
type ErrRefineDeadline struct{}

func (e ErrRefineDeadline) Fault() string {
	return "REDACTED"
}

//
//
type ErrDeclined struct {
	address              NetLocation
	link              net.Conn
	err               error
	id                ID
	isAuthBreakdown     bool
	isReplicated       bool
	isScreened        bool
	isDiscordant    bool
	isMemberDetailsCorrupt bool
	isEgo            bool
}

//
func (e ErrDeclined) Address() NetLocation {
	return e.address
}

func (e ErrDeclined) Fault() string {
	if e.isAuthBreakdown {
		return fmt.Sprintf("REDACTED", e.err)
	}

	if e.isReplicated {
		if e.link != nil {
			return fmt.Sprintf(
				"REDACTED",
				e.link.RemoteAddr().String(),
			)
		}
		if e.id != "REDACTED" {
			return fmt.Sprintf("REDACTED", e.id)
		}
	}

	if e.isScreened {
		if e.link != nil {
			return fmt.Sprintf(
				"REDACTED",
				e.link.RemoteAddr().String(),
				e.err,
			)
		}

		if e.id != "REDACTED" {
			return fmt.Sprintf("REDACTED", e.id, e.err)
		}
	}

	if e.isDiscordant {
		return fmt.Sprintf("REDACTED", e.err)
	}

	if e.isMemberDetailsCorrupt {
		return fmt.Sprintf("REDACTED", e.err)
	}

	if e.isEgo {
		return fmt.Sprintf("REDACTED", e.id)
	}

	return fmt.Sprintf("REDACTED", e.err)
}

//
func (e ErrDeclined) IsAuthBreakdown() bool { return e.isAuthBreakdown }

//
func (e ErrDeclined) IsReplicated() bool { return e.isReplicated }

//
func (e ErrDeclined) IsScreened() bool { return e.isScreened }

//
func (e ErrDeclined) IsDiscordant() bool { return e.isDiscordant }

//
func (e ErrDeclined) IsMemberDetailsCorrupt() bool { return e.isMemberDetailsCorrupt }

//
func (e ErrDeclined) IsEgo() bool { return e.isEgo }

//
//
type ErrRouterReplicatedNodeUID struct {
	ID ID
}

func (e ErrRouterReplicatedNodeUID) Fault() string {
	return fmt.Sprintf("REDACTED", e.ID)
}

//
//
type ErrRouterReplicatedNodeIP struct {
	IP net.IP
}

func (e ErrRouterReplicatedNodeIP) Fault() string {
	return fmt.Sprintf("REDACTED", e.IP.String())
}

//
type ErrRouterEstablishToEgo struct {
	Address *NetLocation
}

func (e ErrRouterEstablishToEgo) Fault() string {
	return fmt.Sprintf("REDACTED", e.Address)
}

type ErrRouterAuthorizationBreakdown struct {
	Called *NetLocation
	Got    ID
}

func (e ErrRouterAuthorizationBreakdown) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Called,
		e.Got,
	)
}

//
type ErrCarrierHalted struct{}

func (e ErrCarrierHalted) Fault() string {
	return "REDACTED"
}

//
type ErrNodeDeletion struct{}

func (e ErrNodeDeletion) Fault() string {
	return "REDACTED"
}

//

type ErrNetLocationNoUID struct {
	Address string
}

func (e ErrNetLocationNoUID) Fault() string {
	return fmt.Sprintf("REDACTED", e.Address)
}

type ErrNetLocationCorrupt struct {
	Address string
	Err  error
}

func (e ErrNetLocationCorrupt) Fault() string {
	return fmt.Sprintf("REDACTED", e.Address, e.Err)
}

type ErrNetLocationSearch struct {
	Address string
	Err  error
}

func (e ErrNetLocationSearch) Fault() string {
	return fmt.Sprintf("REDACTED", e.Address, e.Err)
}

//
//
type ErrPresentlyCallingOrCurrentLocation struct {
	Address string
}

func (e ErrPresentlyCallingOrCurrentLocation) Fault() string {
	return fmt.Sprintf("REDACTED", e.Address)
}
