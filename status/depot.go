package status

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	//
	//
	//
	//
	itemAssignMilestoneDuration = 100000
)

//

func computeAssessorsToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

func computeAgreementParametersToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

func computeIfaceRepliesToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

//

var (
	finalIfaceReplyToken    = []byte("REDACTED")
	inactiveStatusChronizeAltitude = []byte("REDACTED")
)

//

//
//
//
//
type Depot interface {
	//
	//
	FetchOriginatingDatastoreEitherInaugurationRecord(string) (Status, error)
	//
	//
	FetchOriginatingDatastoreEitherOriginPaper(*kinds.OriginPaper) (Status, error)
	//
	Fetch() (Status, error)
	//
	FetchAssessors(int64) (*kinds.AssessorAssign, error)
	//
	FetchCulminateLedgerReply(int64) (*iface.ReplyCulminateLedger, error)
	//
	FetchFinalCulminateLedgerReply(int64) (*iface.ReplyCulminateLedger, error)
	//
	FetchAgreementParameters(int64) (kinds.AgreementSettings, error)
	//
	Persist(Status) error
	//
	PersistCulminateLedgerReply(int64, *iface.ReplyCulminateLedger) error
	//
	Onboard(Status) error
	//
	TrimStatuses(int64, int64, int64) error
	//
	AssignInactiveStatusChronizeAltitude(altitude int64) error
	//
	ObtainInactiveStatusChronizeAltitude() (int64, error)
	//
	Shutdown() error
}

//
type datastoreDepot struct {
	db dbm.DB

	DepotChoices
}

type DepotChoices struct {
	//
	//
	//
	//
	EjectIfaceReplies bool
}

var _ Depot = (*datastoreDepot)(nil)

func EqualsBlank(depot datastoreDepot) (bool, error) {
	status, err := depot.Fetch()
	if err != nil {
		return false, err
	}
	return status.EqualsBlank(), nil
}

//
func FreshDepot(db dbm.DB, choices DepotChoices) Depot {
	return datastoreDepot{db, choices}
}

//
//
func (depot datastoreDepot) FetchOriginatingDatastoreEitherInaugurationRecord(inaugurationRecordRoute string) (Status, error) {
	status, err := depot.Fetch()
	if err != nil {
		return Status{}, err
	}
	if status.EqualsBlank() {
		var err error
		status, err = CreateInaugurationStatusOriginatingRecord(inaugurationRecordRoute)
		if err != nil {
			return status, err
		}
	}

	return status, nil
}

//
//
func (depot datastoreDepot) FetchOriginatingDatastoreEitherOriginPaper(inaugurationPaper *kinds.OriginPaper) (Status, error) {
	status, err := depot.Fetch()
	if err != nil {
		return Status{}, err
	}

	if status.EqualsBlank() {
		var err error
		status, err = CreateInaugurationStatus(inaugurationPaper)
		if err != nil {
			return status, err
		}
	}

	return status, nil
}

//
func (depot datastoreDepot) Fetch() (Status, error) {
	return depot.fetchStatus(statusToken)
}

func (depot datastoreDepot) fetchStatus(key []byte) (status Status, err error) {
	buf, err := depot.db.Get(key)
	if err != nil {
		return status, err
	}
	if len(buf) == 0 {
		return status, nil
	}

	sp := new(strongstatus.Status)

	err = proto.Unmarshal(buf, sp)
	if err != nil {
		//
		strongos.Quit(fmt.Sprintf(`REDACTED:
REDACTED`, err))
	}

	sm, err := OriginatingSchema(sp)
	if err != nil {
		return status, err
	}
	return *sm, nil
}

//
//
func (depot datastoreDepot) Persist(status Status) error {
	return depot.persist(status, statusToken)
}

