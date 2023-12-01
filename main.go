package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhphan97/jwt-authentication-golang/controllers"
	"github.com/thinhphan97/jwt-authentication-golang/database"
	"github.com/thinhphan97/jwt-authentication-golang/middlewares"
)

func main() {
	database.Connect("root:root@tcp(127.0.0.1:3306)/jwt-test?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
