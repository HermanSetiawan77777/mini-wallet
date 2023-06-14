package sql

import (
	"context"
	"herman-technical-julo/internal/log/walletlog"

	"gorm.io/gorm"
)

type WalletLog struct {
	LogId    string `gorm:"primaryKey;autoIncrement;column:log_id"`
	WalletId string `gorm:"column:wallet_id"`
	StatusId int    `gorm:"column:status_id"`
	DateLog  int    `gorm:"column:date_log"`
}

func (WalletLog) TableName() string {
	return "wallet_activation_log"
}

type WalletLogSQLRepository struct {
	db *gorm.DB
}

func NewWalletLogSQLRepository(db *gorm.DB) *WalletLogSQLRepository {
	return &WalletLogSQLRepository{db}
}

func newWalletLogFromServiceModel(data *walletlog.WalletLog) *WalletLog {
	if data == nil {
		return nil
	}

	return &WalletLog{
		WalletId: data.WalletId,
		StatusId: data.StatusId,
	}
}

func (s *WalletLogSQLRepository) Create(ctx context.Context, params *walletlog.WalletLog) error {
	payload := newWalletLogFromServiceModel(params)
	errs := s.db.Create(&payload).Error
	if errs != nil {
		return errs
	}

	return nil
}
