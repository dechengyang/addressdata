#!/usr/bin/env bash

go build main.go

mv main addressinfer

mv addressinfer docker/

cp -r statics docker/

cd docker/

docker build --network host -t allenluo/address-infer:0.1 .

rm addressinfer

rm -rf statics

docker run -ti --name address-infer -p 8080:8080 allenluo/address-infer:0.1

