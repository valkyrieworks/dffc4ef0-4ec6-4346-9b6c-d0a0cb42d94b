package txpool

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/linkedlist"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	schemaspace "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"golang.org/x/sync/semaphore"
)

//
//
//
type Handler struct {
	p2p.FoundationHandler
	settings  *cfg.TxpoolSettings
	txpool *CNCatalogTxpool
	ids     *txpoolIDXDstore

	//
	//
	//
	dynamicEnduringNodesGate    *semaphore.Weighted
	dynamicUnEnduringNodesGate *semaphore.Weighted

	awaitChronize   atomic.Bool
	pauseChronizeChnl chan struct{} //
}

//
func FreshHandler(settings *cfg.TxpoolSettings, txpool *CNCatalogTxpool, awaitChronize bool) *Handler {
	memoryReader := &Handler{
		settings:   settings,
		txpool:  txpool,
		ids:      freshTxpoolIDXDstore(),
		awaitChronize: atomic.Bool{},
	}
	memoryReader.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", memoryReader)
	memoryReader.dynamicEnduringNodesGate = semaphore.NewWeighted(int64(memoryReader.settings.ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes))
	memoryReader.dynamicUnEnduringNodesGate = semaphore.NewWeighted(int64(memoryReader.settings.ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes))

	if awaitChronize {
		memoryReader.awaitChronize.Store(true)
		memoryReader.pauseChronizeChnl = make(chan struct{})
	}
	return memoryReader
}

//
func (memoryReader *Handler) InitializeNode(node p2p.Node) p2p.Node {
	memoryReader.ids.AllocateForeachNode(node)
	return node
}

//
func (memoryReader *Handler) AssignTracer(l log.Tracer) {
	memoryReader.Tracer = l
	memoryReader.txpool.AssignTracer(l)
}

//
func (memoryReader *Handler) UponInitiate() error {
	if memoryReader.AwaitChronize() {
		memoryReader.Tracer.Details("REDACTED")
	}
	if !memoryReader.settings.Multicast {
		memoryReader.Tracer.Details("REDACTED")
	}
	return nil
}

//
//
func (memoryReader *Handler) ObtainConduits() []*p2p.ConduitDefinition {
	greatestTransfer := make([]byte, memoryReader.settings.MaximumTransferOctets)
	clusterSignal := schemaspace.Signal{
		Sum: &schemaspace.Artifact_Trans{
			Txs: &schemaspace.Txs{Txs: [][]byte{greatestTransfer}},
		},
	}

	return []*p2p.ConduitDefinition{
		{
			ID:                  TxpoolConduit,
			Urgency:            5,
			ObtainSignalVolume: clusterSignal.Extent(),
			SignalKind:         &schemaspace.Signal{},
		},
	}
}

//
//
func (memoryReader *Handler) AppendNode(node p2p.Node) {
	if memoryReader.settings.Multicast {
		go func() {
			//
			if !memoryReader.Router.EqualsNodeAbsolute(node.ID()) {
				//
				var nodeGate *semaphore.Weighted
				if node.EqualsEnduring() && memoryReader.settings.ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes > 0 {
					nodeGate = memoryReader.dynamicEnduringNodesGate
				} else if !node.EqualsEnduring() && memoryReader.settings.ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes > 0 {
					nodeGate = memoryReader.dynamicUnEnduringNodesGate
				}

				if nodeGate != nil {
					for node.EqualsActive() {
						//
						//
						contextDeadline, abort := context.WithTimeout(context.TODO(), 30*time.Second)
						//
						//
						err := nodeGate.Acquire(contextDeadline, 1)
						abort()

						if err != nil {
							continue
						}

						//
						defer nodeGate.Release(1)
						break
					}
				}
			}

			memoryReader.txpool.telemetry.DynamicOutgoingLinkages.Add(1)
			defer memoryReader.txpool.telemetry.DynamicOutgoingLinkages.Add(-1)
			memoryReader.multicastTransferProcedure(node)
		}()
	}
}

//
func (memoryReader *Handler) DiscardNode(node p2p.Node, _ any) {
	memoryReader.ids.Recover(node)
	//
}

