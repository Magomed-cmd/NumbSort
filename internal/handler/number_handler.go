package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NumberHandler struct {
	service NumberService
}

type NumberService interface {
	AddAndList(ctx context.Context, value int) ([]int, error)
}

func NewNumberHandler(service NumberService) *NumberHandler {
	return &NumberHandler{service: service}
}

type numberRequest struct {
	Value *int `json:"value" binding:"required"`
}

type numberResponse struct {
	Numbers []int `json:"numbers"`
}

func (h *NumberHandler) AddNumber(c *gin.Context) {
	var req numberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if req.Value == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value is required"})
		return
	}

	values, err := h.service.AddAndList(c.Request.Context(), *req.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, numberResponse{Numbers: values})
}
