package repository

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

type txKey struct{}

func NewDB(dsn string) *DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("connect mysql", zap.Error(err))
	}
	return &DB{DB: db}
}

func (d *DB) ContextWithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func (d *DB) GetTxFromContext(ctx context.Context) *gorm.DB {
	out, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return d.DB
	}
	return out
}
