package settings

import (
	"context"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
)

//
type DaemonSource func(context.Context, *Settings, log.Tracer) (daemon.Daemon, error)

//
type StoreContext struct {
	ID     string
	Settings *Settings
}

//
type StoreSource func(*StoreContext) (dbm.DB, error)

//
//
func StandardStoreSource(ctx *StoreContext) (dbm.DB, error) {
	storeKind := dbm.OriginKind(ctx.Settings.StoreOrigin)

	return dbm.NewStore(ctx.ID, storeKind, ctx.Settings.StoreFolder())
}
