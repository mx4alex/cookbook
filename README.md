# cookbook
Кулинарная книга

## Реализованные запросы

### Блюда
- `GET /dish/` возвращает все блюда кулинарной книги
- `GET /dish/{id}` возвращает всю информацию по выбранному блюду (название, описание, рецепт, время приготовления, ингредиенты)
- `POST /dish/` добавляет новое блюдо в кулинарную книгу
- `PUT /dish/{id}` изменяет информацию о блюде
- `DELETE /dish/{id}` удаляет указанное блюдо
- `GET /dish/cousine/{cousineID}` возвращает блюда из указанной кухни
- `GET /dish/category/{categoryID}` возвращает блюда из указанной категории
- `GET /dish/cousine/category/{cousineID}/{categoryID}` возвращает блюда из указанной кухни и категории
- `GET /dish/search/{text}` возвращает блюда по их названию или ингредиентам 

### Кухня
- `GET /cousine/` возвращает все кухни кулинарной книги
- `POST /cousine/` добавляет новую кухню в кулинарную книгу
- `PUT /cousine/{id}` изменяет информацию о кухне
- `DELETE /cousine/{id}` удаляет указанную кухню

### Категория
- `GET /category/` возвращает все категории кулинарной книги
- `POST /category/` добавляет новую категорию в кулинарную книгу
- `PUT /category/{id}` изменяет информацию о категории
- `DELETE /category/{id}` удаляет указанную категорию

#### GetAllDishes
* Метод: `GET`
* Эндпоинт: `http://localhost:8080/dish/`
* Формат ответа:
```json
  {
      "id": 1,
      "name": "Спагетти Болоньезе",
      "description": "Традиционное итальянское блюдо.",
      "time": 30
  },
  {
      "id": 2,
      "name": "Фруктовый салат",
      "description": "Свежие фрукты в одном блюде.",
      "time": 15
  }
```
#### GetDishInfo
* Метод: `GET`
* Эндпоинт: `http://localhost:8080/dish/{id}`
* Формат ответа:
```json
{
    "id": 1,
    "name": "Фруктовый салат",
    "category_id": 6,
    "cousine_id": 4,
    "description": "Свежие фрукты в одном блюде.",
    "recipe": "Нарежьте разнообразные фрукты (ягоды, яблоки, бананы), подавайте с соусом из меда и мятой.",
    "time": 15,
    "dish_protein": 5,
    "dish_fats": 0,
    "dish_carbohydrates": 68,
    "kilocalories": 292,
    "ingredients": [
        {
            "name": "Бананы",
            "measure_unit": "г",
            "quantity": 200,
            "ingredient_protein": 2,
            "ingredient_fats": 0,
            "ingredient_carbohydrates": 23
        },
        {
            "name": "Клубника",
            "measure_unit": "г",
            "quantity": 100,
            "ingredient_protein": 1,
            "ingredient_fats": 0,
            "ingredient_carbohydrates": 7
        },
        {
            "name": "Мед",
            "measure_unit": "г",
            "quantity": 100,
            "ingredient_protein": 0,
            "ingredient_fats": 0,
            "ingredient_carbohydrates": 15
        }
    ]
}
```

#### AddDish
* Метод: `POST`
* Эндпоинт: `http://localhost:8080/dish/`
* Формат запроса:
```json
{
    "name": "Новое блюдо",
    "category_id": 1,
    "cousine_id": 1,
    "description": "Описание нового блюда",
    "recipe": "Рецепт нового блюда",
    "time": 30,
    "ingredients": [
        {
            "name": "Ингредиент 1",
            "measure_unit": "г",
            "quantity": 50,
            "ingredient_protein": 2,
            "ingredient_fats": 0,
            "ingredient_carbohydrates": 3
        },
        {
            "name": "Ингредиент 2",
            "measure_unit": "г",
            "quantity": 300,
            "ingredient_protein": 1,
            "ingredient_fats": 2,
            "ingredient_carbohydrates": 3
        },
        {
            "name": "Ингредиент 3",
            "measure_unit": "шт",
            "quantity": 200,
            "ingredient_protein": 3,
            "ingredient_fats": 2,
            "ingredient_carbohydrates": 1
        }
    ]
}
```
* Формат ответа:
```json
{
    "id": 3
}
```

#### UpdateDish
* Метод: `PUT`
* Эндпоинт: `http://localhost:8080/dish/{id}`
* Формат запроса:
```json
{
    "name": "Новое название блюда",
    "category_id": 1,
    "cousine_id": 1,
    "description": "Новое описание блюда",
    "recipe": "Новый рецепт блюда",
    "time": 30,
    "ingredients": [
        {
            "name": "Ингредиент 1",
            "measure_unit": "г",
            "quantity": 50,
            "ingredient_protein": 2,
            "ingredient_fats": 0,
            "ingredient_carbohydrates": 3
        },
        {
            "name": "Ингредиент 2",
            "measure_unit": "г",
            "quantity": 300,
            "ingredient_protein": 1,
            "ingredient_fats": 2,
            "ingredient_carbohydrates": 3
        },
        {
            "name": "Ингредиент 3",
            "measure_unit": "шт",
            "quantity": 200,
            "ingredient_protein": 3,
            "ingredient_fats": 2,
            "ingredient_carbohydrates": 1
        }
    ]
}
```

#### DeleteDish
* Метод: `DELETE`
* Эндпоинт: `http://localhost:8080/dish/{id}`

#### DishSearch
* Метод: `GET`
* Эндпоинт: `http://localhost:8080/dish/search/{text}`
* Формат ответа:
```json
[
    {
        "id": 13,
        "name": "Мохито",
        "description": "Классический коктейль.",
        "time": 10
    },
    {
        "id": 7,
        "name": "Тирамису",
        "description": "Итальянский десерт с кофе.",
        "time": 60
    }
]
```

#### GetCategory
* Метод: `GET`
* Энпоинт: `http://localhost:8080/category/`
* Формат ответа:
```json
[
    {
        "id": 1,
        "name": "Напитки",
        "description": "Широкий выбор освежающих напитков."
    },
    {
        "id": 2,
        "name": "Завтраки",
        "description": "Начните день с питательных завтраков."
    }
]
```

#### AddCategory
* Метод: `POST`
* Энпоинт: `http://localhost:8080/category/`
* Формат запроса:
```json
{
    "name": "Новая категория",
    "description": "Описание новой категории."
}
```
* Формат ответа:
```json
{
    "id": 3
}
```

#### UpdateCategory
* Метод: `PUT`
* Энпоинт: `http://localhost:8080/category/{id}`
* Формат запроса:
```json
{
    "name": "Новое название категории",
    "description": "Новое описание категории."
}
```

#### DeleteCategory
* Метод: `PUT`
* Энпоинт: `http://localhost:8080/category/{id}`

#### Примечание
* Формат запросов с кухней аналогичен с запросами категорий.
* Формат запросов блюд с указанием кухни и категории аналогичен с запросом GetAllDishes

