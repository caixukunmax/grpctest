package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpctest/proto_out"
	"net"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *proto_out.HelloRequest) (*proto_out.HelloReply, error) {
	return &proto_out.HelloReply{Message: "hello " + in.Name}, nil
}

const (
	address_port     = ":8080"
)

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", address_port)
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	println("server open sccess")
	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	proto_out.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}
