package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitPg() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			// IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful: true, // Disable color
		},
	)

	dsn := viper.GetString("pg.url")
	fmt.Println("db.url=", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: newLogger,
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
