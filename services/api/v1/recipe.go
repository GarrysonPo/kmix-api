package v1service

import (
	"kmix-api/models"

	u "kmix-api/apiHelpers"
)

type RecipeService struct {
	Recipes []models.Recipe
	Recipe  models.Recipe
}

func (rs *RecipeService) FetchList(count, start int) map[string]interface{} {
	recipes := rs.Recipes
	rows, err := models.GetDB().Query(
		"SELECT recipes.id, title, cooking_time, servings, difficulty.name FROM kmix.recipes "+
			"LEFT JOIN difficulty ON recipes.difficulty = difficulty.id "+
			"LIMIT ? OFFSET ?",
		count,
		start,
	)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var r models.Recipe
		if err := rows.Scan(&r.ID, &r.Title, &r.Dish.CookingTime, &r.Dish.Servings, &r.Dish.Difficulty); err != nil {
			panic(err)
		}
		recipes = append(recipes, r)
	}

	resp := u.Message("success")
	resp["data"] = recipes
	return resp
}

func (rs *RecipeService) FetchItem(id int) map[string]interface{} {
	recipe := rs.Recipe
	err := models.GetDB().QueryRow(
		"SELECT recipes.id, title, cooking_time, servings, difficulty.name FROM kmix.recipes "+
			"LEFT JOIN difficulty ON recipes.difficulty = difficulty.id "+
			"WHERE recipes.id = ?",
		id,
	).Scan(&recipe.ID, &recipe.Title, &recipe.Dish.CookingTime, &recipe.Dish.Servings, &recipe.Dish.Difficulty)
	if err != nil {
		panic(err)
	}

	resp := u.Message("success")
	resp["data"] = recipe
	return resp
}

func (rs *RecipeService) UpdateItem(recipe models.RecipeInput, id int) map[string]interface{} {
	_, err := models.GetDB().Exec(
		"UPDATE recipes SET title = ?, cooking_time = ?, servings = ?, difficulty = ? WHERE (id = ?)",
		recipe.Title, recipe.Dish.CookingTime, recipe.Dish.Servings, recipe.Dish.Difficulty,
		id)
	if err != nil {
		panic(err)
	}

	resp := u.Message("success")
	return resp
}

func (rs *RecipeService) CreateItem(recipe models.RecipeInput) map[string]interface{} {
	res, err := models.GetDB().Exec(
		"INSERT INTO recipes(title, cooking_time, servings, difficulty) VALUES (?, ?, ?, ?); ",
		recipe.Title, recipe.Dish.CookingTime, recipe.Dish.Servings, recipe.Dish.Difficulty,
	)
	if err != nil {
		panic(err)
	}

	id, _ := res.LastInsertId()
	resp := u.Message("success")
	resp["data"] = &models.Recipe{
		ID:    int(id),
		Title: recipe.Title,
		Dish: models.Dish{
			CookingTime: recipe.Dish.CookingTime,
			Servings:    recipe.Dish.CookingTime,
			Difficulty:  "Легко",
		},
	}
	return resp
}

func (rs *RecipeService) DeleteItem(id int) map[string]interface{} {
	_, err := models.GetDB().Exec("DELETE FROM recipes WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	resp := u.Message("success")
	return resp
}
