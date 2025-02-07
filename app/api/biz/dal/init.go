package dal

import (
	"Gomall/app/api/biz/dal/mysql"
	"Gomall/app/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
