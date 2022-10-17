package roles

import (
	"context"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListCommonRequest) (resp *types.ListRoleResponse, err error) {
	// todo: add your logic here and delete this line
	p := types.ListPageItem{
		Limit: req.Limit,
		Page: req.Page,
	}
	roles, err := l.svcCtx.RoleModel.PaginationList(l.ctx, &p)
	if err != nil {
		return nil, errorsx.NewDefaultError(err.Error())
	}
	var roleResp []types.ListRoleItem
	for _, role := range roles{
		roleResp = append(roleResp, types.ListRoleItem{
			Id: role.Id,
			Name: role.Name,
			Pid: role.Pid,
			MaxRole: role.MaxRole,
			Slug: role.Slug,
			CreateAt: role.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt: role.UpdateAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return &types.ListRoleResponse{
		Items: roleResp,
		Paging: p,
	}, nil
}
