package status

import (
	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
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

const ItemCollectionMilestoneDuration = itemCollectionMilestoneDuration

//
//
func ReviseStatus(
	status Status,
	ledgerUUID kinds.LedgerUUID,
	heading *kinds.Heading,
	reply *iface.ReplyCulminateLedger,
	assessorRevisions []*kinds.Assessor,
) (Status, error) {
	return reviseStatus(status, ledgerUUID, heading, reply, assessorRevisions)
}

//
//
func CertifyAssessorRevisions(ifaceRevisions []iface.AssessorRevise, parameters kinds.AssessorParameters) error {
	return certifyAssessorRevisions(ifaceRevisions, parameters)
}

//
//
func PersistAssessorsDetails(db dbm.DB, altitude, finalAltitudeAltered int64, itemAssign *kinds.AssessorAssign) error {
	statusDepot := datastoreDepot{db, DepotChoices{EjectIfaceReplies: false}}
	cluster := statusDepot.db.FreshCluster()
	err := statusDepot.persistAssessorsDetails(altitude, finalAltitudeAltered, itemAssign, cluster)
	if err != nil {
		return err
	}
	err = cluster.PersistChronize()
	if err != nil {
		return err
	}
	return nil
}

func Integer64towOctets(val int64) []byte {
	return integer64towOctets(val)
}

func Integer64fromOctets(val []byte) int64 {
	return integer64fromOctets(val)
}
