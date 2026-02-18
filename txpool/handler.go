package txpool

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/ringlist"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	protomemory "github.com/valkyrieworks/schema/consensuscore/txpool"
	"github.com/valkyrieworks/kinds"
	"golang.org/x/sync/semaphore"
)

//
//
//
type Handler struct {
	p2p.RootHandler
	settings  *cfg.TxpoolSettings
	txpool *CCatalogTxpool
	ids     *txpoolIDXDatastore

	//
	//
	//
	enabledDurableNodesMutex    *semaphore.Weighted
	enabledNotDurableNodesMutex *semaphore.Weighted

	waitAlign   atomic.Bool
	waitAlignChan chan struct{} //
}

//
func NewHandler(settings *cfg.TxpoolSettings, txpool *CCatalogTxpool, waitAlign bool) *Handler {
	memoryReader := &Handler{
		settings:   settings,
		txpool:  txpool,
		ids:      newTxpoolIDXDatastore(),
		waitAlign: atomic.Bool{},
	}
	memoryReader.RootHandler = *p2p.NewRootHandler("REDACTED", memoryReader)
	memoryReader.enabledDurableNodesMutex = semaphore.NewWeighted(int64(memoryReader.settings.ExploratoryMaximumGossipLinkagesToDurableNodes))
	memoryReader.enabledNotDurableNodesMutex = semaphore.NewWeighted(int64(memoryReader.settings.ExploratoryMaximumGossipLinkagesToNotDurableNodes))

	if waitAlign {
		memoryReader.waitAlign.Store(true)
		memoryReader.waitAlignChan = make(chan struct{})
	}
	return memoryReader
}

//
func (memoryReader *Handler) InitNode(node p2p.Node) p2p.Node {
	memoryReader.ids.AllocateForNode(node)
	return node
}

//
func (memoryReader *Handler) AssignTracer(l log.Tracer) {
	memoryReader.Tracer = l
	memoryReader.txpool.AssignTracer(l)
}

//
func (memoryReader *Handler) OnBegin() error {
	if memoryReader.WaitAlign() {
		memoryReader.Tracer.Details("REDACTED")
	}
	if !memoryReader.settings.Multicast {
		memoryReader.Tracer.Details("REDACTED")
	}
	return nil
}

//
//
func (memoryReader *Handler) FetchStreams() []*p2p.StreamDefinition {
	greatestTransfer := make([]byte, memoryReader.settings.MaximumTransferOctets)
	groupMessage := protomemory.Signal{
		Sum: &protomemory.Signal_Trans{
			Txs: &protomemory.Txs{Txs: [][]byte{greatestTransfer}},
		},
	}

	return []*p2p.StreamDefinition{
		{
			ID:                  TxpoolConduit,
			Urgency:            5,
			AcceptSignalVolume: groupMessage.Volume(),
			SignalKind:         &protomemory.Signal{},
		},
	}
}

//
//
func (memoryReader *Handler) AppendNode(node p2p.Node) {
	if memoryReader.settings.Multicast {
		go func() {
			//
			if !memoryReader.Router.IsNodeAbsolute(node.ID()) {
				//
				var nodeMutex *semaphore.Weighted
				if node.IsDurable() && memoryReader.settings.ExploratoryMaximumGossipLinkagesToDurableNodes > 0 {
					nodeMutex = memoryReader.enabledDurableNodesMutex
				} else if !node.IsDurable() && memoryReader.settings.ExploratoryMaximumGossipLinkagesToNotDurableNodes > 0 {
					nodeMutex = memoryReader.enabledNotDurableNodesMutex
				}

				if nodeMutex != nil {
					for node.IsActive() {
						//
						//
						ctxDeadline, revoke := context.WithTimeout(context.TODO(), 30*time.Second)
						//
						//
						err := nodeMutex.Acquire(ctxDeadline, 1)
						revoke()

						if err != nil {
							continue
						}

						//
						defer nodeMutex.Release(1)
						break
					}
				}
			}

			memoryReader.txpool.stats.EnabledOutgoingLinkages.Add(1)
			defer memoryReader.txpool.stats.EnabledOutgoingLinkages.Add(-1)
			memoryReader.multicastTransferProcedure(node)
		}()
	}
}

