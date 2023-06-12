package wallet

import "context"

type WalletRepositoryMock struct {
	GetByCustomerXidFunc func(ctx context.Context, customerXid string) (*Wallet, error)
}

func (m *WalletRepositoryMock) GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error) {
	if m.GetByCustomerXidFunc != nil {
		return m.GetByCustomerXidFunc(ctx, customerXid)
	}

	return &Wallet{}, nil
}

type WalletServiceMock struct {
	GetByCustomerXidFunc func(ctx context.Context, customerXid string) (*Wallet, error)
}

func (m *WalletServiceMock) GetByCustomerXid(ctx context.Context, customerXid string) (*Wallet, error) {
	if m.GetByCustomerXidFunc != nil {
		return m.GetByCustomerXidFunc(ctx, customerXid)
	}

	return &Wallet{}, nil
}