func (depot datastoreDepot) persist(status Status, key []byte) error {
	cluster := depot.db.FreshCluster()
	defer func(cluster dbm.Cluster) {
		err := cluster.Shutdown()
		if err != nil {
			panic(err)
		}
	}(cluster)
	followingAltitude := status.FinalLedgerAltitude + 1
	//
	if followingAltitude == 1 {
		followingAltitude = status.PrimaryAltitude
		//
		//
		if err := depot.persistAssessorsDetails(followingAltitude, followingAltitude, status.Assessors, cluster); err != nil {
			return err
		}
	}
	//
	if err := depot.persistAssessorsDetails(followingAltitude+1, status.FinalAltitudeAssessorsAltered, status.FollowingAssessors, cluster); err != nil {
		return err
	}
	//
	if err := depot.persistAgreementParametersDetails(followingAltitude,
		status.FinalAltitudeAgreementParametersAltered, status.AgreementSettings, cluster); err != nil {
		return err
	}
	if err := cluster.Set(key, status.Octets()); err != nil {
		return err
	}
	if err := cluster.PersistChronize(); err != nil {
		panic(err)
	}
	return nil
}

//
func (depot datastoreDepot) Onboard(status Status) error {
	cluster := depot.db.FreshCluster()
	defer func(cluster dbm.Cluster) {
		err := cluster.Shutdown()
		if err != nil {
			panic(err)
		}
	}(cluster)
	altitude := status.FinalLedgerAltitude + 1
	if altitude == 1 {
		altitude = status.PrimaryAltitude
	}

	if altitude > 1 && !status.FinalAssessors.EqualsVoidEitherBlank() {
		if err := depot.persistAssessorsDetails(altitude-1, altitude-1, status.FinalAssessors, cluster); err != nil {
			return err
		}
	}

	if err := depot.persistAssessorsDetails(altitude, altitude, status.Assessors, cluster); err != nil {
		return err
	}

	if err := depot.persistAssessorsDetails(altitude+1, altitude+1, status.FollowingAssessors, cluster); err != nil {
		return err
	}

	if err := depot.persistAgreementParametersDetails(altitude,
		status.FinalAltitudeAgreementParametersAltered, status.AgreementSettings, cluster); err != nil {
		return err
	}

	if err := cluster.Set(statusToken, status.Octets()); err != nil {
		return err
	}

	if err := cluster.PersistChronize(); err != nil {
		panic(err)
	}

	return cluster.Shutdown()
}

//
//
//
//
//
//
//
//
func (depot datastoreDepot) TrimStatuses(originating int64, to int64, proofLimitAltitude int64) error {
	if originating <= 0 || to <= 0 {
		return fmt.Errorf("REDACTED", originating, to)
	}
	if originating >= to {
		return fmt.Errorf("REDACTED", originating, to)
	}

	itemDetails, err := fetchAssessorsDetails(depot.db, min(to, proofLimitAltitude))
	if err != nil {
		return fmt.Errorf("REDACTED", to, err)
	}
	parametersDetails, err := depot.fetchAgreementParametersDetails(to)
	if err != nil {
		return fmt.Errorf("REDACTED", to, err)
	}

	retainValues := make(map[int64]bool)
	if itemDetails.AssessorAssign == nil {
		retainValues[itemDetails.FinalAltitudeAltered] = true
		retainValues[finalPersistedAltitudeForeach(to, itemDetails.FinalAltitudeAltered)] = true //
	}
	retainParameters := make(map[int64]bool)
	if parametersDetails.AgreementSettings.Equivalent(&commitchema.AgreementSettings{}) {
		retainParameters[parametersDetails.FinalAltitudeAltered] = true
	}

	cluster := depot.db.FreshCluster()
	defer cluster.Shutdown()
	trimmed := uint64(0)

	//
	//
	for h := to - 1; h >= originating; h-- {
		//
		//
		//
		if retainValues[h] {
			v, err := fetchAssessorsDetails(depot.db, h)
			if err != nil || v.AssessorAssign == nil {
				vip, err := depot.FetchAssessors(h)
				if err != nil {
					return err
				}

				pvi, err := vip.TowardSchema()
				if err != nil {
					return err
				}

				v.AssessorAssign = pvi
				v.FinalAltitudeAltered = h

				bz, err := v.Serialize()
				if err != nil {
					return err
				}
				err = cluster.Set(computeAssessorsToken(h), bz)
				if err != nil {
					return err
				}
			}
		} else if h < proofLimitAltitude {
			err = cluster.Erase(computeAssessorsToken(h))
			if err != nil {
				return err
			}
		}
		//
		//

		if retainParameters[h] {
			p, err := depot.fetchAgreementParametersDetails(h)
			if err != nil {
				return err
			}

			if p.AgreementSettings.Equivalent(&commitchema.AgreementSettings{}) {
				parameters, err := depot.FetchAgreementParameters(h)
				if err != nil {
					return err
				}
				p.AgreementSettings = parameters.TowardSchema()

				p.FinalAltitudeAltered = h
				bz, err := p.Serialize()
				if err != nil {
					return err
				}

				err = cluster.Set(computeAgreementParametersToken(h), bz)
				if err != nil {
					return err
				}
			}
		} else {
			err = cluster.Erase(computeAgreementParametersToken(h))
			if err != nil {
				return err
			}
		}

		err = cluster.Erase(computeIfaceRepliesToken(h))
		if err != nil {
			return err
		}
		trimmed++

		//
		if trimmed%1000 == 0 && trimmed > 0 {
			err := cluster.Record()
			if err != nil {
				return err
			}
			cluster.Shutdown()
			cluster = depot.db.FreshCluster()
			defer cluster.Shutdown()
		}
	}

	err = cluster.PersistChronize()
	if err != nil {
		return err
	}

	return nil
}

