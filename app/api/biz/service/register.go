package service

import (
	"Gomall/app/api/infra/rpc"
	"context"

	user "Gomall/app/api/hertz_gen/api/user"

	user_rpc "Gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp_rpc, err := rpc.UserClient.Register(h.Context, &user_rpc.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		return nil, err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp_rpc.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	resp = &user.RegisterResp{
		UserId: resp_rpc.UserId,
	}

	return resp, nil
}
