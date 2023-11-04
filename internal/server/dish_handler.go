package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllDishesHandler(c *gin.Context) {
	dishes, err := h.services.Dish.GetAllDishes()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var outputDishes []DishOutput
	for _, dish := range dishes {
		dishInfo := DishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}

func (h *Handler) GetDishInfoHandler(c *gin.Context) {
	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishInfo, err := h.services.Dish.GetDishInfo(dishID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dishInfo)
}

func (h *Handler) AddDishHandler(c *gin.Context) {
	dish := new(entity.Dish)

	if err := c.BindJSON(dish); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

	id, err := h.services.Dish.AddDish(dish)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	message := statusID {
		ID: id,
	}

	c.JSON(http.StatusOK, message)
}

func (h *Handler) UpdateDishHandler(c *gin.Context) {
	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dish := new(entity.Dish)
	if err := c.BindJSON(dish); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }
	err = h.services.Dish.UpdateDish(dishID, dish)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Блюдо успешно изменено"))
}

func (h *Handler) DeleteDishHandler(c *gin.Context) {
	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err = h.services.Dish.DeleteDish(dishID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Блюдо успешно удалено"))
}


func (h *Handler) GetDishCousineHandler(c *gin.Context) {
	cousineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCousine(cousineID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var outputDishes []DishOutput
	for _, dish := range dishes {
		dishInfo := DishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}

func (h *Handler) GetDishCategoryHandler(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCategory(categoryID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	var outputDishes []DishOutput
	for _, dish := range dishes {
		dishInfo := DishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}

func (h *Handler) GetDishCousineCategoryHandler(c *gin.Context) {
	cousineID, err := strconv.Atoi(c.Param("cousineID"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCousineCategory(cousineID, categoryID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	var outputDishes []DishOutput
	for _, dish := range dishes {
		dishInfo := DishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}