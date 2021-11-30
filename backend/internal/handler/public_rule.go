package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/api/protos/rbac"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
)

type PublicRule struct {
	repo repository.Rbac
	svc  service.Rbac
}

func NewPublicRule(repo repository.Rbac, svc service.Rbac) *PublicRule {
	return &PublicRule{
		repo: repo,
		svc:  svc,
	}
}

func (p *PublicRule) GetRule(ctx context.Context, req *rbac.GetRuleReq) (*rbac.GetRuleResp, error) {
	total, data, err := p.svc.GetRuleList(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		log.Error("service GetRuleList", zap.Error(err))
		return nil, handleErr(err)
	}
	return &rbac.GetRuleResp{
		Total: uint64(total),
		Data:  rulesToPb(data),
	}, nil
}

func (p *PublicRule) SetDetail(ctx context.Context, req *rbac.SetDetailReq) (*empty.Empty, error) {
	err := p.repo.SetRuleDetail(ctx, req.Id, req.Detail)
	if err != nil {
		log.Error("repo SetRuleDetail", zap.Error(err))
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func rulesToPb(in []*model.Rule) []*rbac.Rule {
	out := make([]*rbac.Rule, 0, len(in))
	for _, v := range in {
		out = append(out, ruleToPb(v))
	}
	return out
}

func ruleToPb(in *model.Rule) *rbac.Rule {
	return &rbac.Rule{
		Id:      in.ID,
		Service: in.Service,
		Method:  in.Method,
		Detail:  in.Detail,
	}
}
