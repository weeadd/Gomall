package service

import (
	"Gomall/app/auth/utils"
	auth "Gomall/rpc_gen/kitex_gen/auth"
	"context"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp = &auth.VerifyResp{}

	_, err = utils.ParseToken(req.Token)
	if err != nil {
		resp.Res = false
		return resp, err
	}

	resp.Res = true
	return resp, nil
}
