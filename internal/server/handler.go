package server

import (
	"cookbook/internal/usecase"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "cookbook/docs"
)

type Handler struct {
	services *usecase.Service
	handleTimeout time.Duration
}

func NewHandler(services *usecase.Service, timeout time.Duration) *Handler {
	return &Handler{
		services: services,
		handleTimeout: timeout,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dishes := router.Group("/dish")
	{
		dishes.GET("/", h.GetAllDishesHandler)
		dishes.GET("/:id", h.GetDishInfoHandler)
		dishes.POST("/", h.AddDishHandler)
		dishes.PUT("/:id", h.UpdateDishHandler)
		dishes.DELETE("/:id", h.DeleteDishHandler)

		dishes.GET("/cousine/:cousineID", h.GetDishCousineHandler)
		dishes.GET("/category/:categoryID", h.GetDishCategoryHandler)
		dishes.GET("/cousine/category/:cousineID/:categoryID", h.GetDishCousineCategoryHandler)
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