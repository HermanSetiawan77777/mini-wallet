package sql

import (
	"context"
	sqlgorm "herman-technical-julo/internal/data/sql/gorm"
	"herman-technical-julo/internal/wallet"

	"gorm.io/gorm"
)

type Wallet struct {
	WalletId    string `gorm:"column:wallet_id"`
	CustomerXid string `gorm:"column:customer_xid"`
	StatusId    int    `gorm:"column:status_id"`
	Balance     int    `gorm:"column:balance"`
}

func (Wallet) TableName() string {
	return "wallet"
}

type WalletSQLRepository struct {
	db *gorm.DB
}

func NewWalletSQLRepository(db *gorm.DB) *WalletSQLRepository {
	return &WalletSQLRepository{db}
}

func (w *Wallet) ToServiceModel() *wallet.Wallet {
	return &wallet.Wallet{
		WalletId:    w.WalletId,
		CustomerXid: w.CustomerXid,
		StatusId:    w.StatusId,
		Balance:     w.Balance,
	}
}

func (s *WalletSQLRepository) GetByCustomerXid(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
	var data *Wallet
	db := s.getDatabaseClient(ctx)

	err := db.Where("customer_xid = ?", customerXid).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return data.ToServiceModel(), nil
}

func newWalletFromServiceModel(data *wallet.Wallet) *Wallet {
	if data == nil {
		return nil
	}

	return &Wallet{
		WalletId:    data.WalletId,
		CustomerXid: data.CustomerXid,
		StatusId:    data.StatusId,
		Balance:     data.Balance,
	}
}

func (s *WalletSQLRepository) Create(ctx context.Context, params *wallet.Wallet) error {
	payload := newWalletFromServiceModel(params)
	errs := s.db.Create(&payload).Error
	if errs != nil {
		return errs
	}

	return nil
}

func (s *WalletSQLRepository) getDatabaseClient(ctx context.Context) *gorm.DB {
	db := sqlgorm.GetClientFromContext(ctx)
	if db != nil {
		return db
	}

	return s.db
}
