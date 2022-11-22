package service

import (
	"log"

	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
)

type HelloWorldService interface {
	HelloWorld() string
	GetAllWorld() []entity.HelloWorld
	TestInsertDB() int64
}

type helloWorldService struct {
	worlds []entity.HelloWorld
}

func New() HelloWorldService {
	return &helloWorldService{
		worlds: []entity.HelloWorld{},
	}
}

func (s helloWorldService) HelloWorld() string {
	return "Hello World"
}

func (s helloWorldService) GetAllWorld() []entity.HelloWorld {
	rows, err := database.Database.Query("SELECT id, world FROM helloworld")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// worlds := make([]entity.HelloWorld, 0)
	var worlds []entity.HelloWorld

	for rows.Next() {
		var world entity.HelloWorld
		rows.Scan(&world.Id, &world.World)
		worlds = append(worlds, world)
	}

	if rows.Err(); err != nil {
		log.Fatalln(err)
	}

	return worlds

}

func (s helloWorldService) TestInsertDB() int64 {
	world := "world"
	rs, err := database.Database.Exec("INSERT INTO helloworld(world) VALUES (?)", world)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return id
}
