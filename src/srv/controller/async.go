package controller

import (
	"context"
	"log"
	//"github.com/micro/go-micro"

	cmdconfig "config"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	pb "srv/proto"
	commonpb "srv/proto/common"
	"srv/services/bd_broker"
	"srv/services/proxy"
	"wdyedu.com/micro/cmd-async/src/srv/services/bd_broker"
)

type CmdAsync struct {
}

func NewCmdAsync() (*CmdAsync, error) {
	return &CmdAsync{}, nil
}

func (c *CmdAsync) Action(ctx context.Context, req *pb.CallRequest, res *pb.CallResponse) error {

	//requestId := req.Header["request_id"]

	topic := pb.TypeString{}
	proto.Unmarshal(req.Body["topic"].Value, &topic)

	eventName := pb.TypeString{}
	proto.Unmarshal(req.Body["event_name"].Value, &eventName)

	cmds := pb.TypeString{}
	proto.Unmarshal(req.Body["cmds"].Value, &cmds)

	err := bdbroker.DefaultBrokerPool.PublishMsg(topic.Value, req.Header, []byte(cmds.Value))
	if err != nil {
		res.Status = 500
		res.Msg = "publish msg error"
		log.Println("[request.action_error]", err)
		return nil
	}

	rId := c.MakeRequestId(ctx)
	log.Println("generate_request_id", rId)
	res.Status = 200
	res.Msg = "ok"
	typeId, _ := proto.Marshal(&commonpb.TypeInt64{Value: rId})
	res.Content = make(map[string]*any.Any)
	res.Content["id"] = &any.Any{
		Value: typeId,
	}
	log.Println("push msg content:", res)

	return nil
}

func (c *CmdAsync) MakeRequestId(ctx context.Context) int64 {
	//micro.Action()

	serviceName := cmdconfig.DefaultEnv.AppEnv + "." + cmdconfig.DefaultEnv.SrvNameIdGen
	method := "/go.micro.srv.idgen.Call/Action"
	header := make(map[string]string)

	intval, _ := proto.Marshal(&commonpb.TypeInt32{
		Value: 1,
	})
	body := make(map[string]*any.Any)
	body["type"] = &any.Any{
		Value: intval,
	}
	requestCnt := &commonpb.CallRequest{
		Action: "GetId",
		Header: header,
		Body:   body,
	}

	grpcProxy := proxy.NewGrpcProxy()
	rsp, err := grpcProxy.Request(serviceName, ctx, method, requestCnt)
	if err != nil {
		log.Println("get request id error", rsp, err)
		return 0
	}
	log.Println("idgen rsp=", rsp)

	var id commonpb.TypeInt64
	proto.Unmarshal(rsp.Content["id"].Value, &id)

	return id.Value
}
