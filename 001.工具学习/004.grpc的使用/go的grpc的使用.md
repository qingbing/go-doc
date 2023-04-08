# go的grpc

- [go的grpc](#go的grpc)
  - [1. 参考连接](#1-参考连接)
  - [2. 说明](#2-说明)
    - [2.1 简介描述](#21-简介描述)
    - [2.2 grpc 用于微服务时需要了解的问题](#22-grpc-用于微服务时需要了解的问题)
    - [2.3 了解 protobuf](#23-了解-protobuf)
      - [2.3.1 protobuf 优势](#231-protobuf-优势)
  - [环境准备](#环境准备)
    - [3.1 安装 proto 的编译器](#31-安装-proto-的编译器)
    - [3.2 安装 go 语言的插件](#32-安装-go-语言的插件)
  - [4. grpc 演示](#4-grpc-演示)
    - [4.1 编写 .proto 文件](#41-编写-proto-文件)
    - [4.2 根据 .proto 文件生成相应的 go 代码](#42-根据-proto-文件生成相应的-go-代码)
    - [4.3 对 go 项目安装依赖](#43-对-go-项目安装依赖)
    - [4.4 go 代码](#44-go-代码)
    - [4.4 测试代码](#44-测试代码)

## 1. 参考连接

- grpc 的 go-demo
  - https://github.com/grpc/grpc-go
- grpc官网
  - https://grpc.io/docs/
- protobuf语言挂官网
  - https://protobuf.dev/programming-guides/proto3/


## 2. 说明

### 2.1 简介描述

1. rpc与http是两个完全不同的概念, 没有可比性
2. grpc proto协议文件的基本语法
3. grpc 四种消息传输模式的实践
   - 远程调用（一对一）
   - 流式传输: 
     - 客户端流: client 以流的方式访问 Server 端， Server 端以一元的响应回应 client. 比如： 文件上传
     - 服务端流: client 以一元的请求访问 Server 端， Server 端以流的方式进行回应。 比如： 文件下载
     - 客户端流-服务端流: client 可持续不断的向 server 发送消息，反之亦可

### 2.2 grpc 用于微服务时需要了解的问题

1. 用户鉴权问题
2. grpc 数据传递，类似 http header
3. 拦截器
4. 客户端负载均衡(如果服务已经部署为负载均衡，那么无需客户端负载均衡)
5. 服务的健康检查
6. 数据传输的方式(一元请求或流式请求)
7. 服务之间的认证问题
8. 服务限流的问题，服务接口限流
9. 服务的熔断，通过判断发生错误的次数，对服务做降级
10. 日志追踪

### 2.3 了解 protobuf

1. proto 需要一个编译器
2. 默认不支持go，需要安装编译器插件
3. 生成代码在项目中

#### 2.3.1 protobuf 优势

1. 传输更快，以二进制的方式进行传输
2. 是一种格式，与语言无关，类似json格式，只是其传输是二进制，字段时定义的数字标志，相对体积更小，传输更快
3. 跨语言的，支持多语言
   - C# / .NET
   - C++
   - Dart
   - Go
   - Java
   - Kotlin
   - Node
   - Objective-C
   - PHP
   - Python
   - Ruby

##  环境准备

### 3.1 安装 proto 的编译器

- ubuntu

```bash
# 安装 protoc 可根据 protoc 提示进行安装
# 使用 snap 安装
snap install protobuf --classic
# 或者
sudo apt install -y protoc
```

### 3.2 安装 go 语言的插件

```bash
# 安装 grpc 需要的插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# 确保安装的文件目录在 PATH 变量中
export PATH="$PATH:$(go env GOPATH)/bin"
```

## 4. grpc 演示

### 4.1 编写 .proto 文件

```proto
// proto 文件的解析语法, 不指定默认为 proto2, proto2 的协议比 proto3 的协议更复杂
syntax = "proto3"; 
// 生成 grpc 的代码的引用路径, 引用路径从 mod 开始计算
option go_package = "demo/grpc/stream/proto"; 
// grpc 代码中 servername 的一部分(前缀)， 相当于命名空间
package stream;

// 导入
import "google/protobuf/timestamp.proto"; // 时间类型
import "google/protobuf/any.proto"; // 任意类型

// 定义一个服务
service Greeter {
    // 一元调用
    rpc SayHello (HelloRequest) returns (HelloReplay) {}
    //  客户端流, eg: 上传文件
    rpc SayHelloClientStream(stream HelloRequest) returns (HelloReplay) {}
    // 服务端流, eg: 下载文件
    rpc SayHelloServerStream(HelloRequest) returns (stream HelloReplay) {}
    // 双向流, eg: 机器人克服场景， 一问一答
    rpc SayHelloTwoWayStream(stream HelloRequest) returns (stream HelloReplay) {}
}

enum Gender {
    // 男
    MALE = 0; // 在 proto3 语法里， 第一个值必须时 0
    // 女
    FEMALE = 1;
    // 未知
    UNKNOWN = 3;
}

message Address {
    string provice = 1;
    string city = 2;
}

message HelloRequest {
    // [1-15] 占用一个字节， 数字越大，占用越多， 在使用时尽量不超过15个字段
    // 如果大于15个字段，尽量将使用频繁的字段写在前面， 也有几率构建小的占用
    string name = 1; // 字段的标志
    Gender gender = 2;
    uint32 age = 3;
    google.protobuf.Timestamp birthday = 4; // 时间
    Address addr = 5;
    repeated string hoppy = 7; // repeated 会被转化为数组
    map<string,google.protobuf.Any> data = 8;
    // string sex = 200; // 不能使用， 200 被 reserved
    // string phone = 3; // 不能使用， phone 被 reserved
    reserved 100, 103, 150 to 200; // 保留标记: 不能在定义时使用
    reserved "phone", "email"; // 保留字段: 不能在定义时使用
}

message HelloReplay {
    string message = 1;
}
```

### 4.2 根据 .proto 文件生成相应的 go 代码

- --go_out=. : go文件生成目录， `.` 表示执行命令的当前目录, 会在目录中生成 stream.pb.go 的文件
- --go_opt=paths=source_relative: 表示采用和源路径相同的结构
- --grpc_out=.: 从当前目录的相对位置, 会在目录中生成 stream_grpc.pb.go 的文件
- --go_opt=paths=source_relative: 表示采用和源路径相同的结构

```bash
# 文件时相对目录
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/stream/proto /stream.proto
```

### 4.3 对 go 项目安装依赖

```bash
# 文件生成后，需要使用 "go mod tidy" 对 go 项目进行对齐
go mod tidy
```

### 4.4 go 代码

- grpc 服务端

```go
package main

import (
	"context"
	"demo/grpc/stream/proto"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "")
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReplay, error) {
	log.Printf("Server recv: %+v", in)
	return &proto.HelloReplay{
		Message: "this is a stream server response message",
	}, nil
}

// 客户端流
func (s *server) SayHelloClientStream(stream proto.Greeter_SayHelloClientStreamServer) error {
	list := make([]*proto.HelloRequest, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloReplay{
				Message: fmt.Sprintf("hello client total recv count: %d", len(list)),
			})
		}
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, req)
		fmt.Printf("server recv: %+v\n", req)
	}
	return nil
}

// 服务端流
func (s *server) SayHelloServerStream(in *proto.HelloRequest, stream proto.Greeter_SayHelloServerStreamServer) error {
	fmt.Printf("server recv: %+v\n", in)

	list := []proto.HelloReplay{
		{Message: "1th message\n"},
		{Message: "2th message\n"},
		{Message: "3th message\n"},
		{Message: "4th message\n"},
		{Message: "5th message\n"},
	}

	for _, reply := range list {
		r := reply
		err := stream.Send(&r)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

// 双向流
func (s *server) SayHelloTwoWayStream(stream proto.Greeter_SayHelloTwoWayStreamServer) error {
	i := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		i++
		fmt.Printf("Server recv: %+v\n", req)
		stream.Send(&proto.HelloReplay{
			Message: fmt.Sprintf("%dth request ok", i),
		})
	}
	return nil
}

func main() {
	// 解析命令行参数
	flag.Parse()

	// 启动监听服务
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("%#v", err)
	}
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGreeterServer(s, &server{})
	log.Printf("Server listen at %v\n", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("%#v", err)
	}
}

```

- grpc 客户端

```go
package main

import (
	"context"
	"demo/grpc/stream/proto"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	addr = flag.String("addr", "localhost:50051", "")
)

func main() {
	flag.Parse()
	// 实际应用中， 连接可以采用连接池实现: 创建和销毁连接非常耗时
	// 拨号
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 一元调用
	// sayHello(c)
	// 客户端流
	// sayHelloClientStream(c)
	// 服务端流
	// sayHelloServerStream(c)
	// 双向流
	sayHelloTwoWayStream(c)
}

// 请求结构体
func getRequest() *proto.HelloRequest {
	birthday := timestamppb.New(time.Now())
	any1, _ := anypb.New(&proto.Address{
		Provice: "四川",
		City:    "成都",
	})
	in := &proto.HelloRequest{
		Name:     "stream",
		Gender:   proto.Gender_FEMALE,
		Age:      10,
		Birthday: birthday,
		Addr: &proto.Address{
			Provice: "四川",
			City:    "成都",
		},
		Hoppy: []string{"篮球", "足球"},
		Data: map[string]*anypb.Any{
			"A": any1,
		},
	}
	return in
}

// 一元调用
func sayHello(c proto.GreeterClient) {
	in := getRequest()
	res, err := c.SayHello(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("一元调用, client recv msg: %s\n", res.Message)
}

// 客户端流
func sayHelloClientStream(c proto.GreeterClient) {
	list := []*proto.HelloRequest{
		getRequest(), getRequest(), getRequest(),
	}
	stream, err := c.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, in := range list {
		err := stream.Send(in)
		if err != nil {
			log.Fatal(err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端流, client recv: %v\n", reply.Message)
}

// 服务端流
func sayHelloServerStream(c proto.GreeterClient) {
	in := getRequest()
	stream, err := c.SayHelloServerStream(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			stream.CloseSend()
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("服务端流, client recv： %s", reply.Message)
	}
}

// 双向流
func sayHelloTwoWayStream(c proto.GreeterClient) {
	list := []*proto.HelloRequest{
		getRequest(), getRequest(), getRequest(),
	}
	stream, err := c.SayHelloTwoWayStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, req := range list {
		err := stream.Send(req)
		if err != nil {
			log.Fatal(err)
		}
	}

	done := make(chan struct{})

	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				close(done)
				return
			}
			fmt.Printf("双向流, client recv: %s\n", reply.Message)
		}
	}()

	stream.CloseSend()
	<-done
}
```

### 4.4 测试代码

- 服务端

```bash
bing@QB:/code/go/src/gin-demo$ go run ./grpc/stream/server/server.go 
2023/04/08 17:14:39 Server listen at [::]:50051
Server recv: name:"stream" gender:FEMALE age:10 birthday:{seconds:1680945282 nanos:443279484} addr:{provice:"四川" city:"成都"} hoppy:"篮球" hoppy:"足球" data:{key:"A" value:{[type.googleapis.com/stream.Address]:{provice:"四川" city:"成都"}}}
Server recv: name:"stream" gender:FEMALE age:10 birthday:{seconds:1680945282 nanos:443352446} addr:{provice:"四川" city:"成都"} hoppy:"篮球" hoppy:"足球" data:{key:"A" value:{[type.googleapis.com/stream.Address]:{provice:"四川" city:"成都"}}}
Server recv: name:"stream" gender:FEMALE age:10 birthday:{seconds:1680945282 nanos:443352946} addr:{provice:"四川" city:"成都"} hoppy:"篮球" hoppy:"足球" data:{key:"A" value:{[type.googleapis.com/stream.Address]:{provice:"四川" city:"成都"}}}
```

- 客户端

```bash
bing@QB:/code/go/src/gin-demo$ go run ./grpc/stream/client/client.go 
双向流, client recv: 1th request ok
双向流, client recv: 2th request ok
双向流, client recv: 3th request ok
```
