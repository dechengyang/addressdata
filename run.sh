#!/usr/bin/env bash

go build main.go

docker build -t luozhouyang/address-infer:0.1 .

rm main

docker run -ti --name address-infer -p 8080:8080 luozhouyang/address-infer:0.1

