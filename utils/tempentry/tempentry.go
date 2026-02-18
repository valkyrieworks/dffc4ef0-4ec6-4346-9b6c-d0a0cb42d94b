package tempentry

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	engineconnect "github.com/valkyrieworks/utils/align"
)

const (
	atomicRecordEntryPrefix = "REDACTED"
	//
	//
	atomicRecordEntryMaximumCountClashes = 5
	//
	//
	atomicRecordEntryMaximumCountRecordTries = 1000
	//
	//
	lcgA = 6364136223846793005
	lcgC = 1442695040888963407
	//
	//
	//
	atomicRecordEntryMark = os.O_WRONLY | os.O_CREATE | os.O_SYNC | os.O_TRUNC | os.O_EXCL
)

var (
	atomicRecordEntryRandom   uint64
	atomicRecordEntryRandomMu engineconnect.Lock
)

func recordEntryRandomRegenerate() uint64 {
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
func randomRecordEntryPostfix() string {
	atomicRecordEntryRandomMu.Lock()
	r := atomicRecordEntryRandom
	if r == 0 {
		r = recordEntryRandomRegenerate()
	}

	//
	r = r*lcgA + lcgC

	atomicRecordEntryRandom = r
	atomicRecordEntryRandomMu.Unlock()
	//
	postfix := strconv.Itoa(int(r))
	if string(postfix[0]) == "REDACTED" {
		//
		//
		postfix = strings.Replace(postfix, "REDACTED", "REDACTED", 1)
	}
	return postfix
}

//
//
func RecordEntryAtomic(filename string, data []byte, mode os.FileMode) (err error) {
	//
	//
	//
	//
	//
	//
	var (
		dir = filepath.Dir(filename)
		f   *os.File
	)

	nclash := 0
	//
	//
	//
	i := 0
	for ; i < atomicRecordEntryMaximumCountRecordTries; i++ {
		label := filepath.Join(dir, atomicRecordEntryPrefix+randomRecordEntryPostfix())
		f, err = os.OpenFile(label, atomicRecordEntryMark, mode)
		//
		if os.IsExist(err) {
			//
			//
			if nclash++; nclash > atomicRecordEntryMaximumCountClashes {
				atomicRecordEntryRandomMu.Lock()
				atomicRecordEntryRandom = recordEntryRandomRegenerate()
				atomicRecordEntryRandomMu.Unlock()
			}
			continue
		} else if err != nil {
			return err
		}
		break
	}
	if i == atomicRecordEntryMaximumCountRecordTries {
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

	return os.Rename(f.Name(), filename)
}
