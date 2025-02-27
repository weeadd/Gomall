package order

import (
	"context"

	"Gomall/app/api/biz/service"
	"Gomall/app/api/biz/utils"
	order "Gomall/app/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PlaceOrder .
// @router /order/place [POST]
func PlaceOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.PlaceOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.PlaceOrderResp{}
	resp, err = service.NewPlaceOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListOrder .
// @router /order/list [POST]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.ListOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.ListOrderResp{}
	resp, err = service.NewListOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MarkOrderPaid .
// @router /order/mark [POST]
func MarkOrderPaid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.MarkOrderPaidReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.MarkOrderPaidResp{}
	resp, err = service.NewMarkOrderPaidService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
