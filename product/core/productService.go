package core

import (
	"context"
	"fmt"
	"product/model"
	services "product/services"
)

type productService struct {
}

/*
* 获取商品信息
 */
func (*productService) GetProduct(ctx context.Context, req *services.GetProductReq, resp *services.GetProductResp) error {
	productId := req.Id
	product, err := model.NewProductDao().FindProductByID(productId)
	if err != nil {
		fmt.Println("获取商品信息失败")
		return err
	}
	// 获取商品分类
	categoryNameList, err := model.NewProductCategoryDao().FindCategoryNameByProductID(productId)
	resp.Product = &services.Product{
		Id:          product.ProductId,
		Name:        product.Name,
		Description: product.Description,
		Picture:     product.Picture,
		Price:       product.Price,
		Categories:  categoryNameList,
	}
	return nil
}

/*
* 分页获取商品列表
 */
func (*productService) ListProducts(ctx context.Context, req *services.ListProductsReq, resp *services.ListProductsResp) error {
	// 获取商品列表
	productList, err := model.NewProductDao().ListProducts(req.Page, req.PageSize, req.CategoryName)
	if err != nil {
		fmt.Println("获取商品列表失败")
		return err
	}
	// 获取商品分类
	for _, product := range productList {
		categoryNameList, err := model.NewProductCategoryDao().FindCategoryNameByProductID(product.ProductId)
		if err != nil {
			fmt.Println("获取商品分类失败")
			return err
		}
		resp.Products = append(resp.Products, &services.Product{
			Id:          product.ProductId,
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Categories:  categoryNameList,
		})
	}
	return nil
}

/*
* 查询获取商品列表
 */
func (*productService) SearchProducts(ctx context.Context, req *services.SearchProductsReq, resp *services.SearchProductsResp) error {
	// 获取商品列表
	productList, err := model.NewProductDao().SearchProducts(req.Query)
	if err != nil {
		fmt.Println("查询商品列表失败")
		return err
	}
	// 获取商品分类
	for _, product := range productList {
		categoryNameList, err := model.NewProductCategoryDao().FindCategoryNameByProductID(product.ProductId)
		if err != nil {
			fmt.Println("获取商品分类失败")
			return err
		}
		resp.Results = append(resp.Results, &services.Product{
			Id:          product.ProductId,
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Categories:  categoryNameList,
		})
	}
	return nil
}
