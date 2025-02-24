package service

import (
	"Gomall/app/user/biz/dal/mysql"
	"Gomall/app/user/biz/model"
	"Gomall/app/user/infra/rpc"
	auth_rpc "Gomall/rpc_gen/kitex_gen/auth"
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	resp_rpc, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth_rpc.DeliverTokenReq{
		UserId: int32(row.ID),
	})
	if err != nil {
		return nil, err
	}

	resp = &user.LoginResp{
		UserId: int32(row.ID),
		Token:  resp_rpc.Token,
	}

	return resp, nil
}
