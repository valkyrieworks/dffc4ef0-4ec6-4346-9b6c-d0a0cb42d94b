package core

import (
	"errors"
	"fmt"

	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
func (env *Context) MulticastProof(
	_ *rpctypes.Context,
	ev kinds.Proof,
) (*ctypes.OutcomeMulticastProof, error) {
	if ev == nil {
		return nil, errors.New("REDACTED")
	}

	if err := ev.CertifySimple(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	if err := env.ProofDepository.AppendProof(ev); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return &ctypes.OutcomeMulticastProof{Digest: ev.Digest()}, nil
}
