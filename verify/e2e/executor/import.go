package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/valkyrieworks/utils/log"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/starttime/shipment"
	"github.com/valkyrieworks/kinds"
	"github.com/google/uuid"
)

const operatorDepositoryVolume = 16

//
//
func Import(ctx context.Context, verifychain *e2e.Verifychain) error {
	primaryDeadline := 1 * time.Minute
	delayDeadline := 30 * time.Second
	chanSuccess := make(chan struct{})
	ctx, revoke := context.WithCancel(ctx)
	defer revoke()

	tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", operatorDepositoryVolume))
	launched := time.Now()
	u := [16]byte(uuid.New()) //

	transferChan := make(chan kinds.Tx)
	go importCompose(ctx, transferChan, verifychain, u[:])

	for _, n := range verifychain.Instances {
		if n.TransmitNoImport {
			continue
		}

		for w := 0; w < verifychain.ImportTransferLinkages; w++ {
			go importHandle(ctx, transferChan, chanSuccess, n)
		}
	}

	//
	success := 0
	deadline := primaryDeadline
	for {
		select {
		case <-chanSuccess:
			success++
			if verifychain.ImportMaximumTrans > 0 && success >= verifychain.ImportMaximumTrans {
				tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED",
					success, float64(success)/time.Since(launched).Seconds()))
				return nil
			}
			deadline = delayDeadline
		case <-time.After(deadline):
			return fmt.Errorf("REDACTED", deadline)
		case <-ctx.Done():
			if success == 0 {
				return errors.New("REDACTED")
			}
			tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED",
				success, float64(success)/time.Since(launched).Seconds()))
			return nil
		}
	}
}

//
func importCompose(ctx context.Context, transferChan chan<- kinds.Tx, verifychain *e2e.Verifychain, id []byte) {
	t := time.NewTimer(0)
	defer t.Stop()
	for {
		select {
		case <-t.C:
		case <-ctx.Done():
			close(transferChan)
			return
		}
		t.Reset(time.Second)

		//
		//
		//
		//
		tctx, cf := context.WithTimeout(ctx, time.Second)
		instantiateTransferCluster(tctx, transferChan, verifychain, id)
		cf()
	}
}

//
//
//
func instantiateTransferCluster(ctx context.Context, transferChan chan<- kinds.Tx, verifychain *e2e.Verifychain, id []byte) {
	wg := &sync.WaitGroup{}
	generateChan := make(chan struct{})
	for i := 0; i < operatorDepositoryVolume; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range generateChan {
				tx, err := shipment.NewOctets(&shipment.Shipment{
					Id:          id,
					Volume:        uint64(verifychain.ImportTransferVolumeOctets),
					Ratio:        uint64(verifychain.ImportTransferClusterVolume),
					Linkages: uint64(verifychain.ImportTransferLinkages),
				})
				if err != nil {
					panic(fmt.Sprintf("REDACTED", err))
				}

				select {
				case transferChan <- tx:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	for i := 0; i < verifychain.ImportTransferClusterVolume; i++ {
		select {
		case generateChan <- struct{}{}:
		case <-ctx.Done():
			break
		}
	}
	close(generateChan)
	wg.Wait()
}

//
//
func importHandle(ctx context.Context, transferChan <-chan kinds.Tx, chanSuccess chan<- struct{}, n *e2e.Member) {
	var customer *rpchttp.HTTP
	var err error
	s := struct{}{}
	for tx := range transferChan {
		if customer == nil {
			customer, err = n.Customer()
			if err != nil {
				tracer.Details("REDACTED", "REDACTED", err)
				continue
			}
		}
		if _, err = customer.MulticastTransferAlign(ctx, tx); err != nil {
			continue
		}
		chanSuccess <- s
	}
}
