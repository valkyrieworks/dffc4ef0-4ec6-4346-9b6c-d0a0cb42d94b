package cloudprovider

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/verify/e2e/pkg/invoke"
	"github.com/valkyrieworks/verify/e2e/pkg/platform"
)

var _ platform.Source = (*Source)(nil)

//
type Source struct {
	platform.SourceData
}

//
func (p *Source) Configure() error {
	return nil
}

const ymlSystemd = "REDACTED"

func (p Source) BeginInstances(ctx context.Context, instances ...*e2e.Member) error {
	memberIDXPs := make([]string, len(instances))
	for i, n := range instances {
		memberIDXPs[i] = n.OutsideIP.String()
	}
	if err := p.recordScript(ymlSystemd, true); err != nil {
		return err
	}

	return invokeAnsible(ctx, p.Verifychain.Dir, ymlSystemd, memberIDXPs)
}

func (p Source) HaltVerifychain(ctx context.Context) error {
	memberIDXPs := make([]string, len(p.Verifychain.Instances))
	for i, n := range p.Verifychain.Instances {
		memberIDXPs[i] = n.OutsideIP.String()
	}

	if err := p.recordScript(ymlSystemd, false); err != nil {
		return err
	}
	return invokeAnsible(ctx, p.Verifychain.Dir, ymlSystemd, memberIDXPs)
}

func (p Source) recordScript(yaml string, launching bool) error {
	script := ansibleSystemdOctets(launching)
	//
	//
	err := os.WriteFile(filepath.Join(p.Verifychain.Dir, yaml), []byte(script), 0o644)
	if err != nil {
		return err
	}
	return nil
}

//
//
func ansibleSystemdOctets(launching bool) string {
	beginHalt := "REDACTED"
	if launching {
		beginHalt = "REDACTED"
	}
	script := fmt.Sprintf(`REDACTEDp
REDACTEDl
REDACTEDs
REDACTED:
REDACTEDe

REDACTED:
REDACTEDt
REDACTED:
REDACTEDd
REDACTEDs
REDACTED`, beginHalt)
	return script
}

//
func invokeAnsible(ctx context.Context, dir, script string, memberIDXPs []string, args ...string) error {
	script = filepath.Join(dir, script)
	return invoke.DirectiveDetailed(ctx, append(
		[]string{"REDACTED", script, "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", strings.Join(memberIDXPs, "REDACTED") + "REDACTED"},
		args...)...)
}
