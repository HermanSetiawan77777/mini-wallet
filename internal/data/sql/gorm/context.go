package sqlgorm

import (
	"context"

	"gorm.io/gorm"
)

type Key string

const gormDBKey Key = "gorm_database_client"

func injectClientToContext(parent context.Context, db *gorm.DB) context.Context {
	ctx := context.WithValue(parent, gormDBKey, db)
	return ctx
}

// GetClientFromContext will return a gorm client from the context.
//
// The returned client will be nil if the client in context is not found.
func GetClientFromContext(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(gormDBKey).(*gorm.DB)
	if !ok {
		return nil
	}
	if db == nil {
		return nil
	}

	return db
}
