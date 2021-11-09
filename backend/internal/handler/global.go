package handler

import (
	"context"
	"github.com/Ankr-network/kit/mlog"
	"github.com/gatewayorg/blue-dashboard/api/protos/index"
	"github.com/gatewayorg/blue-dashboard/api/protos/rbac"
	"github.com/gatewayorg/blue-dashboard/api/protos/user"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	log                = mlog.Logger("handler")
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
