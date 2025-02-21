package service

import (
	"Gomall/app/api/infra/rpc"
	"context"

	user "Gomall/app/api/hertz_gen/api/user"

	user_rpc "Gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserInfoService(Context context.Context, RequestContext *app.RequestContext) *GetUserInfoService {
	return &GetUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserInfoService) Run(req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	resp_rpc, err := rpc.UserClient.GetUserInfo(h.Context, &user_rpc.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	resp = &user.UserInfoResp{
		Email:       resp_rpc.Email,
		Permission:  resp_rpc.Permission,
		Role:        resp_rpc.Role,
		Description: resp_rpc.Description,
	}
	return resp, nil
}
