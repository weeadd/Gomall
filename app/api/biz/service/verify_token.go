package service

import (
	"context"

	auth "Gomall/app/api/hertz_gen/api/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type VerifyTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewVerifyTokenService(Context context.Context, RequestContext *app.RequestContext) *VerifyTokenService {
	return &VerifyTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *VerifyTokenService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
