package service

import (
	"context"

	cart "Gomall/app/api/hertz_gen/api/cart"
	"Gomall/app/api/infra/rpc"
	rpccart "Gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.EmptyCart(h.Context, &rpccart.EmptyCartReq{
		UserId: req.UserId,
	})

	if err != nil {
		return nil, err
	}
	return
}
