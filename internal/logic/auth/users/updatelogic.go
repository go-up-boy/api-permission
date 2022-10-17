package users

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"
	"api-permission/internal/utils/hashx"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UpdateLogic) Update(req *types.UpdateUserRequest) (resp *types.UpdateUserResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil || err == sqlc.ErrNotFound {
		return nil, errorsx.NewDefaultError("user does not exist")
	}
	if req.Password != "" {
		user.Password = hashx.BcryptHash(req.Password)
	}
	if req.HeadImg != "" {
		user.Headimg = req.HeadImg
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if l.svcCtx.UserModel.ExistByUsername(l.ctx, user) {
		return nil, errorsx.NewDefaultError("username already exist")
	}
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, errorsx.NewDefaultError("update fail")
	}
	l.svcCtx.UsersRoleModel.UpdateUserRoles(l.ctx, req.Id, req.RoleIds)
	return &types.UpdateUserResponse{
		Id: user.Id,
		HeadImg: user.Headimg,
		Username: user.Username,
	}, nil
}
