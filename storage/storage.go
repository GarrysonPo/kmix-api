package storage

import "kmix-api/models"

var Articles = []models.Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

var Recipes = []models.Recipe{
	{Id: "1", Title: "Title", Dish: models.Dish{CookingTime: 60, Servings: 4, Difficulty: "Easy"}},
	{Id: "2", Title: "Title 2", Dish: models.Dish{CookingTime: 40, Servings: 2, Difficulty: "Hard"}},
}
