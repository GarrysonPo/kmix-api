package v1

import (
	"encoding/json"
	"io/ioutil"
	"kmix-api/models"
	"kmix-api/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func ReturnAllRecipes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Recipes)
}

func ReturnSingleRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, recipe := range storage.Recipes {
		if recipe.Id == id {
			json.NewEncoder(w).Encode(recipe)
		}
	}
}

func CreateNewRecipe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var recipe models.Recipe
	json.Unmarshal(reqBody, &recipe)
	storage.Recipes = append(storage.Recipes, recipe)
	json.NewEncoder(w).Encode(recipe)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updRecipe models.Recipe
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updRecipe)

	for ind, recipe := range storage.Recipes {
		if recipe.Id == id {
			recipe = models.Recipe{
				Id:    updRecipe.Id,
				Title: updRecipe.Title,
				Dish:  updRecipe.Dish,
			}
			storage.Recipes[ind] = recipe
			json.NewEncoder(w).Encode(storage.Recipes[ind])
		}
	}
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, recipe := range storage.Recipes {
		if recipe.Id == id {
			storage.Recipes = append(storage.Recipes[:i], storage.Recipes[i+1:]...)
		}
	}
}
