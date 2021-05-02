# Kmix-api

This project has a simple api for my personal blog about cooking and programming.

## Routes

| # | Route | Method | Type | Description | More |
|-|-|-|-|-|-|
| 1 | /recipes | GET | JSON | Get list of recipes | [More](#1-get-list-of-recipes) |
| 2 | /recipes/`{id}` | GET | JSON | Get a recipe | [More](#2-get-a-recipe) |
| 3 | /recipes | POST | JSON | Create a recipe | [More](#3-create-a-recipe) |
| 4 | /recipes/`{id}` | PUT | JSON | Update a recipe | [More](#4-update-a-recipe) |
| 5 | /recipes/`{id}` | DELETE | JSON | Delete a recipe | [More](#5-delete-a-recipe) |

### 1. Get list of recipes

Request

```http
GET {host}/api/v1/recipes
```
Response

```json
[
  {
    "id": 1,
    "title": "Title",
    "dish": {
      "cookingTime": 60,
      "serving": 4,
      "difficulty": "Easy"
    }
  },
  {
    "id": 2,
    "title": "Title 2",
    "dish": {
      "cookingTime": 40,
      "serving": 2,
      "difficulty": "Hard"
    }
  }
]
```

### 2. Get a recipe

Request

```http
GET {host}/api/v1/recipes/2
```

Response

```json
{
  "id": 2,
  "title": "Title 2",
  "dish": {
    "cookingTime": 40,
    "serving": 2,
    "difficulty": "Hard"
  }
}
```

### 3. Create a recipe

Request

```http
POST {host}/api/v1/recipe
```

Input data

```json
{
    "Title": "Newly Created Recipe",
    "Dish": {
      "CookingTime": 10,
      "Servings": 1,
      "Difficulty": 1
    }
}
```

Response

```json
{
  "Id": 3,
  "Title": "Newly Created Recipe",
  "Dish": {
    "CookingTime": 10,
    "Servings": 1,
    "Difficulty": "Easy"
  }
}
```

### 4. Update a recipe

Request

```http
PUT {host}/api/v1/recipe/1
```

Input data

```json
{
  "Title": "Newly Created Recipe",
  "Dish": {
    "CookingTime": 10,
    "Servings": 1,
    "Difficulty": 1
  }
}
```

Response

```json
{
  "Id": "7",
  "Title": "Updated Recipe",
  "Dish": {
    "CookingTime": 10,
    "Servings": 1,
    "Difficulty": "Easy"
  }
}
```

### 5. Delete a recipe

Request

```http
DELETE {host}/api/v1/recipe/1
```

Response

```json
{
  "status": "success"
}
```