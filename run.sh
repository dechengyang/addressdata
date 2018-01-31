#!/usr/bin/env bash

go build main.go

mv main addressinfer

mv addressinfer docker/

cp -r templates docker/

cd docker/

docker build --network host -t luozhouyang/address-infer:0.1 .

rm addressinfer

rm -rf templates

docker run -ti --name address-infer -p 8080:8080 luozhouyang/address-infer:0.1

