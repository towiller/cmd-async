package proxy

import (
	//"github.com/micro/go-micro"
	//"github.com/micro/go-micro/registry"
	//"github.com/micro/go-micro/cmd"
	"database/sql"
	//"github.com/micro/go-micro/selector"
	"fmt"
	"srv/services/conn_pool"
	"srv/services/kv"
)

var (
	MapMysqlPool = make(map[string]*connpool.MysqlConnPool)
)

type MysqlProxy struct {
	//RegSelector selector.Selector
	//ConsulKv    *kv.ConsulKv
}

func NewMysqlProxy() *MysqlProxy {
	return &MysqlProxy{}
}

func (m *MysqlProxy) Query(serviceName string, sqlstr string) (*sql.Rows, error) {
	mysqlPool := m.GetServicePool(serviceName)
	return mysqlPool.Query(sqlstr)
}

func (m *MysqlProxy) Exec(serviceName string, sqlstr string) (sql.Result, error) {
	mysqlPool := m.GetServicePool(serviceName)
	return mysqlPool.Exec(sqlstr)
}

func (m *MysqlProxy) Begin(serviceName string) (*sql.Tx, error) {
	mysqlPool := m.GetServicePool(serviceName)
	return mysqlPool.Begin()
}

func (m *MysqlProxy) Prepare(serviceName string, sqlstr string) (*sql.Stmt, error) {
	mysqlPool := m.GetServicePool(serviceName)
	return mysqlPool.Prepare(sqlstr)
}

func (m *MysqlProxy) GetServicePool(serviceName string) *connpool.MysqlConnPool {

	key := "services_config/mysql/" + serviceName
	if mysqlPool, ok := MapMysqlPool[key]; ok {
		return mysqlPool
	}

	consulKv := kv.ConsulKvInstnce()
	mysqlConfig := connpool.MysqlOptions{}
	consulKv.GetJsonVal(key, &mysqlConfig)

	fmt.Println(mysqlConfig)

	mysqlPool := connpool.NewMysqlConnPool(
		connpool.MysqlHost(mysqlConfig.Host),
		connpool.MysqlPort(mysqlConfig.Port),
		connpool.MysqlCharset(mysqlConfig.Charset),
		connpool.MysqlDefaultDb(mysqlConfig.DefaultDb),
		connpool.MysqlUserName(mysqlConfig.Username),
		connpool.MysqlPassword(mysqlConfig.Password),
		connpool.MysqlMaxOpenConns(mysqlConfig.MaxOpenConns),
		connpool.MysqlMaxIdleConns(mysqlConfig.MaxIdleConns),
	)

	MapMysqlPool[key] = mysqlPool
	return mysqlPool
}
