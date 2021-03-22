package asynccmd

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"config"
	"net/http"
	"srv/proto/common"
	"srv/services/conn_pool"
	"srv/services/kv"
	"srv/services/proxy"
	"srv/services/queue"
	"strconv"
)

//type Cmd interface {
//	DoCmd(cmdContent interface{})
//}

type AsyncCmd struct {
	kv         *kv.ConsulKv
	errTopic   string
	redisQueue *cmdqueue.RedisQueue
}

var (
	CmdPrefix      = "cmd-async/events/"
	MapRetrySecond = map[int]int64{
		1: 5,
		2: 30,
		3: 300,
		4: 1800,
		5: 3600,
		6: 3600 * 24,
	}
	//CmdMap    = make(map[string]interface{})
	//CmdMap["cmdsql"] = new CmdMysql()
)

func NewCmd(errTopic string, opts ...kv.Option) *AsyncCmd {
	consulKv := kv.NewConsulKv(opts...)

	return &AsyncCmd{
		kv:       consulKv,
		errTopic: errTopic,
		redisQueue: cmdqueue.NewRedisQueue(
			connpool.RedisHost(config.DefaultEnv.RedisHost),
			connpool.RedisPort(config.DefaultEnv.RedisPort),
			connpool.RedisDbCount(config.DefaultEnv.RedisDBCount),
		),
	}
}

func (c *AsyncCmd) ExecCmds(cmds *CmdAgg, headers map[string]string) {

	errCmds := &CmdAgg{}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go c.ExecGrpcCmd(cmds.DoGrpc, headers, wg)
	go c.ExecHttpCmd(cmds.DoHttp, headers, wg)
	go c.ExecRedisCmd(cmds.DoRedis, headers, wg)
	go c.ExecMysqlCmd(cmds.DoMysql, headers, wg)
	wg.Wait()

	fmt.Println(errCmds)
}

func (c *AsyncCmd) GetCmdConfig(name string) *CmdsContent {
	cmdContent := &CmdsContent{}
	cmdKey := CmdPrefix + name
	c.kv.GetJsonVal(cmdKey, cmdContent)

	return cmdContent
}

func (c *AsyncCmd) ExecGrpcCmd(grpcCmds []CmdGrpc, headers map[string]string, wg *sync.WaitGroup) {
	if len(grpcCmds) < 1 {
		log.Println("cmd_continue_grpc")
		wg.Done()
		return
	}

	errCmds := &CmdAgg{
		DoGrpc: []CmdGrpc{},
	}

	grpcProxy := proxy.NewGrpcProxy()
	ctx := context.Background()

	for _, grpcCmd := range grpcCmds {
		var rsp *common_service_srv.CallResponse
		var err error
		if grpcCmd.Host != "" && grpcCmd.Port > 0 {
			rsp, err = grpcProxy.RequestByAddr(grpcCmd.Host, grpcCmd.Port, ctx, grpcCmd.Method, &grpcCmd.Request)
		} else {
			rsp, err = grpcProxy.Request(grpcCmd.ServiceName, ctx, grpcCmd.Method, &grpcCmd.Request)
		}

		if err != nil {
			errCmds.DoGrpc = append(errCmds.DoGrpc, grpcCmd)
			log.Println("[error] cmd_grpc_error", "err=", err.Error(), "heades=", headers, grpcCmd.ServiceName, "method=", grpcCmd.Method, "reuest=", grpcCmd.Request, "response=", rsp)
		} else {
			log.Println("[notice] cmd_grpc_success", grpcCmd.ServiceName, "heades=", headers, "method=", grpcCmd.Method, grpcCmd.Request, "response=", rsp)
		}
	}

	if len(errCmds.DoGrpc) > 0 {
		go c.rePublishCmds(errCmds, headers)
	}
	wg.Done()
}

func (c *AsyncCmd) ExecHttpCmd(httpCmds []CmdHttp, headers map[string]string, wg *sync.WaitGroup) {
	if len(httpCmds) < 1 {
		log.Println("cmd_continue_http")
		wg.Done()
		return
	}

	errCmds := &CmdAgg{
		DoHttp: []CmdHttp{},
	}

	httpProxy := proxy.NewHttpProxy()

	for _, httpCmd := range httpCmds {
		var rsp *http.Response
		var err error
		body := bytes.NewBuffer([]byte(httpCmd.Body))

		if httpCmd.Addr != "" {
			rsp, err = httpProxy.RequestByAddr(httpCmd.Addr, httpCmd.Method, httpCmd.Headers, httpCmd.Uri, body)
		} else {
			rsp, err = httpProxy.Request(httpCmd.ServiceName, httpCmd.Method, httpCmd.Headers, httpCmd.Uri, body)
		}

		if err != nil {
			errCmds.DoHttp = append(errCmds.DoHttp, httpCmd)
			log.Println("[error] cmd_http_error", "heades=", headers, "error=", err.Error(), "http=", httpCmd, "rsp=", rsp)
		} else {
			log.Println("[notice] cmd_http_success", "heades=", headers, "http=", httpCmd)
		}

	}

	if len(errCmds.DoHttp) > 0 {
		log.Println("[error] http_request_err", "heades=", headers, errCmds.DoHttp)
		go c.rePublishCmds(errCmds, headers)
	}
	wg.Done()
	return
}

