package main

import (
	"flag"
	"os"
	"time"

	"github.com/valkyrieworks/security/curve25519"
	"github.com/valkyrieworks/utils/log"
	enginenet "github.com/valkyrieworks/utils/net"
	ctsystem "github.com/valkyrieworks/utils/os"

	"github.com/valkyrieworks/authkey"
)

func main() {
	var (
		addr             = flag.String("REDACTED", "REDACTED", "REDACTED")
		chainID          = flag.String("REDACTED", "REDACTED", "REDACTED")
		privValKeyPath   = flag.String("REDACTED", "REDACTED", "REDACTED")
		privValStatePath = flag.String("REDACTED", "REDACTED", "REDACTED")

		logger = log.NewTMLogger(
			log.NewSyncWriter(os.Stdout),
		).With("REDACTED", "REDACTED")
	)
	flag.Parse()

	logger.Info(
		"REDACTED",
		"REDACTED", *addr,
		"REDACTED", *chainID,
		"REDACTED", *privValKeyPath,
		"REDACTED", *privValStatePath,
	)

	pv := authkey.LoadFilePV(*privValKeyPath, *privValStatePath)

	var dialer authkey.SocketDialer
	protocol, address := enginenet.ProtocolAndAddress(*addr)
	switch protocol {
	case "REDACTED":
		dialer = authkey.DialUnixFn(address)
	case "REDACTED":
		connTimeout := 3 * time.Second //
		dialer = authkey.DialTCPFn(address, connTimeout, curve25519.GenPrivKey())
	default:
		logger.Error("REDACTED", "REDACTED", protocol)
		os.Exit(1)
	}

	sd := authkey.NewSignerDialerEndpoint(logger, dialer)
	ss := authkey.NewSignerServer(sd, *chainID, pv)

	err := ss.Start()
	if err != nil {
		panic(err)
	}

	//
	ctsystem.TrapSignal(logger, func() {
		err := ss.Stop()
		if err != nil {
			panic(err)
		}
	})

	//
	select {}
}
