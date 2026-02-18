package core

import (
	"fmt"

	cm "github.com/valkyrieworks/agreement"
	cometmath "github.com/valkyrieworks/utils/math"
	"github.com/valkyrieworks/p2p"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
//
//
//
func (env *Context) Ratifiers(
	_ *rpctypes.Context,
	levelPointer *int64,
	screenPointer, eachScreenPointer *int,
) (*ctypes.OutcomeRatifiers, error) {
	//
	level, err := env.fetchLevel(env.newestUnsubmittedLevel(), levelPointer)
	if err != nil {
		return nil, err
	}

	ratifiers, err := env.StatusDepot.ImportRatifiers(level)
	if err != nil {
		return nil, err
	}

	sumNumber := len(ratifiers.Ratifiers)
	eachScreen := env.certifyEachScreen(eachScreenPointer)
	screen, err := certifyScreen(screenPointer, eachScreen, sumNumber)
	if err != nil {
		return nil, err
	}

	omitNumber := certifyOmitNumber(screen, eachScreen)

	v := ratifiers.Ratifiers[omitNumber : omitNumber+cometmath.MinimumInteger(eachScreen, sumNumber-omitNumber)]

	return &ctypes.OutcomeRatifiers{
		LedgerLevel: level,
		Ratifiers:  v,
		Number:       len(v),
		Sum:       sumNumber,
	}, nil
}

//
//
//
func (env *Context) ExportAgreementStatus(*rpctypes.Context) (*ctypes.OutcomeExportAgreementStatus, error) {
	//
	nodeConditions := make([]ctypes.NodeStatusDetails, 0, env.P2PNodes.Nodes().Volume())
	var err error
	env.P2PNodes.Nodes().ForEach(func(node p2p.Node) {
		nodeStatus, ok := node.Get(kinds.NodeStatusKey).(*cm.NodeStatus)
		if !ok { //
			return
		}
		nodeStatusJSON, serializeErr := nodeStatus.SerializeJSON()
		if serializeErr != nil {
			err = fmt.Errorf("REDACTED", node.ID(), serializeErr)
			return
		}
		nodeConditions = append(nodeConditions, ctypes.NodeStatusDetails{
			//
			MemberLocation: node.SocketAddress().String(),
			//
			NodeStatus: nodeStatusJSON,
		})
	})
	if err != nil {
		return nil, err
	}

	//
	epochStatus, err := env.AgreementStatus.FetchEpochStatusJSON()
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeExportAgreementStatus{
		EpochStatus: epochStatus,
		Nodes:      nodeConditions,
	}, nil
}

//
//
//
func (env *Context) FetchAgreementStatus(*rpctypes.Context) (*ctypes.OutcomeAgreementStatus, error) {
	//
	bz, err := env.AgreementStatus.FetchEpochStatusBasicJSON()
	return &ctypes.OutcomeAgreementStatus{EpochStatus: bz}, err
}

//
//
//
func (env *Context) AgreementOptions(
	_ *rpctypes.Context,
	levelPointer *int64,
) (*ctypes.OutcomeAgreementOptions, error) {
	//
	//
	level, err := env.fetchLevel(env.newestUnsubmittedLevel(), levelPointer)
	if err != nil {
		return nil, err
	}

	agreementOptions, err := env.StatusDepot.ImportAgreementOptions(level)
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeAgreementOptions{
		LedgerLevel:     level,
		AgreementOptions: agreementOptions,
	}, nil
}
