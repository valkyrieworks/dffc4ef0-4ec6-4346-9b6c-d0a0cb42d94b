package base

import (
	"time"

	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
func (env *Context) Condition(*remoteifacetypes.Env) (*ktypes.OutcomeCondition, error) {
	var (
		initialLedgerAltitude   int64
		initialLedgerDigest     tendermintoctets.HexadecimalOctets
		initialApplicationDigest       tendermintoctets.HexadecimalOctets
		initialLedgerMomentAtomic int64
	)

	if initialLedgerSummary := env.LedgerDepot.FetchFoundationSummary(); initialLedgerSummary != nil {
		initialLedgerAltitude = initialLedgerSummary.Heading.Altitude
		initialApplicationDigest = initialLedgerSummary.Heading.PlatformDigest
		initialLedgerDigest = initialLedgerSummary.LedgerUUID.Digest
		initialLedgerMomentAtomic = initialLedgerSummary.Heading.Moment.UnixNano()
	}

	var (
		newestLedgerDigest     tendermintoctets.HexadecimalOctets
		newestApplicationDigest       tendermintoctets.HexadecimalOctets
		newestLedgerMomentAtomic int64

		newestAltitude = env.LedgerDepot.Altitude()
	)

	if newestAltitude != 0 {
		if newestLedgerSummary := env.LedgerDepot.FetchLedgerSummary(newestAltitude); newestLedgerSummary != nil {
			newestLedgerDigest = newestLedgerSummary.LedgerUUID.Digest
			newestApplicationDigest = newestLedgerSummary.Heading.PlatformDigest
			newestLedgerMomentAtomic = newestLedgerSummary.Heading.Moment.UnixNano()
		}
	}

	//
	//
	var ballotingPotency int64
	if val := env.assessorLocatedAltitude(env.newestPendingAltitude()); val != nil {
		ballotingPotency = val.BallotingPotency
	}

	obtainingAscend := env.AgreementHandler.AwaitChronize()

	return &ktypes.OutcomeCondition{
		PeerDetails: env.Peer2peerCarrier.PeerDetails().(p2p.FallbackPeerDetails),
		ChronizeDetails: ktypes.ChronizeDetails{
			NewestLedgerDigest:     newestLedgerDigest,
			NewestApplicationDigest:       newestApplicationDigest,
			NewestLedgerAltitude:   newestAltitude,
			NewestLedgerMoment:     time.Unix(0, newestLedgerMomentAtomic),
			InitialLedgerDigest:   initialLedgerDigest,
			InitialApplicationDigest:     initialApplicationDigest,
			InitialLedgerAltitude: initialLedgerAltitude,
			InitialLedgerMoment:   time.Unix(0, initialLedgerMomentAtomic),
			ObtainingAscend:          obtainingAscend,
		},
		AssessorDetails: ktypes.AssessorDetails{
			Location:     env.PublicToken.Location(),
			PublicToken:      env.PublicToken,
			BallotingPotency: ballotingPotency,
		},
	}, nil
}

func (env *Context) assessorLocatedAltitude(h int64) *kinds.Assessor {
	valuesUsingHASH, err := env.StatusDepot.FetchAssessors(h)
	if err != nil {
		return nil
	}
	privateItemLocator := env.PublicToken.Location()
	_, val := valuesUsingHASH.ObtainViaLocation(privateItemLocator)
	return val
}
