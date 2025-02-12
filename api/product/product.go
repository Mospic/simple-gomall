package Product

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	mlog "mall/log"
	"mall/middleware/auth"
	"mall/service/product/proto/product"
)

var ProductClient product.ProductCatalogServiceClient
var log *mlog.Log

func Init(engine *gin.Engine) {
	ProductConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "product.rpc",
		},
	})
	ProductClient = product.NewProductCatalogServiceClient(ProductConn.Conn())
	log = mlog.NewLog("ProductAPI")
	group := engine.Group("/Product", auth.ParseToken)
	{
		group.GET("/Get", Get)
		group.GET("/List", List)
		group.GET("/Search", Search)
		group.POST("/Create", Create)
		group.PUT("/Update", Update)
		group.DELETE("/Delete", Delete)
	}
}
