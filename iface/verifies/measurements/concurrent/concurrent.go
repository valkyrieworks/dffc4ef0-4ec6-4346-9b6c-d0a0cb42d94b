package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/valkyrieworks/iface/kinds"
	cometnet "github.com/valkyrieworks/utils/net"
)

func main() {
	link, err := cometnet.Link("REDACTED")
	if err != nil {
		log.Fatal(err.Error())
	}

	//
	go func() {
		tally := 0
		for {
			res := &kinds.Reply{}
			err := kinds.ScanSignal(link, res)
			if err != nil {
				log.Fatal(err.Error())
			}
			tally++
			if tally%1000 == 0 {
				fmt.Println("REDACTED", tally)
			}
		}
	}()

	//
	tally := 0
	for i := 0; ; i++ {
		bufferRecorder := bufio.NewWriter(link)
		req := kinds.ToQueryReverberate("REDACTED")

		err := kinds.RecordSignal(req, bufferRecorder)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = bufferRecorder.Flush()
		if err != nil {
			log.Fatal(err.Error())
		}

		tally++
		if tally%1000 == 0 {
			fmt.Println("REDACTED", tally)
		}
	}
}
