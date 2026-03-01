package testsuite

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

func InitializeSuccession(ctx context.Context, customer abcicustomer.Customer) error {
	sum := 10
	values := make([]kinds.AssessorRevise, sum)
	for i := 0; i < sum; i++ {
		publickey := commitrand.Octets(33)
		potency := commitrand.Int()
		values[i] = kinds.ReviseAssessor(publickey, int64(potency), "REDACTED")
	}
	_, err := customer.InitializeSuccession(ctx, &kinds.SolicitInitializeSuccession{
		Assessors: values,
	})
	if err != nil {
		fmt.Printf("REDACTED", err)
		return err
	}
	fmt.Println("REDACTED")
	return nil
}

func Endorse(ctx context.Context, customer abcicustomer.Customer) error {
	_, err := customer.Endorse(ctx, &kinds.SolicitEndorse{})
	if err != nil {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED", err)
		return err
	}
	fmt.Println("REDACTED")
	return nil
}

func CulminateLedger(ctx context.Context, customer abcicustomer.Customer, transferOctets [][]byte, cipherExpiration []uint32, dataExpiration []byte, digestExpiration []byte) error {
	res, _ := customer.CulminateLedger(ctx, &kinds.SolicitCulminateLedger{Txs: transferOctets})
	platformDigest := res.PlatformDigest
	for i, tx := range res.TransferOutcomes {
		cipher, data, log := tx.Cipher, tx.Data, tx.Log
		if cipher != cipherExpiration[i] {
			fmt.Println("REDACTED")
			fmt.Printf("REDACTED",
				cipher, cipherExpiration, log)
			return errors.New("REDACTED")
		}
		if !bytes.Equal(data, dataExpiration) {
			fmt.Println("REDACTED")
			fmt.Printf("REDACTED",
				data, dataExpiration)
			return errors.New("REDACTED")
		}
	}
	if !bytes.Equal(platformDigest, digestExpiration) {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED", platformDigest, digestExpiration)
		return errors.New("REDACTED")
	}
	fmt.Println("REDACTED")
	return nil
}

func ArrangeNomination(ctx context.Context, customer abcicustomer.Customer, transferOctets [][]byte, transferAnticipated [][]byte, _ []byte) error {
	res, _ := customer.ArrangeNomination(ctx, &kinds.SolicitArrangeNomination{Txs: transferOctets})
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

func HandleNomination(ctx context.Context, customer abcicustomer.Customer, transferOctets [][]byte, conditionExpiration kinds.Responseexecuteitem_Itemstatus) error {
	res, _ := customer.HandleNomination(ctx, &kinds.SolicitHandleNomination{Txs: transferOctets})
	if res.Condition != conditionExpiration {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED",
			res.Condition, conditionExpiration)
		return errors.New("REDACTED")
	}
	fmt.Println("REDACTED")
	return nil
}

func InspectTransfer(ctx context.Context, customer abcicustomer.Customer, transferOctets []byte, cipherExpiration uint32, dataExpiration []byte) error {
	res, _ := customer.InspectTransfer(ctx, &kinds.SolicitInspectTransfer{Tx: transferOctets})
	cipher, data, log := res.Cipher, res.Data, res.Log
	if cipher != cipherExpiration {
		fmt.Println("REDACTED")
		fmt.Printf("REDACTED",
			cipher, cipherExpiration, log)
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
