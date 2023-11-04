package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"context"
)

func (h *Handler) GetAllDishesHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dishes, err := h.services.Dish.GetAllDishes(ctx)
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
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishInfo, err := h.services.Dish.GetDishInfo(ctx, dishID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dishInfo)
}

func (h *Handler) AddDishHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dish := new(entity.Dish)

	if err := c.BindJSON(dish); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

	id, err := h.services.Dish.AddDish(ctx, dish)
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
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

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
	err = h.services.Dish.UpdateDish(ctx, dishID, dish)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Блюдо успешно изменено"))
}

func (h *Handler) DeleteDishHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err = h.services.Dish.DeleteDish(ctx, dishID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Блюдо успешно удалено"))
}


func (h *Handler) GetDishCousineHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCousine(ctx, cousineID)
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
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCategory(ctx, categoryID)
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
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

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

	dishes, err := h.services.Dish.GetDishCousineCategory(ctx, cousineID, categoryID)
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