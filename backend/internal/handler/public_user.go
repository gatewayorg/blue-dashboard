package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/api/protos/user"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"github.com/gatewayorg/blue-dashboard/internal/repository"
	"github.com/gatewayorg/blue-dashboard/internal/service"
	"github.com/gatewayorg/blue-dashboard/pkg/jwt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PublicUser struct {
	repo        repository.User
	svc         service.User
	uidProvider UIDProvider
}

func NewPublicUser(repo repository.User, svc service.User, uidProvider UIDProvider) *PublicUser {
	return &PublicUser{
		repo:        repo,
		svc:         svc,
		uidProvider: uidProvider,
	}
}

func (p *PublicUser) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	id, err := p.svc.Authenticate(ctx, req.Username, req.Password)
	if err != nil {
		return nil, handleErr(ErrAuthorization)
	}
	token, err := jwt.Sign.Sign(id, req.Username)
	if err != nil {
		return nil, handleErr(err)
	}
	return &user.LoginResp{AccessToken: token}, nil
}

func (p *PublicUser) GetList(ctx context.Context, req *user.GetListReq) (*user.GetListResp, error) {
	total, data, err := p.svc.GetList(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, handleErr(err)
	}
	return &user.GetListResp{
		Total: uint64(total),
		Data:  usersToPb(data),
	}, nil
}

func (p *PublicUser) Add(ctx context.Context, req *user.AddReq) (*empty.Empty, error) {
	err := p.svc.Add(ctx, &model.User{
		Username: req.Username,
		Password: req.Passwd,
		Name:     req.Name,
		RoleID:   req.RoleId,
		Enable:   req.Enable,
		CreateAt: time.Now(),
	})
	if err != nil {
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicUser) Update(ctx context.Context, req *user.UpdateReq) (*empty.Empty, error) {
	err := p.svc.Save(ctx, req.Id, &model.UserSave{
		Name:   req.Name,
		RoleID: req.RoleId,
		Enable: req.Enable,
	})
	if err != nil {
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

// undone
func (p *PublicUser) UpdatePwd(ctx context.Context, req *user.UpdatePwdReq) (*empty.Empty, error) {
	err := p.svc.UpdatePwd(ctx, 1, req.OldPasswd, req.NewPasswd)
	if err != nil {
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicUser) SetStatus(ctx context.Context, req *user.SetStatusReq) (*empty.Empty, error) {
	err := p.repo.SetStatus(ctx, req.Id, req.Enable)
	if err != nil {
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicUser) SelectRole(ctx context.Context, req *user.SelectRoleReq) (*empty.Empty, error) {
	err := p.svc.SelectRole(ctx, req.Id, req.RoleId)
	if err != nil {
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func (p *PublicUser) Delete(ctx context.Context, req *user.DeleteReq) (*empty.Empty, error) {
	err := p.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, handleErr(err)
	}
	return &empty.Empty{}, nil
}

func usersToPb(in []*model.UserRole) []*user.User {
	out := make([]*user.User, 0, len(in))
	for _, v := range in {
		out = append(out, userToPb(v))
	}
	return out
}

func userToPb(in *model.UserRole) *user.User {
	return &user.User{
		Id:         in.ID,
		Username:   in.Username,
		Name:       in.Name,
		Enable:     in.Enable,
		CreateAt:   timestamppb.New(in.CreateAt),
		RoleId:     in.RoleID,
		RoleName:   in.RoleName,
		RoleDetail: in.RoleDetail,
	}
}
