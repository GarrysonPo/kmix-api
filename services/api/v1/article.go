package v1service

import (
	"kmix-api/models"

	u "kmix-api/apiHelpers"
)

type ArticleService struct {
	Articles []models.Article
	Article  models.Article
}

func (rs *ArticleService) FetchList(limit, offset int) map[string]interface{} {
	articles := rs.Articles
	rows, err := models.GetDB().Query(
		"SELECT id, title, content FROM kmix.articles "+
			"LIMIT ? OFFSET ?",
		limit,
		offset,
	)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var a models.Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Content); err != nil {
			panic(err)
		}
		articles = append(articles, a)
	}

	resp := u.Message("success")
	resp["data"] = articles
	return resp
}

func (rs *ArticleService) FetchItem(id int) map[string]interface{} {
	article := rs.Article
	err := models.GetDB().QueryRow(
		"SELECT id, title, content FROM articles WHERE id = ?",
		id,
	).Scan(&article.ID, &article.Title, &article.Content)
	if err != nil {
		panic(err)
	}

	resp := u.Message("success")
	resp["data"] = article
	return resp
}

func (rs *ArticleService) UpdateItem(article models.ArticleInput, id int) map[string]interface{} {
	_, err := models.GetDB().Exec(
		"UPDATE articles SET title = ?, content = ? WHERE (id = ?)",
		article.Title, article.Content,
		id,
	)
	if err != nil {
		panic(err)
	}

	resp := u.Message("success")
	return resp
}

func (rs *ArticleService) CreateItem(article models.ArticleInput) map[string]interface{} {
	res, err := models.GetDB().Exec(
		"INSERT INTO articles(title, content) VALUES (?, ?);",
		article.Title, article.Content,
	)
	if err != nil {
		panic(err)
	}

	id, _ := res.LastInsertId()
	resp := u.Message("success")
	resp["data"] = &models.Article{
		ID:      int(id),
		Title:   article.Title,
		Content: article.Content,
	}
	return resp
}

func (rs *ArticleService) DeleteItem(id int) map[string]interface{} {
	_, err := models.GetDB().Exec("DELETE FROM articles WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	resp := u.Message("success")
	return resp
}
