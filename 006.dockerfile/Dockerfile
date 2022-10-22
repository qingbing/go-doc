# docker run -i -t go/test:v1 .
# 编译阶段
FROM golang:1.18.7-alpine AS build
ADD . /data/
WORKDIR /data/
RUN export GO111MODULE=on
RUN export GOSUMDB=off
RUN unset GOPATH
RUN go env -w GOPROXY=https://goproxy.cn
RUN go build -o demo main.go

# 容器化阶段
FROM scratch as demo
#FROM alpine as demo
COPY --from=build /data/demo /bin/hello
WORKDIR /bin/
EXPOSE 9090
CMD [ "/bin/hello" ]
