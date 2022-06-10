package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

type User struct {
	id   string
	name string
	msg  chan string
}

var lock sync.RWMutex

//全局map保存用户
var allUsers = make(map[string]User)

//去哪聚管道，用于接收任何人发来的信息
var massage = make(chan string, 10)

func main() {
	/*	//创建服务器
		//监听
		//建立连接
		//启动处理业务逻辑go协程
	*/

	//1.创建
	listener, err := net.Listen("tcp", ":8384")
	if err != nil {
		fmt.Println("net.listener err: ", err)
		return
	}

	//监听广播管道信息
	go broadcast()

	fmt.Println("服务器启动成功!")

	for {
		//2.监听
		fmt.Println("连接监听中...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			return
		}

		//3.建立连接
		fmt.Println("建立连接成功！")

		//4.处理业务的go协程
		go hander(conn)
	}
}

//具体业务
func hander(conn net.Conn) {
	fmt.Println("启动业务...")

	clientAddress := conn.RemoteAddr().String()
	i := strings.Index(clientAddress, ":")
	rs := []rune(clientAddress)
	address := string(rs[i+1:])

	newUser := User{
		id:   address,
		name: randc(address),
		msg:  make(chan string, 10),
	}
	lock.Lock()
	//添加到map
	allUsers[newUser.id] = newUser
	lock.Unlock()

	var isQuit, restTime = make(chan bool), make(chan bool)

	//用于重启计时器的管道

	go watch(&newUser, conn, isQuit, restTime)

	//向massage写入数据，广播
	massage <- fmt.Sprintf("   massage接收: [%s]:[%s] 上线了！\n", newUser.id, newUser.name)

	//msg数据返回客户端
	go writeToClient(&newUser, conn)

	for {
		buf := make([]byte, 1024)
		//读取客户端发送过来的请求数据
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println(newUser.name, " ctrl+c准备退出！")
			isQuit <- true
		}

		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}

		fmt.Printf("接收%s的信息:%s,长度:%v\n", newUser.id, string(buf[:cnt-1]), cnt)

		scanner := string(buf[:cnt-1])
		//业务逻辑处理
		//1.查询用户
		//接收数据是否为who
		if cnt-1 == 3 && scanner == "who" {
			fmt.Println(">>> UserList")

			//存储用户信息
			var userInfos []string
			lock.RLock()
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("id: %s | name: %s", user.id, user.name)
				userInfos = append(userInfos, userInfo)
			}
			lock.RUnlock()
			//切片拼接成字符串
			str := strings.Join(userInfos, "\n")
			newUser.msg <- str
		} else if strings.Index(scanner, "rename") == 0 {
			indexT := strings.Index(scanner, " ")
			newName := scanner[indexT+1:]
			newUser.name = newName
			lock.Lock()
			allUsers[newUser.id] = newUser
			lock.Unlock()
			newUser.msg <- "success！"
		} else if strings.Index(scanner, "to") == 0 {
			indexF := strings.Index(scanner, " ")
			indexT := strings.LastIndex(scanner, " ")
			userName := scanner[indexF+1 : indexT]
			msg := scanner[indexT+1:]
			go toUser(&newUser, userName, msg)
		} else {
			massage <- scanner
		}

		restTime <- true
	}
}

//发送给指定用户
func toUser(fromUser *User, name string, msg string) {
	user := User{}
	lock.RLock()
	var flag bool
	for _, u := range allUsers {
		if u.name == name {
			user = u
			flag = true
		}
	}
	lock.RUnlock()
	if !flag {
		fmt.Println("发送给用户不存在！")
		fromUser.msg <- fmt.Sprintf("发送给的用户%s不存在", name)
		return
	}
	str := "[" + fromUser.name + "]发来信息: " + msg
	user.msg <- str
	fromUser.msg <- fmt.Sprintf("信息发送成功！")
}

//随机命名
func randc(address string) string {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	var nameStr []rune
	for i := 0; i < len(address); i++ {
		nameStr = append(nameStr, rune(r.Intn(26)+'a'-1))
	}
	//fmt.Println("-------------", string(nameStr), "====", nameStr)
	return string(nameStr)
}

//广播消息
func broadcast() {
	fmt.Println("广播...")
	defer fmt.Println("广播退出")
	for {
		fmt.Println("监听message中...")
		info := <-massage
		fmt.Println("广播给所有用户:\n", info)
		lock.Lock()
		for _, user := range allUsers {
			user.msg <- info
		}
		lock.Unlock()
	}
}

//监听自己的msg管道，负责将数据往回写(客户端)
func writeToClient(user *User, conn net.Conn) {
	fmt.Printf("监听user%s自己的msg --------\n", user.name)
	for date := range user.msg {
		fmt.Printf("写回客户id:%s[ name:%s ]数据:\n{\n  %s\n  }\n", user.id, user.name, date)
		//写回客户端
		_, _ = conn.Write([]byte(date))
	}
}

//退出
func watch(user *User, conn net.Conn, isQuit, resetTime <-chan bool) {
	defer fmt.Println(user.name, "退出成功！")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s exit already!", user.name)
			lock.Lock()
			delete(allUsers, user.id)
			lock.Unlock()
			massage <- logoutInfo
			conn.Close()
			return
		case <-time.After(600 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout exit already!", user.name)
			lock.Lock()
			delete(allUsers, user.id)
			lock.Unlock()
			massage <- logoutInfo
			conn.Close()
			return
		case <-resetTime:
			fmt.Printf("%s重置计时器！\n", user.name)
		}
	}
}
