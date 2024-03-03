package model

import (
	postgre "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=goapi port=5432 sslmode=disable TimeZone=Europe/Bucharest"
	db, err := gorm.Open(postgre.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	return db
}
