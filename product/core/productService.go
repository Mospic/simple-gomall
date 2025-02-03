package core

import (
	"context"
	"fmt"
	"product/model"
	services "product/services"
	"strconv"
)

type ProductService struct {
}

/*
* 获取商品信息
 */
func (*ProductService) GetProduct(ctx context.Context, req *services.GetProductReq, resp *services.GetProductResp) error {
	Id := req.Id
	product, err := model.NewProductDao().FindProductByID(Id)
	if err != nil {
		fmt.Println("获取商品信息失败")
		return err
	}
	// 获取商品分类
	categoryNameList, err := model.NewProductCategoryDao().FindCategoryNameByProductID(Id)
	resp.Product = &services.Product{
		Id:          product.Id,
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
func (*ProductService) ListProducts(ctx context.Context, req *services.ListProductsReq, resp *services.ListProductsResp) error {
	// 获取商品列表
	productList, err := model.NewProductDao().ListProducts(req.Page, req.PageSize, req.CategoryName)
	if err != nil {
		fmt.Println("获取商品列表失败")
		return err
	}
	// 获取商品分类
	for _, product := range productList {
		categoryNameList, err := model.NewProductCategoryDao().FindCategoryNameByProductID(product.Id)
		if err != nil {
			fmt.Println("获取商品分类失败")
			return err
		}
		resp.Products = append(resp.Products, &services.Product{
			Id:          product.Id,
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
func (*ProductService) SearchProducts(ctx context.Context, req *services.SearchProductsReq, resp *services.SearchProductsResp) error {
	// 获取商品列表
	productList, err := model.NewProductDao().SearchProducts(req.Query)
	if err != nil {
		fmt.Println("查询商品列表失败")
		return err
	}
	// 获取商品分类
	for _, product := range productList {
		categoryNameList, err := model.NewProductCategoryDao().FindCategoryNameByProductID(product.Id)
		if err != nil {
			fmt.Println("获取商品分类失败")
			return err
		}
		resp.Results = append(resp.Results, &services.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Categories:  categoryNameList,
		})
	}
	return nil
}

func (*ProductService) CreateProducts(ctx context.Context, req *services.CreateProductsReq, resp *services.CreateProductsResp) error {
	res := make([]uint32, len(req.Products))
	for i, v := range req.Products {
		tx := model.DB.Begin()
		p := model.Product{
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
			Stock:       int32(req.Stock[i]),
		}

		id, err := model.NewProductDao().CreateProduct(&p)
		if err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		tx.Commit()
		fmt.Println("create product id:" + strconv.Itoa(int(id)))
		res[i] = uint32(id)
	}
	resp.ProductId = res
	return nil
}

func (*ProductService) DeleteProducts(ctx context.Context, req *services.DeleteProductsReq, resp *services.DeleteProductsResp) error {
	res := make([]uint32, len(req.ProductId))
	for i, v := range req.ProductId {
		tx := model.DB.Begin()
		err := model.NewProductDao().DeleteProduct(v)

		if err != nil {
			tx.Rollback()
			res[i] = 0
			return err
		}

		tx.Commit()
		fmt.Println("delete product id:" + strconv.Itoa(v))
		res[i] = v
	}
	resp.ProductId = res
	return nil
}

func (*ProductService) UpdateProducts(ctx context.Context, req *services.UpdateProductsReq, resp *services.UpdateProductsResp) error {
	return nil
}
