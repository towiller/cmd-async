package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	//"wdyedu.com/micro/cmd-async/srv/services/proxy"
	"github.com/garyburd/redigo/redis"
	"srv/services/proxy"
	//asyncmd "wdyedu.com/micro/cmd-async/srv/services/cmd"
	"strings"
)

func TestReadList() {

	serviceName := "redis.dev.saas"
	redisProxy := proxy.NewRedisProxy()
	setRes, _ := redisProxy.Do(serviceName, "RPUSH", "cmd-async_test_list", "demo")
	fmt.Println(setRes)

	getRes, getErr := redisProxy.Do(serviceName, "LPOP", "cmd-async_test_list")
	fmt.Println(redis.String(getRes, getErr))
}

func TestReadSet() {
	serviceName := "redis.dev.saas"
	redisProxy := proxy.NewRedisProxy()

	setRes, _ := redisProxy.Do(serviceName, "SADD", "cmd-async_test_set", "test_cnt")
	fmt.Println(setRes)

	//list array
	getRes, getErr := redisProxy.Do(serviceName, "SMEMBERS", "cmd-async_test_set")
	fmt.Println(redis.Strings(getRes, getErr))

	/*
		args := make([]interface{}, 10)
		modCmd := asyncmd.RedisCmdStr{
			CmdName: "SADD",
			Args:    args,
		}
	*/
	cmdStr := "incrby cmd-async_incrby 1"
	cmdArr := strings.Split(cmdStr, " ")
	cmdInterface := make([]interface{}, len(cmdArr))
	for k, v := range cmdArr {
		fmt.Println("k=", k)
		cmdInterface[k] = v
	}
	fmt.Println(strings.ToUpper(cmdArr[0]))
	fmt.Println(cmdInterface...)

	ret2, err2 := redisProxy.Do(serviceName, "INCRBY", "cmd-async_incrby", "1")
	fmt.Println(redis.Int(ret2, err2))

	res1, err1 := redisProxy.Do(serviceName, strings.ToUpper(cmdArr[0]), cmdInterface[1:]...)
	fmt.Println(redis.Int(res1, err1))
	//fmt.Println(redis(getRes, getErr))
}

func TestReadZSet() {
	serviceName := "redis.dev.saas"
	redisProxy := proxy.NewRedisProxy()

	setRes, err := redisProxy.Do(serviceName, "ZADD", "cmd-async_test_zset", "2", "test")
	fmt.Println(redis.Bool(setRes, err))

	getRes, gErr := redisProxy.Do(serviceName, "ZRANGE", "cmd-async_test_zset", "0", "-1", "WITHSCORES")
	fmt.Println(redis.Strings(getRes, gErr))
	fmt.Println(redis.StringMap(getRes, gErr))
}

func main() {
	service := micro.NewService()
	service.Init()
	fmt.Println((*cmd.DefaultCmd.Options().Registry), (*cmd.DefaultCmd.Options().Registry).Options().Addrs)
	TestReadList()
	TestReadSet()
	TestReadZSet()
}
