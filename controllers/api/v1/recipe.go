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

func FetchAllRecipes(w http.ResponseWriter, r *http.Request) {
	var service v1s.RecipeService
	start, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil {
		start = 0
	}
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		count = 10
	}

	resp := service.FetchList(count, start)
	u.Respond(w, resp)
}

func FetchItemRecipe(w http.ResponseWriter, r *http.Request) {
	var service v1s.RecipeService
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	resp := service.FetchItem(id)
	u.Respond(w, resp)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var service v1s.RecipeService
	var newRecipe models.RecipeInput

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newRecipe)

	resp := service.CreateItem(newRecipe)
	u.Respond(w, resp)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var service v1s.RecipeService
	var updRecipe models.RecipeInput
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updRecipe)

	resp := service.UpdateItem(updRecipe, id)
	u.Respond(w, resp)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	var service v1s.RecipeService
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	resp := service.DeleteItem(id)
	u.Respond(w, resp)
}
