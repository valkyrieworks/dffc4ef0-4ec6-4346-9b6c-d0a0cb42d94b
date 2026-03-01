package scratchfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

const (
	indivisiblePersistRecordHeading = "REDACTED"
	//
	//
	indivisiblePersistRecordMaximumCountDisagreements = 5
	//
	//
	indivisiblePersistRecordMaximumCountPersistEndeavors = 1000
	//
	//
	rngAN = 6364136223846793005
	rngCN = 1442695040888963407
	//
	//
	//
	indivisiblePersistRecordMarker = os.O_WRONLY | os.O_CREATE | os.O_SYNC | os.O_TRUNC | os.O_EXCL
)

var (
	indivisiblePersistRecordArbitrary   uint64
	indivisiblePersistRecordArbitraryMutex commitchronize.Exclusion
)

func persistRecordArbitraryRegenerate() uint64 {
	//
	//
	//
	//
	//
	//
	//
	return uint64(time.Now().UnixNano() + int64(os.Getpid()<<20))
}

//
//
//
func arbitraryPersistRecordEnding() string {
	indivisiblePersistRecordArbitraryMutex.Lock()
	r := indivisiblePersistRecordArbitrary
	if r == 0 {
		r = persistRecordArbitraryRegenerate()
	}

	//
	r = r*rngAN + rngCN

	indivisiblePersistRecordArbitrary = r
	indivisiblePersistRecordArbitraryMutex.Unlock()
	//
	ending := strconv.Itoa(int(r))
	if string(ending[0]) == "REDACTED" {
		//
		//
		ending = strings.Replace(ending, "REDACTED", "REDACTED", 1)
	}
	return ending
}

//
//
func PersistRecordIndivisible(recordname string, data []byte, mode os.FileMode) (err error) {
	//
	//
	//
	//
	//
	//
	var (
		dir = filepath.Dir(recordname)
		f   *os.File
	)

	nclash := 0
	//
	//
	//
	i := 0
	for ; i < indivisiblePersistRecordMaximumCountPersistEndeavors; i++ {
		alias := filepath.Join(dir, indivisiblePersistRecordHeading+arbitraryPersistRecordEnding())
		f, err = os.OpenFile(alias, indivisiblePersistRecordMarker, mode)
		//
		if os.IsExist(err) {
			//
			//
			if nclash++; nclash > indivisiblePersistRecordMaximumCountDisagreements {
				indivisiblePersistRecordArbitraryMutex.Lock()
				indivisiblePersistRecordArbitrary = persistRecordArbitraryRegenerate()
				indivisiblePersistRecordArbitraryMutex.Unlock()
			}
			continue
		} else if err != nil {
			return err
		}
		break
	}
	if i == indivisiblePersistRecordMaximumCountPersistEndeavors {
		return fmt.Errorf("REDACTED", i)
	}

	//
	defer os.Remove(f.Name())
	defer f.Close()

	if n, err := f.Write(data); err != nil {
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	//
	//
	f.Close()

	return os.Rename(f.Name(), recordname)
}
