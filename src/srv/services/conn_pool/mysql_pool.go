package connpool

import (
	"database/sql"
	//"fmt"
	"log"
	"strconv"
	"time"

	//"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnPool struct {
	MysqlConf *MysqlOptions
	MysqlDB   *sql.DB
}

func NewMysqlConnPool(opts ...MysqlOption) *MysqlConnPool {
	mysqlConf := &MysqlOptions{
		Charset:      "utf8",
		MaxOpenConns: 100, //最大打开连接数
		MaxIdleConns: 50,  //设置闲置的连接数
	}

	for _, m := range opts {
		m(mysqlConf)
	}

	return &MysqlConnPool{
		MysqlConf: mysqlConf,
	}
}

func (p *MysqlConnPool) MysqlConn() *sql.DB {
	var conStr string

	MysqlConfig := p.MysqlConf
	if p.MysqlDB != nil {
		return p.MysqlDB
	}

	if MysqlConfig.Username != "" {
		conStr += MysqlConfig.Username + ":" + MysqlConfig.Password
	}
	conStr += "@tcp(" + MysqlConfig.Host + ":" + strconv.Itoa(MysqlConfig.Port) + ")/"

	conStr += MysqlConfig.DefaultDb + "?charset=" + MysqlConfig.Charset
	conStr += "&timeout=" + time.Second.String()

	log.Println("[notice] conn str", conStr)
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Println("[error] mysql_conn_error:", conStr)
	}
	db.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
	db.SetConnMaxLifetime(time.Second * 300)
	p.MysqlDB = db
	return p.MysqlDB
}

func (p *MysqlConnPool) Query(sqlStr string) (*sql.Rows, error) {
	rows, err := p.MysqlConn().Query(sqlStr)
	if err != nil {
		log.Println("[error] mysql_query_error", err)
	}
	defer rows.Close()
	return rows, err
}

func (p *MysqlConnPool) Prepare(sqlStr string) (*sql.Stmt, error) {
	rows, err := p.MysqlConn().Prepare(sqlStr)
	if err != nil {
		log.Println("[error] mysql_prepare_error", err)
	}
	defer rows.Close()
	return rows, err
}

func (p *MysqlConnPool) Exec(sqlStr string) (sql.Result, error) {
	rows, err := p.MysqlConn().Exec(sqlStr)
	if err != nil {
		log.Println("[error] mysql_exec_error", err)
	}
	//defer rows.Close()
	return rows, err
}

func (p *MysqlConnPool) Begin() (*sql.Tx, error) {
	return p.MysqlConn().Begin()
}
