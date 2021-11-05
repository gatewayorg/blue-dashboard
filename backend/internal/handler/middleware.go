package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func Authorization(whiteList ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		for _, v := range whiteList {
			if v == info.FullMethod {
				return handler(ctx, req)
			}
		}

		ruleId, ok := repository.RuleIDMap[info.FullMethod]
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
		if !user.Enable {
			return nil, handleErr(service.ErrUserDisable)
		}

		role, err := repository.RbacRepo.FindRoleById(ctx, user.RoleID)
		if err != nil {
			log.Error("role find by Id", zap.Uint64("id", id), zap.Error(err))
			return nil, handleErr(err)
		}

		if !role.Enable {
			return nil, handleErr(service.ErrRoleDisable)
		}

		for _, v := range role.RuleIDs {
			if v == ruleId {
				return handler(ctx, req)
			}
		}
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}
}

func RecordRequestUrl(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("request data",
			zap.String("mothod", r.Method),
			zap.Any("url", r.URL),
			zap.Any("header", r.Header),
		)
		handler.ServeHTTP(w, r)
	})
}
