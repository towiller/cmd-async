package main

import (
	"fmt"
	//"net/url"
	"config"
	"srv/services/proxy"
)

func main() {
	httpProxy := proxy.NewHttpProxy()
	serviceName := config.DefaultEnv.AppRestName + "." + config.DefaultEnv.AppVersion
	header := map[string]string{
		"request_id": "112311",
	}

	//url.Values{
	//	"event_name": {""},
	//}
	fmt.Println(httpProxy, serviceName, header)

	//httpProxy.Request(serviceName, "POST"，header,"11")
	//httpProxy.Request("")
}
