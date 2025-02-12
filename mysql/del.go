package main

import (
	"log"
	"mall/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:yzl222888@tcp(127.0.0.1:3306)/simple_mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 删除所有表
	err = db.Migrator().DropTable(
		&model.User{},
		&model.Categories{},
		&model.Product{},
		&model.Cart{},
		&model.Order{},
		&model.Address{},
		&model.Message{},
	)
	if err != nil {
		log.Fatal("删除表失败:", err)
	}

	log.Println("所有表删除成功！")
}
