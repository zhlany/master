package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"zhl/src/consul/pb"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type Children struct {
	pb.UnimplementedSayNameServer
}

func (this *Children) SayHello(ctx context.Context, t *pb.Person) (*pb.Person, error) {
	return t, nil
}

func registerWithConsul(port int) {
	//grpc服务注册到consuls上
	//1.初始化consul配置
	config := api.DefaultConfig()
	//2.创建consul对象
	client, _ := api.NewClient(config)

	// 服务注册配置
	registration := &api.AgentServiceRegistration{
		ID:      "greeter-server-1",
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc and consul",
		Port:    port,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			//作用：检查服务是否健康
			CheckID: "grpc consul checkID",
			GRPC:    "127.0.0.1:" + string(port), //这里的端口一定要和监听的端口一致
			//TCP:                            "127.0.0.1:8800",
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		log.Fatalf("Consul注册失败: %v", err)
	}
	log.Println("服务成功注册到Consul")
}

func main() {
	lis, err := net.Listen("tcp", ":0") // 自动获取可用端口
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	port := lis.Addr().(*net.TCPAddr).Port

	//4.注册到服务上
	go registerWithConsul(port) // 异步注册服务

	///////////////////////////////远程调用///////////////
	//初始化grpc对象
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))
	//监听
	fmt.Println("监听中...")
	listener, err := net.Listen("tcp", lis.Addr().String())
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	//启动服务
	grpcServer.Serve(listener)
}
