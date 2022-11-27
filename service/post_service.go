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
	UpdatePost(id int, post dto.UpdatePostDto) (int64, error)
	DeletePost(id int) (int64, error)
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
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category_id, &post.Author_id, &post.Created_At, &post.Updated_At)
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

	return posts
}

func (p postService) CreatePost(post dto.CreatePostDto) (int, error) {
	c_post, err := database.Database.Exec("INSERT INTO post (title, content, category_id, author_id) VALUES (?, ?, ?, ?)", post.Title, post.Content, post.Category_id, post.Author_id)

	if err != nil {
		return 0, err
	}

	Id, err := c_post.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(Id), nil
}

func (p postService) UpdatePost(id int, post dto.UpdatePostDto) (int64, error) {
	if len(post.Title) != 0 {
		_, err := database.Database.Exec("UPDATE post SET title = ? WHERE id = ?", post.Title, id)
		if err != nil {
			return 0, err
		}
	}

	if len(post.Content) != 0 {
		_, err := database.Database.Exec("UPDATE post SET content = ? WHERE id = ?", post.Content, id)
		if err != nil {
			return 0, err
		}
	}

	if post.Category_id != 0 {
		_, err := database.Database.Exec("UPDATE post SET category_id = ? WHERE id = ?", post.Category_id, id)
		if err != nil {
			return 0, err
		}
	}

	return int64(id), nil
}

func (p postService) DeletePost(id int) (int64, error) {
	rs, err := database.Database.Exec("DELETE FROM post WHERE id = ?", id)

	if err != nil {
		return 0, err
	}

	return rs.RowsAffected()
}
