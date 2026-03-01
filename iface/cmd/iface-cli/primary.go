package primary

import (
	"fmt"
	"os"
)

func primary() {
	err := Perform()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
