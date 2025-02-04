package conf

import (
	"cart/model"
)

func Init() {
	// 连接数据库
	model.Database("root:123456@tcp(115.159.2.14:3307)/simple_mall?charset=utf8&parseTime=True&loc=Local")
}
