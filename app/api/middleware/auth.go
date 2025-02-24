package middleware

import (
	"Gomall/app/api/biz/utils"
	"Gomall/app/api/infra/rpc"
	auth_rpc "Gomall/rpc_gen/kitex_gen/auth"
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")

		if token == "" {
			utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("Token is required"))
			c.Abort()
			return
		}

		resp_rpc, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth_rpc.VerifyTokenReq{
			Token: token,
		})
		if err != nil {
			utils.SendErrResponse(ctx, c, consts.StatusOK, err)
			c.Abort()
			return
		}

		if !resp_rpc.Res {
			utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("Invalid Token"))
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
