package core

import (
	"errors"
	"fmt"
	"strings"

	"github.com/valkyrieworks/p2p"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//
//
func (env *Context) NetDetails(*rpctypes.Context) (*ctypes.OutcomeNetDetails, error) {
	nodes := make([]ctypes.Node, 0, env.P2PNodes.Nodes().Volume())
	var err error
	env.P2PNodes.Nodes().ForEach(func(node p2p.Node) {
		memberDetails, ok := node.MemberDetails().(p2p.StandardMemberDetails)
		if !ok {
			err = fmt.Errorf("REDACTED", node.ID(), node.MemberDetails())
			return
		}
		nodes = append(nodes, ctypes.Node{
			MemberDetails:         memberDetails,
			IsOutgoing:       node.IsOutgoing(),
			LinkageState: node.Status(),
			DistantIP:         node.DistantIP().String(),
		})
	})
	if err != nil {
		return nil, err
	}
	//
	//
	//
	return &ctypes.OutcomeNetDetails{
		Observing: env.P2PCarrier.IsObserving(),
		Observers: env.P2PCarrier.Observers(),
		NNodes:    len(nodes),
		Nodes:     nodes,
	}, nil
}

//
func (env *Context) RiskyCallOrigins(_ *rpctypes.Context, origins []string) (*ctypes.OutcomeCallOrigins, error) {
	if len(origins) == 0 {
		return &ctypes.OutcomeCallOrigins{}, errors.New("REDACTED")
	}
	env.Tracer.Details("REDACTED", "REDACTED", origins)
	if err := env.P2PNodes.CallNodesAsync(origins); err != nil {
		return &ctypes.OutcomeCallOrigins{}, err
	}
	return &ctypes.OutcomeCallOrigins{Log: "REDACTED"}, nil
}

//
//
func (env *Context) RiskyCallNodes(
	_ *rpctypes.Context,
	nodes []string,
	durable, absolute, internal bool,
) (*ctypes.OutcomeCallNodes, error) {
	if len(nodes) == 0 {
		return &ctypes.OutcomeCallNodes{}, errors.New("REDACTED")
	}

	ids, err := fetchIDXDatastore(nodes)
	if err != nil {
		return &ctypes.OutcomeCallNodes{}, err
	}

	env.Tracer.Details("REDACTED", "REDACTED", nodes, "REDACTED",
		durable, "REDACTED", absolute, "REDACTED", internal)

	if durable {
		if err := env.P2PNodes.AppendDurableNodes(nodes); err != nil {
			return &ctypes.OutcomeCallNodes{}, err
		}
	}

	if internal {
		if err := env.P2PNodes.AppendInternalNodeIDXDatastore(ids); err != nil {
			return &ctypes.OutcomeCallNodes{}, err
		}
	}

	if absolute {
		if err := env.P2PNodes.AppendAbsoluteNodeIDXDatastore(ids); err != nil {
			return &ctypes.OutcomeCallNodes{}, err
		}
	}

	if err := env.P2PNodes.CallNodesAsync(nodes); err != nil {
		return &ctypes.OutcomeCallNodes{}, err
	}

	return &ctypes.OutcomeCallNodes{Log: "REDACTED"}, nil
}

//
//
func (env *Context) Origin(*rpctypes.Context) (*ctypes.OutcomeOrigin, error) {
	if len(env.generateSegments) > 1 {
		return nil, errors.New("REDACTED")
	}

	return &ctypes.OutcomeOrigin{Origin: env.GeneratePaper}, nil
}

func (env *Context) OriginSegmented(_ *rpctypes.Context, segment uint) (*ctypes.OutcomeOriginSegment, error) {
	if env.generateSegments == nil {
		return nil, fmt.Errorf("REDACTED")
	}

	if len(env.generateSegments) == 0 {
		return nil, fmt.Errorf("REDACTED")
	}

	id := int(segment)

	if id > len(env.generateSegments)-1 {
		return nil, fmt.Errorf("REDACTED", len(env.generateSegments)-1, id)
	}

	return &ctypes.OutcomeOriginSegment{
		SumSegments: len(env.generateSegments),
		SegmentAmount: id,
		Data:        env.generateSegments[id],
	}, nil
}

func fetchIDXDatastore(nodes []string) ([]string, error) {
	ids := make([]string, 0, len(nodes))

	for _, node := range nodes {

		spl := strings.Split(node, "REDACTED")
		if len(spl) != 2 {
			return nil, p2p.ErrNetLocationNoUID{Address: node}
		}
		ids = append(ids, spl[0])

	}
	return ids, nil
}
