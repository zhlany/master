package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type Student struct {
}

func (this *Student) Add(date string, resp *string) error {
	*resp = date + " success!"
	return nil
}

func main() {
	/*	//1.注册RPC服务
		err := rpc.RegisterName("student", new(Student))
		if err != nil {
			fmt.Println("rpc.RegisterName err: ", err)
			return
		}*/
	RegisterService(new(Student))

	//2.监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("监听中...")
	//连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() err :", err)
		return
	}
	defer conn.Close()
	fmt.Println("success ...")
	//绑定连接
	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
}
