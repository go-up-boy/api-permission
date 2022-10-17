package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RolesModel = (*customRolesModel)(nil)

type (
	// RolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRolesModel.
	RolesModel interface {
		rolesModel
	}

	customRolesModel struct {
		*defaultRolesModel
	}
)

// NewRolesModel returns a model for the database table.
func NewRolesModel(conn sqlx.SqlConn) RolesModel {
	return &customRolesModel{
		defaultRolesModel: newRolesModel(conn),
	}
}
