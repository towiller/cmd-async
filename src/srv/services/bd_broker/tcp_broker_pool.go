package bdbroker

import (
	"log"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/silenceper/pool"
)

var (
	DefaultBrokerPool = NewBrokerPool()
)

type BrokerPool struct {
	Pool    pool.Pool
	Options []Option
}

func NewBrokerPool(opts ...Option) *BrokerPool {
	PoolFacyory := func() (interface{}, error) {
		bdBroker := NewBroker(opts...)
		err := bdBroker.Conn()
		log.Println("[warning] kafka_conn", err)
		return bdBroker, err
	}

	PoolClose := func(bdBroker interface{}) error {
		log.Println("[warning] kafka_close:", bdBroker)
		return bdBroker.(BdBrokerMethod).Close()
	}

	PoolPing := func(bdBroker interface{}) error { return nil }
	poolConfig := &pool.Config{
		InitialCap: 5,
		MaxCap:     10,
		Factory:    PoolFacyory,
		Close:      PoolClose,
		Ping:       PoolPing,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 30000 * time.Second,
	}

	p, _ := pool.NewChannelPool(poolConfig)

	return &BrokerPool{
		Pool:    p,
		Options: opts,
	}
}

func (b *BrokerPool) ReConnectPool() {
	//b.Pool = nil
	PoolFacyory := func() (interface{}, error) {
		bdBroker := NewBroker(b.Options...)
		err := bdBroker.Conn()
		log.Println("[warning] kafka_conn", err)
		return bdBroker, err
	}

	PoolClose := func(bdBroker interface{}) error {
		log.Println("[warning] kafka_close:", bdBroker)
		return bdBroker.(BdBrokerMethod).Close()
	}

	PoolPing := func(bdBroker interface{}) error { return nil }
	poolConfig := &pool.Config{
		InitialCap: 5,
		MaxCap:     10,
		Factory:    PoolFacyory,
		Close:      PoolClose,
		Ping:       PoolPing,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 30000 * time.Second,
	}

	log.Println("reConnKafka", b.Options)
	p, _ := pool.NewChannelPool(poolConfig)
	b.Pool = p
}

func (b *BrokerPool) PublishMsg(topic string, headers map[string]string, body []byte) error {
	bdBroker, err := b.Pool.Get()
	if err != nil {
		log.Println("[error] pool get error ", err)
		b.ReConnectPool()
		bdBroker, _ = b.Pool.Get()
	}

	pushErr := bdBroker.(BdBrokerMethod).PublishMsg(topic, headers, body)
	b.Pool.Put(bdBroker)
	return pushErr
}

func (b *BrokerPool) ConsumerMsg(topic string, handler broker.Handler) (broker.Subscriber, error) {
	bdBroker, err := b.Pool.Get()
	if err != nil {
		log.Println("[error] pool get error ", err)
		b.ReConnectPool()
		bdBroker, _ = b.Pool.Get()
	}

	ret, subErr := bdBroker.(BdBrokerMethod).ConsumerMsg(topic, handler)
	b.Pool.Put(bdBroker)
	return ret, subErr
}
