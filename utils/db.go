package utils

import (
	"database/sql"

	//匿名导入第三方
	_ "github.com/go-sql-driver/mysql"
)

//Db 数据库
var Db *sql.DB
var err error

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/carsys")
	if err != nil {
		panic(err.Error())
	}
}
