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
