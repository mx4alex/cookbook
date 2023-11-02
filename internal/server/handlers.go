package server

import (
	"cookbook/internal/entity"
	"cookbook/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	services *usecase.Service
}

func NewHandler(services *usecase.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	dishes := router.Group("/dish")
	{
		dishes.GET("/", h.GetAllDishesHandler)
		dishes.GET("/:id", h.GetDishInfoHandler)
		dishes.POST("/", h.AddDishHandler)
		dishes.PUT("/:id", h.UpdateDishHandler)
		dishes.DELETE("/:id", h.DeleteDishHandler)
	}

	return router
}

func (h *Handler) GetAllDishesHandler(c *gin.Context) {
	dishes, err := h.services.Dish.GetAllDishes()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var outputDishes []DishOutput
	for _, dish := range dishes {
		outputDish := DishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, outputDish)
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
	err := h.services.Dish.AddDish(dish)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Блюдо успешно добавлено"))
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
