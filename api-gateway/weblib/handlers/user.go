package handlers

import (
	"api-gateway/pkg/utils"
	user "api-gateway/services/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	PanicIfUserError(err)
	fmt.Printf("User Response:%v", userResp)
	if userResp != nil && userResp.UserId > 0 { //做一下保护，返回的UserId应该大于0
		_, err = utils.GenerateToken(userResp.UserId)
		PanicIfUserError(err)
		//返回
		ginCtx.JSON(http.StatusOK, user.RegisterResp{
			UserId: userResp.UserId,
		})
	} else {
		ginCtx.JSON(400, gin.H{"error": "Invalid User"})
	}

}

// 用户登录
func Login(ginCtx *gin.Context) {
	var userReq user.LoginReq
	if err := ginCtx.ShouldBindJSON(&userReq); err != nil {
		fmt.Println(err)
		ginCtx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(user.UserService)
	userResp, err := userService.Login(context.Background(), &userReq)
	PanicIfUserError(err)

	//生成token
	if userResp != nil && userResp.UserId > 0 {
		_, err = utils.GenerateToken(userResp.UserId)
		PanicIfUserError(err)

		ginCtx.JSON(http.StatusOK, user.LoginResp{
			UserId: userResp.UserId,
		})
	} else {
		ginCtx.JSON(400, gin.H{"error": "Invalid User"})
	}

}

// // 获取用户的详细信息
func UserInfo(ginCtx *gin.Context) {
	//var userReq user.DouyinUserRequest
	////将获取到的user_id转换成int类型
	//user_id, err := strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)
	//PanicIfUserError(err)
	//
	//userReq.UserId = user_id
	//userReq.Token = ginCtx.Query("token")
	//
	//userService := ginCtx.Keys["userService"].(user.UserService)
	//userResp, err := userService.UserInfo(context.Background(), &userReq)
	//PanicIfUserError(err)
	//
	//ginCtx.JSON(http.StatusOK, user.DouyinUserResponse{
	//	StatusCode: userResp.StatusCode,
	//	StatusMsg:  userResp.StatusMsg,
	//	User:       userResp.User,
	//})
}
