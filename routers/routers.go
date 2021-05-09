package routers

import (
	apiControllerV1 "kmix-api/controllers/api/v1"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	api := r.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	api1 := api.PathPrefix("/v1").Subrouter()

	api1.HandleFunc("/articles", apiControllerV1.FetchAllArticles).Methods(http.MethodGet)
	api1.HandleFunc("/articles/{id}", apiControllerV1.FetchItemArticle).Methods(http.MethodGet)
	api1.HandleFunc("/articles", apiControllerV1.CreateArticle).Methods(http.MethodPost)
	api1.HandleFunc("/articles/{id}", apiControllerV1.UpdateArticle).Methods(http.MethodPut)
	api1.HandleFunc("/articles/{id}", apiControllerV1.DeleteArticle).Methods(http.MethodDelete)

	api1.HandleFunc("/recipes", apiControllerV1.FetchAllRecipes).Methods(http.MethodGet)
	api1.HandleFunc("/recipes/{id}", apiControllerV1.FetchItemRecipe).Methods(http.MethodGet)
	api1.HandleFunc("/recipes", apiControllerV1.CreateRecipe).Methods(http.MethodPost)
	api1.HandleFunc("/recipes/{id}", apiControllerV1.UpdateRecipe).Methods(http.MethodPut)
	api1.HandleFunc("/recipes/{id}", apiControllerV1.DeleteRecipe).Methods(http.MethodDelete)

	return r
}
