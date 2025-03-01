package service

import (
	"context"

	payment "Gomall/app/api/hertz_gen/api/payment"
	"Gomall/app/api/infra/rpc"
	payment_rpc "Gomall/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/hertz/pkg/app"
)

type ChargeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChargeService(Context context.Context, RequestContext *app.RequestContext) *ChargeService {
	return &ChargeService{RequestContext: RequestContext, Context: Context}
}

func (h *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	resp_rpc, err := rpc.PaymentClient.Charge(h.Context, &payment_rpc.ChargeReq{
		Amount:  req.Amount,
		UserId:  req.UserId,
		OrderId: req.OrderId,
		CreditCard: &payment_rpc.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	})
	if err != nil {
		return nil, err
	}
	resp = &payment.ChargeResp{}
	resp.TransactionId = resp_rpc.TransactionId
	// resp = &payment.ChargeResp{
	// 	TransactionId: resp_rpc.TransactionId,
	// }

	return resp, nil
}
