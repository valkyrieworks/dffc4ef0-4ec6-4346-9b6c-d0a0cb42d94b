package platform

import (
	"context"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
)

//
//
type Supplier interface {
	//
	//
	Configure() error

	//
	//
	//
	InitiatePeers(context.Context, ...*e2e.Peer) error

	//
	HaltSimnet(context.Context) error

	//
	ObtainFrameworkData() *e2e.FrameworkData
}

type SupplierData struct {
	Simnet            *e2e.Simnet
	FrameworkData e2e.FrameworkData
}

//
func (pd SupplierData) ObtainFrameworkData() *e2e.FrameworkData {
	return &pd.FrameworkData
}
