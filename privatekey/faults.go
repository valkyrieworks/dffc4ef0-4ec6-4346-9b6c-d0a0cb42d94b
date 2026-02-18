package privatekey

import (
	"errors"
	"fmt"
)

//
type TerminusDeadlineFault struct{}

//
func (e TerminusDeadlineFault) Fault() string   { return "REDACTED" }
func (e TerminusDeadlineFault) Deadline() bool   { return true }
func (e TerminusDeadlineFault) Interim() bool { return true }

//
var (
	ErrLinkageDeadline = TerminusDeadlineFault{}
	ErrNoLinkage      = errors.New("REDACTED")
	ErrFetchDeadline       = errors.New("REDACTED")
	ErrRecordDeadline      = errors.New("REDACTED")
)

//
//
type DistantNotaryFault struct {
	//
	Code        int
	Summary string
}

func (e *DistantNotaryFault) Fault() string {
	return fmt.Sprintf("REDACTED", e.Code, e.Summary)
}
