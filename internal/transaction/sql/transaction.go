package sql

import (
	"context"
	sqlgorm "herman-technical-julo/internal/data/sql/gorm"
	"herman-technical-julo/internal/transaction"
	"time"

	"gorm.io/gorm"
)

type ViewTransactionWallet struct {
	WalletId        string    `gorm:"column:wallet_id"`
	TransactionId   string    `gorm:"column:transaction_id"`
	TransactionBy   string    `gorm:"column:transaction_by"`
	TransactionType string    `gorm:"column:transaction_type"`
	Status          string    `gorm:"column:status"`
	TransactionDate time.Time `gorm:"column:transaction_date"`
	Amount          int       `gorm:"column:amount"`
	ReferenceId     string    `gorm:"column:reference_id"`
}

func (ViewTransactionWallet) TableName() string {
	return "transaction_wallet"
}

type TransactionWalletSQLRepository struct {
	db *gorm.DB
}

func NewTransactionWalletSQLRepository(db *gorm.DB) *TransactionWalletSQLRepository {
	return &TransactionWalletSQLRepository{db}
}

func (v *ViewTransactionWallet) ToServiceModelTransaction() *transaction.ViewTransactionWallet {
	return &transaction.ViewTransactionWallet{
		WalletId:        v.WalletId,
		TransactionId:   v.TransactionId,
		TransactionBy:   v.TransactionBy,
		Status:          v.Status,
		TransactionDate: v.TransactionDate,
		Amount:          v.Amount,
		ReferenceId:     v.ReferenceId,
		TransactionType: v.TransactionType,
	}
}

func ToServiceSliceTransactionWallet(datas []*ViewTransactionWallet) []*transaction.ViewTransactionWallet {
	tr := []*transaction.ViewTransactionWallet{}

	for _, data := range datas {
		tr = append(tr, data.ToServiceModelTransaction())
	}

	return tr
}

func (s *TransactionWalletSQLRepository) ViewMyTransactionWallet(ctx context.Context, walletId string) ([]*transaction.ViewTransactionWallet, error) {
	var data []*ViewTransactionWallet
	db := s.getDatabaseClient(ctx)
	if walletId == "" {
		return nil, nil
	}
	err := db.Model(&ViewTransactionWallet{}).Select("transaction_wallet.transaction_id, transaction_wallet.wallet_id, transaction_wallet.transaction_type, transaction_status.status, transaction_wallet.transaction_by, transaction_wallet.transaction_date, transaction_wallet.amount, transaction_wallet.reference_id").Joins("INNER join transaction_status ON transaction_wallet.status=transaction_status.status_id").Where("wallet.wallet_id = ?", walletId).Order("transaction_wallet.transaction_date asc").Scan(&data).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return ToServiceSliceTransactionWallet(data), nil
}

func (s *TransactionWalletSQLRepository) getDatabaseClient(ctx context.Context) *gorm.DB {
	db := sqlgorm.GetClientFromContext(ctx)
	if db != nil {
		return db
	}

	return s.db
}
