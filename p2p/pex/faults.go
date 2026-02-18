package pex

import (
	"errors"
	"fmt"

	"github.com/valkyrieworks/p2p"
)

type ErrAddressRegistryNotForwardable struct {
	Address *p2p.NetLocation
}

func (err ErrAddressRegistryNotForwardable) Fault() string {
	return fmt.Sprintf("REDACTED", err.Address)
}

type errAddressRegistryAgedLocationNewSegment struct {
	Address     *p2p.NetLocation
	ContainerUID int
}

func (err errAddressRegistryAgedLocationNewSegment) Fault() string {
	return fmt.Sprintf("REDACTED"+
		"REDACTED",
		err.Address, err.ContainerUID)
}

type ErrAddressRegistryEgo struct {
	Address *p2p.NetLocation
}

func (err ErrAddressRegistryEgo) Fault() string {
	return fmt.Sprintf("REDACTED", err.Address)
}

type ErrAddressRegistryInternal struct {
	Address *p2p.NetLocation
}

func (err ErrAddressRegistryInternal) Fault() string {
	return fmt.Sprintf("REDACTED", err.Address)
}

func (err ErrAddressRegistryInternal) InternalAddress() bool {
	return true
}

type ErrAddressRegistryInternalOrigin struct {
	Src *p2p.NetLocation
}

func (err ErrAddressRegistryInternalOrigin) Fault() string {
	return fmt.Sprintf("REDACTED", err.Src)
}

func (err ErrAddressRegistryInternalOrigin) InternalAddress() bool {
	return true
}

type ErrAddressRegistryNullAddress struct {
	Address *p2p.NetLocation
	Src  *p2p.NetLocation
}

func (err ErrAddressRegistryNullAddress) Fault() string {
	return fmt.Sprintf("REDACTED", err.Address, err.Src)
}

type ErrAddressRegistryCorruptAddress struct {
	Address    *p2p.NetLocation
	AddressErr error
}

func (err ErrAddressRegistryCorruptAddress) Fault() string {
	return fmt.Sprintf("REDACTED", err.Address, err.AddressErr)
}

//
type ErrLocationProhibited struct {
	Address *p2p.NetLocation
}

func (err ErrLocationProhibited) Fault() string {
	return fmt.Sprintf("REDACTED", err.Address)
}

//
var ErrUninvitedCatalog = errors.New("REDACTED")
