/**
.

:
>
*/

package primary

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	cs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func primary() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "REDACTED")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	defer f.Close()

	journalRecord, err := os.OpenFile(os.Args[2], os.O_EXCL|os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	defer journalRecord.Close()

	//
	//
	//
	br := bufio.NewReaderSize(f, int(2*kinds.LedgerFragmentExtentOctets))
	dec := cs.FreshJournalSerializer(journalRecord)

	for {
		signalJSN, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
		//
		if strings.HasPrefix(string(signalJSN), "REDACTED") {
			continue
		}

		var msg cs.ScheduledJournalSignal
		err = strongmindjson.Decode(signalJSN, &msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		err = dec.Serialize(&msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
	}
}
