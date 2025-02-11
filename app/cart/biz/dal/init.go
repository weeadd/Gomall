package dal

import (
	"Gomall/app/cart/biz/dal/mysql"
	"Gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
