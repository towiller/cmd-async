#使用说明

>1.1 配置.env文件**

```

APP_ENV=dev  #环境
APP_VERSION=v1 #当前运行版本
APP_SRVNAME=micro.cmd-async.srv #服务名称
APP_RESTNAME=micro.cmd-async #restful服务名称
SRVNAME_IDGEN=micro.idgen.srv.v1  #依赖ID生成服务名称
KAFKA_ADDRESS=127.0.0.1:9091
KAFKA_ERR_TOPIC=cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.async_cmd_error
KAFKA_GROUP_ID=user_group_check
REDIS_ERR_TOPIC=dev.async_cmd_error
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_AUTH=
REDIS_DEFAULTDB=5
REDIS_DBCOUNT=16
```

> 1.2 配置consul的kafka消费topics

```
curl -X PUT  -H "Content-type:application/json"  -d '["cc7514bff90b4f0c9bc22787de7f6e77__topic.dev.cmd-async.common"]' "http://127.0.0.1:8500/v1/kv/cmd-async/topics"
```

>2 启动服务

>2.1 启动consumer服务(可以启动多个consumer服务)

```
go run srv/consumer.go --registry_address=127.0.0.1:8500
```

>2.2 启动grpc服务

```
go run srv/main.go --registry_address=127.0.0.1:8500
```

>2.3 如需调用http请求，启动restful接口服务

```$javascript
go run rest/rest.go --registry_address=127.0.0.1:8500
```

>3 调用说明 

>3.1 异步调用接口
>
>* 接口地址：/async/
>* 请求方式 POST
>* 参数说明

参数名称 | 数据类型 | 取值范围 | 说明
--: |:-----:|:-----:|---
event_name|string|string(1-32)|事件名称
topic_name|string|string(1-100)|topic名称
cmds|json|string(1-100)|执行命令、json字符串 
cmds.do_mysql|[]object|object|命令do_mysql
cmds.do_mysql.[].service_name|string|string|服务名称
cmds.do_mysql.[].sqls|[]string|[]string|执行sql
cmds.do_mysql.[].is_transacation|bool|true,false|是否为事务
cmds.do_http|[]object|object{}|http请求
cmds.do_http.[].service_name|string|string|服务名称:dev.micro.idgen.srv.v1
cmds.do_http.[].method|string|string|请求方式HTTP_METHOD
cmds.do_http.[].headers|map[string][]string||["Content-type":["application/json"]]
cmds.do_http.[].body|string|string|http请求body内容
cmds.do_http.[].uri|string|string|/do/anything?a=1&b=2
cmds.do_grpc.[].service_name|string|string|服务名称
cmds.do_grpc.[].method|string|string|请求方法
cmds.do_grpc.[].request|pb.CallRequest| pb.CallRequest|grpc请求结构体
cmds.do_redis.[].service_name|string|string|redis服务名称
cmds.do_redis.[].select_db|int|int|redis选择库
cmds.do_redis.[].cmd|string|string|redis执行命令

>* 返回数据:

```
{"status":100,"msg":"ok","content":{"id":1000}}
```

>* 请求示例

```
curl -X POST -H "Content-type:application/json" -d '{"do_mysql":[{"service_name":"mysql.dev.saas","sqls":["insert into saas.coupon (`shop_id`, `coupon_type`, `name`, `money`, `use_type`, `created_at`, `updated_at`) values(\"158205\", \"1\", \"\u6d4b\u8bd5\", \"1\", \"1\", current_timestamp(), current_timestamp())"],"is_transaction":true}],"do_grpc":[{"cmd_name":"do_grpc","service_name":"dev.micro.idgen.srv.v1","method":"\/go.micro.srv.idgen.Call\/Action","request":{"action":"GetId","header":{"header1":"headerCnt"},"body":{"type":{"type_url":"type_int32","value":"CAE="}}}}],"do_http":[{"service_name":"go.micro.api.example","method":"GET","headers":{"Content-type":["application\/json;charset=utf8"]},"body":"asssssss","uri":"\/greeter?name=22"}],"do_redis":[{"service_name":"redis.dev.saas","select_db":1,"cmd":"incrby cmd-async_incrby 1"}]}' http://127.0.0.1:50358/async/
```
>3.2 事务处理 - http
>
> * 接口地址: /tcc/http
> 
> * 请求方式: POST
> 
> * 参数说明
> 
> 参数名称|参数类型|取值范围|说明
> --: |:-----:|:-----:|---
> event_name|string|string(1-64)|事件名称
> version|string|string|版本号用于锁定
> topic|string|string|string(1-256)|topic名称
> transations|string|string|事务数据
> transations.[].service_name|string||consul中注册的服务名称
> transations.[].header|map[string]string||头信息header
> 


>4 设置mysql、redis配置
>>4.1 mysql配置修改

```
{
    "host":"127.0.0.1",   //服务请求地址
    "port":3306,             //端口
    "username":"root",       //用户名
    "password":"123456",//密码
    "default_db":"saas",     //默认db
    "charset":"utf8",        //编码
    "max_open_conns":1000,   //最高连接数
    "max_idle_conns":500     //最大闲置连接数
}
```

```
curl -X PUT -d '{"host":"127.0.0.1","port":3306,"username":"root","password":"111","default_db":"saas","charset":"utf8","max_open_conns":1000,"max_idle_conns":500}'  "http://127.0.0.1:8500/v1/kv/services_config/mysql/mysql.dev.saas"
```

>>4.2 redis配置修改

```
{
    "host":"192.168.1.18",   //服务地址
    "port":6379,             //端口号
    "password":"",           //密码
    "max_idle":300,          //最大连接数
    "idle_timeout":100,      //最大闲置时间
}
```

```
curl -X PUT -d '{"host":"127.0.0.1","port":6379,"password":"","max_idle":1000,"idle_timeout":100,"select_db":0}' "http://127.0.0.1:8500/v1/kv/services_config/redis/redis.dev.saas"
```
