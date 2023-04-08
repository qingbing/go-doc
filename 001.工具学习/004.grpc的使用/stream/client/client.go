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
