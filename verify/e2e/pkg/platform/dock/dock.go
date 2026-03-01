package dock

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"text/template"

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
//
func (p *Supplier) Configure() error {
	construct, err := dockArrangeOctets(p.Simnet)
	if err != nil {
		return err
	}
	//
	//
	err = os.WriteFile(filepath.Join(p.Simnet.Dir, "REDACTED"), construct, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func (p Supplier) InitiatePeers(ctx context.Context, peers ...*e2e.Peer) error {
	peerIdentifiers := make([]string, len(peers))
	for i, n := range peers {
		peerIdentifiers[i] = n.Alias
	}
	return InvokeArrange(ctx, p.Simnet.Dir, append([]string{"REDACTED", "REDACTED"}, peerIdentifiers...)...)
}

func (p Supplier) HaltSimnet(ctx context.Context) error {
	return InvokeArrange(ctx, p.Simnet.Dir, "REDACTED")
}

//
//
func dockArrangeOctets(simnet *e2e.Simnet) ([]byte, error) {
	//
	layout, err := template.New("REDACTED").Parse(`REDACTED'
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
	err = layout.Execute(&buf, simnet)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//
func InvokeArrange(ctx context.Context, dir string, arguments ...string) error {
	return invoke.Directive(ctx, append(
		[]string{"REDACTED", "REDACTED", "REDACTED", filepath.Join(dir, "REDACTED")},
		arguments...)...)
}

//
func InvokeArrangeEmission(ctx context.Context, dir string, arguments ...string) ([]byte, error) {
	return invoke.DirectiveEmission(ctx, append(
		[]string{"REDACTED", "REDACTED", "REDACTED", filepath.Join(dir, "REDACTED")},
		arguments...)...)
}

//
func InvokeArrangeDetailed(ctx context.Context, dir string, arguments ...string) error {
	return invoke.DirectiveDetailed(ctx, append(
		[]string{"REDACTED", "REDACTED", "REDACTED", filepath.Join(dir, "REDACTED")},
		arguments...)...)
}

//
func Invoke(ctx context.Context, arguments ...string) error {
	return invoke.Directive(ctx, append([]string{"REDACTED"}, arguments...)...)
}
