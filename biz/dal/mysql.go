package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	CodePasteDB *gorm.DB
)

func InitDB(config map[string]string) {
	dsn := GetDBConf(config)
	var err error
	CodePasteDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 配置，不让表名自动加s
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
}

func GetDBConf(config map[string]string) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["user"],
		config["password"],
		config["host"],
		config["port"],
		config["db_name"])
}
