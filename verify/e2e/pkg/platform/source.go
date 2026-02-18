package platform

import (
	"context"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
)

//
//
type Source interface {
	//
	//
	Configure() error

	//
	//
	//
	BeginInstances(context.Context, ...*e2e.Member) error

	//
	HaltVerifychain(context.Context) error

	//
	FetchPlatformData() *e2e.PlatformData
}

type SourceData struct {
	Verifychain            *e2e.Verifychain
	PlatformData e2e.PlatformData
}

//
func (pd SourceData) FetchPlatformData() *e2e.PlatformData {
	return &pd.PlatformData
}
