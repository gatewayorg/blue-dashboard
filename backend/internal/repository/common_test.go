package repository

import (
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testRepo    *DB
	dsn         = "root:123@tcp(127.0.0.1:3306)/dashboard?charset=utf8mb4&parseTime=True"
	testGateway Gateway
	testRbac    Rbac
	testUser    User
)

func TestMain(m *testing.M) {
	testRepo = NewDB(dsn)
	testRepo.Debug()

	testGateway = NewGateway(testRepo)
	testRbac = NewRbac(testRepo)
	testUser = NewUser(testRepo)
	m.Run()

}

func TestAutoMigrate(t *testing.T) {
	err := testRepo.AutoMigrate(&model.GatewayMetrics{}, &model.Rule{}, &model.Role{}, &model.User{})
	assert.NoError(t, err)
}
