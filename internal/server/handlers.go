package server

import (
	"cookbook/internal/entity"
	"cookbook/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	cookInteractor *usecase.CookInteractor
}

func NewHandler(cookInteractor *usecase.CookInteractor) *Handler {
	return &Handler{
		cookInteractor: cookInteractor,
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
	dishes, err := h.cookInteractor.GetAllDishes()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dishes)
}

func (h *Handler) GetDishInfoHandler(c *gin.Context) {
	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishInfo, err := h.cookInteractor.GetDishInfo(dishID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dishInfo)
}

func (h *Handler) AddDishHandler(c *gin.Context) {
	dish := new(entity.DishInput)

	if err := c.BindJSON(dish); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }
	err := h.cookInteractor.AddDish(dish)
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

	dish := new(entity.DishInput)
	if err := c.BindJSON(dish); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }
	err = h.cookInteractor.UpdateDish(dishID, dish)
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
	
	err = h.cookInteractor.DeleteDish(dishID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Блюдо успешно удалено"))
}
