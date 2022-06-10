package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//检测服务器端在注册rpc对象是否合法-----服务器端封装

// MyInterface 创建接口，在接口方法定义方法原型（映射原定义方法，server.Add）
type MyInterface interface {
	// Add 对应服务器端方法
	Add(string, *string) error
}

// RegisterService 调用方法时，需要给i 传参，参数应该是应实现了
func RegisterService(i MyInterface) {
	rpc.RegisterName("student", i)
}

//-----------------client封装
//像调用本地一样调用函数

type MyClient struct {
	c *rpc.Client
}

// InitClient 使用c来调用Call,需要初始化
func InitClient(addr string) MyClient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return MyClient{conn}
}

func (this *MyClient) Add(data string, reply *string) error {
	return this.c.Call("student.Add", data, &reply)
}
