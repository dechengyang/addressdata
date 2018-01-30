#!/usr/bin/env bash

go build main.go

mv main addressinfer

docker rmi luozhouyang/address-infer:0.1

docker build -t luozhouyang/address-infer:0.1 .

rm addressinfer

docker run -ti --name address-infer -p 8080:8080 luozhouyang/address-infer:0.1

