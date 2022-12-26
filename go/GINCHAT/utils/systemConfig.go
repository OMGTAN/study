package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
)

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

func InitRedis() {
	url := viper.GetString("redis.url")
	port := viper.GetString("redis.url")
	Rdb = redis.NewClient(&redis.Options{
		Addr:     url + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

const (
	PublishKey = "websocket"
)

// publish 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	err = Rdb.Publish(ctx, channel, msg).Err()
	return err
}

// subscribe 订阅消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Rdb.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("receive msg : ", msg)
	if msg != nil {
		return msg.Payload, err
	} else {
		return "", err
	}
}
