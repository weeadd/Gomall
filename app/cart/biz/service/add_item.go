package service

import (
	"Gomall/app/cart/biz/dal/mysql"
	"Gomall/app/cart/biz/model"
	"Gomall/app/cart/infra/rpc"
	cart "Gomall/rpc_gen/kitex_gen/cart"
	"Gomall/rpc_gen/kitex_gen/product"
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AddItemService struct {
	ctx context.Context
}

// NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// 检查 req 是否为 nil
	if req == nil {
		klog.Error("req is nil")
		return nil, kerrors.NewBizStatusError(40000, "req is nil")
	}

	// 检查 req.Item 是否为 nil
	if req.Item == nil {
		klog.Error("req.Item is nil")
		return nil, kerrors.NewBizStatusError(40001, "req.Item is nil")
	}

	// 检查 rpc.ProductClient 是否为 nil
	if rpc.ProductClient == nil {
		klog.Error("rpc.ProductClient is nil")
		return nil, kerrors.NewBizStatusError(50000, "rpc.ProductClient is not initialized")
	}

	// 调用 RPC 获取商品信息
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		klog.Errorf("GetProduct RPC failed: %v", err)
		return nil, err
	}

	// 检查 productResp.Product 是否为 nil
	if productResp == nil || productResp.Product == nil {
		klog.Error("product not found")
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}

	// 检查 mysql.DB 是否为 nil
	if mysql.DB == nil {
		klog.Error("mysql.DB is nil")
		return nil, kerrors.NewBizStatusError(50000, "mysql.DB is not initialized")
	}

	// 添加商品到购物车
	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProudctId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		klog.Errorf("AddItem failed: %v", err)
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	// 返回成功响应
	return &cart.AddItemResp{}, nil
}