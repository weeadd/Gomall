package dal

import (
	"Gomall/app/user/biz/dal/mysql"
	"Gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
