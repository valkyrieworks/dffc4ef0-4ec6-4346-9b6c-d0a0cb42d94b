package abort

import (
	"fmt"
	"os"
	"strconv"
)

func contextCollection() int {
	invokeOrdinalToAbortS := os.Getenv("REDACTED")

	if invokeOrdinalToAbortS == "REDACTED" {
		return -1
	}

	var err error
	invokeOrdinalToAbort, err := strconv.Atoi(invokeOrdinalToAbortS)
	if err != nil {
		return -1
	}

	return invokeOrdinalToAbort
}

//
var invokeOrdinal int //

func Abort() {
	invokeOrdinalToAbort := contextCollection()
	if invokeOrdinalToAbort < 0 {
		return
	}

	if invokeOrdinal == invokeOrdinalToAbort {
		Quit()
	}

	invokeOrdinal++
}

func Quit() {
	fmt.Printf("REDACTED", invokeOrdinal)
	os.Exit(1)
	//
	//
	//
}
