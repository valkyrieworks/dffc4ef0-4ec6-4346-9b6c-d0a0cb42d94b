package kinds

import (
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault"
	cryptocode "github.com/valkyrieworks/vault/codec"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

//

//
//
var Tm2schema = tm2schema{}

type tm2schema struct{}

func (tm2schema) Heading(heading *Heading) engineproto.Heading {
	return engineproto.Heading{
		Release: heading.Release,
		LedgerUID: heading.LedgerUID,
		Level:  heading.Level,
		Time:    heading.Time,

		FinalLedgerUid: heading.FinalLedgerUID.ToSchema(),

		FinalEndorseDigest: heading.FinalEndorseDigest,
		DataDigest:       heading.DataDigest,

		RatifiersDigest:     heading.RatifiersDigest,
		FollowingRatifiersDigest: heading.FollowingRatifiersDigest,
		AgreementDigest:      heading.AgreementDigest,
		ApplicationDigest:            heading.ApplicationDigest,
		FinalOutcomesDigest:    heading.FinalOutcomesDigest,

		ProofDigest:    heading.ProofDigest,
		RecommenderLocation: heading.RecommenderLocation,
	}
}

func (tm2schema) Ratifier(val *Ratifier) iface.Ratifier {
	return iface.Ratifier{
		Location: val.PublicKey.Location(),
		Energy:   val.PollingEnergy,
	}
}

func (tm2schema) LedgerUID(ledgerUID LedgerUID) engineproto.LedgerUID {
	return engineproto.LedgerUID{
		Digest:          ledgerUID.Digest,
		SegmentAssignHeading: Tm2schema.SegmentAssignHeading(ledgerUID.SegmentAssignHeading),
	}
}

func (tm2schema) SegmentAssignHeading(heading SegmentAssignHeading) engineproto.SegmentAssignHeading {
	return engineproto.SegmentAssignHeading{
		Sum: heading.Sum,
		Digest:  heading.Digest,
	}
}

//
func (tm2schema) RatifierModify(val *Ratifier) iface.RatifierModify {
	pk, err := cryptocode.PublicKeyToSchema(val.PublicKey)
	if err != nil {
		panic(err)
	}
	return iface.RatifierModify{
		PublicKey: pk,
		Energy:  val.PollingEnergy,
	}
}

//
func (tm2schema) RatifierRefreshes(values *RatifierAssign) []iface.RatifierModify {
	ratifiers := make([]iface.RatifierModify, values.Volume())
	for i, val := range values.Ratifiers {
		ratifiers[i] = Tm2schema.RatifierModify(val)
	}
	return ratifiers
}

//
func (tm2schema) NewRatifierModify(publickey vault.PublicKey, energy int64) iface.RatifierModify {
	publickeyIface, err := cryptocode.PublicKeyToSchema(publickey)
	if err != nil {
		panic(err)
	}
	return iface.RatifierModify{
		PublicKey: publickeyIface,
		Energy:  energy,
	}
}

//

//
//
var Schema2tm = schema2tm{}

type schema2tm struct{}

func (schema2tm) RatifierRefreshes(values []iface.RatifierModify) ([]*Ratifier, error) {
	cometValues := make([]*Ratifier, len(values))
	for i, v := range values {
		pub, err := cryptocode.PublicKeyFromSchema(v.PublicKey)
		if err != nil {
			return nil, err
		}
		cometValues[i] = NewRatifier(pub, v.Energy)
	}
	return cometValues, nil
}
