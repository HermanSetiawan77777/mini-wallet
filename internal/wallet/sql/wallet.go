package sql

import (
	"context"
	sqlgorm "herman-technical-julo/internal/data/sql/gorm"
	"herman-technical-julo/internal/wallet"
	"time"

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

type LinkedWalletDetail struct {
	WalletId    string    `gorm:"column:wallet_id"`
	CustomerXid string    `gorm:"column:customer_xid"`
	Balance     string    `gorm:"column:balance"`
	DateLog     time.Time `gorm:"column:date_log"`
	Status      string    `gorm:"column:status"`
}

func (LinkedWalletDetail) TableName() string {
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

func (w *LinkedWalletDetail) ToLinkedServiceModel() *wallet.WalletDetail {
	return &wallet.WalletDetail{
		WalletId:    w.WalletId,
		CustomerXid: w.CustomerXid,
		Status:      w.Status,
		DateLog:     w.DateLog,
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

func (s *WalletSQLRepository) GetByWalletId(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
	var data *Wallet
	db := s.getDatabaseClient(ctx)

	err := db.Where("wallet_id = ?", customerXid).First(&data).Error
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

func (c *LinkedWalletDetail) ToServiceModelDetail() *wallet.WalletDetail {
	return &wallet.WalletDetail{
		WalletId:    c.WalletId,
		CustomerXid: c.CustomerXid,
		Status:      c.Status,
		DateLog:     c.DateLog,
	}
}

func (s *WalletSQLRepository) GetByLinkedWallet(ctx context.Context, walletId string) (*wallet.WalletDetail, error) {
	var data *LinkedWalletDetail
	db := s.getDatabaseClient(ctx)
	if walletId == "" {
		return nil, nil
	}
	err := db.Model(&LinkedWalletDetail{}).Select("wallet.wallet_id, wallet.customer_xid, wallet.balance, wallet_activation_log.date_log, status_wallet.status").Joins("INNER join wallet_activation_log on wallet.wallet_id = wallet_activation_log.wallet_id").Joins("INNER join status_wallet on status_wallet.status_id = wallet.status_id").Where("wallet.wallet_id = ?", walletId).Order("wallet_activation_log.date_log desc").Scan(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data.ToLinkedServiceModel(), nil
}

func (s *WalletSQLRepository) getDatabaseClient(ctx context.Context) *gorm.DB {
	db := sqlgorm.GetClientFromContext(ctx)
	if db != nil {
		return db
	}

	return s.db
}

func (s *WalletSQLRepository) Update(ctx context.Context, payload *wallet.Wallet) error {
	db := s.getDatabaseClient(ctx)
	newWallet := newWalletFromServiceModel(payload)

	err := db.Where("wallet_id = ?", payload.WalletId).Select("*").Updates(&newWallet).Error
	if err != nil {
		return err
	}

	return nil
}
