package kinds

import engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"

//
func IsBallotKindSound(t engineproto.AttestedMessageKind) bool {
	switch t {
	case engineproto.PreballotKind, engineproto.PreendorseKind:
		return true
	default:
		return false
	}
}

var attestedMessageKindToBriefLabel = map[engineproto.AttestedMessageKind]string{
	engineproto.UnclearKind:   "REDACTED",
	engineproto.PreballotKind:   "REDACTED",
	engineproto.PreendorseKind: "REDACTED",
	engineproto.NominationKind:  "REDACTED",
}

//
func AttestedMessageKindToBriefString(t engineproto.AttestedMessageKind) string {
	if briefLabel, ok := attestedMessageKindToBriefLabel[t]; ok {
		return briefLabel
	}
	return "REDACTED"
}
