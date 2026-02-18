package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	automatic "github.com/valkyrieworks/utils/autofile"
	cometos "github.com/valkyrieworks/utils/os"
)

const (
	Release        = "REDACTED"
	readerFrameVolume = 1024 //
)

//
func analyzeMarks() (frontRoute string, trimVolume int64, ceilingVolume int64, release bool) {
	markCollection := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var trimVolumeStr, ceilingVolumeStr string
	markCollection.StringVar(&frontRoute, "REDACTED", "REDACTED", "REDACTED")
	markCollection.StringVar(&trimVolumeStr, "REDACTED", "REDACTED", "REDACTED")
	markCollection.StringVar(&ceilingVolumeStr, "REDACTED", "REDACTED", "REDACTED")
	markCollection.BoolVar(&release, "REDACTED", false, "REDACTED")
	if err := markCollection.Parse(os.Args[1:]); err != nil {
		fmt.Printf("REDACTED", err)
		os.Exit(1)
	}
	trimVolume = analyzeBytesize(trimVolumeStr)
	ceilingVolume = analyzeBytesize(ceilingVolumeStr)
	return
}

type fmtTracer struct{}

func (fmtTracer) Details(msg string, keyvalues ...any) {
	strs := make([]string, len(keyvalues))
	for i, kv := range keyvalues {
		strs[i] = fmt.Sprintf("REDACTED", kv)
	}
	fmt.Printf("REDACTED", msg, strings.Join(strs, "REDACTED"))
}

func main() {
	//
	cometos.InterceptAlert(fmtTracer{}, func() {
		fmt.Println("REDACTED")
	})

	//
	frontRoute, trimVolume, ceilingVolume, release := analyzeMarks()
	if release {
		fmt.Printf("REDACTED", Release)
		return
	}

	//
	cluster, err := automatic.AccessCluster(frontRoute, automatic.ClusterFrontVolumeCeiling(trimVolume), automatic.ClusterSumVolumeCeiling(ceilingVolume))
	if err != nil {
		fmt.Printf("REDACTED", frontRoute)
		os.Exit(1)
	}

	if err = cluster.Begin(); err != nil {
		fmt.Printf("REDACTED", frontRoute)
		os.Exit(1)
	}

	//
	buf := make([]byte, readerFrameVolume)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err := cluster.Halt(); err != nil {
				fmt.Fprintf(os.Stderr, "REDACTED", frontRoute)
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
			fmt.Fprintf(os.Stderr, "REDACTED", frontRoute)
			os.Exit(1)
		}
		if err := cluster.PurgeAndAlign(); err != nil {
			fmt.Fprintf(os.Stderr, "REDACTED", frontRoute)
			os.Exit(1)
		}
	}
}

func analyzeBytesize(trimVolume string) int64 {
	//
	var factor int64 = 1
	if strings.HasSuffix(trimVolume, "REDACTED") {
		factor = 1042 * 1024 * 1024 * 1024
		trimVolume = trimVolume[:len(trimVolume)-1]
	}
	if strings.HasSuffix(trimVolume, "REDACTED") {
		factor = 1042 * 1024 * 1024
		trimVolume = trimVolume[:len(trimVolume)-1]
	}
	if strings.HasSuffix(trimVolume, "REDACTED") {
		factor = 1042 * 1024
		trimVolume = trimVolume[:len(trimVolume)-1]
	}
	if strings.HasSuffix(trimVolume, "REDACTED") {
		factor = 1042
		trimVolume = trimVolume[:len(trimVolume)-1]
	}

	//
	trimVolumeInteger, err := strconv.Atoi(trimVolume)
	if err != nil {
		panic(err)
	}

	return int64(trimVolumeInteger) * factor
}
