package wallet

import (
	"context"
	"herman-technical-julo/internal/log/walletlog"
	"time"
)

type Wallet struct {
	WalletId    string `json:"wallet_id"`
	CustomerXid string `json:"customer_xid"`
	StatusId    int    `json:"status_id"`
	Balance     int    `json:"balance"`
}

type EnabledWalletDetail struct {
	WalletId    string    `json:"id"`
	CustomerXid string    `json:"owned_by"`
	Status      string    `json:"status"`
	DateLog     time.Time `json:"enabled_at"`
	Balance     int       `json:"balance"`
}

type DisableWalletDetail struct {
	WalletId    string    `json:"id"`
	CustomerXid string    `json:"owned_by"`
	Status      string    `json:"status"`
	DateLog     time.Time `json:"disabled_at"`
	Balance     int       `json:"balance"`
}

type WalletRepository interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
	Create(ctx context.Context, params *Wallet) error
	GetByLinkedWallet(ctx context.Context, walletId string) (*EnabledWalletDetail, error)
	Update(ctx context.Context, payload *Wallet) error
	GetByWalletId(ctx context.Context, WalletId string) (*Wallet, error)
}

type WalletIService interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
	InitializeWallet(ctx context.Context, params *InitializeWalletParam) error
	GetByLinkedWallet(ctx context.Context, walletId string) (*EnabledWalletDetail, error)
	EnableWallet(ctx context.Context, payload *EnableDisableWalletParam) (*EnabledWalletDetail, error)
	GetByWalletId(ctx context.Context, walletId string) (*Wallet, error)
	UpdateBalanceWallet(ctx context.Context, payload *UpdateBalanceWalletParam) error
	DisableWallet(ctx context.Context, params *EnableDisableWalletParam) (*DisableWalletDetail, error)
}

type WalletService struct {
	repository       WalletRepository
	walletLogService walletlog.WalletLogIService
}

func NewWalletService(
	repo WalletRepository,
	walletLogService walletlog.WalletLogIService) *WalletService {
	return &WalletService{repo, walletLogService}
}

func (w *WalletService) GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error) {
	wallet, err := w.repository.GetByCustomerXid(ctx, customerXid)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (w *WalletService) GetByWalletId(ctx context.Context, walletId string) (*Wallet, error) {
	wallet, err := w.repository.GetByWalletId(ctx, walletId)
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

func (s *WalletService) GetByLinkedWallet(ctx context.Context, walletId string) (*EnabledWalletDetail, error) {
	checkWallet, err := s.repository.GetByWalletId(ctx, walletId)
	if err != nil {
		return nil, err
	}
	if checkWallet.StatusId == 1 {
		return nil, ErrWalletDeactive
	}
	wallet, err := s.repository.GetByLinkedWallet(ctx, walletId)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *WalletService) EnableWallet(ctx context.Context, params *EnableDisableWalletParam) (*EnabledWalletDetail, error) {
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

	err = s.walletLogService.Create(ctx, &walletlog.CreateWalletLogParam{
		WalletId: params.WalletId,
		Status:   2,
	})
	if err != nil {
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

func (s *WalletService) UpdateBalanceWallet(ctx context.Context, payload *UpdateBalanceWalletParam) error {
	err := payload.Validate()
	if err != nil {
		return err
	}
	targetWallet, err := s.repository.GetByWalletId(ctx, payload.WalletId)
	if err != nil {
		return err
	}
	if targetWallet == nil {
		return ErrWalletIdNotExist
	}
	if targetWallet.StatusId == 1 {
		return ErrWalletDeactive
	}
	targetWallet.Balance = payload.Balance

	if err = s.repository.Update(ctx, targetWallet); err != nil {
		return err
	}

	return nil
}

func (s *WalletService) DisableWallet(ctx context.Context, params *EnableDisableWalletParam) (*DisableWalletDetail, error) {
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
	if targetWallet.StatusId == 1 {
		return nil, ErrWalletAlreadyDisabled
	}
	targetWallet.StatusId = 1
	if targetWallet.StatusId == 0 {
		return nil, ErrStatusZero
	}

	err = s.walletLogService.Create(ctx, &walletlog.CreateWalletLogParam{
		WalletId: params.WalletId,
		Status:   1,
	})
	if err != nil {
		return nil, err
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

	payload := &DisableWalletDetail{
		WalletId:    currentWallet.WalletId,
		CustomerXid: currentWallet.CustomerXid,
		Status:      currentWallet.Status,
		DateLog:     currentWallet.DateLog,
		Balance:     currentWallet.Balance,
	}

	return payload, nil
}
