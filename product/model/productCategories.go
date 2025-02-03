package model

import (
	"sync"
	"time"
)

type ProductCategory struct {
	Id        int32  `gorm:"primary_key;auto_increment"`
	ProductId uint32 `gorm:"not null"`
	//CategoryId   int32     `gorm:"not null"`
	CategoryName string    `gorm:"default:(-):not null"`
	CreateAt     time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
	DeleteAt     time.Time `gorm:"default:NULL"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}

type ProductCategoryDao struct {
}

var productCategoryDao *ProductCategoryDao
var productCategoryOnce sync.Once

func NewProductCategoryDao() *ProductCategoryDao {
	productCategoryOnce.Do(
		func() {
			productCategoryDao = &ProductCategoryDao{}
		})
	return productCategoryDao
}

/* 根据 productID 查找 category Name */
func (d *ProductCategoryDao) FindCategoryNameByProductID(id uint32) ([]string, error) {
	var productCategoryList []ProductCategory
	var categoryNameList []string
	result := DB.Where("product_id = ?", id).Find(&productCategoryList)
	err := result.Error
	if err != nil {
		return nil, err
	}
	for _, productCategory := range productCategoryList {
		categoryNameList = append(categoryNameList, productCategory.CategoryName)
	}
	return categoryNameList, nil
}

/* 根据 category Name 模糊查找 productID */
func (d *ProductCategoryDao) FindProductIDByCategoryName(categoryName string) ([]uint32, error) {
	var productCategoryList []ProductCategory
	var productIdList []uint32
	result := DB.Where("category_name like ?", "%"+categoryName+"%").Find(&productCategoryList)
	err := result.Error
	if err != nil {
		return nil, err
	}
	for _, productCategory := range productCategoryList {
		productIdList = append(productIdList, productCategory.ProductId)
	}
	return productIdList, nil
}
