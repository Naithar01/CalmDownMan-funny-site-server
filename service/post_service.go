package service

import "github.com/Naithar01/CalmDownMan-funny-site-server/entity"

type PostService interface {
}

type postService struct {
	posts []entity.Post
}

func NewPostService() PostService {
	return &postService{
		posts: []entity.Post{},
	}
}
