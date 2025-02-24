package mysql

import (
	"Gomall/app/order/biz/model"
	"Gomall/app/order/conf"
	"fmt"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		if err := DB.AutoMigrate( //nolint:errcheck
			&model.Order{},
			&model.OrderItem{},
		); err != nil {
			klog.Errorf("auto migrate failed, err:%v", err)
		}
	}
}
