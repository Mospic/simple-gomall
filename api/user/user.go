package user

import (
	//"github.com/casbin/casbin/v2"
	//"github.com/casbin/mysql-adapter/v2"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	mlog "mall/log"
	"mall/middleware/auth"
	"mall/service/user/proto/user"
)

var (
	UserClient user.UserServiceClient
	log        *mlog.Log
	//enforcer   *casbin.Enforcer
)

func Init(engine *gin.Engine) {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "user.rpc",
		},
	})
	UserClient = user.NewUserServiceClient(conn.Conn())
	log = mlog.NewLog("UserAPI")
	group := engine.Group("/User")
	{
		group.POST("/Register", Register)
		group.POST("/Login", Login)
		group.POST("/Logout", auth.ParseToken, Logout)
		group.GET("/Info", auth.ParseToken, Info)
		group.GET("/Message", auth.ParseToken, Message)
		group.DELETE("/Delete", auth.ParseToken, Delete)
		group.PUT("/Update", auth.ParseToken, Update)
	}
}

//// DBConnection -> return db instance
//func DBConnection() (*gorm.DB, error) {
//	USER := "root"
//	PASS := "root"
//	HOST := "localhost"
//	PORT := "3306"
//	DBNAME := "casbin-golang"
//
//	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
//
//	return gorm.Open(mysql.Open(url))
//
//}
