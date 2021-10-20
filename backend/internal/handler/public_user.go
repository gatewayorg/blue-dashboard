package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/api/protos/user"
	"github.com/golang/protobuf/ptypes/empty"
)

type PublicUser struct {
}

func (p *PublicUser) Login(ctx context.Context, req *user.LoginReq) (*empty.Empty, error) {
	return nil, nil
}
func (p *PublicUser) SelectRole(ctx context.Context, req *user.SelectRoleReq) (*empty.Empty, error) {
	return nil, nil
}
func (p *PublicUser) GetList(ctx context.Context, _ *empty.Empty) (*user.GetListResp, error) {
	return nil, nil
}
