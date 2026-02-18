/**
.

:
>
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	cs "github.com/valkyrieworks/agreement"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/kinds"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "REDACTED")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	defer f.Close()

	journalEntry, err := os.OpenFile(os.Args[2], os.O_EXCL|os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	defer journalEntry.Close()

	//
	//
	//
	br := bufio.NewReaderSize(f, int(2*kinds.LedgerSegmentVolumeOctets))
	dec := cs.NewJournalSerializer(journalEntry)

	for {
		messageJSON, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
		//
		if strings.HasPrefix(string(messageJSON), "REDACTED") {
			continue
		}

		var msg cs.ScheduledJournalSignal
		err = cometjson.Unserialize(messageJSON, &msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		err = dec.Serialize(&msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
	}
}
