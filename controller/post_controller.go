package controller

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity/dto"
	"github.com/Naithar01/CalmDownMan-funny-site-server/service"
)

type PostController interface {
	GetAllPost() ([]entity.Post, error)
	CreatePost(post dto.CreatePostDto, userid int) (int, error)
	UpdatePost(id int, post dto.UpdatePostDto) (int64, error)
	DeletePost(id int) (int64, error)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (p postController) GetAllPost() ([]entity.Post, error) {
	return p.postService.GetAllPost()
}

func (p postController) CreatePost(post dto.CreatePostDto, userid int) (int, error) {
	return p.postService.CreatePost(post, userid)
}

func (p postController) UpdatePost(id int, post dto.UpdatePostDto) (int64, error) {
	return p.postService.UpdatePost(id, post)
}

func (p postController) DeletePost(id int) (int64, error) {
	return p.postService.DeletePost(id)
}
