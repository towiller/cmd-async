package connpool

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"math"
	"strconv"
)

type ConnRedis struct {
	redisConf *RedisOptions
}

func NewRedisCon(args ...RedisOption) (*ConnRedis, error) {
	redisConf := &RedisOptions{
		MaxIdle:     1000,
		IdleTimeout: 100 * time.Second,
		DbCount:     10,
	}

	for _, c := range args {
		c(redisConf)
	}

	return &ConnRedis{
		redisConf: redisConf,
	}, nil
}

func (c *ConnRedis) getPool() *redis.Pool {
	redisConf := c.redisConf
	return &redis.Pool{
		MaxIdle:     redisConf.MaxIdle,
		IdleTimeout: redisConf.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				redisConf.Host+":"+strconv.Itoa(redisConf.Port),
				redis.DialReadTimeout(time.Second*5),
				redis.DialConnectTimeout(time.Second*1),
			)
			if err != nil {
				return nil, err
			}

			if redisConf.Password != "" {
				if _, err := c.Do("AUTH", redisConf.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", redisConf.SelectDB); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := c.Do("PING")
			return err
		},
	}
}

func (c *ConnRedis) Do(cmdName string, args ...interface{}) (interface{}, error) {
	conn := c.getPool().Get()
	defer conn.Close()
	ret, err := conn.Do(cmdName, args...)
	if err != nil {
		log.Println("[redis_do_error]", cmdName, err, args)
		return nil, err
	}
	return ret, err
}

func (c *ConnRedis) SelectWithKey(keyName string) (interface{}, error) {
	conn := c.getPool().Get()
	defer conn.Close()

	lastChar := keyName[len(keyName)-1] //获取最后一个字符串的ASCII值
	m := math.Mod(float64(lastChar), float64(c.redisConf.DbCount))
	return c.Do("SELECT", m)
}

func (c *ConnRedis) DoBoolRet(cmdName string, args ...interface{}) (bool, error) {
	ret, err := c.Do(cmdName, args...)
	if err != nil || ret != "OK" {
		return false, nil
	}

	return true, nil
}

func (c *ConnRedis) DoStringRet(cmdName string, args ...interface{}) (string, error) {
	ret, err := c.Do(cmdName, args...)
	return redis.String(ret, err)
}

func (c *ConnRedis) DoIntRet(cmdName string, args ...interface{}) (int, error) {
	ret, err := c.Do(cmdName, args...)
	return redis.Int(ret, err)
}

func (c *ConnRedis) DoInt64Ret(cmdName string, args ...interface{}) (int64, error) {
	ret, err := c.Do(cmdName, args...)
	return redis.Int64(ret, err)
}

func (c *ConnRedis) DoUint64Ret(cmdName string, args ...interface{}) (uint64, error) {
	ret, err := c.Do(cmdName, args...)
	return redis.Uint64(ret, err)
}
