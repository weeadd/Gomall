package service

import (
	"Gomall/app/user/biz/dal/mysql"
	"Gomall/app/user/biz/model"
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"

	"gorm.io/gorm"
)

type UpdateUserInfoService struct {
	ctx context.Context
} // NewUpdateUserInfoService new UpdateUserInfoService
func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	_, err = model.GetById(mysql.DB, req.UserId)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == nil {
		resp = &user.UpdateUserInfoResp{
			Res: false,
		}
		return resp, nil
	}

	new_user_info := model.User{
		Description: req.Description,
	}
	err = model.UpdateById(mysql.DB, req.UserId, &new_user_info)
	if err != nil {
		return nil, err
	}

	resp = &user.UpdateUserInfoResp{
		Res: true,
	}

	return resp, nil
}
