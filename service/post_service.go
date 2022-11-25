package service

import (
	"log"

	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity/dto"
)

type PostService interface {
	GetAllPost() []entity.Post
	CreatePost(dto.CreatePostDto) (int, error)
}

type postService struct {
	posts []entity.Post
}

func NewPostService() PostService {
	return &postService{
		posts: []entity.Post{},
	}
}

func (p postService) GetAllPost() []entity.Post {
	rows, err := database.Database.Query("SELECT * FROM post")

	defer rows.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	var posts []entity.Post

	for rows.Next() {
		var post entity.Post
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category_id, &post.Created_At, &post.Updated_At)
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

	return posts
}

func (p postService) CreatePost(post dto.CreatePostDto) (int, error) {
	c_post, err := database.Database.Exec("INSERT INTO post (title, content, category_id) VALUES (?, ?, ?)", post.Title, post.Content, post.Category_id)

	if err != nil {
		return 0, err
	}

	Id, err := c_post.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(Id), nil
}