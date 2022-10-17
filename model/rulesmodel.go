package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RulesModel = (*customRulesModel)(nil)

type (
	// RulesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRulesModel.
	RulesModel interface {
		rulesModel
	}

	customRulesModel struct {
		*defaultRulesModel
	}
)

// NewRulesModel returns a model for the database table.
func NewRulesModel(conn sqlx.SqlConn) RulesModel {
	return &customRulesModel{
		defaultRulesModel: newRulesModel(conn),
	}
}
