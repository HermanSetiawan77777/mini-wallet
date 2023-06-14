package walletlog

type CreateWalletLogParam struct {
	WalletId string `json:"wallet_id"`
	Status   int    `json:"status"`
}

func (p *CreateWalletLogParam) Validate() error {
	if p.WalletId == "" {
		return ErrWalletIdNil
	}
	if p.Status == 0 {
		return ErrStatusZero
	}
	return nil
}
