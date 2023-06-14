package transaction

import "time"

type CreateTransactionParam struct {
	TransactionId   string    `json:"transaction_id"`
	WalletId        string    `json:"wallet_id"`
	Status          int       `json:"status"`
	TransactionType string    `json:"transaction_type"`
	TransactionBy   string    `json:"transaction_by"`
	TransactionDate time.Time `json:"transaction_at"`
	Amount          int       `json:"amount"`
	ReferenceId     string    `json:"reference_id"`
}

func (p *CreateTransactionParam) Validate() error {
	if p.WalletId == "" {
		return ErrWalletIdNil
	}
	if p.Amount == 0 {
		return ErrAmountZero
	}
	if p.ReferenceId == "" {
		return ErrReferenceIdNil
	}
	return nil
}
