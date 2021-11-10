package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() (err error) {
	DB, err = gorm.Open("mysql", "root:root@(127.0.0.1:3307)/")
	if err != nil {
		log.Printf("MySQL Connection Error : %s", err.Error())
		return
	}
	err = DB.Exec("CREATE DATABASE IF NOT EXISTS todolist").Error
	if err != nil {
		log.Printf("CREATE DATABASE Error : %s", err.Error())
		return
	}
	dsn := "root:root@(127.0.0.1:3307)/todolist?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("MySQL Error : %s", err.Error())
		return
	}
	DB.AutoMigrate(&Todo{})
	DB.LogMode(true)
	return
}
