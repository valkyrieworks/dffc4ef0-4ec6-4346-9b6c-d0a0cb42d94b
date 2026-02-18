package status

import (
	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
//
//
//
//
//
//

const ValueCollectionMilestoneCadence = valueCollectionMilestoneCadence

//
//
func ModifyStatus(
	status Status,
	ledgerUID kinds.LedgerUID,
	heading *kinds.Heading,
	reply *iface.ReplyCompleteLedger,
	ratifierRefreshes []*kinds.Ratifier,
) (Status, error) {
	return modifyStatus(status, ledgerUID, heading, reply, ratifierRefreshes)
}

//
//
func CertifyRatifierRefreshes(ifaceRefreshes []iface.RatifierModify, options kinds.RatifierOptions) error {
	return certifyRatifierRefreshes(ifaceRefreshes, options)
}

//
//
func PersistRatifiersDetails(db dbm.DB, level, finalLevelModified int64, valueCollection *kinds.RatifierAssign) error {
	statusDepot := storeDepot{db, DepotSettings{DropIfaceReplies: false}}
	group := statusDepot.db.NewGroup()
	err := statusDepot.persistRatifiersDetails(level, finalLevelModified, valueCollection, group)
	if err != nil {
		return err
	}
	err = group.RecordAlign()
	if err != nil {
		return err
	}
	return nil
}

func Int64toOctets(val int64) []byte {
	return int64toOctets(val)
}

func Int64fromOctets(val []byte) int64 {
	return int64fromOctets(val)
}
