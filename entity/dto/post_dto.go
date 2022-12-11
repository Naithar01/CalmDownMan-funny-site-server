package dto

type CreatePostDto struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category_id int    `json:"category_id"`
}

type UpdatePostDto struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category_id int    `json:"category_id"`
}

type FindPostQuery struct {
	Category string `json:"category"`
	Postid   int    `json:"postid"`
}
