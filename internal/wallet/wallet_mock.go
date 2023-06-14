package wallet

import "context"

type WalletRepositoryMock struct {
	GetByCustomerXidFunc  func(ctx context.Context, customerXid string) (*Wallet, error)
	CreateFunc            func(ctx context.Context, params *Wallet) error
	GetByLinkedWalletFunc func(ctx context.Context, walletId string) (*EnabledWalletDetail, error)
	UpdateFunc            func(ctx context.Context, payload *Wallet) error
	GetByWalletIdFunc     func(ctx context.Context, WalletId string) (*Wallet, error)
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

func (m *WalletRepositoryMock) GetByLinkedWallet(ctx context.Context, walletId string) (*EnabledWalletDetail, error) {
	if m.GetByLinkedWalletFunc != nil {
		return m.GetByLinkedWalletFunc(ctx, walletId)
	}

	return &EnabledWalletDetail{}, nil
}

func (m *WalletRepositoryMock) Update(ctx context.Context, payload *Wallet) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, payload)
	}

	return nil
}

func (m *WalletRepositoryMock) GetByWalletId(ctx context.Context, WalletId string) (*Wallet, error) {
	if m.GetByWalletIdFunc != nil {
		return m.GetByWalletIdFunc(ctx, WalletId)
	}

	return &Wallet{}, nil
}

type WalletServiceMock struct {
	GetByCustomerXidFunc    func(ctx context.Context, customerXid string) (*Wallet, error)
	InitializeWalletFunc    func(ctx context.Context, params *InitializeWalletParam) error
	GetByLinkedWalletFunc   func(ctx context.Context, walletId string) (*EnabledWalletDetail, error)
	EnableWalletFunc        func(ctx context.Context, payload *EnableDisableWalletParam) (*EnabledWalletDetail, error)
	GetByWalletIdFunc       func(ctx context.Context, walletId string) (*Wallet, error)
	UpdateBalanceWalletFunc func(ctx context.Context, payload *UpdateBalanceWalletParam) error
	DisableWalletFunc       func(ctx context.Context, params *EnableDisableWalletParam) (*DisableWalletDetail, error)
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

func (m *WalletServiceMock) GetByLinkedWallet(ctx context.Context, walletId string) (*EnabledWalletDetail, error) {
	if m.GetByLinkedWalletFunc != nil {
		return m.GetByLinkedWalletFunc(ctx, walletId)
	}

	return &EnabledWalletDetail{}, nil
}

func (m *WalletServiceMock) EnableWallet(ctx context.Context, payload *EnableDisableWalletParam) (*EnabledWalletDetail, error) {
	if m.EnableWalletFunc != nil {
		return m.EnableWalletFunc(ctx, payload)
	}

	return &EnabledWalletDetail{}, nil
}

func (m *WalletServiceMock) GetByWalletId(ctx context.Context, WalletId string) (*Wallet, error) {
	if m.GetByWalletIdFunc != nil {
		return m.GetByWalletIdFunc(ctx, WalletId)
	}

	return &Wallet{}, nil
}

func (m *WalletServiceMock) UpdateBalanceWallet(ctx context.Context, payload *UpdateBalanceWalletParam) error {
	if m.UpdateBalanceWalletFunc != nil {
		return m.UpdateBalanceWalletFunc(ctx, payload)
	}

	return nil
}

func (m *WalletServiceMock) DisableWallet(ctx context.Context, payload *EnableDisableWalletParam) (*DisableWalletDetail, error) {
	if m.DisableWalletFunc != nil {
		return m.DisableWalletFunc(ctx, payload)
	}

	return &DisableWalletDetail{}, nil
}