//
//
func (memoryReader *Handler) Accept(e p2p.Wrapper) {
	memoryReader.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", e.Signal)
	switch msg := e.Signal.(type) {
	case *schemaspace.Txs:
		if memoryReader.AwaitChronize() {
			memoryReader.Tracer.Diagnose("REDACTED", "REDACTED", msg)
			return
		}

		schemaTrans := msg.ObtainTrans()
		if len(schemaTrans) == 0 {
			memoryReader.Tracer.Failure("REDACTED", "REDACTED", e.Src)
			return
		}
		transferDetails := TransferDetails{OriginatorUUID: memoryReader.ids.ObtainForeachNode(e.Src)}
		if e.Src != nil {
			transferDetails.OriginatorNodeid = e.Src.ID()
		}

		var err error
		for _, tx := range schemaTrans {
			ntx := kinds.Tx(tx)
			err = memoryReader.txpool.InspectTransfer(ntx, nil, transferDetails)
			if err != nil {
				switch {
				case errors.Is(err, FaultTransferInsideStash):
					memoryReader.Tracer.Diagnose("REDACTED", "REDACTED", ntx.Text())
				case errors.As(err, &FaultTxpoolEqualsComplete{}):
					//
					memoryReader.Tracer.Diagnose(err.Error())
				default:
					memoryReader.Tracer.Details("REDACTED", "REDACTED", ntx.Text(), "REDACTED", err)
				}
			}
		}
	default:
		memoryReader.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", e.Signal)
		memoryReader.Router.HaltNodeForeachFailure(e.Src, fmt.Errorf("REDACTED", e.Signal))
		return
	}

	//
}

func (memoryReader *Handler) ActivateInsideOutputTrans() {
	memoryReader.Tracer.Details("REDACTED")
	if !memoryReader.awaitChronize.CompareAndSwap(true, false) {
		return
	}

	//
	if memoryReader.settings.Multicast {
		close(memoryReader.pauseChronizeChnl)
	}
}

func (memoryReader *Handler) AwaitChronize() bool {
	return memoryReader.awaitChronize.Load()
}

//
type NodeStatus interface {
	ObtainAltitude() int64
}

//
func (memoryReader *Handler) multicastTransferProcedure(node p2p.Node) {
	//
	if memoryReader.AwaitChronize() {
		select {
		case <-memoryReader.pauseChronizeChnl:
			//
		case <-memoryReader.Exit():
			return
		}
	}

	var nodeStatus NodeStatus
	//
	//
	//
	//
	//
	for {
		if ps, ok := node.Get(kinds.NodeStatusToken).(NodeStatus); ok {
			nodeStatus = ps
			break
		}
		//
		time.Sleep(NodeOvertakeSnoozeDurationMSEC * time.Millisecond)
	}

	nodeUUID := memoryReader.ids.ObtainForeachNode(node)
	var following *linkedlist.CNComponent
	for {
		//
		if !memoryReader.EqualsActive() || !node.EqualsActive() {
			return
		}

		//
		//
		//
		if following == nil {
			select {
			case <-memoryReader.txpool.TransPauseChannel(): //
				if following = memoryReader.txpool.TransLeading(); following == nil {
					continue
				}
			case <-node.Exit():
				return
			case <-memoryReader.Exit():
				return
			}
		}

		//
		//
		//
		//
		//
		//
		memoryTransfer := following.Datum.(*txpoolTransfer)
		if nodeStatus.ObtainAltitude() < memoryTransfer.Altitude()-1 {
			time.Sleep(NodeOvertakeSnoozeDurationMSEC * time.Millisecond)
			continue
		}

		//
		//

		if !memoryTransfer.equalsOriginator(nodeUUID) {
			triumph := node.Transmit(p2p.Wrapper{
				ConduitUUID: TxpoolConduit,
				Signal:   &schemaspace.Txs{Txs: [][]byte{memoryTransfer.tx}},
			})
			if !triumph {
				time.Sleep(NodeOvertakeSnoozeDurationMSEC * time.Millisecond)
				continue
			}
		}

		select {
		case <-following.FollowingPauseChnl():
			//
			following = following.Following()
		case <-node.Exit():
			return
		case <-memoryReader.Exit():
			return
		}
	}
}
