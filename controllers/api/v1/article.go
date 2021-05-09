package v1

import (
	"encoding/json"
	"io/ioutil"
	"kmix-api/models"
	"net/http"
	"strconv"

	u "kmix-api/apiHelpers"
	v1s "kmix-api/services/api/v1"

	"github.com/gorilla/mux"
)

func FetchAllArticles(w http.ResponseWriter, r *http.Request) {
	var service v1s.ArticleService
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	resp := service.FetchList(limit, offset)
	u.Respond(w, resp)
}

func FetchItemArticle(w http.ResponseWriter, r *http.Request) {
	var service v1s.ArticleService
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	resp := service.FetchItem(id)
	u.Respond(w, resp)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var service v1s.ArticleService
	var newArticle models.ArticleInput

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newArticle)

	resp := service.CreateItem(newArticle)
	u.Respond(w, resp)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var service v1s.ArticleService
	var updArticle models.ArticleInput
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updArticle)

	resp := service.UpdateItem(updArticle, id)
	u.Respond(w, resp)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var service v1s.ArticleService
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	resp := service.DeleteItem(id)
	u.Respond(w, resp)
}
