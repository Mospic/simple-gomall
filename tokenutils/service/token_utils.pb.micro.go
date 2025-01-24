// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/proto/token_utils.proto

package service

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

// Api Endpoints for TokenService service

func NewTokenServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TokenService service

type TokenService interface {
	GetIdByToken(ctx context.Context, in *GetIdByTokenRequest, opts ...client.CallOption) (*GetIdByTokenResponse, error)
	GenerateTokenByID(ctx context.Context, in *GenerateTokenByIDRequest, opts ...client.CallOption) (*GenerateTokenByIDResponse, error)
}

type tokenService struct {
	c    client.Client
	name string
}

func NewTokenService(name string, c client.Client) TokenService {
	return &tokenService{
		c:    c,
		name: name,
	}
}

func (c *tokenService) GetIdByToken(ctx context.Context, in *GetIdByTokenRequest, opts ...client.CallOption) (*GetIdByTokenResponse, error) {
	req := c.c.NewRequest(c.name, "TokenService.GetIdByToken", in)
	out := new(GetIdByTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) GenerateTokenByID(ctx context.Context, in *GenerateTokenByIDRequest, opts ...client.CallOption) (*GenerateTokenByIDResponse, error) {
	req := c.c.NewRequest(c.name, "TokenService.GenerateTokenByID", in)
	out := new(GenerateTokenByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TokenService service

type TokenServiceHandler interface {
	GetIdByToken(context.Context, *GetIdByTokenRequest, *GetIdByTokenResponse) error
	GenerateTokenByID(context.Context, *GenerateTokenByIDRequest, *GenerateTokenByIDResponse) error
}

func RegisterTokenServiceHandler(s server.Server, hdlr TokenServiceHandler, opts ...server.HandlerOption) error {
	type tokenService interface {
		GetIdByToken(ctx context.Context, in *GetIdByTokenRequest, out *GetIdByTokenResponse) error
		GenerateTokenByID(ctx context.Context, in *GenerateTokenByIDRequest, out *GenerateTokenByIDResponse) error
	}
	type TokenService struct {
		tokenService
	}
	h := &tokenServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TokenService{h}, opts...))
}

type tokenServiceHandler struct {
	TokenServiceHandler
}

func (h *tokenServiceHandler) GetIdByToken(ctx context.Context, in *GetIdByTokenRequest, out *GetIdByTokenResponse) error {
	return h.TokenServiceHandler.GetIdByToken(ctx, in, out)
}

func (h *tokenServiceHandler) GenerateTokenByID(ctx context.Context, in *GenerateTokenByIDRequest, out *GenerateTokenByIDResponse) error {
	return h.TokenServiceHandler.GenerateTokenByID(ctx, in, out)
}
