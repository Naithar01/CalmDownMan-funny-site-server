package router

import (
	"net/http"
	"strconv"

	"github.com/Naithar01/CalmDownMan-funny-site-server/action"
	"github.com/Naithar01/CalmDownMan-funny-site-server/controller"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity/dto"
	"github.com/Naithar01/CalmDownMan-funny-site-server/middleware"
	"github.com/Naithar01/CalmDownMan-funny-site-server/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	helloWorldService    service.HelloWorldService       = service.New()
	helloWorldController controller.HelloWorldController = controller.New(helloWorldService)
	// Post
	postService    service.PostService       = service.NewPostService()
	postController controller.PostController = controller.NewPostController(postService)
	// User
	userService    service.UserService       = service.NewUserService()
	userController controller.UserController = controller.NewUserController(userService)
)

func InitialApp() *gin.Engine {
	app := gin.Default()

	// Cors
	app.Use(cors.Default())

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

	post := app.Group("/api/post")
	// post.Use(middleware.UserJwtCheckMiddleware)
	{
		post.GET("/", func(c *gin.Context) {
			posts, err := postController.GetAllPost()

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			c.JSON(http.StatusOK, map[string][]entity.PostList{
				"datas": posts,
			})
		})
		post.POST("/", middleware.UserJwtCheckMiddleware, func(c *gin.Context) {
			var post dto.CreatePostDto

			user_tk, err := c.Cookie("access-jwt-token")

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			user_id := action.UserJwtTokenParse(user_tk)

			if err := c.BindJSON(&post); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			create_post_id, err := postController.CreatePost(post, user_id)

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}
			c.JSON(http.StatusOK, map[string]int{
				"created_post_id": create_post_id,
			})

		})
		post.PATCH("/:id", middleware.UserJwtCheckMiddleware, func(c *gin.Context) {
			var post dto.UpdatePostDto

			if err := c.BindJSON(&post); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			id := c.Param("id")

			Id, err := strconv.Atoi(id)

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			updated_post_id, err := postController.UpdatePost(Id, post)

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			c.JSON(http.StatusOK, map[string]int{
				"updated_post_id": int(updated_post_id),
			})
		})
		post.DELETE("/:id", middleware.UserJwtCheckMiddleware, func(c *gin.Context) {
			id := c.Param("id")
			Id, err := strconv.Atoi(id)

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			delete_post, err := postController.DeletePost(Id)

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			c.JSON(http.StatusOK, map[string]int{
				"deleted_post": int(delete_post),
			})
		})
	}

	user := app.Group("/api/user")
	{
		user.GET("/", func(c *gin.Context) {
			users, err := userService.GetAllUser()

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			c.JSON(http.StatusOK, map[string][]entity.User{
				"users": users,
			})
		})

		user.POST("/", func(c *gin.Context) {
			var user dto.CreateUserDto

			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			create_user_id, err := userController.CreateUser(user)

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}
			c.JSON(http.StatusOK, map[string]int{
				"created_user_id": create_user_id,
			})

		})

		user.POST("/login", func(c *gin.Context) {
			var user dto.LoginUserDto

			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			user_jwt, err := userController.LoginUser(user)

			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}

			action.UserLoginSaveJwtCookie(c, user_jwt)

			c.JSON(http.StatusOK, map[string]string{
				"jwt": user_jwt,
			})

		})

	}

	return app
}
