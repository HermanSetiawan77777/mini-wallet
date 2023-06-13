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
	WalletId    string    `json:"id"`
	CustomerXid string    `json:"owned_by"`
	Status      string    `json:"status"`
	DateLog     time.Time `json:"enabled_at"`
	Balance     int       `json:"balance"`
}
type WalletRepository interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
	Create(ctx context.Context, params *Wallet) error
	GetByLinkedWallet(ctx context.Context, walletId string) (*WalletDetail, error)
	Update(ctx context.Context, payload *Wallet) error
	GetByWalletId(ctx context.Context, WalletId string) (*Wallet, error)
}

type WalletIService interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
	InitializeWallet(ctx context.Context, params *InitializeWalletParam) error
	GetByLinkedWallet(ctx context.Context, walletId string) (*WalletDetail, error)
	EnableWallet(ctx context.Context, payload *UpdateWalletParam) (*WalletDetail, error)
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

func (s *WalletService) EnableWallet(ctx context.Context, params *UpdateWalletParam) (*WalletDetail, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	targetWallet, err := s.repository.GetByWalletId(ctx, params.WalletId)
	if err != nil {
		return nil, err
	}
	if targetWallet == nil {
		return nil, ErrWalletIdNotExist
	}
	if targetWallet.StatusId == 2 {
		return nil, ErrWalletAlreadyEnabled
	}
	targetWallet.StatusId = 2
	if targetWallet.StatusId == 0 {
		return nil, ErrStatusZero
	}
	if err = s.repository.Update(ctx, targetWallet); err != nil {
		return nil, err
	}

	currentWallet, err := s.repository.GetByLinkedWallet(ctx, params.WalletId)
	if err != nil {
		return nil, err
	}
	if currentWallet.WalletId == "" {
		return nil, ErrWalletIdNotExist
	}

	return currentWallet, nil
}
