package asynccmd

import (
	pb "srv/proto/common"
)

type CmdMysql struct {
	ServiceName   string   `json:"service_name"`
	Sqls          []string `json:"sqls"`
	IsTransaction bool     `json:"is_transaction"`
}

type CmdHttp struct {
	ServiceName string              `json:"service_name"`
	Addr        string              `json:"addr"`
	Method      string              `json:"method"`
	Headers     map[string][]string `json:"headers"`
	Body        string              `json:"body"`
	Uri         string              `json:"uri"`
}

type CmdGrpc struct {
	ServiceName string         `json:"service_name"`
	Host        string         `json:"host"`
	Port        int            `json:"port"`
	Method      string         `json:"method"`
	Request     pb.CallRequest `json:"request"`
}

type CmdRedis struct {
	ServiceName string `json:"service_name"`
	SelectDb    int    `json:"select_db"`
	Cmd         string `json:"cmd"`
}

type CmdAgg struct {
	DoMysql []CmdMysql `json:"do_mysql"`
	DoGrpc  []CmdGrpc  `json:"do_grpc"`
	DoHttp  []CmdHttp  `json:"do_http"`
	DoRedis []CmdRedis `json:"do_redis"`
}

type CmdRepublishMsg struct {
	Header map[string]string
	Cmds   *CmdAgg
}
