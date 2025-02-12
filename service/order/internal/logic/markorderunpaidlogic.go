package logic

import (
	"context"
	"errors"
	"fmt"
	"mall/model"
	"mall/service/order/internal/svc"
	"mall/service/order/proto/order"
	"mall/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkOrderUnPaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkOrderUnPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkOrderUnPaidLogic {
	return &MarkOrderUnPaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkOrderUnPaidLogic) MarkOrderUnPaid(in *order.MarkOrderUnPaidReq) (*order.MarkOrderUnPaidResp, error) {
	rdb := l.svcCtx.RDB
	db := l.svcCtx.DB
	log := l.svcCtx.Log
	// 加分布式锁，防止支付时订单删除函数将订单删除
	uid, ok := util.GetLock("order:lock"+fmt.Sprint(in.OrderId), rdb, log)
	if !ok {
		return nil, errors.New("time out")
	}
	err := rdb.ZRem(context.Background(), "order:time", in.OrderId).Err()
	if err != nil {
		log.Error("mark order unpaid fail" + err.Error())
		util.UnLock("order:lock"+fmt.Sprint(in.OrderId), rdb, uid)
		return nil, err
	}
	err = db.Model(&model.Order{}).Where("id = ?", in.OrderId).Update("Paid", "False").Error
	if err != nil {
		log.Error("mark order UnPaid:" + err.Error())
		util.UnLock("order:lock"+fmt.Sprint(in.OrderId), rdb, uid)
		return nil, err
	}

	util.UnLock("order:lock"+fmt.Sprint(in.OrderId), rdb, uid)

	return &order.MarkOrderUnPaidResp{}, nil
}
