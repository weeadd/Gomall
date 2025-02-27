package utils

import (
	"context"

	"Gomall/app/api/infra/rpc"
	"Gomall/rpc_gen/kitex_gen/cart"
	apiutils "Gomall/app/api/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// todo edit custom code
	//content["user_id"] = apiutils.GetUserIdFromCtx(ctx)
	var cartNum int
	userId := apiutils.GetUserIdFromCtx(ctx)
	cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: userId})
	if cartResp != nil && cartResp.Items != nil {
		cartNum = len(cartResp.Items)
	}
	content["user_id"] = ctx.Value(apiutils.UserIdKey)
	content["cart_num"] = cartNum
	return content
}
