package proxy

import (
	//"github.com/micro/go-micro/cmd"
	"context"
	//"fmt"
	"log"
	"strconv"
	"time"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/selector"
	"google.golang.org/grpc"
	common "srv/proto/common"
	"srv/services/conn_pool"
)

type RpcProxy struct {
	RegSelector selector.Selector
}

var (
	MapGrpcPool = make(map[string]*connpool.GrpcPool)
)

func NewGrpcProxy() *RpcProxy {
	//fmt.Println(registry.DefaultRegistry)
	r := (*cmd.DefaultCmd.Options().Registry)
	regSelector := selector.NewSelector(selector.Registry(r))

	return &RpcProxy{
		RegSelector: regSelector,
	}
}

func (p *RpcProxy) RequestByAddr(host string, port int, ctx context.Context, method string, request *common.CallRequest, opts ...grpc.CallOption) (*common.CallResponse, error) {
	addrs := host + ":" + strconv.Itoa(port)
	if _, ok := MapGrpcPool[addrs]; ok != true {
		MapGrpcPool[addrs] = connpool.NewGrpcPool(host, port)
	}
	grpcPool, _ := MapGrpcPool[addrs]
	rsp, rerr := grpcPool.Action(ctx, request, opts...)
	if rerr != nil {
		log.Println("[error]", "request_grpc_error", rerr, addrs, request)
	}
	return rsp, rerr
}

/**
 * 通过服务名称请求
 *
 * @param string sname consul中注册的服务名称
 * @param context.Context ctx 请求上下文
 * @param string method 请求方式
 * @param *common.CallRequest request 请求内容
 */
func (p *RpcProxy) Request(sname string, ctx context.Context, method string, request *common.CallRequest, opts ...grpc.CallOption) (*common.CallResponse, error) {
	var rsp *common.CallResponse
	var rerr error

	next, serr := p.RegSelector.Select(sname)
	log.Println("service", sname, next, serr)

	if serr != nil {
		log.Println("selector_select_service_error:", sname, serr)
		return nil, serr
	}

	for i := 0; i < 100; i++ {
		node, err := next()

		if err != nil {
			log.Println("[error] get service node err:", sname, err, "node=", node)
			continue
		}

		addrs := node.Address + ":" + strconv.Itoa(node.Port)
		if _, ok := MapGrpcPool[addrs]; ok != true {
			MapGrpcPool[addrs] = connpool.NewGrpcPool(node.Address, node.Port)
		}

		grpcPool, _ := MapGrpcPool[addrs]
		rsp, rerr = grpcPool.Action(ctx, request, opts...)
		log.Println("[notice]", "get_grpc_node", addrs)

		if rerr == nil {
			return rsp, nil
		} else {
			log.Println("[error]", "request_grpc_error", rerr, addrs, request)
		}

		time.Sleep(time.Millisecond)
	}

	return nil, rerr
}

func (p *RpcProxy) GetClient(addrs string) {
	grpc.Dial(addrs, grpc.WithInsecure())
}
