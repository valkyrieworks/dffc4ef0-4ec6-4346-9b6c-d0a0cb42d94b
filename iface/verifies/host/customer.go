package testsuite

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	abciend "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/kinds"
	engineseed "github.com/valkyrieworks/utils/random"
)

func InitSeries(ctx context.Context, customer abciend.Customer) error {
	sum := 10
	values := make([]kinds.RatifierModify, sum)
	for i := 0; i < sum; i++ {
		publickey := engineseed.Octets(33)
		energy := engineseed.Int()
		values[i] = kinds.ModifyRatifier(publickey, int64(energy), "REDACTED")
	}
	_, err := customer.InitSeries(ctx, &kinds.QueryInitSeries{
		Ratifiers: values,
	})
	if err != nil {
		fmt.Printf("REDACTED", err)
		return err
	}
	fmt.Println("REDACTED")
	return nil
}

func Endorse(ctx context.Context, customer abciend.Customer) error {
	_, err := customer.Endorse(ctx, &kinds.QueryEndorse{})
	if err != nil {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED", err)
		return err
	}
	fmt.Println("REDACTED")
	return nil
}

func CompleteLedger(ctx context.Context, customer abciend.Customer, transferOctets [][]byte, codeExpiration []uint32, dataExpiration []byte, digestExpiration []byte) error {
	res, _ := customer.CompleteLedger(ctx, &kinds.QueryCompleteLedger{Txs: transferOctets})
	applicationDigest := res.ApplicationDigest
	for i, tx := range res.TransOutcomes {
		code, data, log := tx.Code, tx.Data, tx.Log
		if code != codeExpiration[i] {
			fmt.Println("REDACTED")
			fmt.Printf("REDACTED",
				code, codeExpiration, log)
			return errors.New("REDACTED")
		}
		if !bytes.Equal(data, dataExpiration) {
			fmt.Println("REDACTED")
			fmt.Printf("REDACTED",
				data, dataExpiration)
			return errors.New("REDACTED")
		}
	}
	if !bytes.Equal(applicationDigest, digestExpiration) {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED", applicationDigest, digestExpiration)
		return errors.New("REDACTED")
	}
	fmt.Println("REDACTED")
	return nil
}

func ArrangeNomination(ctx context.Context, customer abciend.Customer, transferOctets [][]byte, transferAnticipated [][]byte, _ []byte) error {
	res, _ := customer.ArrangeNomination(ctx, &kinds.QueryArrangeNomination{Txs: transferOctets})
	for i, tx := range res.Txs {
		if !bytes.Equal(tx, transferAnticipated[i]) {
			fmt.Println("REDACTED")
			fmt.Printf("REDACTED",
				tx, transferAnticipated[i])
			return errors.New("REDACTED")
		}
	}
	fmt.Println("REDACTED")
	return nil
}

func HandleNomination(ctx context.Context, customer abciend.Customer, transferOctets [][]byte, stateExpiration kinds.Responseprocessnomination_Nominationstate) error {
	res, _ := customer.HandleNomination(ctx, &kinds.QueryHandleNomination{Txs: transferOctets})
	if res.Status != stateExpiration {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED",
			res.Status, stateExpiration)
		return errors.New("REDACTED")
	}
	fmt.Println("REDACTED")
	return nil
}

func InspectTransfer(ctx context.Context, customer abciend.Customer, transferOctets []byte, codeExpiration uint32, dataExpiration []byte) error {
	res, _ := customer.InspectTransfer(ctx, &kinds.QueryInspectTransfer{Tx: transferOctets})
	code, data, log := res.Code, res.Data, res.Log
	if code != codeExpiration {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED",
			code, codeExpiration, log)
		return errors.New("REDACTED")
	}
	if !bytes.Equal(data, dataExpiration) {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED",
			data, dataExpiration)
		return errors.New("REDACTED")
	}
	fmt.Println("REDACTED")
	return nil
}
