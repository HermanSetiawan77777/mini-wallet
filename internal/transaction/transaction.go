package transaction

import (
	"context"
	"herman-technical-julo/internal/wallet"
	"time"
)

type ViewTransactionWallet struct {
	TransactionId   string    `json:"transaction_id"`
	WalletId        string    `json:"wallet_id"`
	TransactionBy   string    `json:"transaction_by"`
	TransactionType string    `json:"transaction_type"`
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"transaction_at"`
	Amount          int       `json:"amount"`
	ReferenceId     string    `json:"reference_id"`
}
type TransactionWalletRepository interface {
	ViewMyTransactionWallet(ctx context.Context, walletId string) ([]*ViewTransactionWallet, error)
}

type TransactionWalletIService interface {
	ViewMyTransactionWallet(ctx context.Context, walletId string) ([]*ViewTransactionWallet, error)
}

type TransactionWalletService struct {
	repository    TransactionWalletRepository
	walletService wallet.WalletIService
}

func NewTransactionWalletService(
	repo TransactionWalletRepository,
	walletService wallet.WalletIService) *TransactionWalletService {
	return &TransactionWalletService{repo, walletService}
}

func (s *TransactionWalletService) ViewMyTransactionWallet(ctx context.Context, walletId string) ([]*ViewTransactionWallet, error) {
	checkWallet, err := s.walletService.GetByWalletId(ctx, walletId)
	if err != nil {
		return nil, err
	}
	if checkWallet.StatusId == 1 {
		return nil, ErrTransactionWalletDeactive
	}

	transactionWallet, err := s.repository.ViewMyTransactionWallet(ctx, walletId)
	if err != nil {
		return nil, err
	}
	return transactionWallet, nil
}
