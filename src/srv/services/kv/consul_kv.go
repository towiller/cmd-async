package kv

import (
	"context"
	"encoding/json"
	"log"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-sync/data"
	"github.com/micro/go-sync/data/consul"
)

type ConsulKv struct {
	opts       Options
	consulData data.Data
}

type JsonContentKey string

type ParseToStruct func(retJson interface{}, ret interface{})

var (
	_instance       *ConsulKv
	DefaultConsulKv = NewConsulKv()
)

func ConsulKvInstnce(opts ...Option) *ConsulKv {
	if _instance == nil {
		_instance = NewConsulKv(opts...)
	}

	return _instance
}

func NewConsulKv(opts ...Option) *ConsulKv {

	defaultRegistryOptions := (*cmd.DefaultCmd.Options().Registry).Options()
	log.Println("default_registry", (*cmd.DefaultCmd.Options().Registry).Options())
	var defaultAddr string
	if len(defaultRegistryOptions.Addrs) < 1 {
		defaultAddr = "127.0.0.1:8500"
	} else {
		defaultAddr = defaultRegistryOptions.Addrs[0]
	}

	options := Options{
		CheckUser: false,
		Context:   context.Background(),
		Address:   defaultAddr,
	}

	for _, o := range opts {
		o(&options)
	}

	cData := consul.NewData(data.Nodes(options.Address))
	return &ConsulKv{
		opts:       options,
		consulData: cData,
	}
}

func (c *ConsulKv) GetVal(key string) *data.Record {
	readData, err := c.consulData.Read(key)
	if err != nil {
		log.Println("[error]", "get_consul_val_fatal, [key=", key, "]", "[err=", err, "]")
	}

	return readData
}

func (c *ConsulKv) GetJsonVal(keyName string, retJson interface{}) {

	getKey := JsonContentKey(keyName)
	getVal := c.GetContextVal(getKey)

	if getVal == nil {
		readData := c.GetVal(keyName)
		log.Println(readData)
		err := json.Unmarshal(readData.Value, &retJson)
		if err != nil {
			log.Println("[error] consul_get_err:", err, "keyname=", keyName)
		}

		SetJsonContext(keyName, retJson)
		return
	}

	return
}

func (c *ConsulKv) PutVal(key string, val interface{}) error {

	jsonVal, _ := json.Marshal(val)
	record := &data.Record{
		Key:   key,
		Value: jsonVal,
	}
	return c.consulData.Write(record)
}

func (c *ConsulKv) DeleteVal(key string) error {
	return c.consulData.Delete(key)
}

func (c *ConsulKv) GetContextVal(key interface{}) interface{} {
	return c.opts.Context.Value(key)
}
