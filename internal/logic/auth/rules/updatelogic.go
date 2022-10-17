package rules

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"

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

func (l *UpdateLogic) Update(req *types.UpdateRuleRequest) (resp *types.UpdateRuleResponse, err error) {
	// todo: add your logic here and delete this line
	rule, err := l.svcCtx.RuleModel.FindOne(l.ctx, req.Id)
	if err != nil || err == sqlx.ErrNotFound {
		return nil, err
	}
	rule.Name = req.Name
	rule.Uri  = req.Uri
	rule.Slug  = req.Slug
	err = l.svcCtx.RuleModel.Update(l.ctx, rule)
	if err != nil || err == sqlx.ErrNotFound {
		return nil, errorsx.NewDefaultError(err.Error())
	}
	return &types.UpdateRuleResponse{
		Id: rule.Id,
		Uri: rule.Uri,
		Name: rule.Name,
		Slug: rule.Slug,
		CreateAt: rule.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: rule.UpdateAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
