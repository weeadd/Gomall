package service

import (
	"Gomall/app/auth/utils"
	auth "Gomall/rpc_gen/kitex_gen/auth"
	"context"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	token, err := utils.GenerateToken(req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &auth.DeliveryResp{
		Token: token,
	}

	return resp, nil
}
