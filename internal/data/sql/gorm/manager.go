package sqlgorm

import (
	"context"

	"gorm.io/gorm"
)

type GormSQLManager struct {
	db *gorm.DB
}

func NewGormSQLManager(db *gorm.DB) *GormSQLManager {
	return &GormSQLManager{db}
}

// RunInTransaction will inject the database transaction into the context
// and return an error if failed to open transaction or the function in parameter returns error
// To use the DB transaction, use the DB on the context.
//
// Use GetClientFromContext(context) function to retrieve the database transaction from the context.
func (s *GormSQLManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	db := s.db.Begin()
	if db.Error != nil {
		return db.Error
	}
	ctx = injectClientToContext(ctx, db)

	err := fn(ctx)
	if err != nil {
		db.Rollback()
		return err
	}

	db.Commit()
	return nil
}
