// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Register(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error)
	UserInfo(ctx context.Context, in *UserReq, opts ...client.CallOption) (*UserResp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*DeleteResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...client.CallOption) (*LogoutResp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*RegisterResp, error) {
	req := c.c.NewRequest(c.name, "UserService.Register", in)
	out := new(RegisterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error) {
	req := c.c.NewRequest(c.name, "UserService.Login", in)
	out := new(LoginResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserInfo(ctx context.Context, in *UserReq, opts ...client.CallOption) (*UserResp, error) {
	req := c.c.NewRequest(c.name, "UserService.UserInfo", in)
	out := new(UserResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error) {
	req := c.c.NewRequest(c.name, "UserService.Update", in)
	out := new(UpdateResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*DeleteResp, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(DeleteResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Logout(ctx context.Context, in *LogoutReq, opts ...client.CallOption) (*LogoutResp, error) {
	req := c.c.NewRequest(c.name, "UserService.Logout", in)
	out := new(LogoutResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Register(context.Context, *RegisterReq, *RegisterResp) error
	Login(context.Context, *LoginReq, *LoginResp) error
	UserInfo(context.Context, *UserReq, *UserResp) error
	Update(context.Context, *UpdateReq, *UpdateResp) error
	Delete(context.Context, *DeleteReq, *DeleteResp) error
	Logout(context.Context, *LogoutReq, *LogoutResp) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Register(ctx context.Context, in *RegisterReq, out *RegisterResp) error
		Login(ctx context.Context, in *LoginReq, out *LoginResp) error
		UserInfo(ctx context.Context, in *UserReq, out *UserResp) error
		Update(ctx context.Context, in *UpdateReq, out *UpdateResp) error
		Delete(ctx context.Context, in *DeleteReq, out *DeleteResp) error
		Logout(ctx context.Context, in *LogoutReq, out *LogoutResp) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Register(ctx context.Context, in *RegisterReq, out *RegisterResp) error {
	return h.UserServiceHandler.Register(ctx, in, out)
}

func (h *userServiceHandler) Login(ctx context.Context, in *LoginReq, out *LoginResp) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *userServiceHandler) UserInfo(ctx context.Context, in *UserReq, out *UserResp) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}

func (h *userServiceHandler) Update(ctx context.Context, in *UpdateReq, out *UpdateResp) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *DeleteReq, out *DeleteResp) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *userServiceHandler) Logout(ctx context.Context, in *LogoutReq, out *LogoutResp) error {
	return h.UserServiceHandler.Logout(ctx, in, out)
}
