// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	user "Gomall/app/api/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.POST("/delete", append(_deleteuserMw(), user.DeleteUser)...)
		_user.POST("/info", append(_getuserinfoMw(), user.GetUserInfo)...)
		_user.POST("/login", append(_loginMw(), user.Login)...)
		_user.POST("/logout", append(_logoutMw(), user.Logout)...)
		_user.POST("/register", append(_registerMw(), user.Register)...)
		_user.POST("/update", append(_updateuserinfoMw(), user.UpdateUserInfo)...)
	}
}
