#!/bin/bash
PWD=`pwd`
#echo $PWD
export GOPATH=$GOPATH:$PWD
#echo $GOPATH
go build src/main/test.go
