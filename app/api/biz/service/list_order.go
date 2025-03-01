package service

import (
	"context"

	convert "Gomall/app/api/biz/utils"
	order "Gomall/app/api/hertz_gen/api/order"
	order_rpc "Gomall/rpc_gen/kitex_gen/order"

	"Gomall/app/api/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
)

type ListOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrderService(Context context.Context, RequestContext *app.RequestContext) *ListOrderService {
	return &ListOrderService{RequestContext: RequestContext, Context: Context}
}
func (h *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp_rpc, err := rpc.OrderClient.ListOrder(h.Context, &order_rpc.ListOrderReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	// 手动转换 Orders 字段
	var orders []*order.Order
	for _, o := range resp_rpc.Orders {
		// 手动转换每个 Order 的字段
		orders = append(orders, &order.Order{
			OrderItems:   convert.ConvertOrderItemsBack(o.OrderItems), // 转换 OrderItems
			OrderId:      o.OrderId,                                   // 直接赋值 OrderId
			UserId:       o.UserId,                                    // 直接赋值 UserId
			UserCurrency: o.UserCurrency,                              // 直接赋值 UserCurrency
			Address:      convert.ConvertAddressBack(o.Address),       // 转换 Address
			Email:        o.Email,                                     // 直接赋值 Email
			CreatedAt:    o.CreatedAt,                                 // 直接赋值 CreatedAt
		})
	}

	resp = &order.ListOrderResp{
		Orders: orders,
	}

	return resp, nil
}
