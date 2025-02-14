package service

import (
	"Gomall/app/api/infra/rpc"
	"context"

	user "Gomall/app/api/hertz_gen/api/user"

	user_rpc "Gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp_rpc, err := rpc.UserClient.Login(h.Context, &user_rpc.LoginReq{
		Email:    req.Email,
		Password: req.Password,
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

	resp = &user.LoginResp{
		UserId: resp_rpc.UserId,
	}

	return resp, nil
}
