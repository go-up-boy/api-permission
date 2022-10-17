package rules

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

func (l *ViewLogic) View(req *types.ViewRuleRequest) (resp *types.ViewRuleResponse, err error) {
	// todo: add your logic here and delete this line
	rule, err := l.svcCtx.RuleModel.FindOne(l.ctx, req.Id)
	if err != nil || err == sqlx.ErrNotFound {
		return resp, errorsx.NewDefaultError("rule does not exist")
	}
	return &types.ViewRuleResponse{
		Id: rule.Id,
		Uri: rule.Uri,
		Name: rule.Name,
		Slug: rule.Slug,
		CreateAt: rule.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: rule.UpdateAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
