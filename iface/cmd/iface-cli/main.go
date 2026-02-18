package main

import (
	"fmt"
	"os"
)

func main() {
	err := Perform()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
