package handler

import (
	"fmt"
	"net/http"

	"github.com/H0wZy/go.learnings.restapi/schemas"
	"github.com/gin-gonic/gin"
)

// UpdateOpeningHandler
// @Summary UpdateOpeningHandler
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [patch]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("error binding request: %v", err.Error())
		responseError(ctx, http.StatusBadRequest, false, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		responseError(ctx, http.StatusBadRequest, false, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		responseError(ctx, http.StatusBadRequest, false, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		responseError(ctx, http.StatusNotFound, false, fmt.Errorf("opening with id: %s not found", id).Error())
		return
	}

	// Update fields
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		responseError(ctx, http.StatusInternalServerError, false, fmt.Errorf("failed to update opening: %v", err).Error())
		return
	}

	responseSuccess(ctx, "update-opening", request, true)
}
