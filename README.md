# cookbook
Кулинарная книга

## Реализованные запросы
- `GET /dish/` возвращает все блюда кулинарной книги
- `GET /dish/{name}` возвращает всю информацию по выбранному блюду (название, описание, рецепт, время приготовления, ингредиенты)
- `POST /dish/` добавляет новое блюдо в кулинарную книгу

#### GetAllDishes
* Метод: `GET`
* Эндпоинт: `http://localhost:8080/dish/`
* Формат ответа:
```json
  {
      "name": "Спагетти Болоньезе",
      "description": "Традиционное итальянское блюдо.",
      "time": 30
  },
  {
      "name": "Фруктовый салат",
      "description": "Свежие фрукты в одном блюде.",
      "time": 15
  }
```
#### GetDishInfo
* Метод: `GET`
* Эндпоинт: `http://localhost:8080/dish/{name}`
* Формат ответа:
```json
{
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
