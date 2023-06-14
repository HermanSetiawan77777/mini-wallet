package walletlog

import (
	"context"
	"time"
)

type WalletLog struct {
	LogId    int       `json:"log_id"`
	WalletId string    `json:"wallet_id"`
	StatusId int       `json:"status_id"`
	DateLog  time.Time `json:"date_log"`
}
type WalletLogRepository interface {
	Create(ctx context.Context, params *WalletLog) error
}

type WalletLogIService interface {
	Create(ctx context.Context, params *CreateWalletLogParam) error
}

type WalletLogService struct {
	repository WalletLogRepository
}

func NewWalletLogService(
	repo WalletLogRepository) *WalletLogService {
	return &WalletLogService{repo}
}

func (s *WalletLogService) Create(ctx context.Context, params *CreateWalletLogParam) error {
	err := params.Validate()
	if err != nil {
		return err
	}
	data := &WalletLog{
		WalletId: params.WalletId,
		StatusId: params.Status,
	}
	err = s.repository.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
