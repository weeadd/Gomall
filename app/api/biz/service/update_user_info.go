package service

import (
	"Gomall/app/api/infra/rpc"
	"context"

	user_rpc "Gomall/rpc_gen/kitex_gen/user"

	user "Gomall/app/api/hertz_gen/api/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Run(req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	resp_rpc, err := rpc.UserClient.UpdateUserInfo(h.Context, &user_rpc.UpdateUserInfoReq{
		UserId:      req.UserId,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	resp = &user.UpdateResp{
		Res: resp_rpc.Res,
	}
	return resp, nil
}
