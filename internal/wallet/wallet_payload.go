package wallet

type InitializeWalletParam struct {
	WalletId    string `json:"wallet_id"`
	CustomerXid string `json:"customer_xid"`
	StatusId    int    `json:"status_id"`
	Balance     int    `json:"balance"`
}

func (p *InitializeWalletParam) Validate() error {
	if p.WalletId == "" {
		return ErrWalletIdNil
	}
	if p.CustomerXid == "" {
		return ErrCustomerXidNil
	}
	return nil
}

type EnableDisableWalletParam struct {
	WalletId string `json:"wallet_id"`
	StatusId int    `json:"status_id"`
	Balance  int    `json:"balance"`
}

func (p *EnableDisableWalletParam) Validate() error {
	if p.WalletId == "" {
		return ErrWalletIdNil
	}
	return nil
}

type UpdateBalanceWalletParam struct {
	WalletId string `json:"wallet_id"`
	Balance  int    `json:"balance"`
}

func (p *UpdateBalanceWalletParam) Validate() error {
	if p.WalletId == "" {
		return ErrWalletIdNil
	}
	if p.Balance == 0 {
		return ErrAmountZero
	}
	return nil
}
