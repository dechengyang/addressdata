FROM ubuntu:16.04

LABEL maintainer="ZhouYang Luo<stupidme.me.lzy@gmail.com>"

RUN mkdir /root/files

COPY main /root

WORKDIR /root

EXPOSE 8088

CMD ["./main"]