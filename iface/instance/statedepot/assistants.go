package statedepot

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	cryptography "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
)

//
//
func ArbitraryItem() kinds.AssessorRevise {
	publickey := commitrand.Octets(32)
	potency := commitrand.Uint16() + 1
	v := kinds.ReviseAssessor(publickey, int64(potency), "REDACTED")
	return v
}

//
//
//
//
func ArbitraryValues(cnt int) []kinds.AssessorRevise {
	res := make([]kinds.AssessorRevise, cnt)
	for i := 0; i < cnt; i++ {
		res[i] = ArbitraryItem()
	}
	return res
}

//
//
//
func InitializeTokvalDepot(ctx context.Context, app *Platform) error {
	_, err := app.InitializeSuccession(ctx, &kinds.SolicitInitializeSuccession{
		Assessors: ArbitraryValues(1),
	})
	return err
}

//
func FreshTransfer(key, datum string) []byte {
	return []byte(strings.Join([]string{key, datum}, "REDACTED"))
}

func FreshUnpredictableTransfer(extent int) []byte {
	if extent < 4 {
		panic("REDACTED")
	}
	return FreshTransfer(commitrand.Str(2), commitrand.Str(extent-3))
}

func FreshUnpredictableTrans(n int) [][]byte {
	txs := make([][]byte, n)
	for i := 0; i < n; i++ {
		txs[i] = FreshUnpredictableTransfer(10)
	}
	return txs
}

func FreshTransferOriginatingUUID(i int) []byte {
	return []byte(fmt.Sprintf("REDACTED", i, i))
}

//
//
func CreateItemAssignModifyTransfer(publickey security.CommonToken, potency int64) []byte {
	pk, err := cryptography.PublicTokenOriginatingSchema(publickey)
	if err != nil {
		panic(err)
	}
	publicTxt := base64.StdEncoding.EncodeToString(pk.Octets())
	publicKindTxt := pk.Kind()
	return []byte(fmt.Sprintf("REDACTED", AssessorHeading, publicKindTxt, publicTxt, potency))
}
