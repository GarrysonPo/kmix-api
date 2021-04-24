package models

type Recipe struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Dish  Dish   `json:"dish"`
}

type Dish struct {
	CookingTime int    `json:"cookingTime"`
	Servings    int    `json:"servings"`
	Difficulty  string `json:"difficulty"`
}
