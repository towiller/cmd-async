package bdbroker

import (
	//"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/codec"
	"github.com/micro/go-micro/codec/json"
	"sync"
)

type ConsumerHandler func(message *HandlerMessage) error

type HandlerMessage struct {
	val      *broker.Message
	msg      *sarama.ConsumerMessage
	topic    string
	consumer *cluster.Consumer
	pc       *cluster.PartitionConsumer
}

func (h *HandlerMessage) Topic() string {
	return h.topic
}

func (h *HandlerMessage) Val() *broker.Message {
	return h.val
}

func (h *HandlerMessage) Ack() error {
	h.consumer.MarkOffset(h.msg, "")
	return nil
}

func (h *HandlerMessage) ConsumerMessage() *sarama.ConsumerMessage {
	return h.msg
}

type BdConsumer struct {
	consumer *cluster.Consumer
	codec    codec.Marshaler
}

func NewConsumer(opts ...Option) *BdConsumer {
	options := Options{
		Addrs: []string{
			"127.0.0.1:9091",
		},
		EnableTls: true,
	}

	for _, o := range opts {
		o(&options)
	}

	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	config.Version = sarama.V0_10_1_0
	config.Net.TLS.Enable = options.EnableTls
	if config.Net.TLS.Enable {
		WithTls(clientPemPath, clientKeyPath, caPemPath)(&options)
		config.Net.TLS.Config = options.TlsConfig
	}
	consumer, err := cluster.NewConsumer(options.Addrs, options.GroupId, options.Topics, config)

	if err != nil {
		sarama.Logger.Println("consumer error", err)
	}

	return &BdConsumer{
		consumer: consumer,
		codec:    json.Marshaler{},
	}
}

func (b *BdConsumer) Consumer(handler ConsumerHandler, wg sync.WaitGroup) {
	consumer := b.consumer
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case part, ok := <-consumer.Partitions():
			if !ok {
				return
			}

			// start a separate goroutine to consume messages
			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					var getVal broker.Message
					b.codec.Unmarshal(msg.Value, &getVal)
					err := handler(&HandlerMessage{
						val:      &getVal,
						msg:      msg,
						topic:    msg.Topic,
						consumer: consumer,
						pc:       &pc,
					})
					if err != nil {
						sarama.Logger.Println("error:", err)
					}
				}
			}(part)
		case <-signals:
			return
		}
	}

	wg.Done()
}
