package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"github.com/golang/protobuf/ptypes/any"
	common "srv/proto/common"
	"srv/services/proxy"

	"github.com/gogo/protobuf/proto"

	"context"
	"fmt"
	"log"
	"time"
)

func TestAny() {
	header := make(map[string]string)
	body := make(map[string]*any.Any)
	body["type"] = &any.Any{Value: []byte("Test")}

	req := common.CallRequest{
		Action: "GetId",
		Header: header,
		Body:   body,
	}

	m, _ := proto.Marshal(&req)
	fmt.Println(req, m)

	typeInt := &common.TypeInt32{
		Value: 1,
	}
	intval, _ := proto.Marshal(typeInt)

	body1 := make(map[string]*any.Any)
	body1["type"] = &any.Any{
		TypeUrl: "TypeInt32",
		Value:   intval,
	}
	fmt.Println("type_content:", body1["type"], "msgType", proto.MessageName(typeInt))

	req1 := common.CallRequest{
		Action: "GetId",
		Header: header,
		Body:   body1,
	}
	m1, _ := proto.Marshal(&req1)
	fmt.Println(req1, m1)

	readMsg := common.CallRequest{}
	proto.Unmarshal(m1, &readMsg)
	fmt.Println("unmarshal", readMsg)

	getBody := readMsg.Body
	if getType, ok := getBody["type"]; ok {
		getTypeVal := &common.TypeInt32{}
		fmt.Println("getType", proto.Unmarshal(getType.Value, getTypeVal), getTypeVal.Value)
	}
}

func TestCallAsyncCmd() {
	//topic := "cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.cmd-async.common"
	//eventName := "test_cmd"

	serviceName := "dev.micro.cmd-async.srv.v1"
	method := "/cmd.async.srv.Call/Action"
	header := make(map[string]string)
	header["request_id"] = "11111111"

	body := make(map[string]*any.Any)
	topicVal, _ := proto.Marshal(&common.TypeString{
		Value: "cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.cmd-async.common",
	})
	body["topic"] = &any.Any{
		Value: topicVal,
	}

	eventNameVal, _ := proto.Marshal(&common.TypeString{
		Value: "test_cmd",
	})
	body["event_name"] = &any.Any{
		Value: eventNameVal,
	}

	cmdsVal, _ := proto.Marshal(&common.TypeString{
		Value: "{\"do_mysql\":[{\"service_name\":\"mysql.dev.saas\",\"sqls\":[\"insert into saas.coupon (`shop_id`, `coupon_type`, `name`, `money`, `use_type`, `created_at`, `updated_at`) values(\\\"158205\\\", \\\"1\\\", \\\"\\u6d4b\\u8bd5\\\", \\\"1\\\", \\\"1\\\", current_timestamp(), current_timestamp())\"],\"is_transaction\":true}],\"do_grpc\":[{\"cmd_name\":\"do_grpc\",\"service_name\":\"dev.micro.idgen.srv.v1\",\"method\":\"\\/go.micro.srv.idgen.Call\\/Action\",\"request\":{\"action\":\"GetId\",\"header\":{\"header1\":\"headerCnt\"},\"body\":{\"type\":{\"type_url\":\"type_int32\",\"value\":\"CAE=\"}}}}],\"do_http\":[{\"service_name\":\"go.micro.api.example\",\"method\":\"GET\",\"headers\":{\"Content-type\":[\"application\\/json;charset=utf8\"]},\"body\":\"asssssss\",\"uri\":\"\\/greeter?name=22\"}],\"do_redis\":[{\"service_name\":\"redis.dev.saas\",\"select_db\":1,\"cmd\":\"incrby cmd-async_incrby 1\"}]}",
	})
	body["cmds"] = &any.Any{
		Value: cmdsVal,
	}

	req := common.CallRequest{
		Action: "",
		Header: header,
		Body:   body,
	}

	ctx := context.Background()
	grpcProxy := proxy.NewGrpcProxy()
	log.Println(serviceName, ctx, method, req)
	rsp, err := grpcProxy.Request(serviceName, ctx, method, &req)
	log.Println("test_grpc_asyncmd", rsp, err)

	mysqlCmds := "{\"do_grpc\":null,\"do_http\":null,\"do_redis\":null,\"do_mysql\":{\"service_name\":\"mysql.dev.saas\",\"sqls\":[\"insert into saas.coupon (`shop_id`, `coupon_type`, `name`, `money`, `use_type`, `created_at`, `updated_at`) values(\\\"158205\\\", \\\"1\\\", \\\"\\u6d4b\\u8bd5\\\", \\\"1\\\", \\\"1\\\", current_timestamp(), current_timestamp())\"],\"is_transaction\":true}}"
	mysqlCmdValues, _ := proto.Marshal(&common.TypeString{
		Value: mysqlCmds,
	})
	body["cmds"] = &any.Any{
		Value: mysqlCmdValues,
	}

	rsp1, err1 := grpcProxy.Request(serviceName, ctx, method, &req)
	log.Println("test_grpc_asyncmd", rsp1, err1)
	/*
		req := &asynccmd.CmdsContent{
			Topic:topic,
		}
	*/
	/*
		cmdCnt := &asynccmd.CmdsContent{
			Topic:topic,
		}
	*/
}

func TestMysqlCmd() {
	serviceName := "dev.micro.cmd-async.v1"
	method := "/cmd.async.srv.Call/Action"
	cmds := "{\"do_grpc\":null,\"do_http\":null,\"do_redis\":null,\"do_mysql\":{\"service_name\":\"mysql.dev.saas\",\"sqls\":[\"insert into saas.coupon (`shop_id`, `coupon_type`, `name`, `money`, `use_type`, `created_at`, `updated_at`) values(\\\"158205\\\", \\\"1\\\", \\\"\\u6d4b\\u8bd5\\\", \\\"1\\\", \\\"1\\\", current_timestamp(), current_timestamp())\"],\"is_transaction\":true}}"

	header := make(map[string]string)
	body := make(map[string]*any.Any)

	cmdsVal, _ := proto.Marshal(&common.TypeString{
		Value: cmds,
	})
	body["cmds"] = &any.Any{
		Value: cmdsVal,
	}
	req := common.CallRequest{
		Action: "",
		Header: header,
		Body:   body,
	}

	ctx := context.Background()
	grpcProxy := proxy.NewGrpcProxy()
	rsp, err := grpcProxy.Request(serviceName, ctx, method, &req)
	log.Println("test_grpc_asyncmd", rsp, err)
	log.Println("testMysqlRequest", req, "body", body)
}

func init() {

}

func main() {
	//TestAny()
	service := micro.NewService(
		micro.Name("dev.micro.cmd-async11"),
		micro.Version("v1"),
		micro.Registry(consul.NewRegistry()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	TestCallAsyncCmd()
	//TestMysqlCmd()
	return

	registry.DefaultRegistry = consul.NewRegistry()

	fmt.Println((*cmd.DefaultCmd.Options().Registry))

	fmt.Println("default registry:", registry.DefaultRegistry)

	serviceName := "dev.micro.cmd-async"
	fmt.Println(serviceName)

	method := "common.service.srv.Call/Action"
	ctx := context.Background()
	header := make(map[string]string)
	header["test"] = "test"
	requestCnt := &common.CallRequest{
		Action: "send.call",
		Header: header,
	}

	log.Println("ctx:", ctx, "method:", method, "request_cnt:", requestCnt)
	grpcProxy := proxy.NewGrpcProxy()
	rsp, err := grpcProxy.Request(serviceName, ctx, method, requestCnt)

	fmt.Println(rsp, err)
}
