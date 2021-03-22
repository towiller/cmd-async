package main

import (
	//"fmt"
	//"wdyedu.com/micro/cmd-async/srv/services/cmd"
	//"wdyedu.com/micro/cmd-async/srv/services/kv"
	"encoding/json"
	"fmt"
	"time"

	"bytes"
	"context"
	"database/sql"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	pb "srv/proto/common"
	cmd "srv/services/cmd"
	"srv/services/proxy"
	"strings"
)

var (
	mapGrpcType = map[string]interface{}{
		"type_bool":     &pb.TypeBool{},
		"type_string":   &pb.TypeString{},
		"type_int32":    &pb.TypeInt32{},
		"type_uint32":   &pb.TypeUint32{},
		"type_sfixed32": &pb.TypeSfixed32{},
		"type_int64":    &pb.TypeInt64{},
		"type_uint64":   &pb.TypeUint64{},
		"type_sint64":   &pb.TypeSint64{},
		"type_sfixed64": &pb.TypeSfixed64{},
		"type_double":   &pb.TypeDouble{},
		"type_float":    &pb.TypeFloat{},
		"type_bytes":    &pb.TypeBytes{},
	}
)

type M struct {
	A []string
	B []int
}

func TestGrpcIdGenCmd() {
	log.Println("testGrpc:")
	cmdStr := "{\"service_name\":\"dev.micro.cmd-async\",\"method\":\"\\/cmd.async.srv.Call\\/Action\",\"request\":{\"action\":\"GetId\",\"header\":{\"header1\":\"headerCnt\"},\"body\":{\"type\":{\"type_url\":\"type_int32\",\"value\":\"CAE=\"}}}}"
	grpcCmd := &cmd.CmdGrpc{}
	json.Unmarshal([]byte(cmdStr), grpcCmd)

	ctx := context.Background()
	grpcProxy := proxy.NewGrpcProxy()
	rsp, err := grpcProxy.Request(grpcCmd.ServiceName, ctx, grpcCmd.Method, &grpcCmd.Request)
	fmt.Println("cmd", grpcCmd, "rsp", rsp, "error", err)
	log.Println("testGrpc end")
}

func TestHttpCallCmd() {
	log.Println("testHttp")
	cmdStr := "{\"service_name\":\"go.micro.api.greeter\",\"method\":\"GET\",\"headers\":{\"Content-type\":[\"application\\/json;charset=utf8\"]},\"body\":\"asssssss\",\"uri\":\"\\/greeter?name=22\"}"
	httpCmd := &cmd.CmdHttp{}
	json.Unmarshal([]byte(cmdStr), httpCmd)

	httpProxy := proxy.NewHttpProxy()
	body := bytes.NewBuffer([]byte(httpCmd.Body))
	//body := bytes.NewBuffer([]byte("Post this data"))

	rsp, err := httpProxy.Request(httpCmd.ServiceName, httpCmd.Method, httpCmd.Headers, httpCmd.Uri, body)
	rspCnt, rErr := ioutil.ReadAll(rsp.Body)
	fmt.Println("body:", rsp.Body, err, "rspCnt:", string(rspCnt), rErr)
	log.Println("testGrpc end")
}

func TestMysqlCallCmd() {
	log.Println("test mysql")
	cmdStr := "{\"service_name\":\"mysql.dev.saas\",\"sqls\":[\"insert into saas.coupon (`shop_id`, `coupon_type`, `name`, `money`, `use_type`, `created_at`, `updated_at`) values('158205', '1', '\\u6d4b\\u8bd5', '1', '1', current_timestamp(), current_timestamp())\"],\"is_transaction\":true}"
	mysqlCmd := &cmd.CmdMysql{}
	json.Unmarshal([]byte(cmdStr), mysqlCmd)

	fmt.Println(mysqlCmd)

	mysqlProxy := proxy.NewMysqlProxy()
	var transaction *sql.Tx
	var terr error
	if mysqlCmd.IsTransaction {
		transaction, terr = mysqlProxy.Begin(mysqlCmd.ServiceName)
		if terr != nil {
			log.Println("transation start error")
		}
	}

	for _, sql := range mysqlCmd.Sqls {
		result, err := mysqlProxy.Exec(mysqlCmd.ServiceName, sql)
		if err != nil {
			transaction.Rollback()
		}
		log.Println(result.RowsAffected())
	}

	if mysqlCmd.IsTransaction && terr == nil {
		transaction.Commit()
	}

	log.Println("test mysql end")
}

func TestRedisCmd() {
	log.Println("test redis")
	cmdStr := "{\"service_name\":\"redis.dev.saas\",\"select_db\":1,\"cmd\":\"incrby cmd-async_incrby 1\"}"
	redisCmd := &cmd.CmdRedis{}
	json.Unmarshal([]byte(cmdStr), redisCmd)

	redisProxy := proxy.NewRedisProxy()
	cmdArr := strings.Split(redisCmd.Cmd, " ")

	cmdInterface := make([]interface{}, len(cmdArr))
	for k, v := range cmdArr {
		fmt.Println("k=", k)
		cmdInterface[k] = v
	}

	ret, err := redisProxy.Do(redisCmd.ServiceName, strings.ToUpper(cmdArr[0]), cmdInterface[1:]...)
	fmt.Println(ret, err)
	log.Println("test redis end")
}

func TestGrpcCmd() {
	typeInt := &pb.TypeInt32{
		Value: 1,
	}
	intval, _ := proto.Marshal(typeInt)

	body1 := make(map[string]*any.Any)
	body1["type"] = &any.Any{
		TypeUrl: "type_int32",
		Value:   intval,
	}

	req := pb.CallRequest{
		Action: "GetId",
		Header: map[string]string{"test": "test"},
		Body:   body1,
	}

	result, _ := json.Marshal(req)
	fmt.Println("result", string(result))

	req1 := pb.CallRequest{}
	//json.Unmarshal(result, &req1)

	cmdStr := "{\"action\":\"GetId\",\"header\":{\"header1\":\"headerCnt\"},\"body\":{\"type\":{\"type_url\":\"type_int32\",\"value\":\"CAE=\"}}}"
	json.Unmarshal([]byte(cmdStr), &req1)
	fmt.Println("req_body_cnt", req1.Body["type"].Value)

	getTypeVal := pb.TypeInt32{}
	proto.Unmarshal(req1.Body["type"].Value, &getTypeVal)

	fmt.Println("type=", getTypeVal.Value)
}

func TestCmdAgg() {

	//proto.Marshal()
}

func main() {
	service := micro.NewService(
		micro.Name("dev.micro.cmd-async11"),
		micro.Version("v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	TestGrpcIdGenCmd()
	TestGrpcCmd()
	TestHttpCallCmd()
	TestRedisCmd()
	TestMysqlCallCmd()

	TestCmdAgg()

	m := &M{
		A: []string{"111"},
		B: []int{1},
	}
	log.Println(m.A)
}
