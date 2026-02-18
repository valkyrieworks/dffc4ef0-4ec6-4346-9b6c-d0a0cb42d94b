package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"reflect"

	"github.com/valkyrieworks/iface/kinds"
	cometnet "github.com/valkyrieworks/utils/net"
)

func main() {
	link, err := cometnet.Link("REDACTED")
	if err != nil {
		log.Fatal(err.Error())
	}

	//
	tally := 0
	for i := 0; ; i++ {
		req := kinds.ToQueryReverberate("REDACTED")
		_, err := createQuery(link, req)
		if err != nil {
			log.Fatal(err.Error())
		}
		tally++
		if tally%1000 == 0 {
			fmt.Println(tally)
		}
	}
}

func createQuery(link io.ReadWriter, req *kinds.Query) (*kinds.Reply, error) {
	bufferRecorder := bufio.NewWriter(link)

	//
	err := kinds.RecordSignal(req, bufferRecorder)
	if err != nil {
		return nil, err
	}
	err = kinds.RecordSignal(kinds.ToQueryPurge(), bufferRecorder)
	if err != nil {
		return nil, err
	}
	err = bufferRecorder.Flush()
	if err != nil {
		return nil, err
	}

	//
	res := &kinds.Reply{}
	err = kinds.ScanSignal(link, res)
	if err != nil {
		return nil, err
	}
	outputPurge := &kinds.Reply{}
	err = kinds.ScanSignal(link, outputPurge)
	if err != nil {
		return nil, err
	}
	if _, ok := outputPurge.Item.(*kinds.Reply_Purge); !ok {
		return nil, fmt.Errorf("REDACTED", reflect.TypeOf(outputPurge))
	}

	return res, nil
}
