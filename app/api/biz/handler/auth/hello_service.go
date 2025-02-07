package auth

import (
	"Gomall/app/api/biz/utils"
	"context"

	auth "Gomall/app/api/hertz_gen/api/auth"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Method1 .
// @router /hello [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &auth.HelloResp{}
	// resp, err = service.NewMethod1Service(ctx, c).Run(&req)
	resp.RespBody = "Hello"
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
