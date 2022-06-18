package main

import (
	context "context"
	"fmt"
	"net"
	"zhl/src/grpc_demo/pb"

	"google.golang.org/grpc"
)

type Children struct {
}

func (this *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	return t, nil
}
func main() {
	//初始化grpcd对象
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))
	//监听
	fmt.Println("监听中...")
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	//启动服务
	grpcServer.Serve(listener)
}
