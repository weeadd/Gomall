package user

import (
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Logout(ctx context.Context, req *user.LogoutReq, callOptions ...callopt.Option) (resp *user.LogoutResp, err error) {
	resp, err = defaultClient.Logout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Logout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *user.UserInfoReq, callOptions ...callopt.Option) (resp *user.UserInfoResp, err error) {
	resp, err = defaultClient.GetUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = defaultClient.UpdateUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *user.DeleteUserReq, callOptions ...callopt.Option) (resp *user.DeleteUserResp, err error) {
	resp, err = defaultClient.DeleteUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
