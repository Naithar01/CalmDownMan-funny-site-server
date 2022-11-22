package controller

import (
	"log"
	"net/http"

	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK,
		"HelloWorld",
	)
}

func TestInsertDB(c *gin.Context) {
	world := "world"
	rs, err := database.Database.Exec("INSERT INTO helloworld(world) VALUES (?)", world)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
