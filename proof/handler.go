package proof

import (
	"fmt"
	"time"

	"github.com/cosmos/gogoproto/proto"

	ringlist "github.com/valkyrieworks/utils/ringlist"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

const (
	ProofConduit = byte(0x38)

	maximumMessageVolume = 1048576 //

	//
	//
	//
	//
	multicastProofCadenceS = 10
	//
	nodeReprocessSignalCadenceMillis = 100
)

//
type Handler struct {
	p2p.RootHandler
	eventpool   *Depository
	eventBus *kinds.EventBus
}

//
func NewHandler(eventpool *Depository) *Handler {
	evR := &Handler{
		eventpool: eventpool,
	}
	evR.RootHandler = *p2p.NewRootHandler("REDACTED", evR)
	return evR
}

//
func (evR *Handler) AssignTracer(l log.Tracer) {
	evR.Tracer = l
	evR.eventpool.AssignTracer(l)
}

//
//
func (evR *Handler) FetchStreams() []*p2p.StreamDefinition {
	return []*p2p.StreamDefinition{
		{
			ID:                  ProofConduit,
			Urgency:            6,
			AcceptSignalVolume: maximumMessageVolume,
			SignalKind:         &engineproto.ProofCatalog{},
		},
	}
}

//
func (evR *Handler) AppendNode(node p2p.Node) {
	go evR.multicastProofProcedure(node)
}

//
//
func (evR *Handler) Accept(e p2p.Packet) {
	proofs, err := proofCatalogFromSchema(e.Signal)
	if err != nil {
		evR.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", err)
		evR.Router.HaltNodeForFault(e.Src, err)
		return
	}

	for _, ev := range proofs {
		err := evR.eventpool.AppendProof(ev)
		switch err.(type) {
		case *kinds.ErrCorruptProof:
			evR.Tracer.Fault(err.Error())
			//
			evR.Router.HaltNodeForFault(e.Src, err)
			return
		case nil:
		default:
			//
			evR.Tracer.Fault("REDACTED", "REDACTED", proofs, "REDACTED", err)
		}
	}
}

//
func (evR *Handler) AssignEventBus(b *kinds.EventBus) {
	evR.eventBus = b
}

//
//
//
//
//
//
func (evR *Handler) multicastProofProcedure(node p2p.Node) {
	var following *ringlist.CComponent
	for {
		//
		//
		//
		if following == nil {
			select {
			case <-evR.eventpool.ProofWaitChan(): //
				if following = evR.eventpool.ProofHead(); following == nil {
					continue
				}
			case <-node.Exit():
				return
			case <-evR.Exit():
				return
			}
		} else if !node.IsActive() || !evR.IsActive() {
			return
		}

		ev := following.Item.(kinds.Proof)
		proofs := evR.arrangeProofSignal(node, ev)
		if len(proofs) > 0 {
			evR.Tracer.Diagnose("REDACTED", "REDACTED", ev, "REDACTED", node)
			evp, err := proofCatalogToSchema(proofs)
			if err != nil {
				panic(err)
			}

			success := node.Transmit(p2p.Packet{
				StreamUID: ProofConduit,
				Signal:   evp,
			})
			if !success {
				time.Sleep(nodeReprocessSignalCadenceMillis * time.Millisecond)
				continue
			}
		}

		afterChan := time.After(time.Second * multicastProofCadenceS)
		select {
		case <-afterChan:
			//
			//
			following = nil
		case <-following.FollowingWaitChan():
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
func (evR Handler) arrangeProofSignal(
	node p2p.Node,
	ev kinds.Proof,
) (proofs []kinds.Proof) {
	//
	evtLevel := ev.Level()
	nodeStatus, ok := node.Get(kinds.NodeStatusKey).(NodeStatus)
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
		nodeLevel   = nodeStatus.FetchLevel()
		options       = evR.eventpool.Status().AgreementOptions.Proof
		eraCountLedgers = nodeLevel - evtLevel
	)

	if nodeLevel <= evtLevel { //
		return nil
	} else if eraCountLedgers > options.MaximumDurationCountLedgers { //

		//
		//
		evR.Tracer.Details("REDACTED",
			"REDACTED", nodeLevel,
			"REDACTED", evtLevel,
			"REDACTED", options.MaximumDurationCountLedgers,
			"REDACTED", evR.eventpool.Status().FinalLedgerTime,
			"REDACTED", options.MaximumDurationPeriod,
			"REDACTED", node,
		)

		return nil
	}

	//
	return []kinds.Proof{ev}
}

//
type NodeStatus interface {
	FetchLevel() int64
}

//
//
func proofCatalogToSchema(proofs []kinds.Proof) (*engineproto.ProofCatalog, error) {
	evi := make([]engineproto.Proof, len(proofs))
	for i := 0; i < len(proofs); i++ {
		ev, err := kinds.ProofToSchema(proofs[i])
		if err != nil {
			return nil, err
		}
		evi[i] = *ev
	}
	epl := engineproto.ProofCatalog{
		Proof: evi,
	}
	return &epl, nil
}

func proofCatalogFromSchema(m proto.Message) ([]kinds.Proof, error) {
	lm := m.(*engineproto.ProofCatalog)

	proofs := make([]kinds.Proof, len(lm.Proof))
	for i := 0; i < len(lm.Proof); i++ {
		ev, err := kinds.ProofFromSchema(&lm.Proof[i])
		if err != nil {
			return nil, err
		}
		proofs[i] = ev
	}

	for i, ev := range proofs {
		if err := ev.CertifySimple(); err != nil {
			return nil, fmt.Errorf("REDACTED", i, err)
		}
	}

	return proofs, nil
}
