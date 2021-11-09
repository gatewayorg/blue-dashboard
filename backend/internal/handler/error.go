package handler

import (
	"errors"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var (
	ErrInvalidContext = errors.New("context without authentication")
	ErrAuthorization  = errors.New("authentication failed")
)

func handleErr(err error) error {
	switch err {
	case service.ErrUserInUsing:
		return status.Error(codes.AlreadyExists, err.Error())
	case ErrInvalidContext:
		return status.Errorf(codes.Unauthenticated, "InvalidContext:%v", err)
	case ErrAuthorization:
		return status.Errorf(codes.Unauthenticated, "%v", err)
	case gorm.ErrRecordNotFound:
		return status.Errorf(codes.NotFound, "not found")
	case service.ErrUserDisable:
		return status.Errorf(codes.Unauthenticated, "%v", err)
	case service.ErrRoleDisable:
		return status.Errorf(codes.Unauthenticated, "%v", err)
	case service.ErrVerifyPassword:
		return status.Errorf(codes.InvalidArgument, "%v", err)
	case service.ErrRoleNotFound:
		return status.Errorf(codes.NotFound, "%v", err)
	}
	return status.Errorf(codes.Internal, "Internal Error")
}
