package repository

import (
	"github.com/Ankr-network/kit/mlog"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"go.uber.org/zap"
)

var (
	MysqlClient *DB
	log         = mlog.Logger("repository")

	GatewayRepo Gateway
	RbacRepo    Rbac
	UserRepo    User
)

func InitGlobal(dsn string) {
	MysqlClient = NewDB(dsn)
	err := MysqlClient.AutoMigrate(model.GatewayMetrics{}, model.Role{}, model.Rule{}, model.User{})
	if err != nil {
		log.Fatal("db migrate", zap.Error(err))
	}

	GatewayRepo = NewGateway(MysqlClient)
	RbacRepo = NewRbac(MysqlClient)
	UserRepo = NewUser(MysqlClient)
}
