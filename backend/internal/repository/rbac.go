package repository

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/pkg/fields"
)

type Rbac interface {
	GetRoleAll(ctx context.Context) ([]*model.Role, error)
	GetRole(ctx context.Context, page, pageSize int) ([]*model.Role, error)
	GetRoleCount(ctx context.Context) (int64, error)
	AddRole(ctx context.Context, role *model.Role) (id uint64, err error)
	FindRoleById(ctx context.Context, id uint64) (*model.Role, error)
	SaveRole(ctx context.Context, role *model.Role) error
	SelectRole(ctx context.Context, id uint64, rules []uint64) error
	SetRoleStatus(ctx context.Context, id uint64, enable bool) error
	DeleteRole(ctx context.Context, id uint64) error

	GetRule(ctx context.Context, page, pageSize int) ([]*model.Rule, error)
	GetRuleAll(ctx context.Context) ([]*model.Rule, error)
	GetRuleCount(ctx context.Context) (int64, error)
	AddRule(ctx context.Context, rule *model.Rule) error
	DeleteRule(ctx context.Context, id uint64) error
	SetRuleDetail(ctx context.Context, id uint64, detail string) error
}

type rbacImpl struct {
	*DB
}

func NewRbac(db *DB) *rbacImpl {
	return &rbacImpl{DB: db}
}

func (r *rbacImpl) GetRoleAll(ctx context.Context) ([]*model.Role, error) {
	roles := new([]*model.Role)
	res := r.GetTxFromContext(ctx).Find(roles)
	return *roles, res.Error
}

func (r *rbacImpl) GetRole(ctx context.Context, page, pageSize int) ([]*model.Role, error) {
	roles := new([]*model.Role)
	res := r.GetTxFromContext(ctx).Limit(pageSize).Offset((page - 1) * pageSize).Find(roles)
	return *roles, res.Error
}

func (r *rbacImpl) GetRoleCount(ctx context.Context) (num int64, err error) {
	res := r.GetTxFromContext(ctx).Model(&model.Role{}).Count(&num)
	return num, res.Error
}

func (r *rbacImpl) AddRole(ctx context.Context, role *model.Role) (id uint64, err error) {
	err = r.GetTxFromContext(ctx).Create(role).Error
	return role.ID, err
}

func (r *rbacImpl) FindRoleById(ctx context.Context, id uint64) (*model.Role, error) {
	var role model.Role
	res := r.GetTxFromContext(ctx).Where("id = ?", id).First(&role)
	return &role, res.Error
}

func (r *rbacImpl) SaveRole(ctx context.Context, role *model.Role) error {
	return r.GetTxFromContext(ctx).Save(role).Error
}

func (r *rbacImpl) SelectRole(ctx context.Context, id uint64, ruleIDs []uint64) error {
	return r.GetTxFromContext(ctx).Model(&model.Role{}).Where("id = ?", id).Update("rule_ids", fields.Uint64s(ruleIDs)).Error
}

func (r *rbacImpl) SetRoleStatus(ctx context.Context, id uint64, enable bool) error {
	return r.GetTxFromContext(ctx).Model(&model.Role{}).Where("id = ?", id).Update("enable", enable).Error
}

func (r *rbacImpl) DeleteRole(ctx context.Context, id uint64) error {
	return r.GetTxFromContext(ctx).Delete(&model.Role{}, id).Error
}

// ----- rule ------
func (r *rbacImpl) GetRule(ctx context.Context, page, pageSize int) ([]*model.Rule, error) {
	rules := new([]*model.Rule)
	res := r.GetTxFromContext(ctx).Limit(pageSize).Offset((page - 1) * pageSize).Find(rules)
	return *rules, res.Error
}

func (r *rbacImpl) GetRuleAll(ctx context.Context) ([]*model.Rule, error) {
	rules := new([]*model.Rule)
	res := r.GetTxFromContext(ctx).Find(rules)
	return *rules, res.Error
}

func (r *rbacImpl) GetRuleCount(ctx context.Context) (num int64, err error) {
	res := r.GetTxFromContext(ctx).Model(&model.Rule{}).Count(&num)
	return num, res.Error
}

func (r *rbacImpl) AddRule(ctx context.Context, rule *model.Rule) error {
	return r.GetTxFromContext(ctx).Create(rule).Error
}

func (r *rbacImpl) DeleteRule(ctx context.Context, id uint64) error {
	return r.GetTxFromContext(ctx).Delete(&model.Rule{}, id).Error
}

func (r *rbacImpl) SetRuleDetail(ctx context.Context, id uint64, detail string) error {
	return r.GetTxFromContext(ctx).Model(&model.Rule{}).Where("id = ?", id).Update("detail", detail).Error
}
