package mysql

import (
	"Gomall/app/cart/biz/model"
	"Gomall/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
