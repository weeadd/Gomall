package service

import (
	"Gomall/app/api/infra/rpc"
	"context"

	convert "Gomall/app/api/biz/utils"
	order "Gomall/app/api/hertz_gen/api/order"
	order_rpc "Gomall/rpc_gen/kitex_gen/order"

	"github.com/cloudwego/hertz/pkg/app"
)

type PlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *PlaceOrderService {
	return &PlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp_rpc, err := rpc.OrderClient.PlaceOrder(h.Context, &order_rpc.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address:      convert.ConvertAddress(req.Address),
		Email:        req.Email,
		OrderItems:   convert.ConvertOrderItems(req.OrderItems),
	})
	if err != nil {
		return nil, err
	}

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: resp_rpc.Order.OrderId,
		},
	}

	return resp, nil
}
