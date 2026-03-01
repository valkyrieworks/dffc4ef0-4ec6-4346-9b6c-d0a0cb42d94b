package verify

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
func AgreementSettings() *kinds.AgreementSettings {
	c := kinds.FallbackAgreementSettings()
	//
	c.Iface.BallotAdditionsActivateAltitude = 1
	return c
}
