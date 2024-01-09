package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:255;index:idx_name,unique"`
}

func before() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=fly_go port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func test1(db *gorm.DB) {
	// 迁移 schema
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic("migrate table error")
	}
	// Create
	//now := time.Now()
	//timeString := fmt.Sprintf("%d", now)
	//db.Create(&User{Name: "xiaoqiang" + timeString})
}

func main() {
	db := before()
	test1(db)

	var user User
	db.First(&user, 1)
	fmt.Println(user)
	u := &User{
		Name: "xiaoming223",
	}
	u.ID = 5
	db.Model(&User{}).Where("id=?", 5).Updates(u)
}
