package auth

import (
	"Gomall/app/api/biz/service"
	"Gomall/app/api/biz/utils"
	"context"

	auth "Gomall/app/api/hertz_gen/api/auth"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DeliverToken .
// @router /auth/deliver [POST]
func DeliverToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.DeliverTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &auth.DeliveryResp{}
	resp, err = service.NewDeliverTokenService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// VerifyToken .
// @router /auth/verify [POST]
func VerifyToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.VerifyTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &auth.VerifyResp{}
	resp, err = service.NewVerifyTokenService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
