package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kmix-api/models"
	"kmix-api/routers"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router *mux.Router

func TestMain(m *testing.M) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"),
	)
	models.InitDB(dataSourceName)
	defer models.CloseDB()

	router = routers.SetupRouter()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestFetchAllRecipes(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/api/v1/recipes", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	checkResponseCode(t, http.StatusOK, w.Code)
}

func TestFetchItemRecipe(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/recipes/1", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	checkResponseCode(t, http.StatusOK, w.Code)
}

func TestCreateRecipe(t *testing.T) {
	recipeInp := models.RecipeInput{
		Title: "Newly Created Recipe",
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
	checkResponseCode(t, http.StatusOK, w.Code)
}

func TestUpdateRecipe(t *testing.T) {
	recipeUpd := models.RecipeInput{
		Title: "Title new",
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
	checkResponseCode(t, http.StatusOK, w.Code)
}

func TestDeleteRecipe(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/recipes/10", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	checkResponseCode(t, http.StatusOK, w.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
