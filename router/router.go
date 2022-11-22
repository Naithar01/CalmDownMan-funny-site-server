package router

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/controller"
	"github.com/gin-gonic/gin"
)

func InitialApp() *gin.Engine {
	app := gin.Default()

	app.GET("/", controller.HelloWorld)
	app.GET("/test/db", controller.TestInsertDB)

	return app
}
