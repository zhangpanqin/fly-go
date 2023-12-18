package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:255;index:idx_name,unique"`
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=fly_go port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// 迁移 schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("migrate table error")
	}

	// Create
	now := time.Now()
	timeString := fmt.Sprintf("%d", now)
	db.Create(&User{Name: "xiaoqiang" + timeString})

	var user User
	db.First(&user, 1)
	fmt.Println(user)
}
