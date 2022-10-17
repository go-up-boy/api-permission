package users

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"api-permission/internal/svc"
	"api-permission/internal/types"

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

func (l *ListLogic) List(req *types.ListCommonRequest) (resp *types.ListUserResponse, err error) {
	// todo: add your logic here and delete this line
	var userItems []types.ListUserItem
	p := types.ListPageItem{
		Page: req.Page,
		Limit: req.Limit,
	}
	users, err := l.svcCtx.UserModel.PaginationList(l.ctx, &p)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	for _, v := range users{
		userItems = append(userItems, types.ListUserItem{
			Id: v.Id,
			HeadImg: v.Headimg,
			Username: v.Username,
			CreateAt: v.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt: v.UpdateAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return &types.ListUserResponse{
		Items: userItems,
		Paging: p,
	}, nil
}
