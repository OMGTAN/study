package dao

import (
	"blog/common"
	"database/sql"
)

var DB *sql.DB

func sqlOpen() {
	var err error
	DB, err = sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=blog sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;

	//还可以是这种方式打开
	//db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	common.PrintErr(err)
}

func init() {
	sqlOpen()
}
