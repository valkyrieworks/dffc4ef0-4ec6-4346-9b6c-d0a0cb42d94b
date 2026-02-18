package status

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	cometmath "github.com/valkyrieworks/utils/math"
	cometos "github.com/valkyrieworks/utils/os"
	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	//
	//
	//
	valueCollectionMilestoneCadence = 100000
)

//

func computeRatifiersKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

func computeAgreementOptionsKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

func computeIfaceRepliesKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

//

var (
	finalIfaceReplyKey    = []byte("REDACTED")
	inactiveStatusAlignLevel = []byte("REDACTED")
)

//

//
//
//
//
type Depot interface {
	//
	//
	ImportFromStoreOrOriginEntry(string) (Status, error)
	//
	//
	ImportFromStoreOrOriginPaper(*kinds.OriginPaper) (Status, error)
	//
	Import() (Status, error)
	//
	ImportRatifiers(int64) (*kinds.RatifierAssign, error)
	//
	ImportCompleteLedgerReply(int64) (*iface.ReplyCompleteLedger, error)
	//
	ImportFinalCompleteLedgerReply(int64) (*iface.ReplyCompleteLedger, error)
	//
	ImportAgreementOptions(int64) (kinds.AgreementOptions, error)
	//
	Persist(Status) error
	//
	PersistCompleteLedgerReply(int64, *iface.ReplyCompleteLedger) error
	//
	Onboard(Status) error
	//
	TrimConditions(int64, int64, int64) error
	//
	CollectionInactiveStatusAlignLevel(level int64) error
	//
	FetchInactiveStatusAlignLevel() (int64, error)
	//
	End() error
}

//
type storeDepot struct {
	db dbm.DB

	DepotSettings
}

type DepotSettings struct {
	//
	//
	//
	//
	DropIfaceReplies bool
}

var _ Depot = (*storeDepot)(nil)

func IsEmpty(depot storeDepot) (bool, error) {
	status, err := depot.Import()
	if err != nil {
		return false, err
	}
	return status.IsEmpty(), nil
}

//
func NewDepot(db dbm.DB, options DepotSettings) Depot {
	return storeDepot{db, options}
}

//
//
func (depot storeDepot) ImportFromStoreOrOriginEntry(originEntryRoute string) (Status, error) {
	status, err := depot.Import()
	if err != nil {
		return Status{}, err
	}
	if status.IsEmpty() {
		var err error
		status, err = CreateOriginStatusFromEntry(originEntryRoute)
		if err != nil {
			return status, err
		}
	}

	return status, nil
}

//
//
func (depot storeDepot) ImportFromStoreOrOriginPaper(originPaper *kinds.OriginPaper) (Status, error) {
	status, err := depot.Import()
	if err != nil {
		return Status{}, err
	}

	if status.IsEmpty() {
		var err error
		status, err = CreateOriginStatus(originPaper)
		if err != nil {
			return status, err
		}
	}

	return status, nil
}

//
func (depot storeDepot) Import() (Status, error) {
	return depot.importStatus(statusKey)
}

func (depot storeDepot) importStatus(key []byte) (status Status, err error) {
	buf, err := depot.db.Get(key)
	if err != nil {
		return status, err
	}
	if len(buf) == 0 {
		return status, nil
	}

	sp := new(cometstatus.Status)

	err = proto.Unmarshal(buf, sp)
	if err != nil {
		//
		cometos.Quit(fmt.Sprintf(`REDACTED:
REDACTED`, err))
	}

	sm, err := FromSchema(sp)
	if err != nil {
		return status, err
	}
	return *sm, nil
}

//
//
func (depot storeDepot) Persist(status Status) error {
	return depot.persist(status, statusKey)
}

