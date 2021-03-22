package main

import (
	"bytes"
	//"wdyedu.com/micro/cmd-async/srv/services/proxy"
	"fmt"
	//"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"srv/services/conn_pool"
	"time"
)

func TestPool() {
	httpProxy := connpool.NewHttpPool()
	header := make(http.Header)
	header.Set("test", "test")
	body := bytes.NewBuffer([]byte("Post this data"))
	//header.Set("Content-type", "application/json")

	//fmt.Println(io.Reader("11"))

	for i := 0; i < 100; i++ {
		rsp, err := httpProxy.Request("GET", header, "https://blog.csdn.net/u013870094/article/details/78731460/", body)
		log.Println("rsp", rsp, "err", err)

		rspCnt, rErr := ioutil.ReadAll(rsp.Body)
		fmt.Println("content-length:", rsp.ContentLength, ",body:", string(rspCnt), rErr, err)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	TestPool()
}
