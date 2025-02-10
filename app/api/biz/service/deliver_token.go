package service

import (
	"context"

	auth "Gomall/app/api/hertz_gen/api/auth"

	"github.com/cloudwego/hertz/pkg/app"
)

type DeliverTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeliverTokenService(Context context.Context, RequestContext *app.RequestContext) *DeliverTokenService {
	return &DeliverTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *DeliverTokenService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	//defer func() {
	//hlog.CtxInfof(h.Context, "req = %+v", req)
	//hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp = &auth.DeliveryResp{}

	resp.Token = "ok"

	return resp, nil
}
