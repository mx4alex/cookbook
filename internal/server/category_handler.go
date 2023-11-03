package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

func (h *Handler) GetCategoryHandler(c *gin.Context) {
	categories, err := h.services.Category.GetCategories()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var outputCategories []entity.Category
	for _, category := range categories {
		categoryInfo := entity.Category {
			ID: category.ID,
			Name: category.Name,
			Description: category.Description,
		}
		outputCategories = append(outputCategories, categoryInfo)
	}

	c.JSON(http.StatusOK, outputCategories)
}

func (h *Handler) AddCategoryHandler(c *gin.Context) {
	category := new(entity.Category)

	if err := c.BindJSON(category); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

	id, err := h.services.Category.AddCategory(category)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	message := statusID {
		ID: id,
	}

	c.JSON(http.StatusOK, message)
}

func (h *Handler) UpdateCategoryHandler(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	category := new(entity.Category)
	if err := c.BindJSON(category); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }
	
	err = h.services.Category.UpdateCategory(categoryID, category)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Категория успешно изменена"))
}

func (h *Handler) DeleteCategoryHandler(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err = h.services.Category.DeleteCategory(categoryID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Категория успешно удалена"))
}