package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UsersRoleRelationModel = (*customUsersRoleRelationModel)(nil)

type (
	// UsersRoleRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersRoleRelationModel.
	UsersRoleRelationModel interface {
		usersRoleRelationModel
	}

	customUsersRoleRelationModel struct {
		*defaultUsersRoleRelationModel
	}
)

// NewUsersRoleRelationModel returns a model for the database table.
func NewUsersRoleRelationModel(conn sqlx.SqlConn) UsersRoleRelationModel {
	return &customUsersRoleRelationModel{
		defaultUsersRoleRelationModel: newUsersRoleRelationModel(conn),
	}
}
