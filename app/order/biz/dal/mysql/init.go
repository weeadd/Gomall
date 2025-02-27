package mysql

import (
	"Gomall/app/order/biz/model"
	"Gomall/app/order/conf"
	"fmt"
	"log"
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
	// dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	fmt.Println("Connecting to MySQL with DSN:", conf.GetConf().MySQL.DSN)
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {

		log.Printf("Error while connecting to MySQL: %v", err)

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
