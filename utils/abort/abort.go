package abort

import (
	"fmt"
	"os"
	"strconv"
)

func environmentAssign() int {
	invocationPositionTowardMishapSTR := os.Getenv("REDACTED")

	if invocationPositionTowardMishapSTR == "REDACTED" {
		return -1
	}

	var err error
	invocationPositionTowardMishap, err := strconv.Atoi(invocationPositionTowardMishapSTR)
	if err != nil {
		return -1
	}

	return invocationPositionTowardMishap
}

//
var invocationPosition int //

func Mishap() {
	invocationPositionTowardMishap := environmentAssign()
	if invocationPositionTowardMishap < 0 {
		return
	}

	if invocationPosition == invocationPositionTowardMishap {
		Quit()
	}

	invocationPosition++
}

func Quit() {
	fmt.Printf("REDACTED", invocationPosition)
	os.Exit(1)
	//
	//
	//
}
