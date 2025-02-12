// 初始化数据库表结构
package main

import (
	"mall/model"
)
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

func main() {
	dsn := "root:yzl222888@tcp(127.0.0.1:3306)/simple_mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}
	err = db.SetupJoinTable(&model.Cart{}, "Products", &model.CartProducts{})
	if err != nil {
		panic(err)
	}
	err = db.SetupJoinTable(&model.Order{}, "Products", &model.OrderProducts{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Categories{}, &model.Product{}, &model.Cart{}, &model.Order{}, &model.Address{}, model.Message{})
	if err != nil {
		panic(err)
	}
}
