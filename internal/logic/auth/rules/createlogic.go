package rules

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

func (l *CreateLogic) Create(req *types.AddRuleRequest) (resp *types.AddRuleResponse, err error) {
	// todo: add your logic here and delete this line
	var rule  = model.Rules{
		Uri: req.Uri,
		Name: req.Name,
		Slug: req.Slug,
	}
	result, err := l.svcCtx.RuleModel.Insert(l.ctx, &rule)
	if err != nil {
		return nil, errorsx.NewDefaultError("create fail: slug may be repeated")
	}
	lastId, err := result.RowsAffected()
	if err != nil || lastId <= 0 {
		return nil, errorsx.NewDefaultError("create fail")
	}
	return &types.AddRuleResponse{
		Uri: rule.Uri,
		Name: rule.Name,
		Slug: req.Slug,
	}, nil
}