//

//
//
//
//
func TransferOutcomesDigest(transferOutcomes []*iface.InvokeTransferOutcome) []byte {
	return kinds.FreshOutcomes(transferOutcomes).Digest()
}

//
//
//
func (depot datastoreDepot) FetchCulminateLedgerReply(altitude int64) (*iface.ReplyCulminateLedger, error) {
	if depot.EjectIfaceReplies {
		return nil, FaultCulminateLedgerRepliesNegationStored
	}

	buf, err := depot.db.Get(computeIfaceRepliesToken(altitude))
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		return nil, FaultNegativeIfaceRepliesForeachAltitude{altitude}
	}

	reply := new(iface.ReplyCulminateLedger)
	err = reply.Decode(buf)
	//
	//
	//
	//
	//
	//
	//
	if err != nil || reply.PlatformDigest == nil {
		//
		//
		heritageAnswer := new(strongstatus.HeritageIfaceReplies)
		if err := heritageAnswer.Decode(buf); err != nil {
			//
			//
			return nil, FaultIfaceReplyTaintedEitherBlueprintAlterationForeachAltitude{Altitude: altitude, Err: err}
		}
		//
		//
		//
		return replyCulminateLedgerOriginatingHeritage(heritageAnswer), nil
	}

	//

	return reply, nil
}

//
//
//
//
//
//
func (depot datastoreDepot) FetchFinalCulminateLedgerReply(altitude int64) (*iface.ReplyCulminateLedger, error) {
	bz, err := depot.db.Get(finalIfaceReplyToken)
	if err != nil {
		return nil, err
	}

	if len(bz) == 0 {
		return nil, errors.New("REDACTED")
	}

	details := new(strongstatus.IfaceRepliesDetails)
	err = details.Decode(bz)
	if err != nil {
		strongos.Quit(fmt.Sprintf(`REDACTEDs
REDACTED`, err))
	}

	//
	if altitude != details.ObtainAltitude() {
		return nil, fmt.Errorf("REDACTED", altitude, details.ObtainAltitude())
	}

	//
	//
	//
	if details.ReplyCulminateLedger == nil {
		//
		if details.HeritageIfaceReplies == nil {
			panic("REDACTED")
		}
		return replyCulminateLedgerOriginatingHeritage(details.HeritageIfaceReplies), nil
	}

	return details.ReplyCulminateLedger, nil
}

//
//
//
//
//
//
func (depot datastoreDepot) PersistCulminateLedgerReply(altitude int64, reply *iface.ReplyCulminateLedger) error {
	var dtrans []*iface.InvokeTransferOutcome
	//
	for _, tx := range reply.TransferOutcomes {
		if tx != nil {
			dtrans = append(dtrans, tx)
		}
	}
	reply.TransferOutcomes = dtrans

	//
	//
	if !depot.EjectIfaceReplies {
		bz, err := reply.Serialize()
		if err != nil {
			return err
		}
		if err := depot.db.Set(computeIfaceRepliesToken(altitude), bz); err != nil {
			return err
		}
	}

	//
	//
	reply := &strongstatus.IfaceRepliesDetails{
		ReplyCulminateLedger: reply,
		Altitude:                altitude,
	}
	bz, err := reply.Serialize()
	if err != nil {
		return err
	}

	return depot.db.AssignChronize(finalIfaceReplyToken, bz)
}

