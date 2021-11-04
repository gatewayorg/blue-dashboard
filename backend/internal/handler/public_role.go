package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/api/protos/rbac"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PublicRole struct {
	repo repository.Rbac
	svc  service.Rbac
}

func NewPublicRole(repo repository.Rbac, svc service.Rbac) *PublicRole {
	return &PublicRole{
		repo: repo,
		svc:  svc,
	}
}

func (p *PublicRole) GetRole(ctx context.Context, req *rbac.GetRoleReq) (*rbac.GetRoleResp, error) {
	total, data, err := p.svc.GetRoleList(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		log.Error("service GetRoleList", zap.Error(err))
		return nil, handleErr(err)
	}
	return &rbac.GetRoleResp{
		Total: uint64(total),
		Data:  rolesToPb(data),
	}, nil
}

func (p *PublicRole) Add(ctx context.Context, req *rbac.AddRoleReq) (*empty.Empty, error) {
	_, err := p.repo.AddRole(ctx, &model.Role{
		Name:     req.Name,
		Detail:   req.Detail,
		Enable:   req.Enable,
		CreateAt: time.Now(),
	})
	if err != nil {
		log.Error("repo AddRole", zap.Any("data", req), zap.Error(err))
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicRole) Update(ctx context.Context, req *rbac.UpdateRoleReq) (*empty.Empty, error) {
	err := p.svc.SaveRole(ctx, req.Id, &model.RoleSave{
		Name:   req.Name,
		Detail: req.Detail,
		Enable: req.Enable,
	})
	if err != nil {
		log.Error("repo SaveRole", zap.Any("data", req), zap.Error(err))
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicRole) SetStatus(ctx context.Context, req *rbac.SetStatusReq) (*empty.Empty, error) {
	err := p.repo.SetRoleStatus(ctx, req.Id, req.Enable)
	if err != nil {
		log.Error("repo SetRoleStatus", zap.Any("data", req), zap.Error(err))
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicRole) Del(ctx context.Context, req *rbac.DelRoleReq) (*empty.Empty, error) {
	err := p.svc.DeleteRole(ctx, req.Id)
	if err != nil {
		log.Error("repo DeleteRole", zap.Any("data", req), zap.Error(err))
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicRole) SelectRule(ctx context.Context, req *rbac.RoleRule) (*empty.Empty, error) {
	err := p.repo.SelectRole(ctx, req.RoleId, req.Rules)
	if err != nil {
		log.Error("repo SelectRole", zap.Any("data", req), zap.Error(err))
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func rolesToPb(in []*model.Role) []*rbac.Role {
	out := make([]*rbac.Role, 0, len(in))
	for _, v := range in {
		out = append(out, roleToPb(v))
	}
	return out
}

func roleToPb(in *model.Role) *rbac.Role {
	return &rbac.Role{
		Id:         in.ID,
		Name:       in.Name,
		Detail:     in.Detail,
		CreateTime: timestamppb.New(in.CreateAt),
		Enable:     in.Enable,
		RuleIds:    in.RuleIDs,
	}
}
