FROM ubuntu:16.04

LABEL maintainer="stupidme.me.lzy@gmail.com"

RUN apt-get update && apt-get install --no-install-recommends -y \
    python3.5 python3-pip \
    && pip3 install setuptools \
    && pip3 install tensorflow==1.4.1 \
    && pip3 install jieba3k \
    && apt-get remove python-pip3 \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/ \
    && mkdir -p /home/templates

COPY addressinfer /home
COPY templates /home/templates

EXPOSE 8080

CMD ["/home/addressinfer"]