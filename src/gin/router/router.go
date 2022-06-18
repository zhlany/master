package router

import "github.com/gin-gonic/gin"

/*
InitRouter 路由初始化
GET /FindUsers 请求将由 controller 下的 FindUsers 方法处理
*/
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/findStudent", Student)
	return router
}

func Student(context *gin.Context) {

}
