package service

import (
	"context"
	"errors"
	"strconv"

	cart "Gomall/app/api/hertz_gen/api/cart"
	rpccart "Gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "Gomall/rpc_gen/kitex_gen/product"

	"Gomall/app/api/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	//"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.GetCartReq) (resp map[string]any, err error) {
	var items []map[string]string

	if req == nil {
        hlog.CtxErrorf(h.Context, "req is nil")
        return nil, errors.New("req is nil")
    }

	if rpc.CartClient == nil {
		hlog.CtxErrorf(h.Context, "rpc.CartClient is nil")
		return nil, errors.New("rpc.CartClient is not initialized")
	}

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: req.UserId,
	})

	if carts == nil {
		hlog.CtxErrorf(h.Context, "carts is nil")
		return nil, errors.New("carts is nil")
	}
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

	if carts.Items == nil {
		hlog.CtxWarnf(h.Context, "carts.Items is nil")
		return utils.H{
			"items": []map[string]string{},
			"total": "0.00",
		}, nil
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
