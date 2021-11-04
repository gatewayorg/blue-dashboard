package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ankr-network/kit/mlog"
	"github.com/gatewayorg/blue-dashboard/api/protos/index"
	"github.com/gatewayorg/blue-dashboard/api/protos/rbac"
	"github.com/gatewayorg/blue-dashboard/api/protos/user"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"github.com/gatewayorg/blue-dashboard/pkg/password"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"time"
)

var (
	log                = mlog.Logger("handler")
	RuleIDMap          = make(map[string]uint64)
	uidProvider        UIDProvider
	PublicIndexHandler *PublicIndex
	PublicRoleHandler  *PublicRole
	PublicRuleHandler  *PublicRule
	PublicUserHandler  *PublicUser
)

func GlobalInit() {
	uidProvider = NewUIDProvider()

	PublicIndexHandler = NewPublicIndex(service.MetricsSvc)
	PublicRoleHandler = NewPublicRole(repository.RbacRepo, service.RbacSvc)
	PublicRuleHandler = NewPublicRule(repository.RbacRepo, service.RbacSvc)
	PublicUserHandler = NewPublicUser(repository.UserRepo, service.UserSvc, uidProvider)
}

func RegisterGRPC(s *grpc.Server) {
	index.RegisterPublicIndexServer(s, PublicIndexHandler)
	rbac.RegisterPublicRoleServer(s, PublicRoleHandler)
	rbac.RegisterPublicRuleServer(s, PublicRuleHandler)
	user.RegisterPublicUserServer(s, PublicUserHandler)
}

func MustRegisterREST(mux *runtime.ServeMux, grpcListAddress string) {
	err := index.RegisterPublicIndexHandlerFromEndpoint(context.Background(), mux, "localhost"+grpcListAddress, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("register public index rest error", zap.Error(err))
	}

	err = rbac.RegisterPublicRoleHandlerFromEndpoint(context.Background(), mux, "localhost"+grpcListAddress, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("register public role rest error", zap.Error(err))
	}

	err = rbac.RegisterPublicRuleHandlerFromEndpoint(context.Background(), mux, "localhost"+grpcListAddress, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("register public rule rest error", zap.Error(err))
	}

	err = user.RegisterPublicUserHandlerFromEndpoint(context.Background(), mux, "localhost"+grpcListAddress, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("register public user rest error", zap.Error(err))
	}
}

func RegisterRule(s *grpc.Server, exempts ...string) {
	ctx := context.Background()
	rules, err := repository.RbacRepo.GetRuleAll(ctx)
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
				if err = repository.RbacRepo.AddRule(ctx, rule); err != nil {
					log.Fatal("repo AddRule", zap.Any("data", rule), zap.Error(err))
				}
			}
		}
	}

	rules, err = repository.RbacRepo.GetRuleAll(ctx)
	if err != nil {
		log.Fatal("repo GetRuleAll", zap.Error(err))
	}

	ruleIDs := make([]uint64, 0, len(rules))
	for _, v := range rules {
		RuleIDMap[fmt.Sprintf("/%s/%s", v.Service, v.Method)] = v.ID
		ruleIDs = append(ruleIDs, v.ID)
	}

	role, err := repository.RbacRepo.FindRoleById(ctx, 1)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			role = &model.Role{
				ID:       1,
				Name:     "admin",
				Enable:   true,
				RuleIDs:  ruleIDs,
				CreateAt: time.Now(),
			}
			if err = repository.RbacRepo.AddRole(ctx, role); err != nil {
				log.Fatal("add role", zap.Any("data", role), zap.Error(err))
			}
		} else {
			log.Fatal("find role by id ", zap.Uint64("id", 1), zap.Error(err))
		}
	}
	err = repository.RbacRepo.SelectRole(ctx, role.ID, ruleIDs)
	if err != nil {
		log.Fatal("select role ", zap.Error(err))
	}
}

func RegisterUser(username, passwd string) {
	if username == "" || passwd == "" {
		return
	}
	ctx := context.Background()
	hash, err := password.Hash(passwd)
	if err != nil {
		log.Fatal(" encryption passwd", zap.String("passwd", passwd), zap.Error(err))
	}
	var userInfo *model.User
	userInfo, err = repository.UserRepo.FindByID(ctx, 1)
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
			if err = repository.UserRepo.Add(ctx, userInfo); err != nil {
				log.Fatal("add user", zap.Any("data", userInfo), zap.Error(err))
			}
		} else {
			log.Fatal("find user by id ", zap.Uint64("id", 1), zap.Error(err))
		}
	} else {
		userInfo.Username = username
		userInfo.Password = hash
		userInfo.Name = username
		if err = repository.UserRepo.Save(ctx, userInfo); err != nil {
			log.Fatal("save user", zap.Any("data", userInfo), zap.Error(err))
		}
	}

}
