#!/usr/bin/env bash

scp data/*.json allen@116.85.35.192:/home/allen/files

scp data/*.csv allen@116.85.35.192:/home/allen/files

go build main.go

docker build --network host -t lzy/addressdata:0.1 .

docker run -d \
    --name addressdata \
    -p 8088:8088 \
    -v /home/allen/files:/root/files
    allenluo/addressdata:0.1

