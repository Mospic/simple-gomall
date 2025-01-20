package model

import (
	"sync"
	"time"
)

type Product struct {
	ProductId   int32     `gorm:"primary_key;auto_increment"`
	Name        string    `gorm:"default:(-):not null"`
	Description string    `gorm:"default:(-)"`
	Picture     string    `gorm:"default:(-)"`
	Price       float32   `gorm:"default:0"`
	Stock       int32     `gorm:"default:0"`
	CreateAt    time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeleteAt    time.Time `gorm:"default:NULL"`
}

func (Product) TableName() string {
	return "product"
}

type ProductDao struct {
}

var productDao *ProductDao
var productOnce sync.Once

func NewProductDao() *ProductDao {
	productOnce.Do(
		func() {
			productDao = &ProductDao{}
		})
	return productDao
}
