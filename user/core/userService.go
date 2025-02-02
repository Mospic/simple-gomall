package core

import (
	"context"
	"fmt"
	"time"
	"user/model"
	services "user/services"
	"user/utils/sha256"
)

type UserService struct {
}

/*
*
用户登录，service层，
req:用户名，密码
resp:
*/
func (*UserService) Login(ctx context.Context, req *services.LoginReq, resp *services.LoginResp) error {

	email := req.Email
	password := req.Password
	////判断用户名和密码是否为空
	if email == "" || password == "" {
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	user, err := model.NewUserDao().FindUserByEmail(email)
	if err != nil {
		return err
	}

	////判断密码是否正确
	if sha256.Sha256(password) != user.Password {
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	resp.UserId = user.UserId
	return nil
}

/**
注册 service层
req:用户名，密码
resp:新生成的userId，
*/
//TODO 有查询是否有用户和插入数据，最好做个事务，防止有同名用户
func (*UserService) Register(ctx context.Context, req *services.RegisterReq, resp *services.RegisterResp) error {
	//在req中获取用户名和密码
	email := req.Email
	password := req.Password
	confirmPassword := req.ConfirmPassword
	//
	////用户名和密码为空，返回
	if email == "" || password == "" {
		resp.UserId = -1
		resp.Token = ""
		return nil
	}
	// password do not match confirm password
	if password != confirmPassword {
		resp.UserId = -1
		resp.Token = ""
		return nil
	}
	//
	////调用数据库方法，查询是否有同名实体
	if user, err := model.NewUserDao().FindUserByEmail(email); err == nil {
		resp.UserId = user.UserId
		return nil
	}
	//
	////创建一个dao层User实体
	user := &model.User{
		Email:    email,
		Password: sha256.Sha256(password),
		Name:     email,
		CreateAt: time.Now(),
	}
	//
	////调用数据库方法，创建一个新的User实体
	_, err := model.NewUserDao().CreateUser(user)
	if err != nil {
		resp.UserId = -1
		resp.Token = ""
		return err
	}
	//
	////根据用户名，查询新用户的userId，作为返回值返回
	user, _ = model.NewUserDao().FindUserByEmail(email)
	//tokenService := services.NewTokenService("rpcTokenService", services.Client())
	//generateRes, err := tokenService.GenerateTokenByID(ctx, generateReq)
	//
	////补充resp
	//resp.StatusCode = 0
	//resp.StatusMsg = "注册成功"
	resp.UserId = user.UserId
	fmt.Println("Success register user:", user)
	resp.Token = ""
	return nil
}

/**
用户信息查询：
req: userID
resp: User类
*/
//TODO，查出来的数据可以放在缓存里, 这部分还没做，等购物车等功能出来再做
func (*UserService) UserInfo(ctx context.Context, req *services.UserReq, resp *services.UserResp) error {
	user, err := model.NewUserDao().FindUserByID(req.UserId)
	if err != nil {
		resp.User = nil
		return err
	}
	resp.User = BuildProtoUser(user)
	return nil
}

/*
构建Service层user
*/
func BuildProtoUser(item *model.User) *services.User {
	user := services.User{
		UserId: item.UserId,
		Email:  item.Email,
	}
	return &user
}
