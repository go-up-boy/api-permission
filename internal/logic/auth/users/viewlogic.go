package users

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"api-permission/internal/utils/errorsx"

	"api-permission/internal/svc"
	"api-permission/internal/types"

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

func (l *ViewLogic) View(req *types.ViewUserRequest) (resp *types.ViewUserResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil || err == sqlc.ErrNotFound {
		return nil, errorsx.NewDefaultError("user does not exist")
	}
	return &types.ViewUserResponse{
		Id: user.Id,
		HeadImg: user.Headimg,
		Username: user.Username,
		CreateAt: user.CreateAt.String(),
		UpdateAt: user.UpdateAt.Time.String(),
	}, nil
}
