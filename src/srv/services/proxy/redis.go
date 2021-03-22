package proxy

import (
	"srv/services/conn_pool"
	"srv/services/kv"
)

var (
	MapRedisPool = make(map[string]*connpool.ConnRedis)
)

type RedisProxy struct {
}

func NewRedisProxy() *RedisProxy {
	return &RedisProxy{}
}

func (m *RedisProxy) Do(serviceName string, cmdName string, args ...interface{}) (interface{}, error) {
	redisPool := m.GetServicePool(serviceName)
	return redisPool.Do(cmdName, args...)
}

func (m *RedisProxy) GetServicePool(serviceName string) *connpool.ConnRedis {

	key := "services_config/redis/" + serviceName
	//log.Println(key)
	if redisPool, ok := MapRedisPool[key]; ok {
		return redisPool
	}

	consulKv := kv.ConsulKvInstnce()
	redisConfig := connpool.RedisOptions{}
	consulKv.GetJsonVal(key, &redisConfig)

	redisPool, _ := connpool.NewRedisCon(
		connpool.RedisHost(redisConfig.Host),
		connpool.RedisPort(redisConfig.Port),
		connpool.RedisPassword(redisConfig.Password),
		connpool.RedisSelectDB(redisConfig.SelectDB),
		connpool.RedisIdleTimeout(redisConfig.IdleTimeout),
		connpool.RedisMaxIdle(redisConfig.MaxIdle),
	)

	MapRedisPool[key] = redisPool
	return redisPool
}
