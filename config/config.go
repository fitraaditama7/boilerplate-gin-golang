package config

import "github.com/jinzhu/gorm"

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "maman:123459@tcp(127.0.0.1:3306)/product?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	return db
}
