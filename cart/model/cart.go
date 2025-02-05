package model

import (
	"encoding/json"
	"time"
)

type Cart struct {
	UserID     int64           `gorm:"index" json:"user_id"`
	Products   json.RawMessage `gorm:"type:json" json:"products"`
	TotalPrice float64         `gorm:"type:decimal(10,2)" json:"total_price"`
	CreateAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

type CartProduct struct {
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int64   `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

func AddCartItem(userID int64, product CartProduct) error {
	now := time.Now()
	var cart Cart
	err := DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		// 购物车不存在，创建购物车
		products := []CartProduct{product}
		productsJSON, _ := json.Marshal(products)
		cart = Cart{
			UserID:     userID,
			Products:   productsJSON,
			TotalPrice: product.UnitPrice * float64(product.Quantity),
			CreateAt:   now,
			UpdateAt:   now,
		}
		return DB.Create(&cart).Error
	}
	var products []CartProduct
	json.Unmarshal(cart.Products, &products)

	found := false
	for i, p := range products {
		if p.ProductID == product.ProductID {
			products[i].Quantity += product.Quantity
			found = true
			break
		}
	}
	if !found {
		products = append(products, product)
	}

	totalPrice := 0.0
	for _, p := range products {
		totalPrice += float64(p.Quantity) * p.UnitPrice
	}

	productsJSON, _ := json.Marshal(products)
	return DB.Model(&cart).Updates(Cart{
		Products:   productsJSON,
		TotalPrice: totalPrice,
	}).Error
}

func RemoveCartItem(userID, productID int64) error {
	var cart Cart

	err := DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return err
	}

	var products []CartProduct
	json.Unmarshal(cart.Products, &products)

	newProducts := []CartProduct{}
	totalPrice := 0.0
	for _, p := range products {
		if p.ProductID != productID {
			newProducts = append(newProducts, p)
			totalPrice += float64(p.Quantity) * p.UnitPrice
		}
	}

	productsJSON, _ := json.Marshal(newProducts)
	return DB.Model(&cart).Updates(Cart{
		Products:   productsJSON,
		TotalPrice: totalPrice,
	}).Error
}

func GetCartItems(userID int64) ([]CartProduct, float64, error) {
	var cart Cart
	err := DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, 0, err
	}

	var products []CartProduct
	json.Unmarshal(cart.Products, &products)

	return products, cart.TotalPrice, nil
}
