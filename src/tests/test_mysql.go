package main

import (
	"database/sql"
	"fmt"
	"srv/services/proxy"
)

func TestMysql() {
	sqlStr := "insert into saas.coupon (`shop_id`, `coupon_type`, `name`, `money`, `use_type`, `created_at`, `updated_at`) values('158205', '1', '测试', '1', '1', current_timestamp(), current_timestamp())"

	db, connErr := sql.Open("mysql", "root:17dayup2017@tcp(192.168.1.19:3306)/saas?charset=utf8")
	fmt.Println(db, connErr)
	rsult, err := db.Exec(sqlStr)
	fmt.Println(rsult, err)
	rsultLastId, _ := rsult.LastInsertId()
	fmt.Println(rsultLastId)

	serviceName := "mysql.dev.saas"
	mysqlProxy := proxy.NewMysqlProxy()

	sqlResult, err := mysqlProxy.Exec(serviceName, sqlStr)

	lastId, lErr := sqlResult.LastInsertId()
	rowsAffected, _ := sqlResult.RowsAffected()
	fmt.Println("sqlStr: lastId", lastId, ", rows_affected", rowsAffected, ", err", err, "lerr", lErr, "rowsAffected")
}

func main() {
	TestMysql()
}
