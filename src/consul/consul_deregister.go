package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

//注销consul
func main() {
	//1.初始化consul配置
	config := api.DefaultConfig()
	//2.创建consul对象
	consulClient, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}
	//3.注销服务
	consulClient.Agent().ServiceDeregister("go_micro")
}
