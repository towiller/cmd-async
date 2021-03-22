package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/micro/go-micro/broker"
	"strconv"
	//"sync"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins-old/registry/consul"
	"srv/services/bd_broker"
	"srv/services/cmd"
)

var (
	brokerAddr    = "127.0.0.1:9091"
	topic         = "cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.cmd-async.common"
	enableTLS     = true
	clientPemPath = "./config/baidu-kafka/client.pem"
	clientKeyPath = "./config/baidu-kafka/client.key"
	caPemPath     = "./config/baidu-kafka/ca.pem"
)

func TestErrMsg() {
	msg := &asynccmd.CmdAgg{
		DoMysql: []asynccmd.CmdMysql{
			{
				ServiceName:   "",
				Sqls:          []string{"insert into saas.coupon (`shop_id`, `coupon_type`, `name`, `money`, `use_type`, `created_at`, `updated_at`) values(\\\"158205\\\", \\\"1\\\", \\\"name\\\", \\\"1\\\", \\\"1\\\", current_timestamp(), current_timestamp())"},
				IsTransaction: true,
			},
		},
	}

	jsonData, err := json.Marshal(msg)
	fmt.Println(string(jsonData[:]), err)

	headers := map[string]string{}
	topic := "cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.async_cmd_error"
	bdbroker.DefaultBrokerPool.PublishMsg(topic, headers, jsonData)
}

func TestBDKafka() {
	bdBroker := bdbroker.NewBroker(
		bdbroker.Addrs([]string{
			"127.0.0.1:9091",
		}),
	)
	//bdBroker.PublishMsg()
	fmt.Println(bdBroker)
}

func TestBDBrokerPool() {
	brokerPool := bdbroker.NewBrokerPool()
	topic := "cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.cmd-async.common"
	headers := make(map[string]string)

	for i := 0; i < 20; i++ {
		body := []byte("body cnt byte" + strconv.Itoa(i))
		headers["Content-type"] = "application/json;charset=utf8"
		err := brokerPool.PublishMsg(topic, headers, body)
		fmt.Println("push msg:", err)
	}
}

func TestBrokerConsumer() {
	brokerPool := bdbroker.NewBrokerPool()
	topic := "cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.cmd-async.common"

	fmt.Println(brokerPool)

	ret, err := brokerPool.ConsumerMsg(topic, func(p broker.Publication) error {
		sarama.Logger.Println("header", p.Message().Header)
		log.Println("body", p.Message().Body, "header", p.Message().Header, "topic", p.Topic())
		p.Ack()
		return nil
	})
	fmt.Println("subcribe msg:", ret.Topic(), err)
}

func TestConsumer() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_1_0
	config.Net.TLS.Enable = enableTLS
	if enableTLS {
		config.Net.TLS.Config = configTLS()
	}

	consumer, err := sarama.NewConsumer([]string{brokerAddr}, config)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	log.Println("consumer:", consumer)
}

func main() {

	service := micro.NewService(
		micro.Name("dev.micro.cmd-async11"),
		micro.Version("v1"),
		micro.Registry(consul.NewRegistry()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	sarama.Logger = log.New(os.Stderr, "[async]", log.LstdFlags)
	TestErrMsg()
	//TestBDKafka()
	//TestBDBrokerPool()
	//TestConsumer()
	//TestBrokerConsumer())
	return

	sarama.Logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_1_0
	config.Net.TLS.Enable = enableTLS
	if enableTLS {
		config.Net.TLS.Config = configTLS()
	}

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5
	fmt.Println(config.Net.TLS, brokerAddr)

	producer, err := sarama.NewSyncProducer([]string{brokerAddr}, config)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	numOfRecords := 10
	for i := 0; i < numOfRecords; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("%d-hello kafka", i)),
		}
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("message sent, topic:", topic, "partition:", partition, "offset:", offset)
	}

}

func configTLS() (t *tls.Config) {
	checkFile(clientPemPath)
	checkFile(clientKeyPath)
	checkFile(caPemPath)

	clientPem, err := tls.LoadX509KeyPair(clientPemPath, clientKeyPath)
	if err != nil {
		log.Panic(err)
	}

	caPem, err := ioutil.ReadFile(caPemPath)
	if err != nil {
		log.Panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caPem)
	t = &tls.Config{
		Certificates:       []tls.Certificate{clientPem},
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	return t
}

func checkFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Size() == 0 {
		panic("Please replace " + path + " with your own. ")
	}
}
