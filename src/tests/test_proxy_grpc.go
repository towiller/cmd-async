package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"context"
	"encoding/json"
	"math/rand"
	//"net/http"

	"github.com/micro/go-micro"
	//microcmd "github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/selector"

	//"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	//"google.golang.org/grpc/status"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	//"github.com/micro/protobuf/proto"
	"github.com/micro/protobuf/ptypes/any"
	common "srv/proto/common"
	cmd "srv/services/cmd"
)

func TestGrpcCall() {

}

func main() {

	service := micro.NewService(
		micro.Name("dev.micro.cmd-async11"),
		micro.Version("v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	registry.DefaultRegistry = consul.NewRegistry()
	requestStr := "{\"service_name\":\"dev.micro.cmd-async\",\"action\":\"/common.service.srv.Call/Action\",\"header\":{\"Content-type\":\"text\\/application;charset=utf8\"},\"body\":\"\\\\[84 104 105 115 32 105 115 32 116 101 115 116\\\\]\"}"
	fmt.Println([]byte("This is test"))
	var cmdCnt cmd.CmdGrpc
	json.Unmarshal([]byte(requestStr), &cmdCnt)
	log.Println(cmdCnt)

	//ctx := context.Background()
	fmt.Println("default registry:", registry.DefaultRegistry)

	var registrySer []*registry.Service
	for i := 0; i < 10; i++ {
		registrySer, _ = registry.DefaultRegistry.GetService(cmdCnt.ServiceName)
	}

	//se := microcmd.DefaultCmd.Options().Selector

	serviceSelector := selector.NewSelector(selector.Registry(registry.DefaultRegistry))

	next, _ := serviceSelector.Select(cmdCnt.ServiceName)
	fmt.Println("next:", next)
	for j := 0; j < 100; j++ {
		node, nerr := next()
		fmt.Println("node", "address:", node.Address, "port:", node.Port, nerr)
	}

	var getNode *registry.Node
	for _, s := range registrySer {
		l := len(s.Nodes)
		if l < 1 {
			fmt.Print("no nodes")
		}

		r := rand.Intn(l)

		for k, n := range s.Nodes {
			if k == r {
				getNode = n
				fmt.Println(n.Address)
			}
		}
	}

	address := getNode.Address + ":" + strconv.Itoa(getNode.Port)
	fmt.Println(address)
	conn, gerr := grpc.Dial(address, grpc.WithInsecure())

	defer func() {
		if gerr != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", address, cerr)
			}
			return
		}
		go func() {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", address, cerr)
			}
		}()
	}()

	ctx := context.Background()
	client := common.NewCallClient(conn)

	//mux := runtime.NewServeMux()
	//req := &http.Request{}
	//inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
	//rctx, err := runtime.AnnotateContext(ctx, mux, req)
	var metadata runtime.ServerMetadata

	/*
		body := make(map[string]*any.Any)
		test := &any.Any{
			TypeUrl: "string",
			Value:   []byte("This is test"),
		}
		println("getVal", test.GetValue(), test)
		body["test"] = test
	*/

	v := make(map[string]*any.Any)
	fmt.Println(v)
	/*
		a := new(any.Any)
		a.Value = []byte("This is test")
		v["a"] = a
		v["b"] = a
	*/

	//cRequest := &common.CallRequest{
	//	Action: cmdCnt.Action,
	//	Header: cmdCnt.Header,
	//Body:   v,
	//}

	//var reply
	//err := conn.Invoke(ctx, cmdCnt.Action, cRequest, &reply)

	rsp, err := client.Action(ctx, cRequest, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	fmt.Println("rsp=:", rsp, err)

	//fmt.Println("test_get_service:", registrySer, rErr)
	//fmt.Println("random count:", rand.Intn(2))

	/*
		conn, err := grpc.Dial("")
		if err != nil {
			fmt.Println(err)
		}
		defer func() {
			if err != nil {
				if cerr := conn.Close(); cerr != nil {
					grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
				}
				return
			}
			go func() {
				<-ctx.Done()
				if cerr := conn.Close(); cerr != nil {
					grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
				}
			}()
		}()
	*/
	//callService := proto.NewCallService("dev.micro.cmd-async", service.Client())
	//res, err := callService.Action(ctx, &proto.CallRequest{
	//	Action: cmdCnt.Action,
	//	Header: cmdCnt.Header,
	//	Body:   cmdCnt.Body,
	//})

	//fmt.Println(microcmd.DefaultCmd.App())
	//fmt.Println(service.Server().Options().Registry)
	//fmt.Println(registry.DefaultRegistry.Options())

	//fmt.Println("call res:", res, err)
	//if err := service.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
