package base

import (
	"errors"
	"fmt"

	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
func (env *Context) MulticastProof(
	_ *remoteifacetypes.Env,
	ev kinds.Proof,
) (*ktypes.OutcomeMulticastProof, error) {
	if ev == nil {
		return nil, errors.New("REDACTED")
	}

	if err := ev.CertifyFundamental(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	if err := env.ProofHub.AppendProof(ev); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return &ktypes.OutcomeMulticastProof{Digest: ev.Digest()}, nil
}
