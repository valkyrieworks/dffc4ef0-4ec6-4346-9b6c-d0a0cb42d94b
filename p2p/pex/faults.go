package pex

import (
	"errors"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

type FaultLocationRegisterUnDirectable struct {
	Location *p2p.NetworkLocator
}

func (err FaultLocationRegisterUnDirectable) Failure() string {
	return fmt.Sprintf("REDACTED", err.Location)
}

type faultLocationRegisterAgedLocatorFreshSegment struct {
	Location     *p2p.NetworkLocator
	SegmentUUID int
}

func (err faultLocationRegisterAgedLocatorFreshSegment) Failure() string {
	return fmt.Sprintf("REDACTED"+
		"REDACTED",
		err.Location, err.SegmentUUID)
}

type FaultLocationRegisterEgo struct {
	Location *p2p.NetworkLocator
}

func (err FaultLocationRegisterEgo) Failure() string {
	return fmt.Sprintf("REDACTED", err.Location)
}

type FaultLocationRegisterSecluded struct {
	Location *p2p.NetworkLocator
}

func (err FaultLocationRegisterSecluded) Failure() string {
	return fmt.Sprintf("REDACTED", err.Location)
}

func (err FaultLocationRegisterSecluded) SecludedLocation() bool {
	return true
}

type FaultLocationRegisterSecludedOrigin struct {
	Src *p2p.NetworkLocator
}

func (err FaultLocationRegisterSecludedOrigin) Failure() string {
	return fmt.Sprintf("REDACTED", err.Src)
}

func (err FaultLocationRegisterSecludedOrigin) SecludedLocation() bool {
	return true
}

type FaultLocationRegisterVoidLocation struct {
	Location *p2p.NetworkLocator
	Src  *p2p.NetworkLocator
}

func (err FaultLocationRegisterVoidLocation) Failure() string {
	return fmt.Sprintf("REDACTED", err.Location, err.Src)
}

type FaultLocationRegisterUnfitLocation struct {
	Location    *p2p.NetworkLocator
	LocationFault error
}

func (err FaultLocationRegisterUnfitLocation) Failure() string {
	return fmt.Sprintf("REDACTED", err.Location, err.LocationFault)
}

//
type FaultLocatorProhibited struct {
	Location *p2p.NetworkLocator
}

func (err FaultLocatorProhibited) Failure() string {
	return fmt.Sprintf("REDACTED", err.Location)
}

//
var FaultUnpromptedCatalog = errors.New("REDACTED")
