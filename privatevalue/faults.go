package privatevalue

import (
	"errors"
	"fmt"
)

//
type GatewayDeadlineFailure struct{}

//
func (e GatewayDeadlineFailure) Failure() string   { return "REDACTED" }
func (e GatewayDeadlineFailure) Deadline() bool   { return true }
func (e GatewayDeadlineFailure) Interim() bool { return true }

//
var (
	FaultLinkageDeadline = GatewayDeadlineFailure{}
	FaultNegativeLinkage      = errors.New("REDACTED")
	FaultRetrieveDeadline       = errors.New("REDACTED")
	FaultPersistDeadline      = errors.New("REDACTED")
)

//
//
type RemoteEndorserFailure struct {
	//
	Cipher        int
	Characterization string
}

func (e *RemoteEndorserFailure) Failure() string {
	return fmt.Sprintf("REDACTED", e.Cipher, e.Characterization)
}
