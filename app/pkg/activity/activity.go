package activity

import (
	"context"
	"log"

	"github.com/hyorimitsu/hello-temporal.io/app/pkg/domain"
)

func Withdraw(_ context.Context, transferDetails domain.TransferDetails) error {
	log.Printf(
		"Withdrawing: [ID=%s, Ammount=%d, From=%s, To=%s]\n",
		transferDetails.ID,
		transferDetails.Amount,
		transferDetails.From,
		transferDetails.To,
	)
	return nil
}

func Deposit(_ context.Context, transferDetails domain.TransferDetails) error {
	log.Printf(
		"Depositing: [ID=%s, Ammount=%d, From=%s, To=%s]\n",
		transferDetails.ID,
		transferDetails.Amount,
		transferDetails.From,
		transferDetails.To,
	)
	// Switch comment out to simulate activity errors
	//return errors.New("unable to deposit")
	return nil
}
