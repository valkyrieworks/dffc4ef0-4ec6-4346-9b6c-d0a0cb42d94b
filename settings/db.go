package settings

import (
	"context"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

//
type FacilitySupplier func(context.Context, *Settings, log.Tracer) (facility.Facility, error)

//
type DatastoreScope struct {
	ID     string
	Settings *Settings
}

//
type DatastoreSupplier func(*DatastoreScope) (dbm.DB, error)

//
//
func FallbackDatastoreSupplier(ctx *DatastoreScope) (dbm.DB, error) {
	datastoreKind := dbm.OriginKind(ctx.Settings.DatastoreRepository)

	return dbm.FreshDatastore(ctx.ID, datastoreKind, ctx.Settings.DatastorePath())
}
