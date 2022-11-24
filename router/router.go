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
	// Post
	postService    service.PostService       = service.NewPostService()
	postController controller.PostController = controller.NewPostController(postService)
)

func InitialApp() *gin.Engine {
	app := gin.Default()

	// group routes
	// localhost:4000/api/hello
	hello := app.Group("/api/hello")
	{
		hello.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, helloWorldController.HelloWorld())
		})
		hello.GET("/worlds", func(c *gin.Context) {
			c.JSON(http.StatusOK, helloWorldController.GetAllWorld())
		})
		hello.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, helloWorldController.TestInsertDB())
		})
	}

	return app
}
