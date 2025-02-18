package service

import (
	"Gomall/app/user/biz/dal/mysql"
	"Gomall/app/user/biz/model"
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	// Finish your business logic.
	row, err := model.GetById(mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}

	resp = &user.UserInfoResp{
		Email:      row.Email,
		Permission: row.Permission,
	}

	return resp, nil
}
