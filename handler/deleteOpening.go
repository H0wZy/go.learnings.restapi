package handler

import (
	"fmt"
	"net/http"

	"github.com/H0wZy/go.learnings.restapi/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, false, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, false, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, false, fmt.Sprintf("failed to delete opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening, true)
}
