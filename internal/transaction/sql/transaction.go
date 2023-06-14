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

type TransactionWallet struct {
	TransactionId   string `gorm:"column:transaction_id"`
	WalletId        string `gorm:"column:wallet_id"`
	Status          int    `gorm:"column:status"`
	TransactionType string `gorm:"column:transaction_type"`
	TransactionBy   string `gorm:"column:transaction_by"`
	Amount          int    `gorm:"column:amount"`
	ReferenceId     string `gorm:"column:reference_id"`
}

func (TransactionWallet) TableName() string {
	return "transaction_wallet"
}

type ViewTransactionDetailWallet struct {
	TransactionId   string    `gorm:"column:transaction_id"`
	TransactionBy   string    `gorm:"column:transaction_by"`
	Status          string    `gorm:"column:status"`
	TransactionDate time.Time `gorm:"column:transaction_date"`
	Amount          int       `gorm:"column:amount"`
	ReferenceId     string    `gorm:"column:reference_id"`
}

func (ViewTransactionDetailWallet) TableName() string {
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
	err := db.Model(&ViewTransactionWallet{}).Select("transaction_wallet.transaction_id, transaction_wallet.wallet_id, transaction_wallet.transaction_type, transaction_status.status, transaction_wallet.transaction_by, transaction_wallet.transaction_date, transaction_wallet.amount, transaction_wallet.reference_id").Joins("INNER join transaction_status ON transaction_wallet.status=transaction_status.status_id").Where("transaction_wallet.wallet_id = ?", walletId).Order("transaction_wallet.transaction_date asc").Scan(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return ToServiceSliceTransactionWallet(data), nil
}

func newTransactionFromServiceModel(data *transaction.TransactionWallet) *TransactionWallet {
	if data == nil {
		return nil
	}

	return &TransactionWallet{
		WalletId:        data.WalletId,
		TransactionId:   data.TransactionId,
		Status:          data.Status,
		TransactionType: data.TransactionType,
		TransactionBy:   data.TransactionBy,
		Amount:          data.Amount,
		ReferenceId:     data.ReferenceId,
	}
}

func (s *TransactionWalletSQLRepository) Create(ctx context.Context, params *transaction.TransactionWallet) error {
	payload := newTransactionFromServiceModel(params)
	errs := s.db.Create(&payload).Error
	if errs != nil {
		return errs
	}

	return nil
}

func (v *ViewTransactionDetailWallet) ToViewDetailTransactionDepositModel() *transaction.ViewTransactionDepositWallet {
	return &transaction.ViewTransactionDepositWallet{
		TransactionId:   v.TransactionId,
		TransactionBy:   v.TransactionBy,
		Status:          v.Status,
		TransactionDate: v.TransactionDate,
		Amount:          v.Amount,
		ReferenceId:     v.ReferenceId,
	}
}

func (s *TransactionWalletSQLRepository) GetViewDetailDepositTransaction(ctx context.Context, transactionId string) (*transaction.ViewTransactionDepositWallet, error) {
	var data *ViewTransactionDetailWallet
	db := s.getDatabaseClient(ctx)
	if transactionId == "" {
		return nil, nil
	}
	err := db.Model(&ViewTransactionDetailWallet{}).Select("transaction_wallet.transaction_id, transaction_status.status, transaction_wallet.transaction_by, transaction_wallet.transaction_date, transaction_wallet.amount, transaction_wallet.reference_id").Joins("INNER join transaction_status ON transaction_wallet.status=transaction_status.status_id").Where("transaction_wallet.transaction_id = ? ", transactionId).Scan(&data).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data.ToViewDetailTransactionDepositModel(), nil
}

func (v *ViewTransactionDetailWallet) ToViewDetailTransactionWithdrawalModel() *transaction.ViewTransactionWithdrawalWallet {
	return &transaction.ViewTransactionWithdrawalWallet{
		TransactionId:   v.TransactionId,
		TransactionBy:   v.TransactionBy,
		Status:          v.Status,
		TransactionDate: v.TransactionDate,
		Amount:          v.Amount,
		ReferenceId:     v.ReferenceId,
	}
}

func (s *TransactionWalletSQLRepository) GetViewDetailWithdrawalTransaction(ctx context.Context, transactionId string) (*transaction.ViewTransactionWithdrawalWallet, error) {
	var data *ViewTransactionDetailWallet
	db := s.getDatabaseClient(ctx)
	if transactionId == "" {
		return nil, nil
	}
	err := db.Model(&ViewTransactionDetailWallet{}).Select("transaction_wallet.transaction_id, transaction_status.status, transaction_wallet.transaction_by, transaction_wallet.transaction_date, transaction_wallet.amount, transaction_wallet.reference_id").Joins("INNER join transaction_status ON transaction_wallet.status=transaction_status.status_id").Where("transaction_wallet.transaction_id = ? ", transactionId).Scan(&data).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data.ToViewDetailTransactionWithdrawalModel(), nil
}

func (s *TransactionWalletSQLRepository) getDatabaseClient(ctx context.Context) *gorm.DB {
	db := sqlgorm.GetClientFromContext(ctx)
	if db != nil {
		return db
	}

	return s.db
}
