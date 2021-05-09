package main_test

import (
	"bytes"
	"encoding/json"
	"kmix-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAllArticles(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestFetchAllArticlesByLimit(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles?limit=3", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestFetchAllArticlesByOffset(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles?offset=3", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestCreateArticle(t *testing.T) {
	articleInp := models.ArticleInput{
		Title:   "New Title",
		Content: "New Content",
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(articleInp)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestUpdateArticle(t *testing.T) {
	articleUpd := models.ArticleInput{
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(articleUpd)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/articles/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestDeleteArticle(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/articles/10", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}
