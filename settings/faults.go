package settings

import (
	"errors"
	"fmt"
)

var (
	FaultBlankRemoteDaemonPiece             = errors.New("REDACTED")
	FaultNegationAmpleRemoteHosts             = errors.New("REDACTED")
	FaultLackingExplorationMoment       = errors.New("REDACTED")
	FaultLackingSegmentSolicitDeadline = errors.New("REDACTED")
	FaultUnfamiliarReportLayout                = errors.New("REDACTED")
	FaultListeningReserveExtentUnfit   = fmt.Errorf("REDACTED", minimumListeningReserveExtent)
)

//
type FaultInsideSegment struct {
	Err     error
	Segment string
}

func (e FaultInsideSegment) Failure() string {
	return fmt.Sprintf("REDACTED", e.Segment, e.Err.Error())
}

func (e FaultInsideSegment) Disclose() error {
	return e.Err
}

type FaultObsoleteChainchronizeEdition struct {
	Edition string
	Permitted []string
}

func (e FaultObsoleteChainchronizeEdition) Failure() string {
	return fmt.Sprintf("REDACTED", e.Edition, e.Permitted)
}

type FaultUnfamiliarChainchronizeEdition struct {
	Edition string
}

func (e FaultUnfamiliarChainchronizeEdition) Failure() string {
	return fmt.Sprintf("REDACTED", e.Edition)
}
