package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"context"
)

func (h *Handler) GetCousineHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousines, err := h.services.Cousine.GetCousines(ctx)
	if err != nil {
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

	c.JSON(http.StatusOK, StatusResponse("Кухня успешно изменена"))
}

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

	c.JSON(http.StatusOK, StatusResponse("Кухня успешно удалена"))
}