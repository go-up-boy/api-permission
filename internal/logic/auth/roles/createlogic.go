package roles

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"
	"api-permission/model"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRoleRequest) (resp *types.CreateRoleResponse, err error) {
	// todo: add your logic here and delete this line
	role := &model.Roles{
		Name: req.Name,
		Pid: req.Pid,
		MaxRole: req.MaxRole,
		Slug: req.Slug,
	}
	result, err := l.svcCtx.RoleModel.Insert(l.ctx, role)
	if err != nil {
		return resp, errorsx.NewDefaultError(err.Error())
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return resp, errorsx.NewDefaultError(err.Error())
	}
	return &types.CreateRoleResponse{
		Id: lastId,
		Name: role.Name,
		Pid: role.Pid,
		MaxRole: role.MaxRole,
		Slug: req.Slug,
	}, nil
}
