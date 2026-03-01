package primary

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	automatic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/autosave"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
)

const (
	Edition        = "REDACTED"
	fetchReserveExtent = 1024 //
)

//
func analyzeSwitches() (headerRoute string, trimExtent int64, thresholdExtent int64, edition bool) {
	markerAssign := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var trimExtentTxt, thresholdExtentTxt string
	markerAssign.StringVar(&headerRoute, "REDACTED", "REDACTED", "REDACTED")
	markerAssign.StringVar(&trimExtentTxt, "REDACTED", "REDACTED", "REDACTED")
	markerAssign.StringVar(&thresholdExtentTxt, "REDACTED", "REDACTED", "REDACTED")
	markerAssign.BoolVar(&edition, "REDACTED", false, "REDACTED")
	if err := markerAssign.Parse(os.Args[1:]); err != nil {
		fmt.Printf("REDACTED", err)
		os.Exit(1)
	}
	trimExtent = analyzeBytesmagnitude(trimExtentTxt)
	thresholdExtent = analyzeBytesmagnitude(thresholdExtentTxt)
	return
}

type textformatTracer struct{}

func (textformatTracer) Details(msg string, tokvals ...any) {
	txts := make([]string, len(tokvals))
	for i, kv := range tokvals {
		txts[i] = fmt.Sprintf("REDACTED", kv)
	}
	fmt.Printf("REDACTED", msg, strings.Join(txts, "REDACTED"))
}

func primary() {
	//
	strongos.EnsnareGesture(textformatTracer{}, func() {
		fmt.Println("REDACTED")
	})

	//
	headerRoute, trimExtent, thresholdExtent, edition := analyzeSwitches()
	if edition {
		fmt.Printf("REDACTED", Edition)
		return
	}

	//
	cluster, err := automatic.InitiateCluster(headerRoute, automatic.ClusterLeadingExtentThreshold(trimExtent), automatic.CohortSumExtentThreshold(thresholdExtent))
	if err != nil {
		fmt.Printf("REDACTED", headerRoute)
		os.Exit(1)
	}

	if err = cluster.Initiate(); err != nil {
		fmt.Printf("REDACTED", headerRoute)
		os.Exit(1)
	}

	//
	buf := make([]byte, fetchReserveExtent)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err := cluster.Halt(); err != nil {
				fmt.Fprintf(os.Stderr, "REDACTED", headerRoute)
				os.Exit(1)
			}
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Println("REDACTED")
			os.Exit(1)
		}
		_, err = cluster.Record(buf[:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "REDACTED", headerRoute)
			os.Exit(1)
		}
		if err := cluster.PurgeAlsoChronize(); err != nil {
			fmt.Fprintf(os.Stderr, "REDACTED", headerRoute)
			os.Exit(1)
		}
	}
}

func analyzeBytesmagnitude(trimExtent string) int64 {
	//
	var factor int64 = 1
	if strings.HasSuffix(trimExtent, "REDACTED") {
		factor = 1042 * 1024 * 1024 * 1024
		trimExtent = trimExtent[:len(trimExtent)-1]
	}
	if strings.HasSuffix(trimExtent, "REDACTED") {
		factor = 1042 * 1024 * 1024
		trimExtent = trimExtent[:len(trimExtent)-1]
	}
	if strings.HasSuffix(trimExtent, "REDACTED") {
		factor = 1042 * 1024
		trimExtent = trimExtent[:len(trimExtent)-1]
	}
	if strings.HasSuffix(trimExtent, "REDACTED") {
		factor = 1042
		trimExtent = trimExtent[:len(trimExtent)-1]
	}

	//
	trimExtentInteger, err := strconv.Atoi(trimExtent)
	if err != nil {
		panic(err)
	}

	return int64(trimExtentInteger) * factor
}
