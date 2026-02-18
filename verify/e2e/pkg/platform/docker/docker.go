package docker

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"text/template"

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
//
func (p *Source) Configure() error {
	arrange, err := dockerAssembleOctets(p.Verifychain)
	if err != nil {
		return err
	}
	//
	//
	err = os.WriteFile(filepath.Join(p.Verifychain.Dir, "REDACTED"), arrange, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func (p Source) BeginInstances(ctx context.Context, instances ...*e2e.Member) error {
	memberLabels := make([]string, len(instances))
	for i, n := range instances {
		memberLabels[i] = n.Label
	}
	return InvokeAssemble(ctx, p.Verifychain.Dir, append([]string{"REDACTED", "REDACTED"}, memberLabels...)...)
}

func (p Source) HaltVerifychain(ctx context.Context) error {
	return InvokeAssemble(ctx, p.Verifychain.Dir, "REDACTED")
}

//
//
func dockerAssembleOctets(verifychain *e2e.Verifychain) ([]byte, error) {
	//
	tmpl, err := template.New("REDACTED").Parse(`REDACTED'
REDACTED:
REDACTED:
REDACTED:
REDACTEDe
REDACTEDe
REDACTED}
REDACTEDe
REDACTED}
REDACTED:
REDACTEDt
REDACTED:
REDACTED}

REDACTED:
REDACTED}
REDACTED:
REDACTED:
REDACTEDe
REDACTED}
REDACTED}
REDACTED}
REDACTEDn
REDACTED}
REDACTEDe
REDACTED:
REDACTED6
REDACTED7
REDACTED}
REDACTED0
REDACTED}
REDACTED0
REDACTED5
REDACTED6
REDACTED:
REDACTEDt
REDACTEDt
REDACTED:
REDACTED:
REDACTED}
REDACTED}

REDACTED:
REDACTED:
REDACTEDe
REDACTEDu
REDACTED}
REDACTED}
REDACTEDn
REDACTED}
REDACTEDe
REDACTED:
REDACTED6
REDACTED7
REDACTED}
REDACTED0
REDACTED}
REDACTED0
REDACTED5
REDACTED6
REDACTED:
REDACTEDt
REDACTEDt
REDACTED:
REDACTED:
REDACTED}
REDACTED}

REDACTED`)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, verifychain)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//
func InvokeAssemble(ctx context.Context, dir string, args ...string) error {
	return invoke.Directive(ctx, append(
		[]string{"REDACTED", "REDACTED", "REDACTED", filepath.Join(dir, "REDACTED")},
		args...)...)
}

//
func InvokeAssembleResult(ctx context.Context, dir string, args ...string) ([]byte, error) {
	return invoke.DirectiveResult(ctx, append(
		[]string{"REDACTED", "REDACTED", "REDACTED", filepath.Join(dir, "REDACTED")},
		args...)...)
}

//
func InvokeAssembleDetailed(ctx context.Context, dir string, args ...string) error {
	return invoke.DirectiveDetailed(ctx, append(
		[]string{"REDACTED", "REDACTED", "REDACTED", filepath.Join(dir, "REDACTED")},
		args...)...)
}

//
func Invoke(ctx context.Context, args ...string) error {
	return invoke.Directive(ctx, append([]string{"REDACTED"}, args...)...)
}
