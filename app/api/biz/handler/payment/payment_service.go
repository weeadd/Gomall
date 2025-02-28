package payment

import (
	"Gomall/app/api/biz/service"
	"Gomall/app/api/biz/utils"
	payment "Gomall/app/api/hertz_gen/api/payment"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// charge .
// @router /payment/charge [POST]
func Charge(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.ChargeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.ChargeResp{}
	resp, err = service.NewChargeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
