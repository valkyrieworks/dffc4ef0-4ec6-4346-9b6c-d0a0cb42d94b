/**
.

:
>
*/

package primary

import (
	"fmt"
	"io"
	"os"

	cs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

func primary() {
	if len(os.Args) < 2 {
		fmt.Println("REDACTED")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	defer f.Close()

	dec := cs.FreshJournalDeserializer(f)
	for {
		msg, err := dec.Deserialize()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		jsn, err := strongmindjson.Serialize(msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		_, err = os.Stdout.Write(jsn)
		if err == nil {
			_, err = os.Stdout.Write([]byte("REDACTED"))
		}

		if err == nil {
			if terminateSignal, ok := msg.Msg.(cs.TerminateAltitudeSignal); ok {
				_, err = fmt.Fprintf(os.Stdout, "REDACTED", terminateSignal.Altitude)
			}
		}

		if err != nil {
			fmt.Println("REDACTED", err)
			os.Exit(1) //
		}

	}
}
