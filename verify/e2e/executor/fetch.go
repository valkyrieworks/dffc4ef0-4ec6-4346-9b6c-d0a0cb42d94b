package primary

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/fetchmoment/content"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/google/uuid"
)

const operatorHubExtent = 16

//
//
func Fetch(ctx context.Context, simnet *e2e.Simnet) error {
	primaryDeadline := 1 * time.Minute
	freezeDeadline := 30 * time.Second
	chnlTriumph := make(chan struct{})
	ctx, abort := context.WithCancel(ctx)
	defer abort()

	tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", operatorHubExtent))
	initiated := time.Now()
	u := [16]byte(uuid.New()) //

	transferChnl := make(chan kinds.Tx)
	go fetchCompose(ctx, transferChnl, simnet, u[:])

	for _, n := range simnet.Peers {
		if n.TransmitNegativeFetch {
			continue
		}

		for w := 0; w < simnet.FetchTransferLinkages; w++ {
			go fetchHandle(ctx, transferChnl, chnlTriumph, n)
		}
	}

	//
	triumph := 0
	deadline := primaryDeadline
	for {
		select {
		case <-chnlTriumph:
			triumph++
			if simnet.FetchMaximumTrans > 0 && triumph >= simnet.FetchMaximumTrans {
				tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED",
					triumph, float64(triumph)/time.Since(initiated).Seconds()))
				return nil
			}
			deadline = freezeDeadline
		case <-time.After(deadline):
			return fmt.Errorf("REDACTED", deadline)
		case <-ctx.Done():
			if triumph == 0 {
				return errors.New("REDACTED")
			}
			tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED",
				triumph, float64(triumph)/time.Since(initiated).Seconds()))
			return nil
		}
	}
}

//
func fetchCompose(ctx context.Context, transferChnl chan<- kinds.Tx, simnet *e2e.Simnet, id []byte) {
	t := time.NewTimer(0)
	defer t.Stop()
	for {
		select {
		case <-t.C:
		case <-ctx.Done():
			close(transferChnl)
			return
		}
		t.Reset(time.Second)

		//
		//
		//
		//
		tempctx, cf := context.WithTimeout(ctx, time.Second)
		generateTransferCluster(tempctx, transferChnl, simnet, id)
		cf()
	}
}

//
//
//
func generateTransferCluster(ctx context.Context, transferChnl chan<- kinds.Tx, simnet *e2e.Simnet, id []byte) {
	wg := &sync.WaitGroup{}
	produceChnl := make(chan struct{})
	for i := 0; i < operatorHubExtent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range produceChnl {
				tx, err := content.FreshOctets(&content.Content{
					Id:          id,
					Extent:        uint64(simnet.FetchTransferExtentOctets),
					Frequency:        uint64(simnet.FetchTransferClusterExtent),
					Linkages: uint64(simnet.FetchTransferLinkages),
				})
				if err != nil {
					panic(fmt.Sprintf("REDACTED", err))
				}

				select {
				case transferChnl <- tx:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	for i := 0; i < simnet.FetchTransferClusterExtent; i++ {
		select {
		case produceChnl <- struct{}{}:
		case <-ctx.Done():
			break
		}
	}
	close(produceChnl)
	wg.Wait()
}

//
//
func fetchHandle(ctx context.Context, transferChnl <-chan kinds.Tx, chnlTriumph chan<- struct{}, n *e2e.Peer) {
	var customer *rpchttpsvc.Httpsvc
	var err error
	s := struct{}{}
	for tx := range transferChnl {
		if customer == nil {
			customer, err = n.Customer()
			if err != nil {
				tracer.Details("REDACTED", "REDACTED", err)
				continue
			}
		}
		if _, err = customer.MulticastTransferChronize(ctx, tx); err != nil {
			continue
		}
		chnlTriumph <- s
	}
}
