package dal

import (
	"Gomall/app/product/biz/dal/mysql"
	"Gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
