package model

import (
	"gorm.io/gorm"
	"sync"
)

type Product struct {
	gorm.Model
	Id          uint32  `gorm:"primary_key;auto_increment"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"default:(-)"`
	Picture     string  `gorm:"default:(-)"`
	Price       float32 `gorm:"default:0"`
	Stock       int32   `gorm:"default:0"`
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

/*
*
根据商品ID查找商品实体
*/
func (d *ProductDao) FindProductByID(id uint32) (*Product, error) {
	product := Product{Id: id}

	result := DB.Where("id = ?", id).First(&product)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

/*
* 根据商品ID列表查询商品实体
 */
func (d *ProductDao) FindProductByIDs(ids []uint32) ([]*Product, error) {
	var productList []*Product
	result := DB.Where("id in (?)", ids).Find(&productList)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return productList, nil
}

/*
* 分页获取商品列表
 */
func (d *ProductDao) ListProducts(page int32, pageSize int64, categoryName string) ([]*Product, error) {
	// 先根据分类名查找商品ID
	productCategoryDao := NewProductCategoryDao()
	productIds, err := productCategoryDao.FindProductIDByCategoryName(categoryName)
	if err != nil {
		return nil, err
	}
	// 再根据商品ID分页获取商品实体
	var productList []*Product
	result := DB.Where("id in (?)", productIds).Offset(int64(page-1) * pageSize).Limit(pageSize).Find(&productList)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return productList, nil
}

/*
* 查询获取商品列表
 */
func (d *ProductDao) SearchProducts(keyword string) ([]*Product, error) {
	var productList []*Product
	result := DB.Where("name like ?", "%"+keyword+"%").Find(&productList)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return productList, nil
}

func (d *ProductDao) CreateProduct(product *Product) (uint32, error) {
	result := DB.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.Id, nil
}

func (d *ProductDao) UpdateProduct(product *Product) (uint32, error) {
	result := DB.Save(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.Id, nil
}

func (d *ProductDao) DeleteProduct(id uint32) error {
	product := Product{Id: id}
	result := DB.Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
