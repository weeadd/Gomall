package service

import (
	"context"

	order "Gomall/app/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type MarkOrderPaidService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMarkOrderPaidService(Context context.Context, RequestContext *app.RequestContext) *MarkOrderPaidService {
	return &MarkOrderPaidService{RequestContext: RequestContext, Context: Context}
}

func (h *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
