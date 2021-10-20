package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/api/protos/admin"
	"github.com/golang/protobuf/ptypes/empty"
)

type PublicAdmin struct{}

func (p *PublicAdmin) CreateUri(ctx context.Context, req *admin.CreateUrlReq) (*empty.Empty, error) {
	return nil, nil
}

func (p *PublicAdmin) GetUriList(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return nil, nil
}
