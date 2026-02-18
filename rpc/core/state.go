package core

import (
	"time"

	cometbytes "github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/p2p"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
//
func (env *Context) Status(*rpctypes.Context) (*ctypes.OutcomeState, error) {
	var (
		oldestLedgerLevel   int64
		oldestLedgerDigest     cometbytes.HexOctets
		oldestApplicationDigest       cometbytes.HexOctets
		oldestLedgerTimeNano int64
	)

	if oldestLedgerMeta := env.LedgerDepot.ImportRootMeta(); oldestLedgerMeta != nil {
		oldestLedgerLevel = oldestLedgerMeta.Heading.Level
		oldestApplicationDigest = oldestLedgerMeta.Heading.ApplicationDigest
		oldestLedgerDigest = oldestLedgerMeta.LedgerUID.Digest
		oldestLedgerTimeNano = oldestLedgerMeta.Heading.Time.UnixNano()
	}

	var (
		newestLedgerDigest     cometbytes.HexOctets
		newestApplicationDigest       cometbytes.HexOctets
		newestLedgerTimeNano int64

		newestLevel = env.LedgerDepot.Level()
	)

	if newestLevel != 0 {
		if newestLedgerMeta := env.LedgerDepot.ImportLedgerMeta(newestLevel); newestLedgerMeta != nil {
			newestLedgerDigest = newestLedgerMeta.LedgerUID.Digest
			newestApplicationDigest = newestLedgerMeta.Heading.ApplicationDigest
			newestLedgerTimeNano = newestLedgerMeta.Heading.Time.UnixNano()
		}
	}

	//
	//
	var pollingEnergy int64
	if val := env.ratifierAtLevel(env.newestUnsubmittedLevel()); val != nil {
		pollingEnergy = val.PollingEnergy
	}

	trappingUp := env.AgreementHandler.WaitAlign() && !env.IsReplicaStyle

	return &ctypes.OutcomeState{
		MemberDetails: env.P2PCarrier.MemberDetails().(p2p.StandardMemberDetails),
		AlignDetails: ctypes.AlignDetails{
			NewestLedgerDigest:     newestLedgerDigest,
			NewestApplicationDigest:       newestApplicationDigest,
			NewestLedgerLevel:   newestLevel,
			NewestLedgerTime:     time.Unix(0, newestLedgerTimeNano),
			OldestLedgerDigest:   oldestLedgerDigest,
			OldestApplicationDigest:     oldestApplicationDigest,
			OldestLedgerLevel: oldestLedgerLevel,
			OldestLedgerTime:   time.Unix(0, oldestLedgerTimeNano),
			TrappingUp:          trappingUp,
		},
		RatifierDetails: ctypes.RatifierDetails{
			Location:     env.PublicKey.Location(),
			PublicKey:      env.PublicKey,
			PollingEnergy: pollingEnergy,
		},
	}, nil
}

func (env *Context) ratifierAtLevel(h int64) *kinds.Ratifier {
	valuesWithH, err := env.StatusDepot.ImportRatifiers(h)
	if err != nil {
		return nil
	}
	privateValueLocation := env.PublicKey.Location()
	_, val := valuesWithH.FetchByLocation(privateValueLocation)
	return val
}
