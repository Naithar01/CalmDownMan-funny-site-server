package service

import (
	"log"
	"time"

	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
)

type HelloWorldService interface {
	HelloWorld() string
	GetAllWorld() []GetAllWorldRequest
	TestInsertDB() int64
}

type helloWorldService struct {
	worlds []entity.HelloWorld
}

type GetAllWorldRequest struct {
	Id         int       `json:"id"`
	World      string    `json:"world"`
	Created_At time.Time `json:"created_at"`
}

func CreateGetAllWorldResponse(world entity.HelloWorld) GetAllWorldRequest {
	return GetAllWorldRequest{
		Id: world.Id, World: world.World, Created_At: world.Created_At,
	}
}

func New() HelloWorldService {
	return &helloWorldService{
		worlds: []entity.HelloWorld{},
	}
}

func (s helloWorldService) HelloWorld() string {
	return "Hello World"
}

func (s helloWorldService) GetAllWorld() []GetAllWorldRequest {
	rows, err := database.Database.Query("SELECT id, world, create_at FROM helloworld")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// worlds := make([]entity.HelloWorld, 0)
	var worlds []GetAllWorldRequest

	for rows.Next() {
		var world entity.HelloWorld
		rows.Scan(&world.Id, &world.World, &world.Created_At)
		ResponseWorld := CreateGetAllWorldResponse(world)
		worlds = append(worlds, ResponseWorld)
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
