package models

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
