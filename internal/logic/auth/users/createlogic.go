package users

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"
	"api-permission/internal/utils/hashx"
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

func (l *CreateLogic) Create(req *types.CreateUserRequest) (resp *types.CreateUserResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != sqlc.ErrNotFound {
		return nil, errorsx.NewDefaultError("user already registered")
	}
	enPwd := hashx.BcryptHash(req.Password)
	if  enPwd == "" {
		return nil, errorsx.NewDefaultError("password encrypt error, please try again")
	}
	user := model.Users{
		Headimg: req.HeadImg,
		Username: req.Username,
		Password: enPwd,
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &user)
	if err != nil {
		return nil, errorsx.NewDefaultError("create fail")
	}
	lastInsertId, _ := result.LastInsertId()
	l.svcCtx.UsersRoleModel.UpdateUserRoles(l.ctx, lastInsertId, req.RoleIds)
	if err != nil {
		return nil, errorsx.NewDefaultError(err.Error())
	}
	return &types.CreateUserResponse{
		Id: lastInsertId,
		HeadImg: user.Headimg,
		Username: user.Username,
	}, nil
}
