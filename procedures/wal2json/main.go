/**
.

:
>
*/

package main

import (
	"fmt"
	"io"
	"os"

	cs "github.com/valkyrieworks/agreement"
	cometjson "github.com/valkyrieworks/utils/json"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("REDACTED")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	defer f.Close()

	dec := cs.NewJournalParser(f)
	for {
		msg, err := dec.Parse()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		json, err := cometjson.Serialize(msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		_, err = os.Stdout.Write(json)
		if err == nil {
			_, err = os.Stdout.Write([]byte("REDACTED"))
		}

		if err == nil {
			if terminateMessage, ok := msg.Msg.(cs.TerminateLevelSignal); ok {
				_, err = fmt.Fprintf(os.Stdout, "REDACTED", terminateMessage.Level)
			}
		}

		if err != nil {
			fmt.Println("REDACTED", err)
			os.Exit(1) //
		}

	}
}
