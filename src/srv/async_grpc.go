package main

import (
	"flag"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"log"
	"os"
	"strings"
	"time"

	cmdconfig "config"
	controller "srv/controller"
	pb "srv/proto"
	"srv/services/bd_broker"
)

// 解决不支持自定义端口号
func MicroAddress(addr string) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.Address(addr))
	}
}

func main() {
	if len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	bdbroker.DefaultBrokerPool = bdbroker.NewBrokerPool(
		bdbroker.Addrs(strings.Split(cmdconfig.DefaultEnv.KafkaAddress, ",")),
	)

	service := micro.NewService(
		//MicroAddress(":9092"),
		micro.Name(cmdconfig.DefaultEnv.AppEnv+"."+cmdconfig.DefaultEnv.AppSrvName+"."+cmdconfig.DefaultEnv.AppVersion),
		micro.Version(cmdconfig.DefaultEnv.AppVersion),
		micro.Registry(consul.NewRegistry()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	cmdAsync, _ := controller.NewCmdAsync()
	pb.RegisterCallHandler(service.Server(), cmdAsync)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
