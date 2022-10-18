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

2. docker compose 部署
```shell
// 进入根目录执行命令
docker build -t api-permission:v1.0.0 .

docker compose up
```
3. 数据库导入
    * 数据库文件在apictl目录下,可自行使用sql文件导入，也可用import_database脚本导入
```shell
#!/bin/bash

HOST="127.0.0.1"
PORT=3309
USERNAME="go-permission"
PASSWORD="ZkWDRyyyRBM22c4A"
DB="go-permission"

if [ -f "./go-permission.sql" ];then
mysql -h${HOST} -P${PORT} -u${USERNAME} ${DB} -p${PASSWORD} -e "source ./go-permission.sql"
mv ./go-permission.sql ./go-permission.sql.tmp
else
echo "数据库文件不存在或已导入"
fi
```

# 云服务接入
1. 阿里云 SLS
https://blog.csdn.net/manwufeilong/article/details/125781421
   ![post](https://s1.ax1x.com/2022/10/18/xsCCkD.png)
   