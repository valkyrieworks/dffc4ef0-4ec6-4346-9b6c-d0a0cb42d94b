package primary

import (
	"bufio"
	"fmt"
	"log"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
)

func primary() {
	link, err := strongmindnet.Relate("REDACTED")
	if err != nil {
		log.Fatal(err.Error())
	}

	//
	go func() {
		tally := 0
		for {
			res := &kinds.Reply{}
			err := kinds.FetchArtifact(link, res)
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
		bufferPersistor := bufio.NewWriter(link)
		req := kinds.TowardSolicitReverberate("REDACTED")

		err := kinds.PersistArtifact(req, bufferPersistor)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = bufferPersistor.Flush()
		if err != nil {
			log.Fatal(err.Error())
		}

		tally++
		if tally%1000 == 0 {
			fmt.Println("REDACTED", tally)
		}
	}
}
