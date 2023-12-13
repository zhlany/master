package main

import (
	"fmt"
	"os"
	"zhl/src/gin/app/controller"
	"zhl/src/gin/databases"
)

func main() {
	currentDir, _ := os.Getwd()
	fmt.Println("程序入口...", currentDir)
	configs, err := databases.Load("src/gin/resources/conf/config.yaml")
	if err != nil {
		fmt.Println("初始化失败！！", err)
		return
	}
	databases.InitDB(configs)
	controller.RandStudent(5)
}
