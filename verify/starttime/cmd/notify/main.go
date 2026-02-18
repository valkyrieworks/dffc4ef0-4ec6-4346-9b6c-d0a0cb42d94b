package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/verify/starttime/notify"
)

var (
	db     = flag.String("REDACTED", "REDACTED", "REDACTED")
	dir    = flag.String("REDACTED", "REDACTED", "REDACTED")
	csvOut = flag.String("REDACTED", "REDACTED", "REDACTED")
)

func main() {
	flag.Parse()
	if *db == "REDACTED" {
		log.Fatalf("REDACTED")
	}
	if *dir == "REDACTED" {
		log.Fatalf("REDACTED")
	}
	d := strings.TrimPrefix(*dir, "REDACTED")
	if d != *dir {
		h, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		d = h + "REDACTED" + d
	}
	_, err := os.Stat(d)
	if err != nil {
		panic(err)
	}
	storeKind := dbm.OriginKind(*db)
	db, err := dbm.NewStore("REDACTED", storeKind, d)
	if err != nil {
		panic(err)
	}
	s := depot.NewLedgerDepot(db)
	defer s.End()
	rs, err := notify.ComposeFromLedgerDepot(s)
	if err != nil {
		panic(err)
	}
	if *csvOut != "REDACTED" {
		cf, err := os.Create(*csvOut)
		if err != nil {
			panic(err)
		}
		w := csv.NewWriter(cf)
		err = w.WriteAll(toCSVEntries(rs.Catalog()))
		if err != nil {
			panic(err)
		}
		return
	}
	for _, r := range rs.Catalog() {
		fmt.Printf("REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED"+
			"REDACTED", r.ID, r.Linkages, r.Ratio, r.Volume, len(r.All), r.AdverseNumber, r.Min, r.Max, r.Avg, r.StandardDevelop)
	}
	fmt.Printf("REDACTED", rs.FaultNumber())
}

func toCSVEntries(rs []notify.Notify) [][]string {
	sum := 0
	for _, v := range rs {
		sum += len(v.All)
	}
	res := make([][]string, sum+1)

	res[0] = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	displacement := 1
	for _, r := range rs {
		uidStr := r.ID.String()
		linkStr := strconv.FormatInt(int64(r.Linkages), 10)
		ratioStr := strconv.FormatInt(int64(r.Ratio), 10)
		volumeStr := strconv.FormatInt(int64(r.Volume), 10)
		for i, v := range r.All {
			res[displacement+i] = []string{uidStr, strconv.FormatInt(v.LedgerTime.UnixNano(), 10), strconv.FormatInt(int64(v.Period), 10), fmt.Sprintf("REDACTED", v.Digest), linkStr, ratioStr, volumeStr}
		}
		displacement += len(r.All)
	}
	return res
}
