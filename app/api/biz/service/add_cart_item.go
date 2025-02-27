package service

import (
	"context"

	cart "Gomall/app/api/hertz_gen/api/cart"
	"Gomall/app/api/infra/rpc"
	apiutils "Gomall/app/api/utils"
	rpccart "Gomall/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *cart.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: apiutils.GetUserIdFromCtx(h.Context),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	return
}
