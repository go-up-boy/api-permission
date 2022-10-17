package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RoleRulesRelationModel = (*customRoleRulesRelationModel)(nil)

type (
	// RoleRulesRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleRulesRelationModel.
	RoleRulesRelationModel interface {
		roleRulesRelationModel
	}

	customRoleRulesRelationModel struct {
		*defaultRoleRulesRelationModel
	}
)

// NewRoleRulesRelationModel returns a model for the database table.
func NewRoleRulesRelationModel(conn sqlx.SqlConn) RoleRulesRelationModel {
	return &customRoleRulesRelationModel{
		defaultRoleRulesRelationModel: newRoleRulesRelationModel(conn),
	}
}
