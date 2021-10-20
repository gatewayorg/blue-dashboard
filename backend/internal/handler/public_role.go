package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/api/protos/rbac"
	"github.com/golang/protobuf/ptypes/empty"
)

type PublicRole struct {
}

func (p *PublicRole) GetList(ctx context.Context, _ *empty.Empty) (*rbac.GetListResp, error) {
	return nil, nil
}
func (p *PublicRole) Add(ctx context.Context, req *rbac.AddRoleReq) (*empty.Empty, error) {
	return nil, nil
}
func (p *PublicRole) Update(ctx context.Context, req *rbac.UpdateRoleReq) (*empty.Empty, error) {
	return nil, nil
}
func (p *PublicRole) Del(ctx context.Context, req *rbac.DelRoleReq) (*empty.Empty, error) {
	return nil, nil
}
func (p *PublicRole) SelectRule(ctx context.Context, req *rbac.RoleRule) (*empty.Empty, error) {
	return nil, nil
}
