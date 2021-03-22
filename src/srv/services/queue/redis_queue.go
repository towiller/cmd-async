package cmdqueue

import (
	//"encoding/json"
	"strconv"
	//"time"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"srv/services/conn_pool"
	"time"
)

type ConsumerMsg struct {
	topic  string
	msg    []byte
	broker *RedisQueue
}

func (c *ConsumerMsg) Ack() error {
	return nil
}

func (c *ConsumerMsg) GetMsg() []byte {
	return c.msg
}

type ConsumerCallback func(*ConsumerMsg) error

type RedisQueue struct {
	RedisPool *connpool.ConnRedis
	NodeId    string
}

func NewRedisQueue(args ...connpool.RedisOption) *RedisQueue {
	redisPool, err := connpool.NewRedisCon(args...)
	if err != nil {
		log.Println("[error]", "connect_redis_queue_error", err, args)
		redisPool, err = connpool.NewRedisCon(args...)
	}
	return &RedisQueue{
		RedisPool: redisPool,
	}
}

func (r *RedisQueue) DelayPublish(topic string, msg interface{}, delaySecond int64) error {
	newTopic := topic + strconv.FormatInt(time.Now().Unix()+delaySecond, 10)
	return r.Publish(newTopic, msg)
}

func (r *RedisQueue) Publish(topic string, msg interface{}) error {
	redisKey := r.GetRedisKey(topic)
	r.RedisPool.SelectWithKey(redisKey)

	d, err := json.Marshal(msg)
	if err != nil {
		log.Println("[error] redis_queue_json_marshal_err", redisKey, msg)
		return err
	}

	log.Println("[notice] redis_lpush_redis", redisKey, d)
	_, e := r.RedisPool.Do("LPUSH", redisKey, d)

	if e != nil {
		log.Println("[error] redis_queue_lpush_err", redisKey, msg)
		return e
	}
	return nil
}

func (r *RedisQueue) Consumer(topic string, cb ConsumerCallback) {

	redisKey := r.GetRedisKey(topic)
	r.RedisPool.SelectWithKey(redisKey)

	for {
		result, e := r.RedisPool.Do("RPOP", redisKey)
		msg, _ := redis.Bytes(result, e)

		//var msg interface{}
		//json.Unmarshal(data, &msg)

		err := cb(&ConsumerMsg{
			topic:  topic,
			msg:    msg,
			broker: r,
		})

		if err != nil { //若处理失败则重新推送至队列
			r.Publish(topic, msg)
		}
	}
}

func (r *RedisQueue) DelayConsumer(topic string, cb ConsumerCallback, startTime int64, endTime int64) {
	timer := time.NewTicker(time.Second * 1)
	var i int64
	var stime int64
	if startTime > 0 {
		stime = startTime
	} else {
		stime = time.Now().Unix()
	}

	for {
		select {
		case <-timer.C:

			if endTime > 0 && stime > endTime {
				log.Println("[error]", "end_consumer", topic, startTime, endTime)
				return
			}
			redisKey := r.GetRedisKey(topic + strconv.FormatInt(stime+i, 10))

			go func(redisKey string) {
				r.RedisPool.SelectWithKey(redisKey)
				for {
					result, e := r.RedisPool.Do("RPOP", redisKey)
					msg, _ := redis.Bytes(result, e)

					if result == nil || e != nil {
						log.Println("[notice] break_rpop", redisKey, result, e)
						break
					}

					err := cb(&ConsumerMsg{
						topic:  topic,
						msg:    msg,
						broker: r,
					})

					if err != nil { //若处理失败则重新推送至队列
						r.DelayPublish(redisKey, msg, 10)
					}
				}

				r.RedisPool.Do("EXPIRE", redisKey, 3600*24*15)
			}(redisKey)

			i++
		}
	}
}

func (r *RedisQueue) GetRedisKey(topic string) string {
	return "async-cmd_topic_" + "_" + topic
}
