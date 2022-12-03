package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserJwtCheckMiddleware(c *gin.Context) {
	token, err := c.Request.Cookie("access-jwt-token")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"Success": "false",
		})

		c.Abort()
		return
	}

	token_value := token.Value

	if token_value == "" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"Success": "false",
		})

		c.Abort()
		return
	}

	c.Next()
}
