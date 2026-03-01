package kinds

import (
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
)

//
type IfaceOutcomes []*iface.InvokeTransferOutcome

//
//
func FreshOutcomes(replies []*iface.InvokeTransferOutcome) IfaceOutcomes {
	res := make(IfaceOutcomes, len(replies))
	for i, d := range replies {
		res[i] = iface.CertainInvokeTransferOutcome(d)
	}
	return res
}

//
func (a IfaceOutcomes) Digest() []byte {
	return hashmap.DigestOriginatingOctetSegments(a.towardOctetSegments())
}

//
func (a IfaceOutcomes) AscertainOutcome(i int) hashmap.Attestation {
	_, attestations := hashmap.AttestationsOriginatingOctetSegments(a.towardOctetSegments())
	return *attestations[i]
}

func (a IfaceOutcomes) towardOctetSegments() [][]byte {
	l := len(a)
	bzs := make([][]byte, l)
	for i := 0; i < l; i++ {
		bz, err := a[i].Serialize()
		if err != nil {
			panic(err)
		}
		bzs[i] = bz
	}
	return bzs
}