func (c *AsyncCmd) ExecRedisCmd(redisCmds []CmdRedis, headers map[string]string, wg *sync.WaitGroup) {
	if len(redisCmds) < 1 {
		log.Println("cmd_continue_redis")
		wg.Done()
		return
	}

	//异常的redis命令
	errCmds := &CmdAgg{
		DoRedis: []CmdRedis{},
	}

	redisProxy := proxy.NewRedisProxy()
	for _, redisCmd := range redisCmds {
		cmdArr := strings.Split(redisCmd.Cmd, " ")
		redisProxy.Do(redisCmd.ServiceName, "SELECT ", redisCmd.SelectDb)
		cmdInterface := make([]interface{}, len(cmdArr))
		for k, v := range cmdArr {
			cmdInterface[k] = v
		}

		ret, err := redisProxy.Do(redisCmd.ServiceName, strings.ToUpper(cmdArr[0]), cmdInterface[1:]...)
		if err != nil {
			errCmds.DoRedis = append(errCmds.DoRedis, redisCmd)
			log.Println("[error] cmd_redis_error", "heades=", headers, "cmd_arr=", cmdArr, "ret=", ret, "err=", err)
		} else {
			log.Println("[notice] cmd_redis_success", "heades=", headers, "cmd_arr=", cmdArr, "ret=", ret)
		}
	}

	if len(errCmds.DoRedis) > 0 {
		log.Println("[error] redis_request_error", "heades=", headers, errCmds.DoRedis)
		go c.rePublishCmds(errCmds, headers)
	}
	wg.Done()
	return
}

func (c *AsyncCmd) ExecMysqlCmd(mysqlCmds []CmdMysql, headers map[string]string, wg *sync.WaitGroup) {
	if len(mysqlCmds) < 1 {
		log.Println("continue mysql")
		wg.Done()
		return
	}

	errCmds := &CmdAgg{
		DoMysql: []CmdMysql{},
	}
	mysqlProxy := proxy.NewMysqlProxy()
	var isStartTransation bool

	for _, mysqlCmd := range mysqlCmds {
		var transaction *sql.Tx
		var terr error
		if mysqlCmd.IsTransaction && !isStartTransation {
			transaction, terr = mysqlProxy.Begin(mysqlCmd.ServiceName)
			if terr != nil {
				log.Println("[error] cmd_mysql", "heades=", headers, "transation start error")
				errCmds.DoMysql = mysqlCmds
				go c.rePublishCmds(errCmds, headers)
				wg.Done()
				return
			}
		}

		for _, sql := range mysqlCmd.Sqls {
			result, err := mysqlProxy.Exec(mysqlCmd.ServiceName, sql)
			if err != nil {
				log.Println("[error]", "exec_sql_error", "heades=", headers, "sql=", sql, "err=", err.Error())

				if mysqlCmd.IsTransaction {
					errCmds.DoMysql = mysqlCmds
					transaction.Rollback()
					go c.rePublishCmds(errCmds, headers)
					wg.Done()
					return
				} else {
					errCmds.DoMysql = append(errCmds.DoMysql, mysqlCmd)
				}
			}

			affectedRow, _ := result.RowsAffected()
			lastId, _ := result.LastInsertId()

			log.Println("[waring]", "exec_sql_success", "heades=", headers, "sql=", sql, "affected_rows=", affectedRow, "last_id=", lastId)
		}

		if mysqlCmd.IsTransaction && terr == nil {
			log.Println("[notice] [cmd_mysql]", "heades=", headers, "transation commit")
			transaction.Commit()
		}
	}

	if len(errCmds.DoMysql) > 0 {
		log.Println("[error] mysql_request_error", "heades=", headers, errCmds.DoMysql)
		go c.rePublishCmds(errCmds, headers)
	}
	wg.Done()
}

func (c *AsyncCmd) rePublishCmds(errCmds *CmdAgg, headers map[string]string) error {
	rCount, ok := headers["retry_count"]
	var retryCount int
	if ok {
		retryCount, _ = strconv.Atoi(rCount)
		retryCount++
	} else {
		retryCount = 1
	}

	trySecond, tok := MapRetrySecond[retryCount]
	if !tok {
		log.Println("[error] stop_retry_publish", headers, errCmds)
		return nil
	}

	headers["retry_count"] = strconv.Itoa(retryCount)
	msg := CmdRepublishMsg{
		Header: headers,
		Cmds:   errCmds,
	}
	err := c.redisQueue.DelayPublish(config.DefaultEnv.KafkaErrTopic, msg, trySecond)

	if err != nil {
		log.Println("[error] publish_msg_error ", "err", err, "repush_msg", errCmds)
	}

	return err
}
