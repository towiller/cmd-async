package main

import (
	//"time"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"config"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"srv/services/conn_pool"

	"srv/services/cmd"
	"srv/services/queue"
)

func TestMain() {
	str := "asdkjgksks"
	strChar := str[len(str)-1]
	m := math.Mod(float64(strChar), 16)
	fmt.Println("m=", m)

	//mod := math.
	//i[0] = strChar[0]

	fmt.Println(strChar)
	fmt.Println(strconv.QuoteToASCII(str))
	redisConn, err := connpool.NewRedisCon(
		connpool.RedisHost(config.DefaultEnv.RedisHost),
		connpool.RedisPort(config.DefaultEnv.RedisPort),
		connpool.RedisDbCount(config.DefaultEnv.RedisDBCount),
	)

	fmt.Println(redisConn)
	if err != nil {
		fmt.Println("print redis conn")
	}

	timeFormat := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("timeformat", timeFormat, "timeNow", time.Now().Unix())

	for i := 0; i < 100; i++ {
		nowUnix := time.Now().Unix() + int64(i)
		timeRedisKey := "cmd-async_error_" + strconv.FormatInt(nowUnix, 10)

		redisConn.SelectWithKey(timeRedisKey)
		d, _ := json.Marshal("test " + timeRedisKey)
		ret, e := redisConn.Do("LPUSH", timeRedisKey, d)
		log.Println("lpush_str: redis_key = ", timeRedisKey, "ret=", ret, e)
	}

	timer := time.NewTicker(time.Second * 1)
	//go func() {
	for {
		select {
		case a := <-timer.C:
			//timeRedisKey
			fmt.Println("unix_time", a.Unix())
			redisKey := "cmd-async_error_" + strconv.FormatInt(a.Unix(), 10)
			redisConn.SelectWithKey(redisKey)
			for {
				result, e := redisConn.Do("RPOP", redisKey)

				s, _ := redis.Bytes(result, e)

				var msg interface{}
				json.Unmarshal(s, &msg)

				if result == nil || e != nil {
					log.Println("break rpop", msg, e)
					break
				}
				log.Println("s = ", msg)
			}

			redisConn.Do("EXPIRE", redisKey, 3600*24*15)
		}
	}
}

func TestRedisQueue() {

	redisQueue := cmdqueue.NewRedisQueue(
		connpool.RedisHost(config.DefaultEnv.RedisHost),
		connpool.RedisPort(config.DefaultEnv.RedisPort),
		connpool.RedisDbCount(config.DefaultEnv.RedisDBCount),
	)
	for i := 0; i < 100; i++ {
		//nowUnix := time.Now().Unix() + int64(i)
		msg := asynccmd.CmdAgg{
			DoMysql: []asynccmd.CmdMysql{},
		}

		redisQueue.DelayPublish("async_error", msg, int64(i)+10)
	}

	redisQueue.DelayConsumer("async_error", func(cmsg *cmdqueue.ConsumerMsg) error {
		cmdAgg := &asynccmd.CmdAgg{}
		json.Unmarshal(cmsg.GetMsg(), cmdAgg)
		log.Println(cmdAgg)
		return nil
	}, 0, 0)
}

func main() {
	TestRedisQueue()
	//}()
}
