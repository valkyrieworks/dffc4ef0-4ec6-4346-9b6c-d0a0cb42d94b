package objectdepot

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/valkyrieworks/iface/kinds"
	cryptography "github.com/valkyrieworks/vault/codec"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/schema/consensuscore/vault"
)

//
//
func RandomValue() kinds.RatifierModify {
	publickey := engineseed.Octets(32)
	energy := engineseed.Uint16() + 1
	v := kinds.ModifyRatifier(publickey, int64(energy), "REDACTED")
	return v
}

//
//
//
//
func RandomValues(cnt int) []kinds.RatifierModify {
	res := make([]kinds.RatifierModify, cnt)
	for i := 0; i < cnt; i++ {
		res[i] = RandomValue()
	}
	return res
}

//
//
//
func InitObjectDepot(ctx context.Context, app *Software) error {
	_, err := app.InitSeries(ctx, &kinds.QueryInitSeries{
		Ratifiers: RandomValues(1),
	})
	return err
}

//
func NewTransfer(key, item string) []byte {
	return []byte(strings.Join([]string{key, item}, "REDACTED"))
}

func NewArbitraryTransfer(volume int) []byte {
	if volume < 4 {
		panic("REDACTED")
	}
	return NewTransfer(engineseed.Str(2), engineseed.Str(volume-3))
}

func NewArbitraryTrans(n int) [][]byte {
	txs := make([][]byte, n)
	for i := 0; i < n; i++ {
		txs[i] = NewArbitraryTransfer(10)
	}
	return txs
}

func NewTransferFromUID(i int) []byte {
	return []byte(fmt.Sprintf("REDACTED", i, i))
}

//
//
func CreateValueCollectionAlterTransfer(publickey vault.PublicKey, energy int64) []byte {
	pk, err := cryptography.PublicKeyFromSchema(publickey)
	if err != nil {
		panic(err)
	}
	publicStr := base64.StdEncoding.EncodeToString(pk.Octets())
	publicKindStr := pk.Kind()
	return []byte(fmt.Sprintf("REDACTED", RatifierPrefix, publicKindStr, publicStr, energy))
}
