package dal

import (
	"Gomall/app/order/biz/dal/mysql"
	// "Gomall/app/order/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
