package controller

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/service"
)

type PostController interface {
	GetAllPost() []entity.Post
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (p postController) GetAllPost() []entity.Post {
	return p.postService.GetAllPost()
}
