package handlers

import (
	token "api-gateway/services/tokenutils"
	user "api-gateway/services/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
func Register(ginCtx *gin.Context) {
	var userReq user.RegisterReq
	var tokenReq token.GenerateTokenByIDRequest
	//获取用户名和密码
	if err := ginCtx.ShouldBindJSON(&userReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(user.UserService)

	//调用user微服务，将context的上下文传入
	userResp, err := userService.Register(context.Background(), &userReq)
	LoggingIfUserError(err)
	if userResp != nil && userResp.UserId > 0 { //做一下保护，返回的UserId应该大于0
		tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
		tokenReq.UserId = userResp.UserId
		GenerateTokenByIDResponse, err := tokenService.GenerateTokenByID(context.Background(), &tokenReq)
		if GenerateTokenByIDResponse != nil {
			jwtToken := GenerateTokenByIDResponse.Token
			LoggingIfUserError(err)
			//返回
			ginCtx.JSON(http.StatusOK, user.RegisterResp{
				UserId: userResp.UserId,
				Token:  jwtToken,
			})
		}

	} else {
		ginCtx.JSON(400, gin.H{"error": "Invalid User"})
	}

}

// 用户登录
func Login(ginCtx *gin.Context) {
	var userReq user.LoginReq
	var tokenReq token.GenerateTokenByIDRequest
	if err := ginCtx.ShouldBindJSON(&userReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(user.UserService)
	userResp, err := userService.Login(context.Background(), &userReq)
	LoggingIfUserError(err)

	//生成token
	if userResp != nil && userResp.UserId > 0 {
		tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
		tokenReq.UserId = userResp.UserId
		GenerateTokenByIDResponse, err := tokenService.GenerateTokenByID(context.Background(), &tokenReq)
		LoggingIfUserError(err)
		if GenerateTokenByIDResponse != nil {
			jwtToken := GenerateTokenByIDResponse.Token
			LoggingIfUserError(err)
			//返回
			ginCtx.JSON(http.StatusOK, user.LoginResp{
				UserId: userResp.UserId,
				Token:  jwtToken,
			})
		}
	} else {
		ginCtx.JSON(400, gin.H{"error": "Invalid User"}) // TODO 拓展
	}

}

// 用户登出
func Logout(ginCtx *gin.Context) {

}

// 更新用户信息
func Update(ginCtx *gin.Context) {
	var userReq user.UpdateReq
	//将获取到的user_id转换成int类型
	if err := ginCtx.ShouldBindJSON(&userReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// 构造一个验证请求
	variRequest := token.VerifyTokenRequest{UserId: userReq.UserId, Token: userReq.Token}
	tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
	_, err := tokenService.VarifyToken(ginCtx, &variRequest)
	if err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 验证通过
	userService := ginCtx.Keys["userService"].(user.UserService)
	userResp, err1 := userService.Update(context.Background(), &userReq)
	LoggingIfUserError(err1)
	//返回
	ginCtx.JSON(http.StatusOK, user.UpdateResp{
		UserId: userResp.UserId,
		Msg:    userResp.Msg,
	})
}

// 删除用户
func DeleteUser(ginCtx *gin.Context) {
	var userReq user.DeleteReq
	//将获取到的user_id转换成int类型
	if err := ginCtx.ShouldBindJSON(&userReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// 构造一个验证请求
	variRequest := token.VerifyTokenRequest{UserId: userReq.UserId, Token: userReq.Token}
	tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
	_, err := tokenService.VarifyToken(ginCtx, &variRequest)
	if err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 验证通过
	userService := ginCtx.Keys["userService"].(user.UserService)
	userResp, err1 := userService.Delete(context.Background(), &userReq)
	LoggingIfUserError(err1)
	ginCtx.JSON(http.StatusOK, gin.H{"Status": userResp})
}

// // 获取用户的详细信息
func UserInfo(ginCtx *gin.Context) {
	var userReq user.UserReq
	var tokenReq token.GetIdByTokenRequest
	//将获取到的user_id转换成int类型
	if err := ginCtx.ShouldBindJSON(&tokenReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
	GetIdByTokenResponse, err := tokenService.GetIdByToken(context.Background(), &tokenReq)
	LoggingIfUserError(err)
	if GetIdByTokenResponse != nil && GetIdByTokenResponse.UserId > 0 {
		userService := ginCtx.Keys["userService"].(user.UserService)
		userReq.UserId = GetIdByTokenResponse.UserId
		userResp, err := userService.UserInfo(context.Background(), &userReq)
		LoggingIfUserError(err)
		if userResp != nil {
			ginCtx.JSON(http.StatusOK, user.UserResp{
				User: userResp.User,
			})
		} else {
			ginCtx.JSON(400, gin.H{"error": "Invalid User"})
		}
	} else {
		ginCtx.JSON(400, gin.H{"error": "Invalid Token"})
	}
}
