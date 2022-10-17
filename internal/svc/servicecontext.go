package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"api-permission/internal/config"
	"api-permission/internal/middleware"
	"api-permission/model"
)

type ServiceContext struct {
	Config config.Config
	UserModel 		model.UsersModel
	RuleModel 		model.RulesModel
	RoleModel 		model.RolesModel
	UsersRoleModel  model.UsersRoleRelationModel
	RoleRulesModel  model.RoleRulesRelationModel
	Permissions 	rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	userModel := model.NewUsersModel(conn)
	ruleModel := model.NewRulesModel(conn)
	roleModel := model.NewRolesModel(conn)
	usersRoleModel := model.NewUsersRoleRelationModel(conn)
	roleRulesModel := model.NewRoleRulesRelationModel(conn)
	return &ServiceContext{
		Config: c,
		Permissions: middleware.NewPermissionsMiddleware(middleware.MustParams{
			UserModel: userModel,
			RuleModel: ruleModel,
			RoleModel: roleModel,
			UsersRoleModel: usersRoleModel,
			RoleRulesModel: roleRulesModel,
		}).Handle,
		UserModel: userModel,
		RuleModel: ruleModel,
		RoleModel: roleModel,
		UsersRoleModel: usersRoleModel,
		RoleRulesModel: roleRulesModel,
	}
}
