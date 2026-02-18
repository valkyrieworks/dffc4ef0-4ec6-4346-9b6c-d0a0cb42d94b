package settings

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyRPCHostRecord             = errors.New("REDACTED")
	ErrNoSufficientRPCHosts             = errors.New("REDACTED")
	ErrInadequateDetectionTime       = errors.New("REDACTED")
	ErrInadequateSegmentQueryDeadline = errors.New("REDACTED")
	ErrUnclearTraceLayout                = errors.New("REDACTED")
	ErrEnrollmentBufferVolumeCorrupt   = fmt.Errorf("REDACTED", minimumEnrollmentBufferVolume)
)

//
type ErrInSegment struct {
	Err     error
	Segment string
}

func (e ErrInSegment) Fault() string {
	return fmt.Sprintf("REDACTED", e.Segment, e.Err.Error())
}

func (e ErrInSegment) Disclose() error {
	return e.Err
}

type ErrObsoleteChainconnectRelease struct {
	Release string
	Permitted []string
}

func (e ErrObsoleteChainconnectRelease) Fault() string {
	return fmt.Sprintf("REDACTED", e.Release, e.Permitted)
}

type ErrUnclearChainconnectRelease struct {
	Release string
}

func (e ErrUnclearChainconnectRelease) Fault() string {
	return fmt.Sprintf("REDACTED", e.Release)
}
