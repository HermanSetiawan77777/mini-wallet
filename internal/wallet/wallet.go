package wallet

import (
	"context"
	"time"
)

type Wallet struct {
	WalletId    string `json:"wallet_id"`
	CustomerXid string `json:"customer_xid"`
	StatusId    int    `json:"status_id"`
	Balance     int    `json:"balance"`
}

type WalletDetail struct {
	WalletId    string    `json:"wallet_id"`
	CustomerXid string    `json:"customer_xid"`
	Status      string    `json:"status"`
	DateLog     time.Time `json:"enabled_at"`
}
type WalletRepository interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
	Create(ctx context.Context, params *Wallet) error
	GetByLinkedWallet(ctx context.Context, walletId string) (*WalletDetail, error)
}

type WalletIService interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
	InitializeWallet(ctx context.Context, params *InitializeWalletParam) error
	GetByLinkedWallet(ctx context.Context, walletId string) (*WalletDetail, error)
}

type WalletService struct {
	repository WalletRepository
}

func NewWalletService(
	repo WalletRepository) *WalletService {
	return &WalletService{repo}
}

func (w *WalletService) GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error) {
	wallet, err := w.repository.GetByCustomerXid(ctx, customerXid)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *WalletService) InitializeWallet(ctx context.Context, params *InitializeWalletParam) error {
	err := params.Validate()
	if err != nil {
		return err
	}
	data := &Wallet{
		WalletId:    params.WalletId,
		Balance:     params.Balance,
		CustomerXid: params.CustomerXid,
		StatusId:    0,
	}
	err = s.repository.Create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *WalletService) GetByLinkedWallet(ctx context.Context, walletId string) (*WalletDetail, error) {
	wallet, err := s.repository.GetByLinkedWallet(ctx, walletId)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
