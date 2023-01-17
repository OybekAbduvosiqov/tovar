package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/OybekAbduvosiqov/tovar/models"

	"github.com/gin-gonic/gin"

	"errors"
	"strconv"
)

// CreateTovar godoc
// @ID CreateTovar
// @Router /tovar [POST]
// @Summary CreateTovar
// @Description CreateTovar
// @Tags Tovar
// @Accept json
// @Produce json
// @Param tovar body models.CreateTovar true "CreateTovarRequestBody"
// @Success 201 {object} models.Tovar "GetTovarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateTovar(c *gin.Context) {

	var tovar models.CreateTovar

	err := c.ShouldBindJSON(&tovar)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Tovar().Insert(context.Background(), &tovar)
	if err != nil {
		log.Println("error whiling create tovar:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.Tovar().GetByID(context.Background(), &models.TovarPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id tovar:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIDTovar godoc
// @ID Get_By_IDTovar
// @Router /tovar/{id} [GET]
// @Summary GetByID Tovar
// @Description GetByID Tovar
// @Tags Tovar
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Tovar "GetByIDTovarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDTovar(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Tovar().GetByID(context.Background(), &models.TovarPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id tovar:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListTovar godoc
// @ID TovarPrimeryKey
// @Router /tovar [GET]
// @Summary Get List Tovar
// @Description Get List Tovar
// @Tags Tovar
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListTovarResponse "GetTovarListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListTovar(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res, err := h.storage.Tovar().GetList(context.Background(), &models.GetListTovarRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list tovar:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateTovar godoc
// @ID UpdateTovar
// @Router /tovar/{id} [PUT]
// @Summary Update Tovar
// @Description Update Tovar
// @Tags Tovar
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param tovar body models.UpdateTovarSwag true "UpdateTovarRequestBody"
// @Success 202 {object} models.Tovar "UpdateTovarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateTovar(c *gin.Context) {

	var (
		tovar models.UpdateTovar
	)

	err := c.ShouldBindJSON(&tovar)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tovar.Id = c.Param("id")

	err = h.storage.Tovar().Update(context.Background(), &models.UpdateTovar{
		Id:          tovar.Id,
		Name:        tovar.Name,
		Description: tovar.Description,
		Price:       tovar.Price,
		Photo:       tovar.Photo,
	})
	if err != nil {
		log.Printf("error whiling update 2: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	// if rowsAffected == 0 {
	// 	log.Printf("error whiling update rows affected: %v", err)
	// 	c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
	// 	return
	// }

	resp, err := h.storage.Tovar().GetByID(context.Background(), &models.TovarPrimeryKey{Id: tovar.Id})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// DeleteTovar godoc
// @ID DeleteTovar
// @Router /tovar/{id} [DELETE]
// @Summary Delete Tovar
// @Description Delete Tovar
// @Tags Tovar
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteTovarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteTovar(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Tovar().Delete(context.Background(), &models.TovarPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete tovar:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "delete tovar")
}
