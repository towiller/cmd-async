#!/bin/bash

basepath=$(cd `dirname $0`;pwd)

export GOPATH=$basepath:$GOPATH

setArgsKey=("env" "registry" "registry_address")
setArgsDefault=("dev" "consul" "127.0.0.1:8500")
getArgs=$*
getArgValue=()

for ((i=0; i<${#setArgsKey[*]}; i++)); do
	setOption="--${setArgsKey[i]}"
	argVal=${setArgsDefault[i]}

	for getArg in $getArgs; do 
		argKey=${getArg%=*}

		if [[ "$setOption" == "$argKey" ]]; then 
			argVal=${getArg#*=}
		fi

		echo $argVal
	done

	getArgValue[i]=$argVal
done

env=${getArgValue[0]}

cp -rf $basepath/src/config/$env.env $basepath/src/.env

mkdir -p $basepath/src/golang.org/
mkdir -p $basepath/src/golang.org/x
if [ ! -d "$basepath/src/golang.org/x/net" ]; then
    git clone https://github.com/golang/net.git $basepath/src/golang.org/x/net
fi

if [ ! -d "$basepath/src/golang.org/x/sys" ]; then
    git clone https://github.com/golang/sys.git $basepath/src/golang.org/x/sys
fi

if [ ! -d "$basepath/src/golang.org/x/text" ]; then
    git clone https://github.com/golang/text.git $basepath/src/golang.org/x/text
fi

mkdir -p $basepath/src/google.golang.org
if [ ! -d "$basepath/src/google.golang.org/grpc" ]; then
    git clone https://github.com/grpc/grpc-go.git $basepath/src/google.golang.org/grpc
fi

if [ ! -d "$basepath/src/google.golang.org/genproto" ]; then
    git clone https://github.com/google/go-genproto.git $basepath/src/google.golang.org/genproto
fi

if [ ! -d "$basepath/src/golang.org/x/tools" ]; then
    git clone https://github.com/golang/tools.git $basepath/src/google.golang.org/x/tools
fi


venderList=(
    "github.com/gogo/protobuf/proto"
    "github.com/0x5010/grpcp"
    "github.com/Shopify/sarama"
    "github.com/garyburd/redigo"
    "github.com/gomodule/redigo/redis"
    "gopkg.in/bsm/sarama-cluster.v2"
    "gopkg.in/ini.v1"
    "github.com/silenceper/pool"
    "github.com/rcrowley/go-metrics"
    "github.com/google/uuid"
    "github.com/pierrec/lz4"
    "github.com/emicklei/go-restful"

    "github.com/bsm/sarama-cluster"
    "github.com/go-sql-driver/mysql"
    "gopkg.in/yaml.v2"


    "github.com/micro/go-log"
    "github.com/micro/go-micro"
    "github.com/micro/go-sync"
    "github.com/micro/go-web"
    "github.com/micro/grpc-go"
)

for ((i=0; i<${#venderList[*]}; i++)); do
     if [ ! -d "$basepath/src/${venderList[$i]}" ]; then
        echo go get -u ${venderList[$i]}
        go get -u ${venderList[$i]}
     fi
done

cd $basepath/src

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o async_consumer srv/consumer.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o async_grpc srv/async_grpc.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o async_rest rest/rest.go

mkdir -p $basepath/output

mv $basepath/src/async_consumer $basepath/output/async_consumer
mv $basepath/src/async_grpc  $basepath/output/async_grpc
mv $basepath/src/async_rest $basepath/output/async_rest

cp -rf $basepath/src/config $basepath/output/
cp $basepath/src/.env $basepath/output/
cp -rf $basepath/bin  $basepath/output/