//
func (memoryReader *Handler) DeleteNode(node p2p.Node, _ any) {
	memoryReader.ids.Recover(node)
	//
}

//
//
func (memoryReader *Handler) Accept(e p2p.Packet) {
	memoryReader.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", e.Signal)
	switch msg := e.Signal.(type) {
	case *protomemory.Txs:
		if memoryReader.WaitAlign() {
			memoryReader.Tracer.Diagnose("REDACTED", "REDACTED", msg)
			return
		}

		schemaTrans := msg.FetchTrans()
		if len(schemaTrans) == 0 {
			memoryReader.Tracer.Fault("REDACTED", "REDACTED", e.Src)
			return
		}
		transferDetails := TransferDetails{EmitterUID: memoryReader.ids.FetchForNode(e.Src)}
		if e.Src != nil {
			transferDetails.EmitterP2pid = e.Src.ID()
		}

		var err error
		for _, tx := range schemaTrans {
			ntx := kinds.Tx(tx)
			err = memoryReader.txpool.InspectTransfer(ntx, nil, transferDetails)
			if err != nil {
				switch {
				case errors.Is(err, ErrTransferInRepository):
					memoryReader.Tracer.Diagnose("REDACTED", "REDACTED", ntx.String())
				case errors.As(err, &ErrTxpoolIsComplete{}):
					//
					memoryReader.Tracer.Diagnose(err.Error())
				default:
					memoryReader.Tracer.Details("REDACTED", "REDACTED", ntx.String(), "REDACTED", err)
				}
			}
		}
	default:
		memoryReader.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", e.Signal)
		memoryReader.Router.HaltNodeForFault(e.Src, fmt.Errorf("REDACTED", e.Signal))
		return
	}

	//
}

func (memoryReader *Handler) ActivateInOutTrans() {
	memoryReader.Tracer.Details("REDACTED")
	if !memoryReader.waitAlign.CompareAndSwap(true, false) {
		return
	}

	//
	if memoryReader.settings.Multicast {
		close(memoryReader.waitAlignChan)
	}
}

func (memoryReader *Handler) WaitAlign() bool {
	return memoryReader.waitAlign.Load()
}

//
type NodeStatus interface {
	FetchLevel() int64
}

//
func (memoryReader *Handler) multicastTransferProcedure(node p2p.Node) {
	//
	if memoryReader.WaitAlign() {
		select {
		case <-memoryReader.waitAlignChan:
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
		if ps, ok := node.Get(kinds.NodeStatusKey).(NodeStatus); ok {
			nodeStatus = ps
			break
		}
		//
		time.Sleep(NodeOvertakePauseCadenceMillis * time.Millisecond)
	}

	nodeUID := memoryReader.ids.FetchForNode(node)
	var following *ringlist.CComponent
	for {
		//
		if !memoryReader.IsActive() || !node.IsActive() {
			return
		}

		//
		//
		//
		if following == nil {
			select {
			case <-memoryReader.txpool.TransWaitChan(): //
				if following = memoryReader.txpool.TransHead(); following == nil {
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
		memoryTransfer := following.Item.(*txpoolTransfer)
		if nodeStatus.FetchLevel() < memoryTransfer.Level()-1 {
			time.Sleep(NodeOvertakePauseCadenceMillis * time.Millisecond)
			continue
		}

		//
		//

		if !memoryTransfer.isEmitter(nodeUID) {
			success := node.Transmit(p2p.Packet{
				StreamUID: TxpoolConduit,
				Signal:   &protomemory.Txs{Txs: [][]byte{memoryTransfer.tx}},
			})
			if !success {
				time.Sleep(NodeOvertakePauseCadenceMillis * time.Millisecond)
				continue
			}
		}

		select {
		case <-following.FollowingWaitChan():
			//
			following = following.Following()
		case <-node.Exit():
			return
		case <-memoryReader.Exit():
			return
		}
	}
}
