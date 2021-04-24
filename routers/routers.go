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

	// articles
	api1.HandleFunc("/articles", apiControllerV1.ReturnAllArticles).Methods("GET")
	api1.HandleFunc("/articles/{id}", apiControllerV1.ReturnSingleArticle).Methods("GET")
	api1.HandleFunc("/articles", apiControllerV1.CreateNewArticle).Methods("POST")
	api1.HandleFunc("/articles/{id}", apiControllerV1.UpdateArticle).Methods("PUT")
	api1.HandleFunc("/articles/{id}", apiControllerV1.DeleteArticle).Methods("DELETE")

	// recipes
	api1.HandleFunc("/recipes", apiControllerV1.ReturnAllRecipes).Methods("GET")
	api1.HandleFunc("/recipes/{id}", apiControllerV1.ReturnSingleRecipe).Methods("GET")
	api1.HandleFunc("/recipes", apiControllerV1.CreateNewRecipe).Methods("POST")
	api1.HandleFunc("/recipes/{id}", apiControllerV1.UpdateRecipe).Methods("PUT")
	api1.HandleFunc("/recipes/{id}", apiControllerV1.DeleteRecipe).Methods("DELETE")

	return r
}
