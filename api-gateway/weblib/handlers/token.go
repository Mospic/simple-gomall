package handlers

import (
	token "api-gateway/services/tokenutils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
func VarifyToken(ginCtx *gin.Context) {
	var tokenVariReq token.VerifyTokenRequest
	//获取id和token
	if err := ginCtx.ShouldBindJSON(&tokenVariReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
	tokenResp, err := tokenService.VarifyToken(ginCtx.Request.Context(), &tokenVariReq)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"info": "验证失败！",
		})
	}
	if tokenResp != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"token": tokenResp.Token,
			"info":  tokenResp.Status,
		})
	}
}
