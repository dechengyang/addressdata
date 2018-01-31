FROM ubuntu:16.04

LABEL maintainer="stupidme.me.lzy@gmail.com"

RUN apt-get update && apt-get install --no-install-recommends -y \
    python3.5 python3-pip \
    && pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple setuptools \
    && pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple tensorflow==1.4.1 \
    && pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple jieba3k \
    && rm -rf ~/.cache/pip \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/ \
    && mkdir -p /home/templates

COPY addressinfer /home
COPY templates /home/templates

EXPOSE 8080

WORKDIR /home/

CMD ["./addressinfer"]