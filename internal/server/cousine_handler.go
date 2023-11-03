package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

func (h *Handler) GetCousineHandler(c *gin.Context) {
	cousines, err := h.services.Cousine.GetCousines()
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
	cousine := new(entity.Cousine)

	if err := c.BindJSON(cousine); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

	id, err := h.services.Cousine.AddCousine(cousine)
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
	
	err = h.services.Cousine.UpdateCousine(cousineID, cousine)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Кухня успешно изменена"))
}

func (h *Handler) DeleteCousineHandler(c *gin.Context) {
	cousineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err = h.services.Cousine.DeleteCousine(cousineID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse("Кухня успешно удалена"))
}