package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
)

var (
	brokerAddr    = "127.0.0.1:9091"
	topic         = "cc7514bff90b4f0c9bc22787de7f6e77__logtest"
	enableTLS     = true
	clientPemPath = "./config/baidu-kafka/client.pem"
	clientKeyPath = "./config/baidu-kafka/client.key"
	caPemPath     = "./config/baidu-kafka/ca.pem"
)

func main() {

	sarama.Logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_1_0
	config.Net.TLS.Enable = enableTLS
	config.Net.TLS.Config = configTLS()

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9091"}, config)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	partitions, err := consumer.Partitions(topic)
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
	}

	wg.Wait()
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
