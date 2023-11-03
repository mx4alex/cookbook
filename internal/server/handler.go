package server

import (
	"cookbook/internal/usecase"
	"github.com/gin-gonic/gin"
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

	cousine := router.Group("/cousine")
	{
		cousine.GET("/", h.GetCousineHandler)
		cousine.POST("/", h.AddCousineHandler)
		cousine.PUT("/:id", h.UpdateCousineHandler)
		cousine.DELETE("/:id", h.DeleteCousineHandler)
	}

	category := router.Group("/category")
	{
		category.GET("/", h.GetCategoryHandler)
		category.POST("/", h.AddCategoryHandler)
		category.PUT("/:id", h.UpdateCategoryHandler)
		category.DELETE("/:id", h.DeleteCategoryHandler)
	}

	return router
}