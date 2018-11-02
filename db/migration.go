package db

import "first-go/model"

func Migrate() {
	db := InitDb()

	db.AutoMigrate(&model.User{})
}
