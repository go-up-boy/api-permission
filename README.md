# api-permission
基于go-zero单体架构开发的API权限管理

# 为什么要这个？
就是对用户请求API路由做权限管理，很多项目中对用户的权限在前端显示隐藏进行控制，这样其实有很大的风险。

# 有什么功能？
* RBAC权限管理（Ok）
* Response统一处理（Ok）
* 用户基础操作（Ok）
* 权限管理文件迁移模式（待开发）

# 常用操作
1. go-zero一键生成api代码
```shell
goctl api go -api ./apictl/admin_users.api -dir .
```