func (depot storeDepot) persist(status Status, key []byte) error {
	group := depot.db.NewGroup()
	defer func(group dbm.Group) {
		err := group.End()
		if err != nil {
			panic(err)
		}
	}(group)
	followingLevel := status.FinalLedgerLevel + 1
	//
	if followingLevel == 1 {
		followingLevel = status.PrimaryLevel
		//
		//
		if err := depot.persistRatifiersDetails(followingLevel, followingLevel, status.Ratifiers, group); err != nil {
			return err
		}
	}
	//
	if err := depot.persistRatifiersDetails(followingLevel+1, status.FinalLevelRatifiersModified, status.FollowingRatifiers, group); err != nil {
		return err
	}
	//
	if err := depot.persistAgreementOptionsDetails(followingLevel,
		status.FinalLevelAgreementOptionsModified, status.AgreementOptions, group); err != nil {
		return err
	}
	if err := group.Set(key, status.Octets()); err != nil {
		return err
	}
	if err := group.RecordAlign(); err != nil {
		panic(err)
	}
	return nil
}

//
func (depot storeDepot) Onboard(status Status) error {
	group := depot.db.NewGroup()
	defer func(group dbm.Group) {
		err := group.End()
		if err != nil {
			panic(err)
		}
	}(group)
	level := status.FinalLedgerLevel + 1
	if level == 1 {
		level = status.PrimaryLevel
	}

	if level > 1 && !status.FinalRatifiers.IsNullOrEmpty() {
		if err := depot.persistRatifiersDetails(level-1, level-1, status.FinalRatifiers, group); err != nil {
			return err
		}
	}

	if err := depot.persistRatifiersDetails(level, level, status.Ratifiers, group); err != nil {
		return err
	}

	if err := depot.persistRatifiersDetails(level+1, level+1, status.FollowingRatifiers, group); err != nil {
		return err
	}

	if err := depot.persistAgreementOptionsDetails(level,
		status.FinalLevelAgreementOptionsModified, status.AgreementOptions, group); err != nil {
		return err
	}

	if err := group.Set(statusKey, status.Octets()); err != nil {
		return err
	}

	if err := group.RecordAlign(); err != nil {
		panic(err)
	}

	return group.End()
}

//
//
//
//
//
//
//
//
func (depot storeDepot) TrimConditions(from int64, to int64, proofLimitLevel int64) error {
	if from <= 0 || to <= 0 {
		return fmt.Errorf("REDACTED", from, to)
	}
	if from >= to {
		return fmt.Errorf("REDACTED", from, to)
	}

	valueDetails, err := importRatifiersDetails(depot.db, min(to, proofLimitLevel))
	if err != nil {
		return fmt.Errorf("REDACTED", to, err)
	}
	optionsDetails, err := depot.importAgreementOptionsDetails(to)
	if err != nil {
		return fmt.Errorf("REDACTED", to, err)
	}

	retainValues := make(map[int64]bool)
	if valueDetails.RatifierAssign == nil {
		retainValues[valueDetails.FinalLevelModified] = true
		retainValues[finalArchivedLevelFor(to, valueDetails.FinalLevelModified)] = true //
	}
	retainOptions := make(map[int64]bool)
	if optionsDetails.AgreementOptions.Equivalent(&engineproto.AgreementOptions{}) {
		retainOptions[optionsDetails.FinalLevelModified] = true
	}

	group := depot.db.NewGroup()
	defer group.End()
	trimmed := uint64(0)

	//
	//
	for h := to - 1; h >= from; h-- {
		//
		//
		//
		if retainValues[h] {
			v, err := importRatifiersDetails(depot.db, h)
			if err != nil || v.RatifierAssign == nil {
				vip, err := depot.ImportRatifiers(h)
				if err != nil {
					return err
				}

				pvi, err := vip.ToSchema()
				if err != nil {
					return err
				}

				v.RatifierAssign = pvi
				v.FinalLevelModified = h

				bz, err := v.Serialize()
				if err != nil {
					return err
				}
				err = group.Set(computeRatifiersKey(h), bz)
				if err != nil {
					return err
				}
			}
		} else if h < proofLimitLevel {
			err = group.Erase(computeRatifiersKey(h))
			if err != nil {
				return err
			}
		}
		//
		//

		if retainOptions[h] {
			p, err := depot.importAgreementOptionsDetails(h)
			if err != nil {
				return err
			}

			if p.AgreementOptions.Equivalent(&engineproto.AgreementOptions{}) {
				options, err := depot.ImportAgreementOptions(h)
				if err != nil {
					return err
				}
				p.AgreementOptions = options.ToSchema()

				p.FinalLevelModified = h
				bz, err := p.Serialize()
				if err != nil {
					return err
				}

				err = group.Set(computeAgreementOptionsKey(h), bz)
				if err != nil {
					return err
				}
			}
		} else {
			err = group.Erase(computeAgreementOptionsKey(h))
			if err != nil {
				return err
			}
		}

		err = group.Erase(computeIfaceRepliesKey(h))
		if err != nil {
			return err
		}
		trimmed++

		//
		if trimmed%1000 == 0 && trimmed > 0 {
			err := group.Record()
			if err != nil {
				return err
			}
			group.End()
			group = depot.db.NewGroup()
			defer group.End()
		}
	}

	err = group.RecordAlign()
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
	return kinds.NewOutcomes(transferOutcomes).Digest()
}

