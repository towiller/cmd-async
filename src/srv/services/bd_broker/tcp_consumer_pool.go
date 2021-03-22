package bdbroker

/**
  暂时用不到，未完待续
*/

import (
	"time"

	//cluster "github.com/bsm/sarama-cluster"
	"github.com/silenceper/pool"
)

type ConsumePoolMethod interface {
	Conn()
	Consumer()
}

type ConsumerPool struct {
	Pool pool.Pool
}

func NewConsumerPool(opts ...Option) *ConsumerPool {
	PoolFacyory := func() (interface{}, error) {
		bdBroker := NewBroker(opts...)
		err := bdBroker.Conn()
		return bdBroker, err
	}

	PoolClose := func(bdBroker interface{}) error { return bdBroker.(BdBrokerMethod).Close() }
	PoolPing := func(bdBroker interface{}) error { return nil }
	poolConfig := &pool.Config{
		InitialCap: 5,
		MaxCap:     30,
		Factory:    PoolFacyory,
		Close:      PoolClose,
		Ping:       PoolPing,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 3000 * time.Second,
	}

	p, _ := pool.NewChannelPool(poolConfig)

	return &ConsumerPool{
		Pool: p,
	}
}
