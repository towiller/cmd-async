package main

import (
	"context"
	"log"
	"time"

	"github.com/emicklei/go-restful"
	//"github.com/micro/go-micro/client"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-web"
	"github.com/micro/grpc-go/logger"

	cmdconfig "config"
	"encoding/json"
	"github.com/micro/go-micro/registry"
	"io/ioutil"
	common "srv/proto/common"
	"srv/services/proxy"
)

type RequestBody struct {
	EventName string                 `json:"event_name"`
	Cmds      map[string]interface{} `json:"cmds"`
	Topic     string                 `json:"topic"`
}

type RestCmd struct {
	ServiceName string
}

// @Summary 异步调用接口
// @Router
// swagger:route POST /async cmd-async
// swagger::
func (c *RestCmd) Action(req *restful.Request, rsp *restful.Response) {
	header := make(map[string]string)
	body := make(map[string]*any.Any)

	bodyByte, _ := ioutil.ReadAll(req.Request.Body)
	requestBody := &RequestBody{}
	jerr := json.Unmarshal(bodyByte, requestBody)
	if jerr != nil {
		log.Println("unmarshal_json_error", jerr)
		rspEntity := map[string]string{
			"status": "500",
			"msg":    jerr.Error(),
		}
		rsp.WriteEntity(rspEntity)
		return
	}

	log.Println("get_json_str", string(bodyByte[:]), requestBody)
	//log.Println(req.Request, req.Request.Body.Read())
	header["request_id"] = req.HeaderParameter("request_id")

	reqCmd, _ := json.Marshal(requestBody.Cmds)
	cmdVal, _ := proto.Marshal(&common.TypeString{
		Value: string(reqCmd[:]),
	})
	body["cmds"] = &any.Any{
		Value: cmdVal,
	}

	evtNameVal, _ := proto.Marshal(&common.TypeString{
		Value: requestBody.EventName,
	})
	body["event_name"] = &any.Any{
		Value: evtNameVal,
	}

	log.Println("eventName", body["event_name"])
	topicVal, _ := proto.Marshal(&common.TypeString{
		Value: requestBody.Topic,
	})
	body["topic"] = &any.Any{
		Value: topicVal,
	}

	grpcRequest := common.CallRequest{
		Action: "async_call",
		Header: header,
		Body:   body,
	}

	ctx := context.Background()
	method := "/cmd.async.srv.Call/Action"

	grpcProxy := proxy.NewGrpcProxy()
	grpcRsp, err1 := grpcProxy.Request(c.ServiceName, ctx, method, &grpcRequest)
	//log.Println(grpcRsp.GetContent()["id"])

	if err1 != nil {
		rspEntity := map[string]string{
			"status": "500",
			"msg":    err1.Error(),
		}
		rsp.WriteEntity(rspEntity)
		return
	}

	var gId int64
	if contentId, ok := grpcRsp.Content["id"]; ok {
		typeId := common.TypeInt64{}
		proto.Unmarshal(contentId.Value, &typeId)
		gId = typeId.Value
	}

	r := make(map[string]interface{})
	r["status"] = 200
	r["msg"] = "ok"
	r["content"] = map[string]int64{
		"id": gId,
	}
	rsp.WriteEntity(r)
	return
}

func main() {
	restServiceName := cmdconfig.DefaultEnv.AppEnv + "." + cmdconfig.DefaultEnv.AppRestName + "." + cmdconfig.DefaultEnv.AppVersion
	service := web.NewService(
		web.Name(restServiceName),
		web.Registry(consul.NewRegistry()),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)
	registry.DefaultRegistry = consul.NewRegistry()

	// Create RESTful handler
	restCmd := &RestCmd{
		ServiceName: cmdconfig.DefaultEnv.AppEnv + "." + cmdconfig.DefaultEnv.AppSrvName + "." + cmdconfig.DefaultEnv.AppVersion,
	}
	service.Init()

	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/async")
	ws.Route(ws.POST("/").To(restCmd.Action))
	wc.Add(ws)

	service.Handle("/", wc)

	if err := service.Run(); err != nil {
		logger.Error("error msg is ", err)
	}
}
