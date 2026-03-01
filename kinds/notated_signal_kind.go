package kinds

import commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"

//
func EqualsBallotKindSound(t commitchema.AttestedSignalKind) bool {
	switch t {
	case commitchema.PreballotKind, commitchema.PreendorseKind:
		return true
	default:
		return false
	}
}

var notatedSignalKindTowardBriefAlias = map[commitchema.AttestedSignalKind]string{
	commitchema.UnfamiliarKind:   "REDACTED",
	commitchema.PreballotKind:   "REDACTED",
	commitchema.PreendorseKind: "REDACTED",
	commitchema.NominationKind:  "REDACTED",
}

//
func AttestedSignalKindTowardBriefText(t commitchema.AttestedSignalKind) string {
	if briefAlias, ok := notatedSignalKindTowardBriefAlias[t]; ok {
		return briefAlias
	}
	return "REDACTED"
}
