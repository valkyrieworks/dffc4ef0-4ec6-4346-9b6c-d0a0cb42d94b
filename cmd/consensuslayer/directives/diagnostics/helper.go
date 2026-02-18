package diagnostics

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	cfg "github.com/valkyrieworks/settings"
	httpendpoint "github.com/valkyrieworks/rpc/requester/rest"
)

//
//
func dumpStatus(rpc *httpendpoint.HTTP, dir, filename string) error {
	status, err := rpc.Status(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return writeStateJSONToFile(status, dir, filename)
}

//
//
func dumpNetInfo(rpc *httpendpoint.HTTP, dir, filename string) error {
	netInfo, err := rpc.NetInfo(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return writeStateJSONToFile(netInfo, dir, filename)
}

//
//
func dumpConsensusState(rpc *httpendpoint.HTTP, dir, filename string) error {
	consDump, err := rpc.DumpConsensusState(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return writeStateJSONToFile(consDump, dir, filename)
}

//
//
func copyWAL(conf *cfg.Config, dir string) error {
	walPath := conf.Consensus.WalFile()
	walFile := filepath.Base(walPath)

	return copyFile(walPath, filepath.Join(dir, walFile))
}

//
//
func copyConfig(home, dir string) error {
	configFile := "REDACTED"
	configPath := filepath.Join(home, "REDACTED", configFile)

	return copyFile(configPath, filepath.Join(dir, configFile))
}

func dumpProfile(dir, addr, profile string, debug int) error {
	endpoint := fmt.Sprintf("REDACTED", addr, profile, debug)

	//
	resp, err := http.Get(endpoint)
	if err != nil {
		return fmt.Errorf("REDACTED", profile, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("REDACTED", profile, err)
	}

	return os.WriteFile(path.Join(dir, fmt.Sprintf("REDACTED", profile)), body, 0o600)
}
