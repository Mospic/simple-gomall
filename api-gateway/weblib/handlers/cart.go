package handlers

import (
	cart "api-gateway/services/cart"
	"context"
	"github.com/gin-gonic/gin"
)

func AddItem(ginCtx *gin.Context) {
	var cartReq cart.AddItemReq
	if err := ginCtx.ShouldBindJSON(&cartReq); err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	cartService := ginCtx.Keys["cartService"].(cart.CartService)
	cartResp, err := cartService.AddItem(context.Background(), &cartReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ginCtx.JSON(200, cartResp)
}

func GetCart(ginCtx *gin.Context) {
	var cartReq cart.GetCartReq
	if err := ginCtx.ShouldBindJSON(&cartReq); err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	cartService := ginCtx.Keys["cartService"].(cart.CartService)
	cartResp, err := cartService.GetCart(context.Background(), &cartReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"error": err.Error()})
	}
	ginCtx.JSON(200, cartResp)
}

func RemoveItem(ginCtx *gin.Context) {
	var cartReq cart.RemoveItemReq
	if err := ginCtx.ShouldBindJSON(&cartReq); err != nil {
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	cartService := ginCtx.Keys["cartService"].(cart.CartService)
	cartResp, err := cartService.RemoveItem(context.Background(), &cartReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ginCtx.JSON(200, cartResp)
}
