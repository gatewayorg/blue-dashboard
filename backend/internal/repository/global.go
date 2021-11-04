package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ankr-network/kit/mlog"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/pkg/password"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"time"
)

var (
	log = mlog.Logger("repository")

	MysqlClient   *DB
	RuleIDMap     = make(map[string]uint64)

	GatewayRepo Gateway
	RbacRepo    Rbac
	UserRepo    User
)

func GlobalInit(dsn string) {
	MysqlClient = NewDB(dsn)
	err := MysqlClient.AutoMigrate(model.GatewayMetrics{}, model.Role{}, model.Rule{}, model.User{})
	if err != nil {
		log.Fatal("db migrate", zap.Error(err))
	}

	GatewayRepo = NewGateway(MysqlClient)
	RbacRepo = NewRbac(MysqlClient)
	UserRepo = NewUser(MysqlClient)
}

func RegisterRule(s *grpc.Server, exempts ...string) {
	ctx := context.Background()
	rules, err := RbacRepo.GetRuleAll(ctx)
	if err != nil {
		log.Fatal("repo GetRuleAll", zap.Error(err))
	}

	rulesMap := make(map[string]struct{}, len(rules))
	for _, rule := range rules {
		rulesMap[fmt.Sprintf("/%s/%s", rule.Service, rule.Method)] = struct{}{}
	}

	exemptMap := make(map[string]struct{}, len(exempts))
	for _, exempt := range exempts {
		exemptMap[exempt] = struct{}{}
	}

	serviceInfo := s.GetServiceInfo()
	for svc, info := range serviceInfo {
		for _, method := range info.Methods {
			fullMethod := fmt.Sprintf("/%s/%s", svc, method.Name)
			if _, ok := exemptMap[fullMethod]; ok {
				continue
			}
			if _, ok := rulesMap[fullMethod]; !ok {
				rule := &model.Rule{
					Service: svc,
					Method:  method.Name,
				}
				if err = RbacRepo.AddRule(ctx, rule); err != nil {
					log.Fatal("repo AddRule", zap.Any("data", rule), zap.Error(err))
				}
			}
		}
	}

	rules, err = RbacRepo.GetRuleAll(ctx)
	if err != nil {
		log.Fatal("repo GetRuleAll", zap.Error(err))
	}
	for _, v := range rules {
		RuleIDMap[fmt.Sprintf("/%s/%s", v.Service, v.Method)] = v.ID
	}

}

func RegisterSuperRole() {
	ctx := context.Background()

	ruleIDs := make([]uint64, 0, len(RuleIDMap))
	for _, id := range RuleIDMap {
		ruleIDs = append(ruleIDs, id)
	}

	role, err := RbacRepo.FindRoleById(ctx, 1)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			role = &model.Role{
				ID:       1,
				Name:     "admin",
				Enable:   true,
				RuleIDs:  ruleIDs,
				CreateAt: time.Now(),
			}
			if _, err = RbacRepo.AddRole(ctx, role); err != nil {
				log.Fatal("add role", zap.Any("data", role), zap.Error(err))
			}
		} else {
			log.Fatal("find role by id ", zap.Uint64("id", 1), zap.Error(err))
		}
	}
	err = RbacRepo.SelectRole(ctx, role.ID, ruleIDs)
	if err != nil {
		log.Fatal("select role ", zap.Error(err))
	}
}

func RegisterSuperUser(username string, passwd string) {
	if username == "" || passwd == "" {
		return
	}
	ctx := context.Background()
	hash, err := password.Hash(passwd)
	if err != nil {
		log.Fatal(" encryption passwd", zap.String("passwd", passwd), zap.Error(err))
	}
	var userInfo *model.User
	userInfo, err = UserRepo.FindByID(ctx, 1)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userInfo = &model.User{
				ID:       1,
				Username: username,
				Password: hash,
				Name:     username,
				RoleID:   1,
				Enable:   true,
				CreateAt: time.Now(),
			}
			if _, err = UserRepo.Add(ctx, userInfo); err != nil {
				log.Fatal("add user", zap.Any("data", userInfo), zap.Error(err))
			}
		} else {
			log.Fatal("find user by id ", zap.Uint64("id", 1), zap.Error(err))
		}
	} else {
		userInfo.Username = username
		userInfo.Password = hash
		userInfo.Name = username
		if err = UserRepo.Save(ctx, userInfo); err != nil {
			log.Fatal("save user", zap.Any("data", userInfo), zap.Error(err))
		}
	}
}
