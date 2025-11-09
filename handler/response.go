package handler

import (
	"fmt"
	"net/http"

	"github.com/H0wZy/go.learnings.restapi/schemas"
	"github.com/gin-gonic/gin"
)

func responseError(ctx *gin.Context, code int, success bool, msg string) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(code, gin.H{
		"statusCode": code,
		"success":    success,
		"message":    msg,
	})
}
func responseSuccess(ctx *gin.Context, opr string, data interface{}, success bool) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"success":    success,
		"message":    fmt.Sprintf("operation from handler: '%s' completed successfully", opr),
		"data":       data,
	})
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Success bool                    `json:"success"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Success bool                    `json:"success"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Success bool                    `json:"success"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Success bool                      `json:"success"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Success bool                    `json:"success"`
	Data    schemas.OpeningResponse `json:"data"`
}
