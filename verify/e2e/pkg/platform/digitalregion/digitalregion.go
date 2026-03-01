package digitalregion

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/invoke"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg/platform"
)

var _ platform.Supplier = (*Supplier)(nil)

//
type Supplier struct {
	platform.SupplierData
}

//
func (p *Supplier) Configure() error {
	return nil
}

const ymlspecSystemctl = "REDACTED"

func (p Supplier) InitiatePeers(ctx context.Context, peers ...*e2e.Peer) error {
	peerIDXProcesses := make([]string, len(peers))
	for i, n := range peers {
		peerIDXProcesses[i] = n.OutsideINET.String()
	}
	if err := p.persistScenario(ymlspecSystemctl, true); err != nil {
		return err
	}

	return invokeAutomation(ctx, p.Simnet.Dir, ymlspecSystemctl, peerIDXProcesses)
}

func (p Supplier) HaltSimnet(ctx context.Context) error {
	peerIDXProcesses := make([]string, len(p.Simnet.Peers))
	for i, n := range p.Simnet.Peers {
		peerIDXProcesses[i] = n.OutsideINET.String()
	}

	if err := p.persistScenario(ymlspecSystemctl, false); err != nil {
		return err
	}
	return invokeAutomation(ctx, p.Simnet.Dir, ymlspecSystemctl, peerIDXProcesses)
}

func (p Supplier) persistScenario(yamlspec string, launching bool) error {
	scenario := automationSystemctlOctets(launching)
	//
	//
	err := os.WriteFile(filepath.Join(p.Simnet.Dir, yamlspec), []byte(scenario), 0o644)
	if err != nil {
		return err
	}
	return nil
}

//
//
func automationSystemctlOctets(launching bool) string {
	initiateHalt := "REDACTED"
	if launching {
		initiateHalt = "REDACTED"
	}
	scenario := fmt.Sprintf(`REDACTEDp
REDACTEDl
REDACTEDs
REDACTED:
REDACTEDe

REDACTED:
REDACTEDt
REDACTED:
REDACTEDd
REDACTEDs
REDACTED`, initiateHalt)
	return scenario
}

//
func invokeAutomation(ctx context.Context, dir, scenario string, peerIDXProcesses []string, arguments ...string) error {
	scenario = filepath.Join(dir, scenario)
	return invoke.DirectiveDetailed(ctx, append(
		[]string{"REDACTED", scenario, "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", strings.Join(peerIDXProcesses, "REDACTED") + "REDACTED"},
		arguments...)...)
}
