package dal

import (
	"Gomall/app/cart/biz/dal/mysql"
)

func Init() {
	//	redis.Init()
	mysql.Init()
}
