package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"database/sql"
	"net/http"
	"context"
)

// @Summary 	GetCategories
// @Tags 		category
// @Description get all categories
// @ID 			get-categories
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} entity.Category
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /category/ [get]
func (h *Handler) GetCategoryHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	categories, err := h.services.Category.GetCategories(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			emptyArray := []interface{}{}
			c.JSON(http.StatusOK, emptyArray)
			return
		}

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

// @Summary 	AddCategory
// @Tags 		category
// @Description add category
// @ID 			add-category
// @Accept  	json
// @Produce  	json
// @Param 		input body categoryInfo true "category information"
// @Success 	200 {object} statusID
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /category/ [post]
func (h *Handler) AddCategoryHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	category := new(entity.Category)

	if err := c.BindJSON(category); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

	id, err := h.services.Category.AddCategory(ctx, category)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	message := statusID {
		ID: id,
	}

	c.JSON(http.StatusOK, message)
}

// @Summary 	UpdateCategory
// @Tags 		category
// @Description update category by id
// @ID 			update-category
// @Accept  	json
// @Produce  	json
// @Param 		id path int true "categoryID"
// @Param 		input body categoryInfo true "category information"
// @Success 	200 {object} statusResponse
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /category/{id} [put]
func (h *Handler) UpdateCategoryHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

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
	
	err = h.services.Category.UpdateCategory(ctx, categoryID, category)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newStatusResponse("Категория успешно изменена"))
}

// @Summary 	DeleteCategory
// @Tags 		category
// @Description delete category by id
// @ID 			delete-category
// @Accept  	json
// @Produce  	json
// @Param 		id path int true "categoryID"
// @Success 	200 {object} statusResponse
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /category/{id} [delete]
func (h *Handler) DeleteCategoryHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err = h.services.Category.DeleteCategory(ctx, categoryID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newStatusResponse("Категория успешно удалена"))
}