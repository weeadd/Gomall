package service

import (
	"Gomall/app/api/infra/rpc"
	user_rpc "Gomall/rpc_gen/kitex_gen/user"
	"context"

	user "Gomall/app/api/hertz_gen/api/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUserService(Context context.Context, RequestContext *app.RequestContext) *DeleteUserService {
	return &DeleteUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUserService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	resp_rpc, err := rpc.UserClient.DeleteUser(h.Context, &user_rpc.DeleteUserReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	resp = &user.DeleteResp{
		Res: resp_rpc.Res,
	}
	return resp, nil
}
