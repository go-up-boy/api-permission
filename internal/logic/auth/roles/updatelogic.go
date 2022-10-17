package roles

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"
	"api-permission/model"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRoleRequest) (resp *types.UpdateRoleResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.RoleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return resp, errorsx.NewDefaultError(err.Error())
	}
	if req.Pid != 0 {
		_, err = l.svcCtx.RoleModel.FindOne(l.ctx, req.Pid)
		if err != nil {
			return resp, errorsx.NewDefaultError(err.Error())
		}
	}
	role := &model.Roles{
		Id: req.Id,
		Name: req.Name,
		Pid: req.Pid,
		MaxRole: req.MaxRole,
		Slug: req.Slug,
	}
	if l.svcCtx.RoleModel.Update(l.ctx, role) != nil {
		return nil, errorsx.NewDefaultError(err.Error())
	}
	return &types.UpdateRoleResponse{
		Id: role.Id,
		Name: role.Name,
		Pid: role.Pid,
		MaxRole: role.MaxRole,
		Slug: role.Slug,
	}, nil
}
