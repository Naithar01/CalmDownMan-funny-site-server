package controller

import (
	"log"
	"net/http"

	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/models"
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

func GetAllWorld(c *gin.Context) {
	rows, err := database.Database.Query("SELECT id, world FROM helloworld")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// worlds := make([]models.HelloWorld, 0)
	var worlds []models.HelloWorld

	for rows.Next() {
		var world models.HelloWorld
		rows.Scan(&world.Id, &world.World)
		worlds = append(worlds, world)
	}

	if rows.Err(); err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"worlds": worlds,
	})

}
