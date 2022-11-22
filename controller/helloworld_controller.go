package controller

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/service"
)

type HelloWorldController interface {
	HelloWorld() string
	GetAllWorld() []service.GetAllWorldRequest
	TestInsertDB() int64
}

type helloWorldController struct {
	helloWorldService service.HelloWorldService
}

func New(helloworldService service.HelloWorldService) HelloWorldController {
	return &helloWorldController{
		helloWorldService: helloworldService,
	}
}

func (c helloWorldController) HelloWorld() string {
	return c.helloWorldService.HelloWorld()
}

func (c helloWorldController) GetAllWorld() []service.GetAllWorldRequest {
	return c.helloWorldService.GetAllWorld()
}

func (c helloWorldController) TestInsertDB() int64 {
	return c.helloWorldService.TestInsertDB()
}
