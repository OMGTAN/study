package main

import (
	"fmt"
	"ginchat/models"
  "gorm.io/gorm"
  "gorm.io/gorm/schema"
  "gorm.io/driver/postgres"
)

func main() {
  dsn := "host=localhost user=postgres password=123456 dbname=ginchat port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 使用单数表名
	},
  })
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移 schema
  db.AutoMigrate(&models.UserBasic{})

  // Create
  user := &models.UserBasic{}
  user.Name="test"
  db.Create(user)

  // Read
  fmt.Println(db.First(user, 1)) // 根据整型主键查找
//   db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

  // Update - 将 product 的 price 更新为 200
  db.Model(user).Update("Password", "200")
  // Update - 更新多个字段
//   db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
//   db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - 删除 product
//   db.Delete(&product, 1)
}