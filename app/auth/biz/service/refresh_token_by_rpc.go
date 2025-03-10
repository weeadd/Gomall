package service

import (
	"Gomall/app/auth/utils"
	auth "Gomall/rpc_gen/kitex_gen/auth"
	"context"
)

type RefreshTokenByRPCService struct {
	ctx context.Context
} // NewRefreshTokenByRPCService new RefreshTokenByRPCService
func NewRefreshTokenByRPCService(ctx context.Context) *RefreshTokenByRPCService {
	return &RefreshTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *RefreshTokenByRPCService) Run(req *auth.RefreshTokenReq) (resp *auth.RefreshTokenResp, err error) {
	_, err = utils.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(req.UserId)
	if err != nil {
		return nil, err
	}

	resp = &auth.RefreshTokenResp{
		Token: token,
	}

	return resp, nil
}
