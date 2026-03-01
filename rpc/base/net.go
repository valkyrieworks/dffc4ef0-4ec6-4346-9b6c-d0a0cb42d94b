package base

import (
	"errors"
	"fmt"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//
//
func (env *Context) NetworkDetails(*remoteifacetypes.Env) (*ktypes.OutcomeNetworkDetails, error) {
	nodes := make([]ktypes.Node, 0, env.Peer2peerNodes.Nodes().Extent())
	var err error
	env.Peer2peerNodes.Nodes().ForeachEvery(func(node p2p.Node) {
		peerDetails, ok := node.PeerDetails().(p2p.FallbackPeerDetails)
		if !ok {
			err = fmt.Errorf("REDACTED", node.ID(), node.PeerDetails())
			return
		}
		nodes = append(nodes, ktypes.Node{
			PeerDetails:         peerDetails,
			EqualsOutgoing:       node.EqualsOutgoing(),
			LinkageCondition: node.Condition(),
			DistantINET:         node.DistantINET().String(),
		})
	})
	if err != nil {
		return nil, err
	}
	//
	//
	//
	return &ktypes.OutcomeNetworkDetails{
		Observing: env.Peer2peerCarrier.EqualsObserving(),
		Observers: env.Peer2peerCarrier.Observers(),
		NTHNodes:    len(nodes),
		Nodes:     nodes,
	}, nil
}

//
func (env *Context) InsecureCallOrigins(_ *remoteifacetypes.Env, origins []string) (*ktypes.OutcomeCallOrigins, error) {
	if len(origins) == 0 {
		return &ktypes.OutcomeCallOrigins{}, errors.New("REDACTED")
	}
	env.Tracer.Details("REDACTED", "REDACTED", origins)
	if err := env.Peer2peerNodes.CallNodesAsyncronous(origins); err != nil {
		return &ktypes.OutcomeCallOrigins{}, err
	}
	return &ktypes.OutcomeCallOrigins{Log: "REDACTED"}, nil
}

//
//
func (env *Context) InsecureCallNodes(
	_ *remoteifacetypes.Env,
	nodes []string,
	enduring, absolute, secluded bool,
) (*ktypes.OutcomeCallNodes, error) {
	if len(nodes) == 0 {
		return &ktypes.OutcomeCallNodes{}, errors.New("REDACTED")
	}

	ids, err := obtainIDXDstore(nodes)
	if err != nil {
		return &ktypes.OutcomeCallNodes{}, err
	}

	env.Tracer.Details("REDACTED", "REDACTED", nodes, "REDACTED",
		enduring, "REDACTED", absolute, "REDACTED", secluded)

	if enduring {
		if err := env.Peer2peerNodes.AppendEnduringNodes(nodes); err != nil {
			return &ktypes.OutcomeCallNodes{}, err
		}
	}

	if secluded {
		if err := env.Peer2peerNodes.AppendSecludedNodeIDXDstore(ids); err != nil {
			return &ktypes.OutcomeCallNodes{}, err
		}
	}

	if absolute {
		if err := env.Peer2peerNodes.AppendAbsoluteNodeIDXDstore(ids); err != nil {
			return &ktypes.OutcomeCallNodes{}, err
		}
	}

	if err := env.Peer2peerNodes.CallNodesAsyncronous(nodes); err != nil {
		return &ktypes.OutcomeCallNodes{}, err
	}

	return &ktypes.OutcomeCallNodes{Log: "REDACTED"}, nil
}

//
//
func (env *Context) Inauguration(*remoteifacetypes.Env) (*ktypes.OutcomeInauguration, error) {
	if len(env.produceSegments) > 1 {
		return nil, errors.New("REDACTED")
	}

	return &ktypes.OutcomeInauguration{Inauguration: env.ProducePaper}, nil
}

func (env *Context) InaugurationSegmented(_ *remoteifacetypes.Env, segment uint) (*ktypes.OutcomeInaugurationSegment, error) {
	if env.produceSegments == nil {
		return nil, fmt.Errorf("REDACTED")
	}

	if len(env.produceSegments) == 0 {
		return nil, fmt.Errorf("REDACTED")
	}

	id := int(segment)

	if id > len(env.produceSegments)-1 {
		return nil, fmt.Errorf("REDACTED", len(env.produceSegments)-1, id)
	}

	return &ktypes.OutcomeInaugurationSegment{
		SumSegments: len(env.produceSegments),
		SegmentNumeral: id,
		Data:        env.produceSegments[id],
	}, nil
}

func obtainIDXDstore(nodes []string) ([]string, error) {
	ids := make([]string, 0, len(nodes))

	for _, node := range nodes {

		spl := strings.Split(node, "REDACTED")
		if len(spl) != 2 {
			return nil, p2p.FaultNetworkLocatorNegativeUUID{Location: node}
		}
		ids = append(ids, spl[0])

	}
	return ids, nil
}
