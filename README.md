# Kmix RESTful API

This project has a simple RESTful API for my personal blog about cooking and programming.

## Table of Contents

1. [Run the project](#run-the-project)
2. [Routes](#routes)

## Run the project

To run the project enter the following command in the terminal:

```bash
> go run main.go
```

For testing, enter the following command in the terminal:

```bash
> cd /tests
> go test
```

## Routes

| # | Route | Method | Type | Description | More |
|-|-|-|-|-|-|
| **Recipes** |
| 1 | /recipes | GET | JSON | Get list of recipes | [More](#1-get-list-of-recipes) |
| 2 | /recipes/`{id}` | GET | JSON | Get a recipe | [More](#2-get-a-recipe) |
| 3 | /recipes | POST | JSON | Create a recipe | [More](#3-create-a-recipe) |
| 4 | /recipes/`{id}` | PUT | JSON | Update a recipe | [More](#4-update-a-recipe) |
| 5 | /recipes/`{id}` | DELETE | JSON | Delete a recipe | [More](#5-delete-a-recipe) |
| **Articles** |
| 6 | /articles | GET | JSON | Get list of articles | [More](#6-get-list-of-articles) |
| 7 | /articles/`{id}` | GET | JSON | Get a article | [More](#7-get-a-article) |
| 8 | /articles | POST | JSON | Create a article | [More](#8-create-a-article) |
| 9 | /articles/`{id}` | PUT | JSON | Update a article | [More](#9-update-a-articles) |
| 10 | /articles/`{id}` | DELETE | JSON | Delete a article | [More](#10-delete-a-article) |


### 1. Get list of recipes

Request

```http
GET {host}/api/v1/recipes
```

Addition params

- offset
- limit

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
    "Title": "Title 3",
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
  "Title": "Title 3",
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
  "Title": "The Title has been updated",
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
  "Id": 1,
  "Title": "The Title has been updated",
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

### 6. Get list of articles

Request

```http
GET {host}/api/v1/articles
```

Addition params

- offset
- limit

Response

```json
[
  {
    "id": 1,
    "title": "Title",
    "content": "Content"
  },
  {
    "id": 2,
    "title": "Title 2",
    "content": "Content 2"
  }
]
```

### 7. Get a article

Request

```http
GET {host}/api/v1/articles/2
```

Response

```json
{
  "id": 2,
  "title": "Title 2",
  "content": "Content 2"
}
```

### 8. Create a article

Request

```http
POST {host}/api/v1/articles
```

Input data

```json
{
  "Title": "Title 3",
  "Content": "Content 3"
}
```

Response

```json
{
  "Id": 3,
  "Title": "Title 3",
  "Content": "Content 3"
}
```

### 9. Update a articles

Request

```http
PUT {host}/api/v1/articles/1
```

Input data

```json
{
  "Title": "The Title has been updated",
  "Content": "Content"
}
```

Response

```json
{
  "Id": 1,
  "Title": "The Title has been updated",
  "Content": "Content"
}
```

### 10. Delete a article

Request

```http
DELETE {host}/api/v1/articles/1
```

Response

```json
{
  "status": "success"
}
```