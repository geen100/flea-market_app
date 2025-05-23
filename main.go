package main

import (
	"gin_fleamarket/controller"
	"gin_fleamarket/infra"
	"gin_fleamarket/middlewares"
	"gin_fleamarket/repository"

	// "gin_fleamarket/models"

	"gin_fleamarket/services"

	"github.com/gin-contrib/cors"
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

	authRepository := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("", ItemController.FindAll)
	itemRouterWithAuth.GET("/:id", ItemController.FindById)
	itemRouterWithAuth.POST("", ItemController.Create)
	itemRouterWithAuth.PUT("/:id", ItemController.Update)
	itemRouterWithAuth.DELETE("/:id", ItemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	r.Run(":8082")

}
