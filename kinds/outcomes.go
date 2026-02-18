package kinds

import (
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/merkle"
)

//
type IfaceOutcomes []*iface.InvokeTransferOutcome

//
//
func NewOutcomes(replies []*iface.InvokeTransferOutcome) IfaceOutcomes {
	res := make(IfaceOutcomes, len(replies))
	for i, d := range replies {
		res[i] = iface.CertainInvokeTransferOutcome(d)
	}
	return res
}

//
func (a IfaceOutcomes) Digest() []byte {
	return merkle.DigestFromOctetSegments(a.toOctetSegments())
}

//
func (a IfaceOutcomes) DemonstrateOutcome(i int) merkle.Attestation {
	_, evidences := merkle.EvidencesFromOctetSegments(a.toOctetSegments())
	return *evidences[i]
}

func (a IfaceOutcomes) toOctetSegments() [][]byte {
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
