package handlers

import (
	"api-gateway/pkg/utils/redis"
	token "api-gateway/services/tokenutils"
	user "api-gateway/services/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 用户注册
func Register(ginCtx *gin.Context) {
	var userReq user.RegisterReq
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
	ginCtx.JSON(200, gin.H{
		"Msg":     "用户创建成功！",
		"user_id": userResp.UserId,
	})
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
			// 将token存入redis
			err := redis.RdbJwt.Set(context.Background(), strconv.Itoa(int(userResp.UserId)), jwtToken, 0).Err()
			if err != nil {
				ginCtx.JSON(400, gin.H{"error": err.Error()})
				return
			}
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
	var userReq user.LogoutReq
	//将获取到的user_id转换成int类型
	if err := ginCtx.ShouldBindJSON(&userReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// 从redis中拿token
	var tokenReq token.GetTokenByRedisRequest
	tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
	tokenReq.UserId = userReq.UserId
	GetTokenByRedisResponse, _ := tokenService.GetTokenByRedis(context.Background(), &tokenReq)
	if GetTokenByRedisResponse != nil {
		RedisToken := GetTokenByRedisResponse.Token
		// 构造一个验证请求
		variRequest := token.VerifyTokenRequest{UserId: userReq.UserId, Token: RedisToken}
		tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
		_, err := tokenService.VarifyToken(ginCtx, &variRequest)
		if err != nil {
			ginCtx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 验证通过
		userService := ginCtx.Keys["userService"].(user.UserService)
		userResp, err := userService.Logout(context.Background(), &userReq)
		if err != nil {
			ginCtx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ginCtx.JSON(200, gin.H{
			"user_id": userResp.UserId,
			"Msg":     userResp.Msg,
		})
	} else {
		ginCtx.JSON(400, gin.H{"error": "用户未登录！"})
	}
}

// 更新用户信息
func Update(ginCtx *gin.Context) {
	var userReq user.UpdateReq
	var tokenReq token.GetTokenByRedisRequest
	//将获取到的user_id转换成int类型
	if err := ginCtx.ShouldBindJSON(&userReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// 从redis中拿token
	tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
	tokenReq.UserId = userReq.UserId
	GetTokenByRedisResponse, _ := tokenService.GetTokenByRedis(context.Background(), &tokenReq)
	if GetTokenByRedisResponse != nil {
		RedisToken := GetTokenByRedisResponse.Token
		// 构造一个验证请求
		variRequest := token.VerifyTokenRequest{UserId: userReq.UserId, Token: RedisToken}
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
	} else {
		ginCtx.JSON(400, gin.H{"error": "该用户未登录"})
	}
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
	// 从redis中拿token
	var tokenReq token.GetTokenByRedisRequest
	tokenService := ginCtx.Keys["tokenService"].(token.TokenService)
	tokenReq.UserId = userReq.UserId
	GetTokenByRedisResponse, _ := tokenService.GetTokenByRedis(context.Background(), &tokenReq)
	if GetTokenByRedisResponse != nil {
		RedisToken := GetTokenByRedisResponse.Token
		// 构造一个验证请求
		variRequest := token.VerifyTokenRequest{UserId: userReq.UserId, Token: RedisToken}
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
	} else {
		ginCtx.JSON(400, gin.H{"error": "该用户未登录"})
	}
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
