package roles

import (
	"context"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewLogic {
	return &ViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewLogic) View(req *types.ViewRoleRequest) (resp *types.ViewRoleResponse, err error) {
	// todo: add your logic here and delete this line
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorsx.NewDefaultError(err.Error())
	}
	return &types.ViewRoleResponse{
		Id: role.Id,
		Name: role.Name,
		Pid: role.Pid,
		MaxRole: role.MaxRole,
		Slug: role.Slug,
		CreateAt: role.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: role.UpdateAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
