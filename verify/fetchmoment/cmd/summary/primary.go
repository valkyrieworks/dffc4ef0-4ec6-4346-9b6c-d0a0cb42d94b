package primary

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/fetchmoment/summary"
)

var (
	db     = flag.String("REDACTED", "REDACTED", "REDACTED")
	dir    = flag.String("REDACTED", "REDACTED", "REDACTED")
	spreadsheetOutput = flag.String("REDACTED", "REDACTED", "REDACTED")
)

func primary() {
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
	datastoreKind := dbm.OriginKind(*db)
	db, err := dbm.FreshDatastore("REDACTED", datastoreKind, d)
	if err != nil {
		panic(err)
	}
	s := depot.FreshLedgerDepot(db)
	defer s.Shutdown()
	rs, err := summary.ComposeOriginatingLedgerDepot(s)
	if err != nil {
		panic(err)
	}
	if *spreadsheetOutput != "REDACTED" {
		cf, err := os.Create(*spreadsheetOutput)
		if err != nil {
			panic(err)
		}
		w := csv.NewWriter(cf)
		err = w.WriteAll(towardSpreadsheetEntries(rs.Catalog()))
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
			"REDACTED", r.ID, r.Linkages, r.Frequency, r.Extent, len(r.All), r.AdverseTally, r.Min, r.Max, r.Avg, r.StandardDevelop)
	}
	fmt.Printf("REDACTED", rs.FailureTally())
}

func towardSpreadsheetEntries(rs []summary.Summary) [][]string {
	sum := 0
	for _, v := range rs {
		sum += len(v.All)
	}
	res := make([][]string, sum+1)

	res[0] = []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"}
	displacement := 1
	for _, r := range rs {
		uuidTxt := r.ID.String()
		linkTxt := strconv.FormatInt(int64(r.Linkages), 10)
		frequencyTxt := strconv.FormatInt(int64(r.Frequency), 10)
		extentTxt := strconv.FormatInt(int64(r.Extent), 10)
		for i, v := range r.All {
			res[displacement+i] = []string{uuidTxt, strconv.FormatInt(v.LedgerMoment.UnixNano(), 10), strconv.FormatInt(int64(v.Interval), 10), fmt.Sprintf("REDACTED", v.Digest), linkTxt, frequencyTxt, extentTxt}
		}
		displacement += len(r.All)
	}
	return res
}
