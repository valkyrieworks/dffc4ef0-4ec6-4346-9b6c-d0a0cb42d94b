package base

import (
	"fmt"

	cm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
//
//
//
func (env *Context) Assessors(
	_ *remoteifacetypes.Env,
	altitudeReference *int64,
	screenReference, everyScreenReference *int,
) (*ktypes.OutcomeAssessors, error) {
	//
	altitude, err := env.obtainAltitude(env.newestPendingAltitude(), altitudeReference)
	if err != nil {
		return nil, err
	}

	assessors, err := env.StatusDepot.FetchAssessors(altitude)
	if err != nil {
		return nil, err
	}

	sumTally := len(assessors.Assessors)
	everyScreen := env.certifyEveryScreen(everyScreenReference)
	screen, err := certifyScreen(screenReference, everyScreen, sumTally)
	if err != nil {
		return nil, err
	}

	omitTally := certifyOmitTally(screen, everyScreen)

	v := assessors.Assessors[omitTally : omitTally+strongarithmetic.MinimumInteger(everyScreen, sumTally-omitTally)]

	return &ktypes.OutcomeAssessors{
		LedgerAltitude: altitude,
		Assessors:  v,
		Tally:       len(v),
		Sum:       sumTally,
	}, nil
}

//
//
//
func (env *Context) ExportAgreementStatus(*remoteifacetypes.Env) (*ktypes.OutcomeExportAgreementStatus, error) {
	//
	nodeStatuses := make([]ktypes.NodeStatusDetails, 0, env.Peer2peerNodes.Nodes().Extent())
	var err error
	env.Peer2peerNodes.Nodes().ForeachEvery(func(node p2p.Node) {
		nodeStatus, ok := node.Get(kinds.NodeStatusToken).(*cm.NodeStatus)
		if !ok { //
			return
		}
		nodeStatusJSN, serializeFault := nodeStatus.SerializeJSN()
		if serializeFault != nil {
			err = fmt.Errorf("REDACTED", node.ID(), serializeFault)
			return
		}
		nodeStatuses = append(nodeStatuses, ktypes.NodeStatusDetails{
			//
			PeerLocator: node.PortLocation().Text(),
			//
			NodeStatus: nodeStatusJSN,
		})
	})
	if err != nil {
		return nil, err
	}

	//
	iterationStatus, err := env.AgreementStatus.ObtainIterationStatusJSN()
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeExportAgreementStatus{
		IterationStatus: iterationStatus,
		Nodes:      nodeStatuses,
	}, nil
}

//
//
//
func (env *Context) ObtainAgreementStatus(*remoteifacetypes.Env) (*ktypes.OutcomeAgreementStatus, error) {
	//
	bz, err := env.AgreementStatus.ObtainIterationStatusPlainJSN()
	return &ktypes.OutcomeAgreementStatus{IterationStatus: bz}, err
}

//
//
//
func (env *Context) AgreementSettings(
	_ *remoteifacetypes.Env,
	altitudeReference *int64,
) (*ktypes.OutcomeAgreementParameters, error) {
	//
	//
	altitude, err := env.obtainAltitude(env.newestPendingAltitude(), altitudeReference)
	if err != nil {
		return nil, err
	}

	agreementParameters, err := env.StatusDepot.FetchAgreementParameters(altitude)
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeAgreementParameters{
		LedgerAltitude:     altitude,
		AgreementSettings: agreementParameters,
	}, nil
}
