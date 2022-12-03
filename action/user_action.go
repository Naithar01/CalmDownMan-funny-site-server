package action

import (
	"time"

	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func UserRegisterHashPassWord(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 3)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func UserLoginCheckPassword(password, userInfo_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userInfo_password), []byte(password))

	if err != nil {
		return false
	}

	return true
}

func UserLoginCreateJwt(userInfo entity.User) (string, error) {
	jwtKey := []byte("JWT_key")

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)

	claims["username"] = userInfo.Username
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tk, err := aToken.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tk, err
}

func UserLoginSaveJwtCookie(c *gin.Context, jwt_token string) {
	c.SetCookie("access-jwt-token", jwt_token, 3600, "/", "localhost", false, true)
}
