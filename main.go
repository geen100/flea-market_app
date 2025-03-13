package main

import (
	"gin_fleamarket/controller"
	"gin_fleamarket/infra"
	"gin_fleamarket/repository"

	// "gin_fleamarket/models"

	"gin_fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 500, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 600, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 700, Description: "説明3", SoldOut: false},
	// }

	// itemRepository := repository.NewItemMemoryRepository(items)
	itemRepository := repository.NewItemRepository(db)
	ItemService := services.NewItemServices(itemRepository)
	ItemController := controller.NewItemController(ItemService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouter.GET("", ItemController.FindAll)
	itemRouter.GET("/:id", ItemController.FindById)
	itemRouter.POST("", ItemController.Create)
	itemRouter.PUT("/:id", ItemController.Update)
	itemRouter.DELETE("/:id", ItemController.Delete)
	r.Run(":8082")

}
