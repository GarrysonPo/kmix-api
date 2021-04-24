package v1

import (
	"encoding/json"
	"io/ioutil"
	"kmix-api/models"
	"kmix-api/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range storage.Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	storage.Articles = append(storage.Articles, article)
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updArticle models.Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updArticle)

	for i, article := range storage.Articles {
		if article.Id == id {
			article = models.Article{
				Id:      updArticle.Id,
				Title:   updArticle.Title,
				Desc:    updArticle.Desc,
				Content: updArticle.Content,
			}
			storage.Articles[i] = updArticle
			json.NewEncoder(w).Encode(storage.Articles[i])
		}
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, article := range storage.Articles {
		if article.Id == id {
			storage.Articles = append(storage.Articles[:i], storage.Articles[i+1:]...)
		}
	}
}
