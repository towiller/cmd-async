package bdbroker

import (
	"github.com/Shopify/sarama"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/broker"
	sc "gopkg.in/bsm/sarama-cluster.v2"
	"plugins/broker/kafka"
)

var (
	clientPemPath = "./config/baidu-kafka/client.pem"
	clientKeyPath = "./config/baidu-kafka/client.key"
	caPemPath     = "./config/baidu-kafka/ca.pem"
)

type BdBrokerMethod interface {
	Conn() error
	Close() error
	PublishMsg(topic string, headers map[string]string, body []byte) error
	ConsumerMsg(topic string, handler broker.Handler) (broker.Subscriber, error)
}

type BdBroker struct {
	options Options
	broker  broker.Broker
	config  *sarama.Config
}

func NewBroker(opts ...Option) *BdBroker {
	options := Options{
		Addrs:     []string{"127.0.0.1:9091"},
		EnableTls: true,
	}

	for _, op := range opts {
		op(&options)
	}

	saramaConfig := sarama.NewConfig()
	clusterConfig := sc.NewConfig()
	clusterConfig.Version = sarama.V0_10_1_0
	saramaConfig.Version = sarama.V0_10_1_0

	saramaConfig.Net.TLS.Enable = options.EnableTls
	clusterConfig.Net.TLS.Enable = options.EnableTls
	clusterConfig.Consumer.Offsets.Initial = options.OffsetId

	//saramaConfig.Consumer.Group.Member.UserData = []byte("cmd-async")
	//clusterConfig.Consumer.Group.Member.UserData = []byte("cmd-async")

	if options.EnableTls {
		WithTls(clientPemPath, clientKeyPath, caPemPath)(&options)
		saramaConfig.Net.TLS.Config = options.TlsConfig
		clusterConfig.Net.TLS.Config = options.TlsConfig
	}
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	saramaConfig.Producer.Retry.Max = 5

	kafkaBroker := synckafka.NewBroker(func(bo *broker.Options) {
		bo.Addrs = options.Addrs
		synckafka.BrokerConfig(saramaConfig)(bo)
		synckafka.ClusterConfig(clusterConfig)(bo)
	})
	return &BdBroker{
		broker:  kafkaBroker,
		options: options,
		config:  saramaConfig,
	}
}

func (b *BdBroker) Conn() error {
	err := b.broker.Connect()
	if err != nil {
		log.Log("[notice] kafka connect error: ", err)
		return err
	}
	return nil
}

func (b *BdBroker) Close() error {
	err := b.broker.Disconnect()
	if err != nil {
		log.Log("kafka close error:", err)
		return err
	}
	return nil
}

func (b *BdBroker) PublishMsg(topic string, headers map[string]string, body []byte) error {
	msg := broker.Message{
		Header: headers,
		Body:   body,
	}
	err := b.broker.Publish(topic, &msg)
	if err != nil {
		log.Fatal("send msg error", err, topic, headers, body)
		return err
	}
	return nil
}

func (b *BdBroker) ConsumerMsg(topic string, handler broker.Handler) (broker.Subscriber, error) {
	return b.broker.Subscribe(topic, handler)
}
