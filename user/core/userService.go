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
		resp.UserId = -1
		resp.Token = ""
		return err
	}

	////判断密码是否正确
	if sha256.Sha256(password) != user.Password {
		resp.UserId = -1
		resp.Token = ""
		return nil
	}

	// 判断该用户是否已经被删除
	if user.DeleteAt == 1 {
		resp.Token = "该用户已经被删除，登录失败!"
		return nil
	}

	resp.UserId = user.Id
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
		resp.UserId = user.Id
		return nil
	}
	//
	////创建一个dao层User实体
	user := &model.User{
		Name:     email,
		Email:    email,
		Password: sha256.Sha256(password),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	//
	////调用数据库方法，创建一个新的User实体，返回用户id和err
	userId, err := model.NewUserDao().CreateUser(user)
	if err != nil {
		resp.UserId = -1
		resp.Token = ""
		return err
	}
	resp.UserId = userId
	fmt.Println("Success register user:", user)
	// 写入token
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

func (*UserService) Update(ctx context.Context, req *services.UpdateReq, resp *services.UpdateResp) error {
	user, err := model.NewUserDao().FindUserByEmail(req.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 判断该用户是否已经被删除
	if user.DeleteAt == 1 {
		resp.Msg = "该用户已经被删除，登录失败"
		return nil
	}
	user.UpdateAt = time.Now()
	user.Name = req.Name
	user.Avatar = req.Avatar
	user.BackgroundImage = req.BackgroundImage
	user.Signature = req.Signature
	user.Password = sha256.Sha256(req.Password)
	info, err1 := model.NewUserDao().UpdateUserInfo(user)
	if err1 != nil {
		resp.Msg = "失败"
		return err1
	}
	resp.Msg = "成功！"
	fmt.Println(info)
	return nil
}

func (*UserService) Delete(ctx context.Context, req *services.DeleteReq, resp *services.DeleteResp) error {
	// 要删除这个人 先看看 这个人是不是已经被删除了
	user, err := model.NewUserDao().FindUserByEmail(req.Email)
	if user != nil {
		resp.UserId = user.Id
		if user.DeleteAt == 1 {
			resp.Msg = "当前用户已被删除，不可重复删除！"
			return nil
		}
	}
	if err != nil {
		resp.Msg = "查询该用户失败！"
		return err
	}
	user.DeleteAt = 1
	user.UpdateAt = time.Now()
	info, err := model.NewUserDao().UpdateUserInfo(user)
	resp.Msg = "当前用户已成功删除！"
	fmt.Println(info)
	if err != nil {
		return err
	}
	return nil
}

/*
构建Service层user
*/
func BuildProtoUser(item *model.User) *services.User {
	user := services.User{
		UserId: item.Id,
		Email:  item.Email,
	}
	return &user
}
