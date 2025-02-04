package core

import (
	"cart/model"
	protos "cart/services"
	"context"
)

type CartService struct{}

// 添加商品
func (*CartService) AddItem(ctx context.Context, req *protos.AddItemReq, resp *protos.AddItemResp) error {
	product := model.CartProduct{
		ProductID: req.ProductId,
		Name:      req.Name,
		Quantity:  req.Quantity,
		UnitPrice: req.UnitPrice,
	}

	err := model.AddCartItem(req.UserId, product)
	if err != nil {
		resp.Success = false
		resp.Message = "添加失败"
		return err
	}

	resp.Success = true
	resp.Message = "添加成功"
	return nil
}

func (*CartService) RemoveItem(ctx context.Context, req *protos.RemoveItemReq, resp *protos.RemoveItemResp) error {
	err := model.RemoveCartItem(req.UserId, req.ProductId)
	if err != nil {
		resp.Success = false
		resp.Message = "删除失败"
		return err
	}

	resp.Success = true
	resp.Message = "删除成功"
	return nil
}

func (*CartService) GetCart(ctx context.Context, req *protos.GetCartReq, resp *protos.GetCartResp) error {
	items, totalPrice, err := model.GetCartItems(req.UserId)
	if err != nil {
		return err
	}

	resp.TotalPrice = totalPrice
	for _, item := range items {
		resp.Items = append(resp.Items, &protos.CartItem{
			ProductId: item.ProductID,
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}
	return nil
}
