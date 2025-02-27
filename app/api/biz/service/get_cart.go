package service

import (
	"context"
	"strconv"

	cart "Gomall/app/api/hertz_gen/api/cart"
	rpccart "Gomall/rpc_gen/kitex_gen/cart"
	"Gomall/app/api/infra/rpc"
	apiutils "Gomall/app/api/utils"
	rpcproduct "Gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var items []map[string]string
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint32(h.Context.Value(apiutils.UserIdKey).(float64)),
	})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.GetProductId()})
		if err != nil {
			continue
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{"Name": p.Name, "Description": p.Description, "Picture": p.Picture, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Qty": strconv.Itoa(int(v.Quantity))})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
