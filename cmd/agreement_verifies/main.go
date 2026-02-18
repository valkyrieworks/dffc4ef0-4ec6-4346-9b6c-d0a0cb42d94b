package main

import (
	"fmt"
	"strings"

	"github.com/snikch/goodman/hooks"
	"github.com/snikch/goodman/transaction"
)

func main() {
	//
	h := hooks.NewHooks()
	host := hooks.NewServer(hooks.NewHooksRunner(h))
	h.BeforeAll(func(t []*transaction.Transaction) {
		fmt.Println(t[0].Name)
	})
	h.BeforeEach(func(t *transaction.Transaction) {
		if strings.HasPrefix(t.Name, "REDACTED") ||
			//
			strings.HasPrefix(t.Name, "REDACTED") ||
			//
			strings.HasPrefix(t.Name, "REDACTED") ||
			//
			//
			strings.HasPrefix(t.Name, "REDACTED") {
			t.Skip = true
			fmt.Printf("REDACTED", t.Name)
		}
	})
	host.Serve()
	defer host.Listener.Close()
	fmt.Print("REDACTED")
}