//

//
//
func (depot datastoreDepot) FetchAssessors(altitude int64) (*kinds.AssessorAssign, error) {
	itemDetails, err := fetchAssessorsDetails(depot.db, altitude)
	if err != nil {
		return nil, FaultNegativeItemAssignForeachAltitude{altitude}
	}
	if itemDetails.AssessorAssign == nil {
		finalPersistedAltitude := finalPersistedAltitudeForeach(altitude, itemDetails.FinalAltitudeAltered)
		itemDetails2, err := fetchAssessorsDetails(depot.db, finalPersistedAltitude)
		if err != nil || itemDetails2.AssessorAssign == nil {
			return nil,
				fmt.Errorf("REDACTED",
					finalPersistedAltitude,
					altitude,
					err,
				)
		}

		vs, err := kinds.AssessorAssignOriginatingSchema(itemDetails2.AssessorAssign)
		if err != nil {
			return nil, err
		}

		vs.AdvanceNominatorUrgency(strongarithmetic.SecureAdaptInteger32(altitude - finalPersistedAltitude)) //
		vi2, err := vs.TowardSchema()
		if err != nil {
			return nil, err
		}

		itemDetails2.AssessorAssign = vi2
		itemDetails = itemDetails2
	}

	vip, err := kinds.AssessorAssignOriginatingSchema(itemDetails.AssessorAssign)
	if err != nil {
		return nil, err
	}

	return vip, nil
}

func finalPersistedAltitudeForeach(altitude, finalAltitudeAltered int64) int64 {
	milestoneAltitude := altitude - altitude%itemAssignMilestoneDuration
	return strongarithmetic.MaximumInt64n(milestoneAltitude, finalAltitudeAltered)
}

//
func fetchAssessorsDetails(db dbm.DB, altitude int64) (*strongstatus.AssessorsDetails, error) {
	buf, err := db.Get(computeAssessorsToken(altitude))
	if err != nil {
		return nil, err
	}

	if len(buf) == 0 {
		return nil, errors.New("REDACTED")
	}

	v := new(strongstatus.AssessorsDetails)
	err = v.Decode(buf)
	if err != nil {
		//
		strongos.Quit(fmt.Sprintf(`REDACTED:
REDACTED`, err))
	}
	//

	return v, nil
}

//
//
//
//
//
func (depot datastoreDepot) persistAssessorsDetails(altitude, finalAltitudeAltered int64, itemAssign *kinds.AssessorAssign, cluster dbm.Cluster) error {
	if finalAltitudeAltered > altitude {
		return errors.New("REDACTED")
	}
	itemDetails := &strongstatus.AssessorsDetails{
		FinalAltitudeAltered: finalAltitudeAltered,
	}
	//
	//
	if altitude == finalAltitudeAltered || altitude%itemAssignMilestoneDuration == 0 {
		pv, err := itemAssign.TowardSchema()
		if err != nil {
			return err
		}
		itemDetails.AssessorAssign = pv
	}

	bz, err := itemDetails.Serialize()
	if err != nil {
		return err
	}

	err = cluster.Set(computeAssessorsToken(altitude), bz)
	if err != nil {
		return err
	}

	return nil
}

//

//

//
func (depot datastoreDepot) FetchAgreementParameters(altitude int64) (kinds.AgreementSettings, error) {
	var (
		blank   = kinds.AgreementSettings{}
		voidschema = commitchema.AgreementSettings{}
	)
	parametersDetails, err := depot.fetchAgreementParametersDetails(altitude)
	if err != nil {
		return blank, fmt.Errorf("REDACTED", altitude, err)
	}

	if parametersDetails.AgreementSettings.Equivalent(&voidschema) {
		parametersDetails2, err := depot.fetchAgreementParametersDetails(parametersDetails.FinalAltitudeAltered)
		if err != nil {
			return blank, fmt.Errorf(
				"REDACTED",
				parametersDetails.FinalAltitudeAltered,
				altitude,
				err,
			)
		}

		parametersDetails = parametersDetails2
	}

	return kinds.AgreementParametersOriginatingSchema(parametersDetails.AgreementSettings), nil
}