//
//
//
func (depot storeDepot) ImportCompleteLedgerReply(level int64) (*iface.ReplyCompleteLedger, error) {
	if depot.DropIfaceReplies {
		return nil, ErrCompleteLedgerRepliesNotSustained
	}

	buf, err := depot.db.Get(computeIfaceRepliesKey(level))
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		return nil, ErrNoIfaceRepliesForLevel{level}
	}

	reply := new(iface.ReplyCompleteLedger)
	err = reply.Unserialize(buf)
	//
	//
	//
	//
	//
	//
	//
	if err != nil || reply.ApplicationDigest == nil {
		//
		//
		pastReply := new(cometstatus.PastIfaceReplies)
		if err := pastReply.Unserialize(buf); err != nil {
			//
			//
			return nil, ErrIfaceReplyTaintedOrBlueprintAlterForLevel{Level: level, Err: err}
		}
		//
		//
		//
		return replyCompleteLedgerFromPast(pastReply), nil
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
func (depot storeDepot) ImportFinalCompleteLedgerReply(level int64) (*iface.ReplyCompleteLedger, error) {
	bz, err := depot.db.Get(finalIfaceReplyKey)
	if err != nil {
		return nil, err
	}

	if len(bz) == 0 {
		return nil, errors.New("REDACTED")
	}

	details := new(cometstatus.IfaceRepliesDetails)
	err = details.Unserialize(bz)
	if err != nil {
		cometos.Quit(fmt.Sprintf(`REDACTEDs
REDACTED`, err))
	}

	//
	if level != details.FetchLevel() {
		return nil, fmt.Errorf("REDACTED", level, details.FetchLevel())
	}

	//
	//
	//
	if details.ReplyCompleteLedger == nil {
		//
		if details.PastIfaceReplies == nil {
			panic("REDACTED")
		}
		return replyCompleteLedgerFromPast(details.PastIfaceReplies), nil
	}

	return details.ReplyCompleteLedger, nil
}

//
//
//
//
//
//
func (depot storeDepot) PersistCompleteLedgerReply(level int64, reply *iface.ReplyCompleteLedger) error {
	var dtrans []*iface.InvokeTransferOutcome
	//
	for _, tx := range reply.TransOutcomes {
		if tx != nil {
			dtrans = append(dtrans, tx)
		}
	}
	reply.TransOutcomes = dtrans

	//
	//
	if !depot.DropIfaceReplies {
		bz, err := reply.Serialize()
		if err != nil {
			return err
		}
		if err := depot.db.Set(computeIfaceRepliesKey(level), bz); err != nil {
			return err
		}
	}

	//
	//
	reply := &cometstatus.IfaceRepliesDetails{
		ReplyCompleteLedger: reply,
		Level:                level,
	}
	bz, err := reply.Serialize()
	if err != nil {
		return err
	}

	return depot.db.CollectionAlign(finalIfaceReplyKey, bz)
}

//

//
//
func (depot storeDepot) ImportRatifiers(level int64) (*kinds.RatifierAssign, error) {
	valueDetails, err := importRatifiersDetails(depot.db, level)
	if err != nil {
		return nil, ErrNoValueCollectionForLevel{level}
	}
	if valueDetails.RatifierAssign == nil {
		finalArchivedLevel := finalArchivedLevelFor(level, valueDetails.FinalLevelModified)
		valueDetail2, err := importRatifiersDetails(depot.db, finalArchivedLevel)
		if err != nil || valueDetail2.RatifierAssign == nil {
			return nil,
				fmt.Errorf("REDACTED",
					finalArchivedLevel,
					level,
					err,
				)
		}

		vs, err := kinds.RatifierCollectionFromSchema(valueDetail2.RatifierAssign)
		if err != nil {
			return nil, err
		}

		vs.AugmentRecommenderUrgency(cometmath.SecureTransformInt32(level - finalArchivedLevel)) //
		vi2, err := vs.ToSchema()
		if err != nil {
			return nil, err
		}

		valueDetail2.RatifierAssign = vi2
		valueDetails = valueDetail2
	}

	vip, err := kinds.RatifierCollectionFromSchema(valueDetails.RatifierAssign)
	if err != nil {
		return nil, err
	}

	return vip, nil
}

func finalArchivedLevelFor(level, finalLevelModified int64) int64 {
	milestoneLevel := level - level%valueCollectionMilestoneCadence
	return cometmath.MaximumInt64(milestoneLevel, finalLevelModified)
}

//
func importRatifiersDetails(db dbm.DB, level int64) (*cometstatus.RatifiersDetails, error) {
	buf, err := db.Get(computeRatifiersKey(level))
	if err != nil {
		return nil, err
	}

	if len(buf) == 0 {
		return nil, errors.New("REDACTED")
	}

	v := new(cometstatus.RatifiersDetails)
	err = v.Unserialize(buf)
	if err != nil {
		//
		cometos.Quit(fmt.Sprintf(`REDACTED:
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
func (depot storeDepot) persistRatifiersDetails(level, finalLevelModified int64, valueCollection *kinds.RatifierAssign, group dbm.Group) error {
	if finalLevelModified > level {
		return errors.New("REDACTED")
	}
	valueDetails := &cometstatus.RatifiersDetails{
		FinalLevelModified: finalLevelModified,
	}
	//
	//
	if level == finalLevelModified || level%valueCollectionMilestoneCadence == 0 {
		pv, err := valueCollection.ToSchema()
		if err != nil {
			return err
		}
		valueDetails.RatifierAssign = pv
	}

	bz, err := valueDetails.Serialize()
	if err != nil {
		return err
	}

	err = group.Set(computeRatifiersKey(level), bz)
	if err != nil {
		return err
	}

	return nil
}

//

//

//
func (depot storeDepot) ImportAgreementOptions(level int64) (kinds.AgreementOptions, error) {
	var (
		empty   = kinds.AgreementOptions{}
		emptypb = engineproto.AgreementOptions{}
	)
	optionsDetails, err := depot.importAgreementOptionsDetails(level)
	if err != nil {
		return empty, fmt.Errorf("REDACTED", level, err)
	}

	if optionsDetails.AgreementOptions.Equivalent(&emptypb) {
		optionsDetail2, err := depot.importAgreementOptionsDetails(optionsDetails.FinalLevelModified)
		if err != nil {
			return empty, fmt.Errorf(
				"REDACTED",
				optionsDetails.FinalLevelModified,
				level,
				err,
			)
		}

		optionsDetails = optionsDetail2
	}

	return kinds.AgreementOptionsFromSchema(optionsDetails.AgreementOptions), nil
}

func (depot storeDepot) importAgreementOptionsDetails(level int64) (*cometstatus.AgreementOptionsDetails, error) {
	buf, err := depot.db.Get(computeAgreementOptionsKey(level))
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		return nil, errors.New("REDACTED")
	}

	optionsDetails := new(cometstatus.AgreementOptionsDetails)
	if err = optionsDetails.Unserialize(buf); err != nil {
		//
		cometos.Quit(fmt.Sprintf(`REDACTED:
REDACTED`, err))
	}
	//

	return optionsDetails, nil
}

//
//
//
//
func (depot storeDepot) persistAgreementOptionsDetails(followingLevel, alterLevel int64, options kinds.AgreementOptions, group dbm.Group) error {
	optionsDetails := &cometstatus.AgreementOptionsDetails{
		FinalLevelModified: alterLevel,
	}

	if alterLevel == followingLevel {
		optionsDetails.AgreementOptions = options.ToSchema()
	}
	bz, err := optionsDetails.Serialize()
	if err != nil {
		return err
	}

	err = group.Set(computeAgreementOptionsKey(followingLevel), bz)
	if err != nil {
		return err
	}

	return nil
}

func (depot storeDepot) CollectionInactiveStatusAlignLevel(level int64) error {
	err := depot.db.CollectionAlign(inactiveStatusAlignLevel, int64toOctets(level))
	if err != nil {
		return err
	}
	return nil
}

//
func (depot storeDepot) FetchInactiveStatusAlignLevel() (int64, error) {
	buf, err := depot.db.Get(inactiveStatusAlignLevel)
	if err != nil {
		return 0, err
	}

	if len(buf) == 0 {
		return 0, errors.New("REDACTED")
	}

	level := int64fromOctets(buf)
	if level < 0 {
		return 0, errors.New("REDACTED")
	}
	return level, nil
}

func (depot storeDepot) End() error {
	return depot.db.End()
}

//
//
func replyCompleteLedgerFromPast(pastReply *cometstatus.PastIfaceReplies) *iface.ReplyCompleteLedger {
	var reply iface.ReplyCompleteLedger
	events := make([]iface.Event, 0)

	if pastReply.DispatchTrans != nil {
		reply.TransOutcomes = pastReply.DispatchTrans
	}

	//
	if pastReply.InitiateLedger != nil {
		if pastReply.InitiateLedger.Events != nil {
			//
			for idx := range pastReply.InitiateLedger.Events {
				pastReply.InitiateLedger.Events[idx].Properties = append(pastReply.InitiateLedger.Events[idx].Properties, iface.EventProperty{
					Key:   "REDACTED",
					Item: "REDACTED",
					Ordinal: false,
				})
			}
			events = append(events, pastReply.InitiateLedger.Events...)
		}
	}
	if pastReply.TerminateLedger != nil {
		if pastReply.TerminateLedger.RatifierRefreshes != nil {
			reply.RatifierRefreshes = pastReply.TerminateLedger.RatifierRefreshes
		}
		if pastReply.TerminateLedger.AgreementArgumentRefreshes != nil {
			reply.AgreementArgumentRefreshes = pastReply.TerminateLedger.AgreementArgumentRefreshes
		}
		if pastReply.TerminateLedger.Events != nil {
			//
			for idx := range pastReply.TerminateLedger.Events {
				pastReply.TerminateLedger.Events[idx].Properties = append(pastReply.TerminateLedger.Events[idx].Properties, iface.EventProperty{
					Key:   "REDACTED",
					Item: "REDACTED",
					Ordinal: false,
				})
			}
			events = append(events, pastReply.TerminateLedger.Events...)
		}
	}

	reply.Events = events

	//
	//
	return &reply
}

func int64fromOctets(bz []byte) int64 {
	v, _ := binary.Varint(bz)
	return v
}

func int64toOctets(i int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, i)
	return buf[:n]
}
