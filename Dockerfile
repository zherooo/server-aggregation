FROM golang:1.24.2-alpine AS builder

# ENV 设置环境变量
ENV GOPATH=/opt/repo
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# COPY 源路径 目标路径
COPY . $GOPATH/src/server-aggregation

# RUN 执行 go build .
RUN cd $GOPATH/src/server-aggregation && go build -gcflags "all=-N -l" -o server-aggregation .

# FROM 基于 alpine:latest
FROM alpine:latest

# RUN 设置代理镜像
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories

# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=builder /opt/repo/src/server-aggregation/dlv /usr/local/bin
COPY --from=builder /opt/repo/src/server-aggregation/server-aggregation /opt

# EXPOSE 设置端口映射
EXPOSE 8080/tcp

# WORKDIR 设置工作目录
WORKDIR /opt

# CMD 设置启动命令(运行的时候指定)
#CMD ["/opt/server-aggregation", "--env", "dev"]