func (depot datastoreDepot) fetchAgreementParametersDetails(altitude int64) (*strongstatus.AgreementParametersDetails, error) {
	buf, err := depot.db.Get(computeAgreementParametersToken(altitude))
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		return nil, errors.New("REDACTED")
	}

	parametersDetails := new(strongstatus.AgreementParametersDetails)
	if err = parametersDetails.Decode(buf); err != nil {
		//
		strongos.Quit(fmt.Sprintf(`REDACTED:
REDACTED`, err))
	}
	//

	return parametersDetails, nil
}

//
//
//
//
func (depot datastoreDepot) persistAgreementParametersDetails(followingAltitude, alterationAltitude int64, parameters kinds.AgreementSettings, cluster dbm.Cluster) error {
	parametersDetails := &strongstatus.AgreementParametersDetails{
		FinalAltitudeAltered: alterationAltitude,
	}

	if alterationAltitude == followingAltitude {
		parametersDetails.AgreementSettings = parameters.TowardSchema()
	}
	bz, err := parametersDetails.Serialize()
	if err != nil {
		return err
	}

	err = cluster.Set(computeAgreementParametersToken(followingAltitude), bz)
	if err != nil {
		return err
	}

	return nil
}

func (depot datastoreDepot) AssignInactiveStatusChronizeAltitude(altitude int64) error {
	err := depot.db.AssignChronize(inactiveStatusChronizeAltitude, integer64towOctets(altitude))
	if err != nil {
		return err
	}
	return nil
}

//
func (depot datastoreDepot) ObtainInactiveStatusChronizeAltitude() (int64, error) {
	buf, err := depot.db.Get(inactiveStatusChronizeAltitude)
	if err != nil {
		return 0, err
	}

	if len(buf) == 0 {
		return 0, errors.New("REDACTED")
	}

	altitude := integer64fromOctets(buf)
	if altitude < 0 {
		return 0, errors.New("REDACTED")
	}
	return altitude, nil
}

func (depot datastoreDepot) Shutdown() error {
	return depot.db.Shutdown()
}

//
//
func replyCulminateLedgerOriginatingHeritage(heritageAnswer *strongstatus.HeritageIfaceReplies) *iface.ReplyCulminateLedger {
	var reply iface.ReplyCulminateLedger
	incidents := make([]iface.Incident, 0)

	if heritageAnswer.DispatchTrans != nil {
		reply.TransferOutcomes = heritageAnswer.DispatchTrans
	}

	//
	if heritageAnswer.InitiateLedger != nil {
		if heritageAnswer.InitiateLedger.Incidents != nil {
			//
			for idx := range heritageAnswer.InitiateLedger.Incidents {
				heritageAnswer.InitiateLedger.Incidents[idx].Properties = append(heritageAnswer.InitiateLedger.Incidents[idx].Properties, iface.IncidentProperty{
					Key:   "REDACTED",
					Datum: "REDACTED",
					Ordinal: false,
				})
			}
			incidents = append(incidents, heritageAnswer.InitiateLedger.Incidents...)
		}
	}
	if heritageAnswer.TerminateLedger != nil {
		if heritageAnswer.TerminateLedger.AssessorRevisions != nil {
			reply.AssessorRevisions = heritageAnswer.TerminateLedger.AssessorRevisions
		}
		if heritageAnswer.TerminateLedger.AgreementArgumentRevisions != nil {
			reply.AgreementArgumentRevisions = heritageAnswer.TerminateLedger.AgreementArgumentRevisions
		}
		if heritageAnswer.TerminateLedger.Incidents != nil {
			//
			for idx := range heritageAnswer.TerminateLedger.Incidents {
				heritageAnswer.TerminateLedger.Incidents[idx].Properties = append(heritageAnswer.TerminateLedger.Incidents[idx].Properties, iface.IncidentProperty{
					Key:   "REDACTED",
					Datum: "REDACTED",
					Ordinal: false,
				})
			}
			incidents = append(incidents, heritageAnswer.TerminateLedger.Incidents...)
		}
	}

	reply.Incidents = incidents

	//
	//
	return &reply
}

func integer64fromOctets(bz []byte) int64 {
	v, _ := binary.Varint(bz)
	return v
}

func integer64towOctets(i int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, i)
	return buf[:n]
}
