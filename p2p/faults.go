package p2p

import (
	"fmt"
	"net"
)

//
type FaultRefineDeadline struct{}

func (e FaultRefineDeadline) Failure() string {
	return "REDACTED"
}

//
//
type FaultDeclined struct {
	location              NetworkLocator
	link              net.Conn
	err               error
	id                ID
	equalsAuthBreakdown     bool
	equalsReplicated       bool
	equalsScreened        bool
	equalsUnmatched    bool
	equalsPeerDetailsUnfit bool
	equalsEgo            bool
}

//
func (e FaultDeclined) Location() NetworkLocator {
	return e.location
}

func (e FaultDeclined) Failure() string {
	if e.equalsAuthBreakdown {
		return fmt.Sprintf("REDACTED", e.err)
	}

	if e.equalsReplicated {
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

	if e.equalsScreened {
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

	if e.equalsUnmatched {
		return fmt.Sprintf("REDACTED", e.err)
	}

	if e.equalsPeerDetailsUnfit {
		return fmt.Sprintf("REDACTED", e.err)
	}

	if e.equalsEgo {
		return fmt.Sprintf("REDACTED", e.id)
	}

	return fmt.Sprintf("REDACTED", e.err)
}

//
func (e FaultDeclined) EqualsAuthBreakdown() bool { return e.equalsAuthBreakdown }

//
func (e FaultDeclined) EqualsReplicated() bool { return e.equalsReplicated }

//
func (e FaultDeclined) EqualsScreened() bool { return e.equalsScreened }

//
func (e FaultDeclined) EqualsUnmatched() bool { return e.equalsUnmatched }

//
func (e FaultDeclined) EqualsPeerDetailsUnfit() bool { return e.equalsPeerDetailsUnfit }

//
func (e FaultDeclined) EqualsEgo() bool { return e.equalsEgo }

//
//
type FaultRouterReplicatedNodeUUID struct {
	ID ID
}

func (e FaultRouterReplicatedNodeUUID) Failure() string {
	return fmt.Sprintf("REDACTED", e.ID)
}

//
//
type FaultRouterReplicatedNodeINET struct {
	IP net.IP
}

func (e FaultRouterReplicatedNodeINET) Failure() string {
	return fmt.Sprintf("REDACTED", e.IP.String())
}

//
type FaultRouterRelateTowardEgo struct {
	Location *NetworkLocator
}

func (e FaultRouterRelateTowardEgo) Failure() string {
	return fmt.Sprintf("REDACTED", e.Location)
}

type FaultRouterAuthorizationBreakdown struct {
	Called *NetworkLocator
	Got    ID
}

func (e FaultRouterAuthorizationBreakdown) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Called,
		e.Got,
	)
}

//
type FaultCarrierTerminated struct{}

func (e FaultCarrierTerminated) Failure() string {
	return "REDACTED"
}

//
type FaultNodeDeletion struct{}

func (e FaultNodeDeletion) Failure() string {
	return "REDACTED"
}

//

type FaultNetworkLocatorNegativeUUID struct {
	Location string
}

func (e FaultNetworkLocatorNegativeUUID) Failure() string {
	return fmt.Sprintf("REDACTED", e.Location)
}

type FaultNetworkLocatorUnfit struct {
	Location string
	Err  error
}

func (e FaultNetworkLocatorUnfit) Failure() string {
	return fmt.Sprintf("REDACTED", e.Location, e.Err)
}

type FaultNetworkLocatorSearch struct {
	Location string
	Err  error
}

func (e FaultNetworkLocatorSearch) Failure() string {
	return fmt.Sprintf("REDACTED", e.Location, e.Err)
}

//
//
type FaultPresentlyCallingEitherPresentLocator struct {
	Location string
}

func (e FaultPresentlyCallingEitherPresentLocator) Failure() string {
	return fmt.Sprintf("REDACTED", e.Location)
}
