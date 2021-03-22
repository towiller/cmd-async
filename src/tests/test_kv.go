package main

import (
	//"os"
	"fmt"
	"srv/services/conn_pool"
	"srv/services/kv"
)

func main() {
	consulKv := kv.NewConsulKv(kv.Address("127.0.0.1:8500"))

	mysqlConfig := connpool.MysqlOptions{}
	consulKv.GetJsonVal("services_config/mysql/mysql.dev.saas", &mysqlConfig)
	fmt.Println(mysqlConfig)

	redisConf := connpool.RedisOptions{}
	consulKv.GetJsonVal("services_config/redis/redis.dev.saas", &redisConf)
	fmt.Println(redisConf)
}
