package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	//"fmt"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	//"github.com/micro/go-micro/codec"
	//"github.com/micro/go-plugins/broker/kafka"
	"encoding/json"
	"github.com/micro/go-micro"
	"srv/services/bd_broker"
	"srv/services/cmd"
	"srv/services/kv"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

var (
	brokerAddr    = "kafka.bj.baidubce.com:9091"
	topic         = "cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.cmd-async.common"
	enableTLS     = true
	clientPemPath = "./config/baidu-kafka/client.pem"
	clientKeyPath = "./config/baidu-kafka/client.key"
	caPemPath     = "./config/baidu-kafka/ca.pem"
)

func TestGroupConsumer() {
	var topics []string

	groupId := "user_group_check"

	service := micro.NewService(
		micro.Name("test-micro-service"),
		micro.Version("v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	kv.DefaultConsulKv.GetJsonVal("/cmd-async/topics", &topics)

	bdConsumer := bdbroker.NewConsumer(
		bdbroker.Topics(topics),
		bdbroker.GroupId(groupId),
	)

	bdConsumer.Consumer(func(msg *bdbroker.HandlerMessage) error {
		cnt := &asynccmd.CmdAgg{}
		json.Unmarshal(msg.Val().Body[:], cnt)

		//sarama.Logger.Println("notice:", "Header", msg.Val().Header, "Body", string(msg.Val().Body[:]))
		return nil
	})
	fmt.Println(bdConsumer)
}

func TestConsumer() {
	sarama.Logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)

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

	partitions, err := consumer.Partitions(topic)
	log.Println("partitions:", partitions)
	if err != nil {
		log.Panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(partitions))

	for _, partition := range partitions {
		log.Println("consume partition:", partition)

		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		go func(partition int32) {
			partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				log.Fatalln(err)
				return
			}

			defer func() {
				wg.Done()
				if err := partitionConsumer.Close(); err != nil {
					log.Fatalln(err)
				}
			}()

		ConsumerLoop:
			for {
				select {
				case err := <-partitionConsumer.Errors():
					if err != nil {
						log.Fatalln("error:", err)
					}
				case msg := <-partitionConsumer.Messages():
					if msg != nil {
						log.Println("topic:", msg.Topic, "partition:", msg.Partition, "offset:", msg.Offset, "message:", string(msg.Value))
					}
				case <-shutdown:
					log.Println("stop consuming partition:", partition)
					break ConsumerLoop
				}
			}
		}(partition)
		wg.Wait()
	}

}

func gorunOne() {
	//wg.Add(1)
	go func() {
		for {
			log.Println("1111")
			time.Sleep(time.Second)
		}
	}()
}

func gorunTwo() {
	go func() {
		for {
			log.Println("222")
			time.Sleep(time.Second)
		}
	}()
}

func main() {
	sarama.Logger = log.New(os.Stderr, "[async]", log.LstdFlags)
	TestGroupConsumer()
	//gorunOne()
	//gorunTwo()

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
