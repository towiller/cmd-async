package proxy

import (
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/selector"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"srv/services/conn_pool"
	"strconv"
)

type HttpProxy struct {
	Pool        *connpool.HttpPool
	RegSelector selector.Selector
}

func NewHttpProxy() *HttpProxy {
	httpPool := connpool.NewHttpPool()
	r := (*cmd.DefaultCmd.Options().Registry)
	regSelector := selector.NewSelector(selector.Registry(r))

	return &HttpProxy{
		Pool:        httpPool,
		RegSelector: regSelector,
	}
}

func (p *HttpProxy) RequestByAddr(addr string, method string, header map[string][]string, uri string, body io.Reader) (*http.Response, error) {
	url := addr + uri
	rsp, rerr := p.Pool.Request(method, header, url, body)
	if rerr == nil {
		strBody, _ := ioutil.ReadAll(body)
		log.Println("http_request_err:", "address:", addr, "method:", method, "url", url, "body", string(strBody))
		return rsp, rerr
	}

	return nil, rerr
}

func (p *HttpProxy) Request(serviceName string, method string, header map[string][]string, uri string, body io.Reader) (*http.Response, error) {

	var rsp *http.Response
	var rerr error
	next, serr := p.RegSelector.Select(serviceName)
	if serr != nil {
		log.Println("selector select service error:", serviceName, serr)
		return nil, serr
	}

	for i := 0; i < 100; i++ {
		node, err := next()
		if err != nil {
			continue
		}

		addrs := "http://" + node.Address + ":" + strconv.Itoa(node.Port)
		url := addrs + uri
		rsp, rerr = p.Pool.Request(method, header, url, body)
		if rerr == nil {
			strBody, _ := ioutil.ReadAll(body)
			log.Println("http_request_err:", "service_name:", serviceName, "method:", method, "url", url, "body", string(strBody))
			return rsp, rerr
		}
	}

	return nil, rerr
}
