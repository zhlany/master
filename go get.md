# zhl-go

#### go get

```go
go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]
```

|    参数     |                   用法说明                   |
| :-------: | :--------------------------------------: |
|    -u     |  在线下载更新指定的模块（包）及依赖包（默认不更新已安装模块），并创建、安装   |
|    -v     |                打印出所下载的包名                 |
|    -d     |              只下载，而不执行创建、安装               |
|    -t     |           同时下载命令行指定包的测试代码（测试包）           |
| -insecure |      允许命令在非安全的scheme（如HTTP）下执行get命令      |
|   -fix    | 在下载代码包后先执行修正动作，而后再进行编译和安装，根据当前GO版本对所下载的模块（包）代码做语法修正 |
|    -f     |            忽略掉对已下载代码包的导入路径的检查            |
|    -x     |           打印输出，get 执行过程中的具体命令            |
|           |                                          |

​	

#### 依赖下载	

```go
//mysql
go get -u github.com/go-sql-driver/mysql

//gorm + mysql
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/mysql
"gorm.io/driver/mysql" // mysql 数据库驱动
"gorm.io/gorm"         // 使用 gorm ，操作数据库的 orm 框架
"gorm.io/gorm/logger"
//gin
go get github.com/gin-gonic/gin
//uuid
go get -u -v github.com/google/uuid
```


​	
​	
​	
​	
​	
​	

