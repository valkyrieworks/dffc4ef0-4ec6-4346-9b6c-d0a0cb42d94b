package proof

import (
	"fmt"
	"time"

	"github.com/cosmos/gogoproto/proto"

	linkedlist "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/linkedlist"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	ProofConduit = byte(0x38)

	maximumSignalExtent = 1048576 //

	//
	//
	//
	//
	multicastProofDurationSTR = 10
	//
	nodeReissueArtifactDurationMSEC = 100
)

//
type Handler struct {
	p2p.FoundationHandler
	incidentpool   *Hub
	incidentPipeline *kinds.IncidentChannel
}

//
func FreshHandler(incidentpool *Hub) *Handler {
	evR := &Handler{
		incidentpool: incidentpool,
	}
	evR.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", evR)
	return evR
}

//
func (evR *Handler) AssignTracer(l log.Tracer) {
	evR.Tracer = l
	evR.incidentpool.AssignTracer(l)
}

//
//
func (evR *Handler) ObtainConduits() []*p2p.ConduitDefinition {
	return []*p2p.ConduitDefinition{
		{
			ID:                  ProofConduit,
			Urgency:            6,
			ObtainSignalVolume: maximumSignalExtent,
			SignalKind:         &commitchema.ProofCatalog{},
		},
	}
}

//
func (evR *Handler) AppendNode(node p2p.Node) {
	go evR.multicastProofProcedure(node)
}

//
//
func (evR *Handler) Accept(e p2p.Wrapper) {
	proofs, err := proofCatalogOriginatingSchema(e.Signal)
	if err != nil {
		evR.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", err)
		evR.Router.HaltNodeForeachFailure(e.Src, err)
		return
	}

	for _, ev := range proofs {
		err := evR.incidentpool.AppendProof(ev)
		switch err.(type) {
		case *kinds.FaultUnfitProof:
			evR.Tracer.Failure(err.Error())
			//
			evR.Router.HaltNodeForeachFailure(e.Src, err)
			return
		case nil:
		default:
			//
			evR.Tracer.Failure("REDACTED", "REDACTED", proofs, "REDACTED", err)
		}
	}
}

//
func (evR *Handler) AssignIncidentChannel(b *kinds.IncidentChannel) {
	evR.incidentPipeline = b
}

//
//
//
//
//
//
func (evR *Handler) multicastProofProcedure(node p2p.Node) {
	var following *linkedlist.CNComponent
	for {
		//
		//
		//
		if following == nil {
			select {
			case <-evR.incidentpool.ProofPauseChnl(): //
				if following = evR.incidentpool.ProofLeading(); following == nil {
					continue
				}
			case <-node.Exit():
				return
			case <-evR.Exit():
				return
			}
		} else if !node.EqualsActive() || !evR.EqualsActive() {
			return
		}

		ev := following.Datum.(kinds.Proof)
		proofs := evR.arrangeProofArtifact(node, ev)
		if len(proofs) > 0 {
			evR.Tracer.Diagnose("REDACTED", "REDACTED", ev, "REDACTED", node)
			evp, err := proofCatalogTowardSchema(proofs)
			if err != nil {
				panic(err)
			}

			triumph := node.Transmit(p2p.Wrapper{
				ConduitUUID: ProofConduit,
				Signal:   evp,
			})
			if !triumph {
				time.Sleep(nodeReissueArtifactDurationMSEC * time.Millisecond)
				continue
			}
		}

		subsequentChnl := time.After(time.Second * multicastProofDurationSTR)
		select {
		case <-subsequentChnl:
			//
			//
			following = nil
		case <-following.FollowingPauseChnl():
			//
			following = following.Following()
		case <-node.Exit():
			return
		case <-evR.Exit():
			return
		}
	}
}

//
//
func (evR Handler) arrangeProofArtifact(
	node p2p.Node,
	ev kinds.Proof,
) (proofs []kinds.Proof) {
	//
	occurenceAltitude := ev.Altitude()
	nodeStatus, ok := node.Get(kinds.NodeStatusToken).(NodeStatus)
	if !ok {
		//
		//
		//
		//
		//
		return nil
	}

	//
	//
	var (
		nodeAltitude   = nodeStatus.ObtainAltitude()
		parameters       = evR.incidentpool.Status().AgreementSettings.Proof
		lifespanCountLedgers = nodeAltitude - occurenceAltitude
	)

	if nodeAltitude <= occurenceAltitude { //
		return nil
	} else if lifespanCountLedgers > parameters.MaximumLifespanCountLedgers { //

		//
		//
		evR.Tracer.Details("REDACTED",
			"REDACTED", nodeAltitude,
			"REDACTED", occurenceAltitude,
			"REDACTED", parameters.MaximumLifespanCountLedgers,
			"REDACTED", evR.incidentpool.Status().FinalLedgerMoment,
			"REDACTED", parameters.MaximumLifespanInterval,
			"REDACTED", node,
		)

		return nil
	}

	//
	return []kinds.Proof{ev}
}

//
type NodeStatus interface {
	ObtainAltitude() int64
}

//
//
func proofCatalogTowardSchema(proofs []kinds.Proof) (*commitchema.ProofCatalog, error) {
	evi := make([]commitchema.Proof, len(proofs))
	for i := 0; i < len(proofs); i++ {
		ev, err := kinds.ProofTowardSchema(proofs[i])
		if err != nil {
			return nil, err
		}
		evi[i] = *ev
	}
	epl := commitchema.ProofCatalog{
		Proof: evi,
	}
	return &epl, nil
}

func proofCatalogOriginatingSchema(m proto.Message) ([]kinds.Proof, error) {
	lm := m.(*commitchema.ProofCatalog)

	proofs := make([]kinds.Proof, len(lm.Proof))
	for i := 0; i < len(lm.Proof); i++ {
		ev, err := kinds.ProofOriginatingSchema(&lm.Proof[i])
		if err != nil {
			return nil, err
		}
		proofs[i] = ev
	}

	for i, ev := range proofs {
		if err := ev.CertifyFundamental(); err != nil {
			return nil, fmt.Errorf("REDACTED", i, err)
		}
	}

	return proofs, nil
}
