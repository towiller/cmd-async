#!/bin/sh

basepath=$(cd `dirname $0`; cd ..; pwd)

export GOPATH=$basepath:$GOPATH

go get -u gopkg.in/yaml.v2

which swg | awk -F ":" '{if($2 != ""|| $1=="swg not found"){print "false"}else{print "true"}}'

#echo $?