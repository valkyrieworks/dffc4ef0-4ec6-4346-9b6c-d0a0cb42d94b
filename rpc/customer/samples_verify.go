package agent_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	"github.com/valkyrieworks/rpc/jsonrpc/kinds"
	rpctest "github.com/valkyrieworks/rpc/verify"
)

func Samplehttp_basic() {
	//
	app := objectdepot.NewInRamSoftware()
	member := rpctest.BeginConsensuscore(app, rpctest.InhibitStdout, rpctest.RebuildSettings)
	defer rpctest.HaltConsensuscore(member)

	//
	rpcAddress := rpctest.FetchSettings().RPC.AcceptLocation
	c, err := rpchttp.New(rpcAddress, "REDACTED")
	if err != nil {
		log.Fatal(err) //
	}

	//
	k := []byte("REDACTED")
	v := []byte("REDACTED")
	tx := bytes.Join([][]byte{
		k,
		[]byte("REDACTED"),
		v,
	}, nil)

	//
	//
	bout, err := c.MulticastTransferEndorse(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	if bout.InspectTransfer.IsErr() || bout.TransOutcome.IsErr() {
		log.Fatal("REDACTED")
	}

	//
	inquiryout, err := c.IfaceInquire(context.Background(), "REDACTED", k)
	if err != nil {
		log.Fatal(err)
	}
	if inquiryout.Reply.IsErr() {
		log.Fatal("REDACTED")
	}
	if !bytes.Equal(inquiryout.Reply.Key, k) {
		log.Fatal("REDACTED")
	}
	if !bytes.Equal(inquiryout.Reply.Item, v) {
		log.Fatal("REDACTED")
	}

	fmt.Println("REDACTED", string(tx))
	fmt.Println("REDACTED", string(inquiryout.Reply.Key))
	fmt.Println("REDACTED", string(inquiryout.Reply.Item))

	//
	//
	//
	//
}

func Samplehttp_segmenting() {
	//
	app := objectdepot.NewInRamSoftware()
	member := rpctest.BeginConsensuscore(app, rpctest.InhibitStdout, rpctest.RebuildSettings)

	//
	rpcAddress := rpctest.FetchSettings().RPC.AcceptLocation
	c, err := rpchttp.New(rpcAddress, "REDACTED")
	if err != nil {
		log.Fatal(err)
	}

	defer rpctest.HaltConsensuscore(member)

	//
	k1 := []byte("REDACTED")
	v1 := []byte("REDACTED")
	tx1 := bytes.Join([][]byte{k1, []byte("REDACTED"), v1}, nil)

	k2 := []byte("REDACTED")
	v2 := []byte("REDACTED")
	tx2 := bytes.Join([][]byte{k2, []byte("REDACTED"), v2}, nil)

	txs := [][]byte{tx1, tx2}

	//
	group := c.NewGroup()

	//
	for _, tx := range txs {
		//
		//
		if _, err := group.MulticastTransferEndorse(context.Background(), tx); err != nil {
			log.Fatal(err) //
		}
	}

	//
	if _, err := group.Transmit(context.Background()); err != nil {
		log.Fatal(err)
	}

	//
	keys := [][]byte{k1, k2}
	for _, key := range keys {
		if _, err := group.IfaceInquire(context.Background(), "REDACTED", key); err != nil {
			log.Fatal(err)
		}
	}

	//
	outcomes, err := group.Transmit(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//
	//
	for _, outcome := range outcomes {
		qr, ok := outcome.(*ctypes.OutcomeIfaceInquire)
		if !ok {
			log.Fatal("REDACTED")
		}
		fmt.Println(string(qr.Reply.Key), "REDACTED", string(qr.Reply.Item))
	}

	//
	//
	//
}

//
func Samplehttp_maxsegmentsize() {
	//
	app := objectdepot.NewInRamSoftware()
	member := rpctest.BeginConsensuscore(app, rpctest.RebuildSettings, rpctest.InhibitStdout, rpctest.MaximumRequestGroupVolume)

	//
	member.Settings().RPC.MaximumQueryClusterVolume = 2

	//
	rpcAddress := rpctest.FetchSettings().RPC.AcceptLocation
	c, err := rpchttp.New(rpcAddress, "REDACTED")
	if err != nil {
		log.Fatal(err)
	}

	defer rpctest.HaltConsensuscore(member)

	//
	group := c.NewGroup()

	for i := 1; i <= 5; i++ {
		if _, err := group.Vitality(context.Background()); err != nil {
			log.Fatal(err) //
		}
	}

	//
	outcomes, err := group.Transmit(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//
	//
	for _, outcome := range outcomes {
		rpcFault, ok := outcome.(*kinds.RPCFault)
		if !ok {
			log.Fatal("REDACTED")
		}
		if !strings.Contains(rpcFault.Data, "REDACTED") {
			fmt.Println("REDACTED")
		} else {
			//
			fmt.Println("REDACTED")
		}
	}

	//
	//
}
