package service

import (
	"Gomall/app/user/biz/dal/mysql"
	"Gomall/app/user/biz/model"
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.

	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" || req.Role == "" {
		return nil, errors.New("Email, Password, ConfirmPassword, Role can not be empty")
	}

	if req.Password != req.ConfirmPassword {
		return nil, errors.New("Password not match")
	}

	if req.Role != "admin" && req.Role != "user" {
		return nil, errors.New("Role not match")
	}

	_, err = model.GetByEmail(mysql.DB, req.Email)
	if err != gorm.ErrRecordNotFound {
		return nil, errors.New("Email already exists")
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
	}

	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
		Permission:     req.Permission,
		Role:           req.Role,
		Description:    req.Description,
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}


	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
