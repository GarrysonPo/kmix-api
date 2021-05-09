package main_test

import (
	"bytes"
	"encoding/json"
	"kmix-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAllRecipes(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recipes", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestFetchAllRecipesByLimit(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recipes?limit=3", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestFetchAllRecipesByOffset(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recipes?offset=3", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestFetchItemRecipe(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recipes/1", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestCreateRecipe(t *testing.T) {
	recipeInp := models.RecipeInput{
		Title: "New Title",
		Dish: models.DishInput{
			CookingTime: 40,
			Servings:    1,
			Difficulty:  1,
		},
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(recipeInp)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/recipes", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestUpdateRecipe(t *testing.T) {
	recipeUpd := models.RecipeInput{
		Title: "Updated Title",
		Dish: models.DishInput{
			CookingTime: 40,
			Servings:    1,
			Difficulty:  1,
		},
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(recipeUpd)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/recipes/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}

func TestDeleteRecipe(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/recipes/10", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	CheckResponseCode(t, http.StatusOK, w.Code)
}
