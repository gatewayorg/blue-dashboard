package repository

import (
	"github.com/Ankr-network/kit/mlog"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlClient *gorm.DB
	log         = mlog.Logger("repository")

	GatewayRepo Gateway
)

func InitGlobal(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("connect mysql", zap.Error(err))
	}
	db.AutoMigrate(model.GatewayMetrics{})

	GatewayRepo = NewGateway(db)
}
