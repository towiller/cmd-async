package main

import (
	"context"
	//"encoding/json"
	"fmt"
	"reflect"

	"srv/services/kv"
)

type contextKey struct{}

type registryKey string

type addContext struct {
	Topic      string        `json:"topic"`
	FatalTopic string        `json:"fatal_topic"`
	Cmds       []interface{} `json:"cmds"`
}

type TypeDecode func()

func main() {

	registryAddress := "192.168.1.18:8500"

	key := registryKey("one")

	a := &addContext{}
	consulData := kv.NewConsulKv(
		kv.Address(registryAddress),
	)
	consulData.GetJsonVal("cmd-async/events/event.sns_push", a)

	ctx := context.WithValue(context.Background(), key, a)
	aType := reflect.TypeOf(a).Elem()

	val := ctx.Value(key).(*addContext)

	//reflect
	fmt.Println(val)
	fmt.Println(aType)
}
