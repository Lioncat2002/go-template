package main

import (
	"backend/controllers"
	"backend/services"

	"github.com/gin-gonic/gin"
)

var PostRoute *gin.RouterGroup

func RunRouter() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	UserRoute := router.Group("/api/user")
	UserRoute.GET("/", controllers.AllUsers)
	UserRoute.GET("/:id", controllers.GetOneUser)
	UserRoute.POST("/", controllers.AddUser)
	UserRoute.POST("/update", controllers.UpdateUser)
	UserRoute.POST("/buycoins", controllers.UpdateCoins)
	//TODO: make route protected
	ItemRoute := router.Group("/api/item")

	ItemRoute.GET("/", controllers.AllItems)
	ItemRoute.GET("/:id", controllers.GetOneItem)
	ItemRoute.POST("/", controllers.CreateItem)
	ItemRoute.POST("/buy", controllers.BuyItem)

	router.Run(":8080")
}
func main() {
	services.ConnectDatabase()
	RunRouter()
}
