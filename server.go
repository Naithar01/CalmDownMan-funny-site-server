package main

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/router"
)

func main() {
	app := router.InitialApp()

	database.ConnectDB()

	app.Run() // localhost 8080
}
