#!/bin/bash

cd cmd && \
go run main.go --registry=etcd --registry_address=172.16.7.16:9002,172.16.7.16:9004,172.16.7.16:9006
