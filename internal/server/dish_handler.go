package server

import (
	"cookbook/internal/entity"
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	"strconv"
	"context"
)

// @Summary 	GetAllDishes
// @Tags 		dish
// @Description get all dishes
// @ID 			get-dishes
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} dishOutput
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/ [get]
func (h *Handler) GetAllDishesHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dishes, err := h.services.Dish.GetAllDishes(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			emptyArray := []interface{}{}
			c.JSON(http.StatusOK, emptyArray)
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var outputDishes []dishOutput
	for _, dish := range dishes {
		dishInfo := dishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}


// @Summary 	GetDishInfo
// @Tags 		dish
// @Description get dish information by id
// @ID 			get-dish-info
// @Accept  	json
// @Produce  	json
// @Param 		id path int true "dishID"
// @Success 	200 {object} entity.Dish
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/{id} [get]
func (h *Handler) GetDishInfoHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishInfo, err := h.services.Dish.GetDishInfo(ctx, dishID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.String(http.StatusNotFound, "Блюдо не найдено")
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dishInfo)
}

// @Summary 	AddDish
// @Tags 		dish
// @Description add dish to cookbook
// @ID 			add-dish
// @Accept  	json
// @Produce  	json
// @Param 		input body entity.Dish true "dish information"
// @Success 	200 {object} statusID
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/ [post]
func (h *Handler) AddDishHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dish := new(entity.Dish)

	if err := c.BindJSON(dish); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

	id, err := h.services.Dish.AddDish(ctx, dish)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	message := statusID {
		ID: id,
	}

	c.JSON(http.StatusOK, message)
}

// @Summary 	UpdateDish
// @Tags 		dish
// @Description update dish information by id
// @ID 			update-dish-info
// @Accept  	json
// @Produce  	json
// @Param 		id path int true "dishID"
// @Param 		input body entity.Dish true "dish information"
// @Success 	200 {object} statusResponse
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/{id} [put]
func (h *Handler) UpdateDishHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

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
	err = h.services.Dish.UpdateDish(ctx, dishID, dish)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newStatusResponse("Блюдо успешно изменено"))
}

// @Summary 	DeleteDish
// @Tags 		dish
// @Description delete dish by id
// @ID 			delete-dish
// @Accept  	json
// @Produce  	json
// @Param 		id path int true "dishID"
// @Success 	200 {object} statusResponse
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/{id} [delete]
func (h *Handler) DeleteDishHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dishID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err = h.services.Dish.DeleteDish(ctx, dishID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newStatusResponse("Блюдо успешно удалено"))
}

// @Summary 	GetDishesByCousine
// @Tags 		dish
// @Description get dishes by cousineID
// @ID 			get-cousine-dishes
// @Accept  	json
// @Produce  	json
// @Param 		cousineID path int true "cousineID"
// @Success 	200 {object} dishOutput
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/cousine/{cousineID} [get]
func (h *Handler) GetDishCousineHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousineID, err := strconv.Atoi(c.Param("cousineID"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCousine(ctx, cousineID)
	if err != nil {
		if err == sql.ErrNoRows {
			emptyArray := []interface{}{}
			c.JSON(http.StatusOK, emptyArray)
			return
		}
		
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var outputDishes []dishOutput
	for _, dish := range dishes {
		dishInfo := dishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}

// @Summary 	GetDishesByCategory
// @Tags 		dish
// @Description get dishes by categoryID
// @ID 			get-category-dishes
// @Accept  	json
// @Produce  	json
// @Param 		categoryID path int true "categoryID"
// @Success 	200 {object} dishOutput
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/category/{categoryID} [get]
func (h *Handler) GetDishCategoryHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCategory(ctx, categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			emptyArray := []interface{}{}
			c.JSON(http.StatusOK, emptyArray)
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	var outputDishes []dishOutput
	for _, dish := range dishes {
		dishInfo := dishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}

// @Summary 	GetDishesByCousineAndCategory
// @Tags 		dish
// @Description get dishes by cousineID and categoryID
// @ID 			get-cousine-category-dishes
// @Accept  	json
// @Produce  	json
// @Param 		cousineID path int true "cousineID"
// @Param 		categoryID path int true "categoryID"
// @Success 	200 {object} dishOutput
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/cousine/category/{cousineID}/{categoryID} [get]
func (h *Handler) GetDishCousineCategoryHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	cousineID, err := strconv.Atoi(c.Param("cousineID"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes, err := h.services.Dish.GetDishCousineCategory(ctx, cousineID, categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			emptyArray := []interface{}{}
			c.JSON(http.StatusOK, emptyArray)
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	var outputDishes []dishOutput
	for _, dish := range dishes {
		dishInfo := dishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}

// @Summary 	GetDishSearch
// @Tags 		dish
// @Description get dishes by name or ingredients
// @ID 			get-dish-search
// @Accept  	json
// @Produce  	json
// @Param 		text path string true "input text"
// @Param 		input body inputText true "dishName or dishIngredients"
// @Success 	200 {object} dishOutput
// @Failure 	400,404 {object} errorResponse
// @Failure 	500 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /dish/search/ [get]
func (h *Handler) GetDishSearchHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	dishes, err := h.services.Dish.GetDishSearch(ctx, c.Param("text"))
	if err != nil {
		if err == sql.ErrNoRows {
			emptyArray := []interface{}{}
			c.JSON(http.StatusOK, emptyArray)
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	var outputDishes []dishOutput
	for _, dish := range dishes {
		dishInfo := dishOutput {
			ID: dish.ID,
			Name: dish.Name,
			Description: dish.Description,
			Time: dish.Time,
		}
		outputDishes = append(outputDishes, dishInfo)
	}

	c.JSON(http.StatusOK, outputDishes)
}