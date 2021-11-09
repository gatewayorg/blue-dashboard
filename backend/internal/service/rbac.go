package service

import (
	"context"
	"errors"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
)

var (
	ErrUserInUsing  = errors.New("user is using")
	ErrRoleDisable  = errors.New("role is disabled")
	ErrRoleNotFound = errors.New("role not found")
)

type Rbac interface {
	GetRuleList(ctx context.Context, page, pageSize int) (total int64, data []*model.Rule, err error)
	GetRoleList(ctx context.Context, page, pageSize int) (total int64, data []*model.Role, err error)
	SaveRole(ctx context.Context, id uint64, role *model.RoleSave) error
	DeleteRole(ctx context.Context, id uint64) error
}

type rbacImpl struct {
	repo     repository.Rbac
	userRepo repository.User
}

func NewRbac(repo repository.Rbac, userRepo repository.User) *rbacImpl {
	return &rbacImpl{repo: repo, userRepo: userRepo}
}

func (r *rbacImpl) GetRuleList(ctx context.Context, page, pageSize int) (total int64, data []*model.Rule, err error) {
	data, err = r.repo.GetRule(ctx, page, pageSize)
	if err != nil {
		return 0, nil, err
	}
	total, err = r.repo.GetRuleCount(ctx)
	if err != nil {
		return 0, nil, err
	}
	return total, data, nil
}

func (r *rbacImpl) GetRoleList(ctx context.Context, page, pageSize int) (total int64, data []*model.Role, err error) {
	data, err = r.repo.GetRole(ctx, page, pageSize)
	if err != nil {
		return 0, nil, err
	}
	total, err = r.repo.GetRoleCount(ctx)
	if err != nil {
		return 0, nil, err
	}
	return total, data, nil
}

func (r *rbacImpl) SaveRole(ctx context.Context, id uint64, role *model.RoleSave) error {
	roleInfo, err := r.repo.FindRoleById(ctx, id)
	if err != nil {
		return err
	}
	roleInfo.Enable = role.Enable
	roleInfo.Name = role.Name
	roleInfo.Detail = role.Detail
	return r.repo.SaveRole(ctx, roleInfo)
}

func (r *rbacImpl) DeleteRole(ctx context.Context, id uint64) error {
	num, err := r.userRepo.GetCountByRoleId(ctx, id)
	if err != nil {
		return err
	}
	if num > 0 {
		return ErrUserInUsing
	}
	return r.repo.DeleteRole(ctx, id)
}
