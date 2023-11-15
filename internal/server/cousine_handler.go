package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"database/sql"
	"net/http"
	"context"
)

// @Summary 	GetCousines
// @Tags 		cousine
// @Description get all cousines
// @ID 			get-cousines
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} entity.Cousine
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /cousine/ [get]
func (h *Handler) GetCousineHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousines, err := h.services.Cousine.GetCousines(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			emptyArray := []interface{}{}
			c.JSON(http.StatusOK, emptyArray)
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var outputCousines []entity.Cousine
	for _, cousine := range cousines {
		cousineInfo := entity.Cousine {
			ID: cousine.ID,
			Name: cousine.Name,
			Description: cousine.Description,
		}
		outputCousines = append(outputCousines, cousineInfo)
	}

	c.JSON(http.StatusOK, outputCousines)
}

// @Summary 	AddCousine
// @Tags 		cousine
// @Description add cousine
// @ID 			add-cousine
// @Accept  	json
// @Produce  	json
// @Param 		input body cousineInfo true "cousine information"
// @Success 	200 {object} statusID
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /cousine/ [post]
func (h *Handler) AddCousineHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousine := new(entity.Cousine)

	if err := c.BindJSON(cousine); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

	id, err := h.services.Cousine.AddCousine(ctx, cousine)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	message := statusID {
		ID: id,
	}

	c.JSON(http.StatusOK, message)
}

// @Summary 	UpdateCousine
// @Tags 		cousine
// @Description update cousine by id
// @ID 			update-cousine
// @Accept  	json
// @Produce  	json
// @Param 		id path int true "cousineID"
// @Param 		input body cousineInfo true "cousine information"
// @Success 	200 {object} statusResponse
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /cousine/{id} [put]
func (h *Handler) UpdateCousineHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	cousine := new(entity.Cousine)
	if err := c.BindJSON(cousine); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }
	
	err = h.services.Cousine.UpdateCousine(ctx, cousineID, cousine)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newStatusResponse("Кухня успешно изменена"))
}

// @Summary 	DeleteCousine
// @Tags 		cousine
// @Description delete cousine by id
// @ID 			delete-cousine
// @Accept  	json
// @Produce  	json
// @Param 		id path int true "cousineID"
// @Success 	200 {object} statusResponse
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /cousine/{id} [delete]
func (h *Handler) DeleteCousineHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err = h.services.Cousine.DeleteCousine(ctx, cousineID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newStatusResponse("Кухня успешно удалена"))
}