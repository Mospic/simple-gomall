package handlers

import (
	product "api-gateway/services/product"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取商品信息
func ProductInfo(ginCtx *gin.Context) {
	var productReq product.GetProductReq
	//获取商品id
	if err := ginCtx.ShouldBindJSON(&productReq); err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// 从gin.Key中取出服务实例
	productService := ginCtx.Keys["productService"].(product.ProductCatalogService)

	//调用product微服务
	productResp, err := productService.GetProduct(context.Background(), &productReq)
	if err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid Product"})
		return
	}
	ginCtx.JSON(http.StatusOK, productResp)
}

// 分页获取商品列表
func ListProducts(ginCtx *gin.Context) {
	var listProductsReq product.ListProductsReq

	//获取商品列表请求参数
	if err := ginCtx.ShouldBindJSON(&listProductsReq); err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// 从gin.Key中取出服务实例
	productService := ginCtx.Keys["productService"].(product.ProductCatalogService)

	//调用product微服务
	listProductsResp, err := productService.ListProducts(context.Background(), &listProductsReq)
	if err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid Product"})
		return
	}
	ginCtx.JSON(http.StatusOK, listProductsResp)
}

// 查询获取商品列表
func SearchProducts(ginCtx *gin.Context) {
	var searchProductsReq product.SearchProductsReq
	//获取商品列表请求参数
	if err := ginCtx.ShouldBindJSON(&searchProductsReq); err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// 从gin.Key中取出服务实例
	productService := ginCtx.Keys["productService"].(product.ProductCatalogService)

	//调用product微服务
	searchProductsResp, err := productService.SearchProducts(context.Background(), &searchProductsReq)
	if err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid Product"})
		return
	}
	ginCtx.JSON(http.StatusOK, searchProductsResp)
}
