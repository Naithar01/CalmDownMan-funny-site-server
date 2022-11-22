package router

import (
	"net/http"

	"github.com/Naithar01/CalmDownMan-funny-site-server/controller"
	"github.com/Naithar01/CalmDownMan-funny-site-server/service"
	"github.com/gin-gonic/gin"
)

var (
	helloWorldService    service.HelloWorldService       = service.New()
	helloWorldController controller.HelloWorldController = controller.New(helloWorldService)
)

func InitialApp() *gin.Engine {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helloWorldController.HelloWorld())
	})
	app.GET("/worlds", func(c *gin.Context) {
		c.JSON(http.StatusOK, helloWorldController.GetAllWorld())
	})
	app.POST("/test/db", func(c *gin.Context) {
		c.JSON(http.StatusOK, helloWorldController.TestInsertDB())
	})
	return app
}
