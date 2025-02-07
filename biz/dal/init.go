package dal

import (
	"Gomall/biz/dal/mysql"
	"Gomall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
