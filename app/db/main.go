package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Open(driver string, dataSource string) error {

	var err error
	db, err = sql.Open(driver, dataSource)
	if err != nil {
		return err
	}
	// open 操作只是确定参数是否合法，
	// 不会创建一个数据库连接,
	// 此处用Ping 函数确定数据库服务是否可以访问
	return db.Ping()
}
