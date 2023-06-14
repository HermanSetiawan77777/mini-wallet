package transaction

import (
	"context"
	"herman-technical-julo/internal/wallet"
	"time"

	"github.com/google/uuid"
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

type TransactionWallet struct {
	TransactionId   string `json:"transaction_id"`
	WalletId        string `json:"wallet_id"`
	Status          int    `json:"status"`
	TransactionType string `json:"transaction_type"`
	TransactionBy   string `json:"transaction_by"`
	Amount          int    `json:"amount"`
	ReferenceId     string `json:"reference_id"`
}

type ViewTransactionDepositWallet struct {
	TransactionId   string    `json:"id"`
	TransactionBy   string    `json:"deposited_by"`
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"deposited_at"`
	Amount          int       `json:"amount"`
	ReferenceId     string    `json:"reference_id"`
}

type ViewTransactionWithdrawalWallet struct {
	TransactionId   string    `json:"id"`
	TransactionBy   string    `json:"withdrawn_by"`
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"withdrawn_at"`
	Amount          int       `json:"amount"`
	ReferenceId     string    `json:"reference_id"`
}

type TransactionWalletRepository interface {
	ViewMyTransactionWallet(ctx context.Context, walletId string) ([]*ViewTransactionWallet, error)
	Create(ctx context.Context, params *TransactionWallet) error
	GetViewDetailDepositTransaction(ctx context.Context, transactionId string) (*ViewTransactionDepositWallet, error)
	GetViewDetailWithdrawalTransaction(ctx context.Context, transactionId string) (*ViewTransactionWithdrawalWallet, error)
}

type TransactionWalletIService interface {
	ViewMyTransactionWallet(ctx context.Context, walletId string) ([]*ViewTransactionWallet, error)
	DepositTransaction(ctx context.Context, params *CreateTransactionParam) (*ViewTransactionDepositWallet, error)
	WithdrawalTransaction(ctx context.Context, params *CreateTransactionParam) (*ViewTransactionWithdrawalWallet, error)
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

func (s *TransactionWalletService) DepositTransaction(ctx context.Context, params *CreateTransactionParam) (*ViewTransactionDepositWallet, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}
	if params.TransactionBy == "" {
		return nil, ErrTransactionByNil
	}
	statusWallet, err := s.walletService.GetByWalletId(ctx, params.WalletId)
	if err != nil {
		return nil, err
	}
	if statusWallet.StatusId == 1 {
		return nil, ErrWalletDeactive
	}

	id := uuid.New()
	idString := id.String()
	data := &TransactionWallet{
		WalletId:        params.WalletId,
		TransactionId:   idString,
		Status:          1,
		TransactionType: "deposit",
		TransactionBy:   params.TransactionBy,
		Amount:          params.Amount,
		ReferenceId:     params.ReferenceId,
	}
	err = s.repository.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	err = s.walletService.UpdateBalanceWallet(ctx, &wallet.UpdateBalanceWalletParam{
		WalletId: params.WalletId,
		Balance:  statusWallet.Balance + params.Amount,
	})
	if err != nil {
		return nil, err
	}

	view, err := s.repository.GetViewDetailDepositTransaction(ctx, idString)
	if err != nil {
		return nil, err
	}

	return view, nil
}

func (s *TransactionWalletService) WithdrawalTransaction(ctx context.Context, params *CreateTransactionParam) (*ViewTransactionWithdrawalWallet, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}
	if params.TransactionBy == "" {
		return nil, ErrTransactionByNil
	}
	statusWallet, err := s.walletService.GetByWalletId(ctx, params.WalletId)
	if err != nil {
		return nil, err
	}
	if statusWallet.StatusId == 1 {
		return nil, ErrWalletDeactive
	}
	if statusWallet.Balance < params.Amount {
		return nil, ErrAmountNotEnough
	}

	id := uuid.New()
	idString := id.String()
	data := &TransactionWallet{
		WalletId:        params.WalletId,
		TransactionId:   idString,
		Status:          1,
		TransactionType: "withdrawal",
		TransactionBy:   params.TransactionBy,
		Amount:          params.Amount,
		ReferenceId:     params.ReferenceId,
	}
	err = s.repository.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	err = s.walletService.UpdateBalanceWallet(ctx, &wallet.UpdateBalanceWalletParam{
		WalletId: params.WalletId,
		Balance:  statusWallet.Balance - params.Amount,
	})
	if err != nil {
		return nil, err
	}

	view, err := s.repository.GetViewDetailWithdrawalTransaction(ctx, idString)
	if err != nil {
		return nil, err
	}

	return view, nil
}
