package rules

import (
	"context"

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

func (l *ListLogic) List(req *types.ListCommonRequest) (resp *types.ListRuleResponse, err error) {
	// todo: add your logic here and delete this line
	var ruleItems []types.ListRuleItem
	p := types.ListPageItem{
		Limit: req.Limit,
		Page: req.Page,
	}
	rules, err := l.svcCtx.RuleModel.PaginationList(l.ctx, &p)
	for _, rule := range rules{
		ruleItems = append(ruleItems, types.ListRuleItem{
			Id: rule.Id,
			Uri: rule.Uri,
			Name: rule.Name,
			Slug: rule.Slug,
			CreateAt: rule.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt: rule.UpdateAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return &types.ListRuleResponse{
		Items: ruleItems,
		Paging: p,
	}, nil
}
