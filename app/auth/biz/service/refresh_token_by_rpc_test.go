package service

import (
	"context"
	"testing"
	auth "Gomall/rpc_gen/kitex_gen/auth"
)

func TestRefreshTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRefreshTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.RefreshTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
