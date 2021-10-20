package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/api/protos/rbac"
	"github.com/golang/protobuf/ptypes/empty"
)

type PublicRule struct {
}

func (p *PublicRule) GetList(ctx context.Context, _ *empty.Empty) (*rbac.GetRuleResp, error) {
	return nil, nil
}
func (p *PublicRule) Del(ctx context.Context, req *rbac.DelRuleReq) (*empty.Empty, error) {
	return nil, nil
}
