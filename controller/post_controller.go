package controller

import "github.com/Naithar01/CalmDownMan-funny-site-server/service"

type PostController interface {
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postController{
		postService: postService,
	}
}
