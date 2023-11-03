# cookbook
Кулинарная книга

## Реализованные запросы

### Блюда
- `GET /dish/` возвращает все блюда кулинарной книги
- `GET /dish/{id}` возвращает всю информацию по выбранному блюду (название, описание, рецепт, время приготовления, ингредиенты)
- `POST /dish/` добавляет новое блюдо в кулинарную книгу
- `PUT /dish/{id}` изменяет информацию о блюде
- `DELETE /dish/{id}` удаляет указанное блюдо

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
    "name": "Спагетти Болоньезе",
    "description": "Традиционное итальянское блюдо.",
    "recipe": "Обжарьте мелко нарезанный лук и чеснок, добавьте мясной фарш, тушите с томатами. Подавайте с отварными спагетти.",
    "time": 30,
    "ingredients": [
        {
            "name": "Лук",
            "measure_unit": "г",
            "quantity": 50
        },
        {
            "name": "Свинина",
            "measure_unit": "г",
            "quantity": 300
        },
        {
            "name": "Помидоры",
            "measure_unit": "шт",
            "quantity": 200
        },
        {
            "name": "Спагетти",
            "measure_unit": "г",
            "quantity": 200
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
            "quantity": 50
        },
        {
            "name": "Ингредиент 2",
            "measure_unit": "г",
            "quantity": 300
        },
        {
            "name": "Ингредиент 3",
            "measure_unit": "шт",
            "quantity": 200
        }
    ]
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
            "quantity": 50
        },
        {
            "name": "Ингредиент 2",
            "measure_unit": "г",
            "quantity": 300
        },
        {
            "name": "Ингредиент 3",
            "measure_unit": "шт",
            "quantity": 200
        }
    ]
}
```

#### DeleteDish
* Метод: `DELETE`
* Эндпоинт: `http://localhost:8080/dish/{id}`


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

