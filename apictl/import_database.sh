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
