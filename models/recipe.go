package models

type Recipe struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Dish  Dish   `json:"dish"`
}

type Dish struct {
	CookingTime int    `json:"cookingTime"`
	Servings    int    `json:"servings"`
	Difficulty  string `json:"difficulty"`
}

type RecipeInput struct {
	Title string    `json:"title"`
	Dish  DishInput `json:"dish"`
}

type DishInput struct {
	CookingTime int `json:"cookingTime"`
	Servings    int `json:"servings"`
	Difficulty  int `json:"difficulty"`
}
