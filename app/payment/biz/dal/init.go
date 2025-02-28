package dal

import (
	"Gomall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
