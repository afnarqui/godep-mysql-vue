package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Connect() (db *xorm.Engine,err error) {
	return xorm.NewEngine("mysql","test:test@tcp(localhost:3306)test?charset=utf8")
}


