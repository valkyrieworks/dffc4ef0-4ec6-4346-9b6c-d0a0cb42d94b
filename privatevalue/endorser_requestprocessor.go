package privatevalue

import (
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	cryptographyproto "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	privatevalueschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func FallbackCertificationSolicitProcessor(
	privateItem kinds.PrivateAssessor,
	req privatevalueschema.Signal,
	successionUUID string,
) (privatevalueschema.Signal, error) {
	var (
		res privatevalueschema.Signal
		err error
	)

	switch r := req.Sum.(type) {
	case *privatevalueschema.Artifact_Publictokensolicit:
		if r.PublicTokenSolicit.ObtainSuccessionUuid() != successionUUID {
			res = shouldEncloseSignal(&privatevalueschema.PublicTokenReply{
				PublicToken: cryptographyproto.CommonToken{}, Failure: &privatevalueschema.RemoteEndorserFailure{
					Cipher: 0, Characterization: "REDACTED",
				},
			})
			return res, fmt.Errorf("REDACTED", r.PublicTokenSolicit.ObtainSuccessionUuid(), successionUUID)
		}

		var publicToken security.PublicToken
		publicToken, err = privateItem.ObtainPublicToken()
		if err != nil {
			return res, err
		}
		pk, err := cryptocode.PublicTokenTowardSchema(publicToken)
		if err != nil {
			res = shouldEncloseSignal(&privatevalueschema.PublicTokenReply{
				PublicToken: cryptographyproto.CommonToken{}, Failure: &privatevalueschema.RemoteEndorserFailure{Cipher: 0, Characterization: err.Error()},
			})
		} else {
			res = shouldEncloseSignal(&privatevalueschema.PublicTokenReply{PublicToken: pk, Failure: nil})
		}

	case *privatevalueschema.Artifact_Notateballotsolicit:
		if r.AttestBallotSolicit.SuccessionUuid != successionUUID {
			res = shouldEncloseSignal(&privatevalueschema.NotatedBallotReply{
				Ballot: commitchema.Ballot{}, Failure: &privatevalueschema.RemoteEndorserFailure{
					Cipher: 0, Characterization: "REDACTED",
				},
			})
			return res, fmt.Errorf("REDACTED", r.AttestBallotSolicit.ObtainSuccessionUuid(), successionUUID)
		}

		ballot := r.AttestBallotSolicit.Ballot

		err = privateItem.AttestBallot(successionUUID, ballot)
		if err != nil {
			res = shouldEncloseSignal(&privatevalueschema.NotatedBallotReply{
				Ballot: commitchema.Ballot{}, Failure: &privatevalueschema.RemoteEndorserFailure{Cipher: 0, Characterization: err.Error()},
			})
		} else {
			res = shouldEncloseSignal(&privatevalueschema.NotatedBallotReply{Ballot: *ballot, Failure: nil})
		}

	case *privatevalueschema.Artifact_Notateproposalsolicit:
		if r.AttestNominationSolicit.ObtainSuccessionUuid() != successionUUID {
			res = shouldEncloseSignal(&privatevalueschema.NotatedNominationReply{
				Nomination: commitchema.Nomination{}, Failure: &privatevalueschema.RemoteEndorserFailure{
					Cipher:        0,
					Characterization: "REDACTED",
				},
			})
			return res, fmt.Errorf("REDACTED", r.AttestNominationSolicit.ObtainSuccessionUuid(), successionUUID)
		}

		nomination := r.AttestNominationSolicit.Nomination

		err = privateItem.AttestNomination(successionUUID, nomination)
		if err != nil {
			res = shouldEncloseSignal(&privatevalueschema.NotatedNominationReply{
				Nomination: commitchema.Nomination{}, Failure: &privatevalueschema.RemoteEndorserFailure{Cipher: 0, Characterization: err.Error()},
			})
		} else {
			res = shouldEncloseSignal(&privatevalueschema.NotatedNominationReply{Nomination: *nomination, Failure: nil})
		}
	case *privatevalueschema.Artifact_Pingsolicit:
		err, res = nil, shouldEncloseSignal(&privatevalueschema.PingReply{})

	default:
		err = fmt.Errorf("REDACTED", r)
	}

	return res, err
}
