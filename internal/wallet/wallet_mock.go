package wallet

import "context"

type WalletRepositoryMock struct {
	GetByCustomerXidFunc  func(ctx context.Context, customerXid string) (*Wallet, error)
	CreateFunc            func(ctx context.Context, params *Wallet) error
	GetByLinkedWalletFunc func(ctx context.Context, walletId string) (*WalletDetail, error)
}

func (m *WalletRepositoryMock) GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error) {
	if m.GetByCustomerXidFunc != nil {
		return m.GetByCustomerXidFunc(ctx, customerXid)
	}

	return &Wallet{}, nil
}

func (m *WalletRepositoryMock) Create(ctx context.Context, params *Wallet) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, params)
	}

	return nil
}

func (m *WalletRepositoryMock) GetByLinkedWallet(ctx context.Context, walletId string) (*WalletDetail, error) {
	if m.GetByLinkedWalletFunc != nil {
		return m.GetByLinkedWalletFunc(ctx, walletId)
	}

	return &WalletDetail{}, nil
}

type WalletServiceMock struct {
	GetByCustomerXidFunc  func(ctx context.Context, customerXid string) (*Wallet, error)
	InitializeWalletFunc  func(ctx context.Context, params *InitializeWalletParam) error
	GetByLinkedWalletFunc func(ctx context.Context, walletId string) (*WalletDetail, error)
}

func (m *WalletServiceMock) GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error) {
	if m.GetByCustomerXidFunc != nil {
		return m.GetByCustomerXidFunc(ctx, customerXid)
	}

	return &Wallet{}, nil
}

func (m *WalletServiceMock) InitializeWallet(ctx context.Context, params *InitializeWalletParam) error {
	if m.InitializeWalletFunc != nil {
		return m.InitializeWalletFunc(ctx, params)
	}

	return nil
}

func (m *WalletServiceMock) GetByLinkedWallet(ctx context.Context, walletId string) (*WalletDetail, error) {
	if m.GetByLinkedWalletFunc != nil {
		return m.GetByLinkedWalletFunc(ctx, walletId)
	}

	return &WalletDetail{}, nil
}
