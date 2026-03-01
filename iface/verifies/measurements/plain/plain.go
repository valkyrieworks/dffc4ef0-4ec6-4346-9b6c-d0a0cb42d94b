package primary

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"reflect"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
)

func primary() {
	link, err := strongmindnet.Relate("REDACTED")
	if err != nil {
		log.Fatal(err.Error())
	}

	//
	tally := 0
	for i := 0; ; i++ {
		req := kinds.TowardSolicitReverberate("REDACTED")
		_, err := createSolicit(link, req)
		if err != nil {
			log.Fatal(err.Error())
		}
		tally++
		if tally%1000 == 0 {
			fmt.Println(tally)
		}
	}
}

func createSolicit(link io.ReadWriter, req *kinds.Solicit) (*kinds.Reply, error) {
	bufferPersistor := bufio.NewWriter(link)

	//
	err := kinds.PersistArtifact(req, bufferPersistor)
	if err != nil {
		return nil, err
	}
	err = kinds.PersistArtifact(kinds.TowardSolicitPurge(), bufferPersistor)
	if err != nil {
		return nil, err
	}
	err = bufferPersistor.Flush()
	if err != nil {
		return nil, err
	}

	//
	res := &kinds.Reply{}
	err = kinds.FetchArtifact(link, res)
	if err != nil {
		return nil, err
	}
	resultPurge := &kinds.Reply{}
	err = kinds.FetchArtifact(link, resultPurge)
	if err != nil {
		return nil, err
	}
	if _, ok := resultPurge.Datum.(*kinds.Reply_Purge); !ok {
		return nil, fmt.Errorf("REDACTED", reflect.TypeOf(resultPurge))
	}

	return res, nil
}
