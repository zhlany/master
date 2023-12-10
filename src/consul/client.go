package main

import (
	"context"
	"fmt"
	"strconv"
	"zhl/src/consul/pb"

	"github.com/hashicorp/consul/api"

	"google.golang.org/grpc"
)

func main() {
	//初始化consul配置
	//1.初始化consul配置
	config := api.DefaultConfig()
	//2.创建consul对象 (可以重新指定也可以默认)
	consulClient, err := api.NewClient(config)
	if err != nil {
		fmt.Println("client api.NewClient err:", err)
		return
	}
	//3.服务发现，获取健康服务 passingOnly:  是否通过
	services, _, err := consulClient.Health().Service("grpc and consul", "consul", true, nil)
	if err != nil {
		fmt.Println("client consulClient.Health().Service err:", err)
		return
	}
	//简单的负载均衡
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	//////////远程调用////////////
	//连接grpc服务
	grpcConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err :", err)
		return
	}
	defer grpcConn.Close()
	//初始化
	grpcClient := pb.NewSayNameClient(grpcConn)
	//初始化对象
	var teacher pb.Person
	teacher.Age = 1
	teacher.Name = "zhl name consul consul client"
	//调用远程服务
	t, err := grpcClient.SayHello(context.TODO(), &teacher)
	fmt.Println(t, err)
}
