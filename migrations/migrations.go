package main

import (
	"gin_fleamarket/infra"
	"gin_fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("faild to migrate database")
	}
}
