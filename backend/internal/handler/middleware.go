package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Authorization(whiteList ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		for _, v := range whiteList {
			if v == info.FullMethod {
				return handler(ctx, req)
			}
		}

		ruleId, ok := RuleIDMap[info.FullMethod]
		if !ok {
			return handler(ctx, req)
		}

		id, err := uidProvider.UID(ctx)
		if err != nil {
			return nil, handleErr(err)
		}

		user, err := repository.UserRepo.FindByID(context.Background(), id)
		if err != nil {
			log.Error("user find by Id", zap.Uint64("id", id), zap.Error(err))
			return nil, handleErr(err)
		}

		role, err := repository.RbacRepo.FindRoleById(ctx, user.RoleID)
		if err != nil {
			log.Error("role find by Id", zap.Uint64("id", id), zap.Error(err))
			return nil, handleErr(err)
		}

		for _, v := range role.RuleIDs {
			if v == ruleId {
				return handler(ctx, req)
			}
		}
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}
}
