package dal

import (
	"Gomall/app/checkout/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
