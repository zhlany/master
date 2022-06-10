package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main01() {
	//1.注册路由
	//(go特有的数据序列化--gob)
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	//(json的数据格式)
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("rpc.Dial err:", err)
		return
	}
	defer conn.Close()
	//2.调用远程函数
	var reply string
	err = conn.Call("student.Add", "远程调用", &reply)
	if err != nil {
		fmt.Println("conn.Call err;", err)
		return
	}
	fmt.Println("=====", reply)
}

func main() {
	myClient := InitClient("127.0.0.1:8800")
	var resp string
	err := myClient.Add("测试...", &resp)
	if err != nil {
		fmt.Println("myClient.Add err:", err)
		return
	}
	fmt.Println(resp, err)
}
