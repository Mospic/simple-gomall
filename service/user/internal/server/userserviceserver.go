// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"mall/service/user/internal/logic"
	"mall/service/user/internal/svc"
	"mall/service/user/proto/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Register(ctx context.Context, in *user.RegisterReq) (*user.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServiceServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServiceServer) Delete(ctx context.Context, in *user.DeleteReq) (*user.DeleteResp, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *UserServiceServer) Update(ctx context.Context, in *user.UpdateReq) (*user.UpdateResp, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *UserServiceServer) Info(ctx context.Context, in *user.InfoReq) (*user.InfoResp, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}

func (s *UserServiceServer) GetMessage(ctx context.Context, in *user.GetMessageReq) (*user.GetMessageResp, error) {
	l := logic.NewGetMessageLogic(ctx, s.svcCtx)
	return l.GetMessage(in)
}
