package service

import (
	"Gomall/app/order/biz/dal/mysql"
	"Gomall/app/order/biz/model"
	"Gomall/rpc_gen/kitex_gen/cart"
	order "Gomall/rpc_gen/kitex_gen/order"
	"context"
	"fmt"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var orderList []*order.Order

	for _, v := range list {
		var itemList []*order.OrderItem
		for _, item := range v.OrderItems {
			itemList = append(itemList, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				},
				Cost: item.Cost,
			})
		}
		orderList = append(orderList, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				Country:       v.Consignee.Country,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: itemList,
		})
	}
	fmt.Println("orderList:", orderList)
	resp = &order.ListOrderResp{
		Orders: orderList,
	}
	return resp, nil
}
