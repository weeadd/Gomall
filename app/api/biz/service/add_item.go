package service

import (
	"context"

	cart "Gomall/app/api/hertz_gen/api/cart"
	"Gomall/app/api/infra/rpc"
	rpccart "Gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddItemService(Context context.Context, RequestContext *app.RequestContext) *AddItemService {
	return &AddItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: req.UserId,
		Item: &rpccart.CartItem{
			ProductId: req.Item.ProductId,
			Quantity:  req.Item.Quantity,
		},
	})
	return
}
