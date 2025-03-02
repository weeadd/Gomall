package service

import (
	"Gomall/app/api/infra/rpc"
	"context"

	convert "Gomall/app/api/biz/utils"
	checkout "Gomall/app/api/hertz_gen/api/checkout"
	checkout_rpc "Gomall/rpc_gen/kitex_gen/checkout"

	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp_rpc, err := rpc.CheckoutClient.Checkout(h.Context, &checkout_rpc.CheckoutReq{
		UserId:     req.UserId,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Email:      req.Email,
		Address:    convert.ConvertCheckoutAddress(req.Address),
		CreditCard: convert.ConvertCreditCard(req.CreditCard),
	})

	if err != nil {
		return nil, err
	}

	resp = &checkout.CheckoutResp{
		OrderId:       resp_rpc.OrderId,
		TransactionId: resp_rpc.TransactionId,
	}

	return resp, nil
}
