package service

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity/dto"
)

type PostService interface {
	GetAllPost() ([]entity.PostList, error)
	CreatePost(post dto.CreatePostDto, userid int) (int, error)
	UpdatePost(id int, post dto.UpdatePostDto) (int64, error)
	DeletePost(id int) (int64, error)
	FindPost(category, postid string) (entity.PostList, error)
}

type postService struct {
	posts []entity.Post
}

func NewPostService() PostService {
	return &postService{
		posts: []entity.Post{},
	}
}

func (p postService) GetAllPost() ([]entity.PostList, error) {
	rows, err := database.Database.Query("SELECT * FROM post")

	defer rows.Close()

	if err != nil {
		return []entity.PostList{}, err
	}

	var posts []entity.PostList

	for rows.Next() {
		// Inner Join 코드로 변경해야함 22.12.11
		var post entity.PostList
		var check_post entity.Post

		var category entity.PostList_Category
		var author entity.PostList_Author

		rows.Scan(&check_post.Id, &check_post.Title, &check_post.Content, &check_post.Category_id, &check_post.Author_id, &check_post.Created_At, &check_post.Updated_At)

		category.GetCategoryInfo(check_post.Category_id)
		author.GetCategoryInfo(check_post.Author_id)

		post.Id = check_post.Id
		post.Title = check_post.Title
		post.Content = check_post.Content
		post.Category = category
		post.Author = author
		post.Created_At = check_post.Created_At
		post.Updated_At = check_post.Updated_At

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return []entity.PostList{}, err
	}

	return posts, nil

}

func (p postService) CreatePost(post dto.CreatePostDto, userid int) (int, error) {
	// 22 - 12 - 3 token payload 로 authord id 저장
	c_post, err := database.Database.Exec("INSERT INTO post (title, content, category_id, author_id) VALUES (?, ?, ?, ?)", post.Title, post.Content, post.Category_id, userid) //  post.Author_id

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
func (p postService) FindPost(category, postid string) (entity.PostList, error) {
	var findPost entity.PostList
	err := database.Database.QueryRow("SELECT p.id, p.title, p.content,p.created_at, p.updated_at ,c.id ,c.title, u.id, u.username FROM post AS p INNER JOIN category AS c ON p.category_id = c.id  INNER JOIN user AS u ON p.author_id = u.id  WHERE p.id = ? AND c.title = ? ", postid, category).Scan(&findPost.Id, &findPost.Title, &findPost.Content, &findPost.Created_At, &findPost.Updated_At, &findPost.Category.Id, &findPost.Category.Title, &findPost.Author.Id, &findPost.Author.Username)

	if err != nil {
		return entity.PostList{}, err
	}

	return findPost, nil
}
