package server

import (
	"cookbook/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
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
		dishes.GET("/:name", h.GetDishInfoHandler)
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
	dishInfo, err := h.cookInteractor.GetDishInfo(c.Param("name"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dishInfo)
}
