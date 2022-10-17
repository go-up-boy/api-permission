package roles

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"api-permission/model"

	"api-permission/internal/svc"
	"api-permission/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRulesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRulesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRulesLogic {
	return &AddRulesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRulesLogic) AddRules(req *types.AddRulesRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	for _, ruleId := range req.Rules{
		_, err = l.svcCtx.RoleRulesModel.FindOneByRoleIdRuleId(l.ctx, req.RoleId, ruleId)
		if err == sqlx.ErrNotFound {
			data := model.RoleRulesRelation{
				RoleId: req.RoleId,
				RuleId: ruleId,
			}
			l.svcCtx.RoleRulesModel.Insert(l.ctx, &data)
		}
	}
	return &types.CommonResponse{
		Success: true,
	}, nil
}
