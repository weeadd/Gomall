package service

import (
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
	return
}
