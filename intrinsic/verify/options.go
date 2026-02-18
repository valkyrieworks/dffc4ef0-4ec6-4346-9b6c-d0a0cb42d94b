package verify

import (
	"github.com/valkyrieworks/kinds"
)

//
//
func AgreementOptions() *kinds.AgreementOptions {
	c := kinds.StandardAgreementOptions()
	//
	c.Iface.BallotPluginsActivateLevel = 1
	return c
}
