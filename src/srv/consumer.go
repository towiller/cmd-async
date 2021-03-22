package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/micro/go-micro"
	//"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/registry/consul"

	cmdconfig "config"
	"github.com/Shopify/sarama"
	"github.com/micro/go-micro/cmd"
	"io/ioutil"
	"srv/services/bd_broker"
	"srv/services/cmd"
	"srv/services/conn_pool"
	"srv/services/kv"
	"srv/services/queue"
	"strconv"
	"strings"
)

var (
	wg sync.WaitGroup
)

func main() {
	var topics []string
	sarama.Logger = log.New(os.Stderr, "[async-cmd]", log.LstdFlags)

	errTopic := cmdconfig.DefaultEnv.KafkaErrTopic

	service := micro.NewService(
		micro.Name("micro.cmd-async.consumer.v1"),
		micro.Registry(consul.NewRegistry()),
	)
	service.Init()

	var kOffset int64
	fileName := "kafka_consumer_offset"
	fileObj, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic("open file error" + fileName)
	}
	defer fileObj.Close()
	fContents, rerr := ioutil.ReadAll(fileObj)
	if rerr != nil {
		panic("read file error" + fileName)
	}
	kOffset, _ = strconv.ParseInt(string(fContents[:]), 10, 64)

	//监听注册的topic
	kv.DefaultConsulKv = kv.NewConsulKv()
	kv.DefaultConsulKv.GetJsonVal("/cmd-async/topics", &topics)
	bdbroker.DefaultBrokerPool = bdbroker.NewBrokerPool(
		bdbroker.Addrs(strings.Split(cmdconfig.DefaultEnv.KafkaAddress, ",")),
		bdbroker.OffsetId(kOffset),
	)

	log.Println((*cmd.DefaultCmd.Options().Registry).Options())
	wg.Add(2)

	go func(wg sync.WaitGroup, fName string) {
		asyncCmd := asynccmd.NewCmd(errTopic)
		bdConsumer := bdbroker.NewConsumer(
			bdbroker.Topics(topics),
			bdbroker.GroupId(cmdconfig.DefaultEnv.KafkaGroupId),
		)

		bdConsumer.Consumer(func(msg *bdbroker.HandlerMessage) error {
			sarama.Logger.Println("consumer_msg:", "offset", msg.ConsumerMessage().Offset, "Header", msg.Val().Header, "Body", string(msg.Val().Body[:]))
			log.Println(msg.ConsumerMessage().Offset)

			strOffset := strconv.FormatInt(msg.ConsumerMessage().Offset, 10)
			fhandle, _ := os.OpenFile(fName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			fhandle.WriteAt([]byte(strOffset), 0)
			defer fhandle.Close()
			fhandle.Close()
			cnt := &asynccmd.CmdAgg{}
			json.Unmarshal(msg.Val().Body[:], cnt)
			asyncCmd.ExecCmds(cnt, msg.Val().Header)

			msg.Ack()
			return nil
		}, wg)
	}(wg, fileName)

	//redis定时队列
	go func(wg sync.WaitGroup, fName string) {
		asyncCmd := asynccmd.NewCmd(errTopic)

		redisQueue := cmdqueue.NewRedisQueue(
			connpool.RedisHost(cmdconfig.DefaultEnv.RedisHost),
			connpool.RedisPort(cmdconfig.DefaultEnv.RedisPort),
			connpool.RedisDbCount(cmdconfig.DefaultEnv.RedisDBCount),
		)

		redisQueue.DelayConsumer(cmdconfig.DefaultEnv.KafkaErrTopic, func(cmsg *cmdqueue.ConsumerMsg) error {
			cnt := &asynccmd.CmdRepublishMsg{}
			json.Unmarshal(cmsg.GetMsg(), cnt)
			log.Println("[notice] repeat_reconsumer_msg:", cnt)
			asyncCmd.ExecCmds(cnt.Cmds, cnt.Header)
			cmsg.Ack()
			return nil
		}, 0, 0)
		wg.Done()
	}(wg, fileName)

	/*
		go func(wg sync.WaitGroup, fName string) {
			asyncCmd := asynccmd.NewCmd(errTopic)
			errBdConsumer := bdbroker.NewConsumer(
				bdbroker.Topics([]string{
					errTopic,
				}),
				bdbroker.GroupId(groupId),
			)

			errBdConsumer.Consumer(func(msg *bdbroker.HandlerMessage) error {
				sarama.Logger.Println("repeat_reconsumer_msg", "offset", msg.ConsumerMessage().Offset, "Header", msg.Val().Header, "Body", string(msg.Val().Body[:]))
				cnt := &asynccmd.CmdAgg{}
				json.Unmarshal(msg.Val().Body[:], cnt)
				log.Println("cnt:", cnt)
				asyncCmd.ExecCmds(cnt, msg.Val().Header)
				msg.Ack()
				strOffset := strconv.FormatInt(msg.ConsumerMessage().Offset, 10)
				fhandle, _ := os.OpenFile(fName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
				fhandle.WriteAt([]byte(strOffset), 0)
				defer fhandle.Close()
				fhandle.Close()
				return nil
			}, wg)
		}(wg, fileName)
	*/

	wg.Wait()
	fileObj.Close()
	//time.Ticker{}
}
