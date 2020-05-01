#!/bin/bash

cd cmd && \
go run main.go --registry=etcd --registry_address=192.168.0.112:9002,192.168.0.112:9004,192.168.0.112:9006
