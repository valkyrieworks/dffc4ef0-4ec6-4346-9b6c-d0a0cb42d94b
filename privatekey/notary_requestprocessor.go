package privatekey

import (
	"fmt"

	"github.com/valkyrieworks/vault"
	cryptocode "github.com/valkyrieworks/vault/codec"
	cryptography "github.com/valkyrieworks/schema/consensuscore/vault"
	privatekeyproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

func StandardVerificationQueryManager(
	privateValue kinds.PrivateRatifier,
	req privatekeyproto.Signal,
	ledgerUID string,
) (privatekeyproto.Signal, error) {
	var (
		res privatekeyproto.Signal
		err error
	)

	switch r := req.Sum.(type) {
	case *privatekeyproto.Signal_Publickeyquery:
		if r.PublicKeyQuery.FetchSeriesUid() != ledgerUID {
			res = shouldEncloseMessage(&privatekeyproto.PublicKeyAnswer{
				PublicKey: cryptography.PublicKey{}, Fault: &privatekeyproto.DistantNotaryFault{
					Code: 0, Summary: "REDACTED",
				},
			})
			return res, fmt.Errorf("REDACTED", r.PublicKeyQuery.FetchSeriesUid(), ledgerUID)
		}

		var publicKey vault.PublicKey
		publicKey, err = privateValue.FetchPublicKey()
		if err != nil {
			return res, err
		}
		pk, err := cryptocode.PublicKeyToSchema(publicKey)
		if err != nil {
			res = shouldEncloseMessage(&privatekeyproto.PublicKeyAnswer{
				PublicKey: cryptography.PublicKey{}, Fault: &privatekeyproto.DistantNotaryFault{Code: 0, Summary: err.Error()},
			})
		} else {
			res = shouldEncloseMessage(&privatekeyproto.PublicKeyAnswer{PublicKey: pk, Fault: nil})
		}

	case *privatekeyproto.Signal_Attestballotquery:
		if r.AttestBallotQuery.SeriesUid != ledgerUID {
			res = shouldEncloseMessage(&privatekeyproto.AttestedBallotAnswer{
				Ballot: engineproto.Ballot{}, Fault: &privatekeyproto.DistantNotaryFault{
					Code: 0, Summary: "REDACTED",
				},
			})
			return res, fmt.Errorf("REDACTED", r.AttestBallotQuery.FetchSeriesUid(), ledgerUID)
		}

		ballot := r.AttestBallotQuery.Ballot

		err = privateValue.AttestBallot(ledgerUID, ballot)
		if err != nil {
			res = shouldEncloseMessage(&privatekeyproto.AttestedBallotAnswer{
				Ballot: engineproto.Ballot{}, Fault: &privatekeyproto.DistantNotaryFault{Code: 0, Summary: err.Error()},
			})
		} else {
			res = shouldEncloseMessage(&privatekeyproto.AttestedBallotAnswer{Ballot: *ballot, Fault: nil})
		}

	case *privatekeyproto.Signal_Attestproposalquery:
		if r.AttestNominationQuery.FetchSeriesUid() != ledgerUID {
			res = shouldEncloseMessage(&privatekeyproto.AttestedNominationAnswer{
				Nomination: engineproto.Nomination{}, Fault: &privatekeyproto.DistantNotaryFault{
					Code:        0,
					Summary: "REDACTED",
				},
			})
			return res, fmt.Errorf("REDACTED", r.AttestNominationQuery.FetchSeriesUid(), ledgerUID)
		}

		nomination := r.AttestNominationQuery.Nomination

		err = privateValue.AttestNomination(ledgerUID, nomination)
		if err != nil {
			res = shouldEncloseMessage(&privatekeyproto.AttestedNominationAnswer{
				Nomination: engineproto.Nomination{}, Fault: &privatekeyproto.DistantNotaryFault{Code: 0, Summary: err.Error()},
			})
		} else {
			res = shouldEncloseMessage(&privatekeyproto.AttestedNominationAnswer{Nomination: *nomination, Fault: nil})
		}
	case *privatekeyproto.Signal_Pingquery:
		err, res = nil, shouldEncloseMessage(&privatekeyproto.PingAnswer{})

	default:
		err = fmt.Errorf("REDACTED", r)
	}

	return res, err
}
