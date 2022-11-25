package controller

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity/dto"
	"github.com/Naithar01/CalmDownMan-funny-site-server/service"
)

type PostController interface {
	GetAllPost() []entity.Post
	CreatePost(dto.CreatePostDto) (int, error)
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

func (p postController) CreatePost(post dto.CreatePostDto) (int, error) {
	return p.postService.CreatePost(post)
}