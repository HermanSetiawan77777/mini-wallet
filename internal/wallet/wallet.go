package wallet

import "context"

type Wallet struct {
	WalletId    string `json:"wallet_id"`
	CustomerXid string `json:"customer_xid"`
	StatusId    int    `json:"status_id"`
	Balance     int    `json:"balance"`
}
type WalletRepository interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
}

type WalletIService interface {
	GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error)
}

type WalletService struct {
	repository WalletRepository
}

func NewWalletService(repo WalletRepository) *WalletService {
	return &WalletService{repo}
}

func (w *WalletService) GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error) {
	wallet, err := w.repository.GetByCustomerXid(ctx, customerXid)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
