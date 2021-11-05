package service

import (
	"context"
	"errors"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/pkg/password"
	"go.uber.org/zap"
)

var (
	ErrVerifyPassword = errors.New("password verification failed")
	ErrUserDisable    = errors.New("user is disabled")
)

type User interface {
	GetList(ctx context.Context, page, pageSize int) (total int64, data []*model.UserRole, err error)
	Add(ctx context.Context, user *model.User) (id uint64, err error)
	Save(ctx context.Context, id uint64, user *model.UserSave) error
	UpdatePwd(ctx context.Context, id uint64, oldPasswd, newPasswd string) error
	SelectRole(ctx context.Context, id uint64, roleID uint64) error
	Authenticate(ctx context.Context, username, password string) (ID uint64, err error)
}

type userImpl struct {
	rbacRepo repository.Rbac
	repo     repository.User
}

func NewUser(repo repository.User, rbacRepo repository.Rbac) *userImpl {
	return &userImpl{
		repo:     repo,
		rbacRepo: rbacRepo,
	}
}

func (u *userImpl) GetList(ctx context.Context, page, pageSize int) (total int64, data []*model.UserRole, err error) {
	total, err = u.repo.GetAllCount(ctx)
	if err != nil {
		log.Error("get user all count", zap.Error(err))
		return
	}
	data, err = u.repo.GetList(ctx, page, pageSize)
	if err != nil {
		log.Error("get list", zap.Int("page", page), zap.Int("pagesize", pageSize), zap.Error(err))
		return
	}
	return
}

func (u *userImpl) Add(ctx context.Context, user *model.User) (id uint64, err error) {
	passwd := user.Password
	user.Password, err = password.Hash(passwd)
	if err != nil {
		log.Error("password hash", zap.String("passwd", passwd), zap.Error(err))
		return 0, err
	}

	return u.repo.Add(ctx, user)
}

func (u *userImpl) Save(ctx context.Context, id uint64, user *model.UserSave) error {
	userInfo, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	userInfo.Name = user.Name
	userInfo.Enable = user.Enable
	userInfo.RoleID = user.RoleID
	return u.repo.Save(ctx, userInfo)
}

func (u *userImpl) UpdatePwd(ctx context.Context, id uint64, oldPasswd, newPasswd string) error {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		log.Error("user find by id", zap.Uint64("id", id), zap.Error(err))
		return err
	}
	err = password.Verify(user.Password, oldPasswd)
	if err != nil {
		return ErrVerifyPassword
	}

	pwd, err := password.Hash(newPasswd)
	if err != nil {
		log.Error("password hash", zap.String("passwd", newPasswd), zap.Error(err))
		return err
	}
	err = u.repo.UpdatePwd(ctx, id, pwd)
	return err
}

func (u *userImpl) SelectRole(ctx context.Context, id uint64, roleID uint64) error {
	_, err := u.rbacRepo.FindRoleById(ctx, roleID)
	if err != nil {
		log.Error("find role by id", zap.Uint64("id", roleID), zap.Error(err))
		return err
	}
	return u.repo.SelectRole(ctx, id, roleID)
}

func (u *userImpl) Authenticate(ctx context.Context, username, passwd string) (ID uint64, err error) {
	user, err := u.repo.FindByUserName(ctx, username)
	if err != nil {
		return 0, err
	}

	if !user.Enable {
		return 0, ErrUserDisable
	}

	if err = password.Verify(user.Password, passwd); err != nil {
		return 0, err
	}

	return user.ID, nil
}
