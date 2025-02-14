package dal

import (
	"Gomall/app/auth/biz/dal/mysql"
	"Gomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
