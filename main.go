package main

import (
	"gin_fleamarket/controller"
	"gin_fleamarket/models"
	"gin_fleamarket/repository"
	"gin_fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 500, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 600, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 700, Description: "説明3", SoldOut: false},
	}

	itemRepository := repository.NewItemMemoryRepository(items)
	ItemService := services.NewItemServices(itemRepository)
	ItemController := controller.NewItemController(ItemService)

	r := gin.Default()
	r.GET("/items", ItemController.FindAll)
	r.GET("/items/:id", ItemController.FindById)
	r.POST("/items", ItemController.Create)
	r.PUT("/items/:id", ItemController.Update)
	r.Run(":8082")

}
