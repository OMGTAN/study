package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitPg() {
	dsn := viper.GetString("pg.url")
	fmt.Println("db.url=", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println("init pg err {}", err)
	}
	DB = db
}

func InitConfig() {

	viper.AddConfigPath("config")
	viper.SetConfigName("app")
	viper.ReadInConfig()
	dsn := viper.GetString("pg.url")
	fmt.Println(dsn)
}
