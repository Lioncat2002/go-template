package main

import (
	"backend/controllers"
	"backend/middlewares"
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
	UserRoute.PUT("/", controllers.AddUser)
	UserRoute.GET("/", controllers.AllUsers)
	UserRoute.POST("/login", controllers.LoginUser)

	UserRoute.Use(middlewares.JwtAuth())
	UserRoute.GET("/one", controllers.GetOneUser)
	UserRoute.PATCH("/update", controllers.UpdateUser)
	UserRoute.DELETE("/delete", controllers.DeleteUser)

	ItemRoute := router.Group("/api/book")
	ItemRoute.GET("/", controllers.AllPosts)
	ItemRoute.GET("/:id", controllers.GetOnePost)
	ItemRoute.GET("/search/:name", controllers.SearchPost)

	ItemRoute.Use(middlewares.JwtAuth())
	ItemRoute.PUT("/", controllers.CreatePost)
	ItemRoute.PUT("/upload/:id", controllers.AddPostUrl)
	router.Run(":8080")
}
func main() {
	services.ConnectDatabase()
	RunRouter()
}
