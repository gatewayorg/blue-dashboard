package repository

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
)

type User interface {
	GetAllCount(ctx context.Context) (int64, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*model.UserRole, error)
	GetCountByRoleId(ctx context.Context, roleID uint64) (int64, error)
	Add(ctx context.Context, user *model.User) error
	Save(ctx context.Context, user *model.User) error
	UpdatePwd(ctx context.Context, id uint64, passwd string) error
	SetStatus(ctx context.Context, id uint64, enable bool) error
	SelectRole(ctx context.Context, id uint64, roleID uint64) error
	FindByUserName(ctx context.Context, username string) (*model.User, error)
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	Delete(ctx context.Context, id uint64) error
}

type userImpl struct {
	*DB
}

func NewUser(db *DB) *userImpl {
	return &userImpl{DB: db}
}

func (u *userImpl) GetAllCount(ctx context.Context) (num int64, err error) {
	err = u.GetTxFromContext(ctx).Model(&model.User{}).Count(&num).Error
	return
}

func (u *userImpl) GetAll(ctx context.Context, page, pageSize int) (userRoles []*model.UserRole, err error) {
	userRoles = []*model.UserRole{}
	err = u.GetTxFromContext(ctx).Model(&model.User{}).Select("users.*, roles.id role_id, roles.name role_name, roles.detail role_detail").
		Joins("left join roles on users.role_id = roles.id").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Scan(&userRoles).Error
	return userRoles, err
}

func (u *userImpl) GetCountByRoleId(ctx context.Context, roleID uint64) (num int64, err error) {
	res := u.GetTxFromContext(ctx).Model(&model.User{}).Where("role_id = ?", roleID).Count(&num)
	if res.Error != nil {
		return 0, res.Error
	}
	return num, nil
}

func (u *userImpl) Add(ctx context.Context, user *model.User) error {
	return u.GetTxFromContext(ctx).Create(user).Error
}

func (u *userImpl) Save(ctx context.Context, user *model.User) error {
	return u.GetTxFromContext(ctx).Save(user).Error
}

func (u *userImpl) UpdatePwd(ctx context.Context, id uint64, passwd string) error {
	return u.GetTxFromContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("password", passwd).Error
}

func (u *userImpl) SetStatus(ctx context.Context, id uint64, enable bool) error {
	return u.GetTxFromContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("enable", enable).Error
}

func (u *userImpl) SelectRole(ctx context.Context, id uint64, roleID uint64) error {
	return u.GetTxFromContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("role_id", roleID).Error
}

func (u *userImpl) FindByUserName(ctx context.Context, username string) (user *model.User, err error) {
	user = new(model.User)
	err = u.GetTxFromContext(ctx).First(&user, "username = ?", username).Error
	return user, err
}

func (u *userImpl) FindByID(ctx context.Context, id uint64) (user *model.User, err error) {
	user = new(model.User)
	err = u.GetTxFromContext(ctx).First(&user, "id = ?", id).Error
	return user, err
}

func (u *userImpl) Delete(ctx context.Context, id uint64) error {
	return u.GetTxFromContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}
