package main

import (
	"context"
	"encoding/json"
	"fmt"
	"zhl/src/grpc_demo/pb"

	"google.golang.org/grpc"
)

func main() {
	//连接grpc服务
	grpcConn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err :", err)
		return
	}
	defer grpcConn.Close()
	//初始化
	grpcClient := pb.NewSayNameClient(grpcConn)
	//初始化对象
	var teacher pb.Teacher
	teacher.Id = 2
	teacher.Name = "zhl ddd"
	//调用远程服务
	t, err := grpcClient.SayHello(context.TODO(), &teacher)
	fmt.Println(t, err)
	tt, _ := json.Marshal(teacher)
	fmt.Println(string(tt))
}
