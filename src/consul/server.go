package main

import (
	"context"
	"fmt"
	"net"
	"zhl/src/consul/pb"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type Children struct {
}

func (this *Children) SayHello(ctx context.Context, t *pb.Person) (*pb.Person, error) {
	return t, nil
}
func main() {
	//grpc服务注册到consuls上
	//1.初始化consul配置
	config := api.DefaultConfig()
	//2.创建consul对象
	consulClient, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}
	//3.配置
	registerService := api.AgentServiceRegistration{
		ID:      "go_micro",
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc and consul",
		Address: "127.0.0.1",
		Port:    8800,
		Check: &api.AgentServiceCheck{
			CheckID:  "grpc consul checkID",
			TCP:      "127.0.0.1:8800",
			Interval: "5s",
			Timeout:  "2s",
		},
	}
	//4.注册到服务上
	consulClient.Agent().ServiceRegister(&registerService)

	///////////////////////////////远程调用///////////////
	//初始化grpc对象
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))
	//监听
	fmt.Println("监听中...")
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	//启动服务
	grpcServer.Serve(listener)
}
