package foodhandler

import (
	"net/http"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	service usecase.Food
}

func Register(group *gin.RouterGroup, service usecase.Food) {
	h := FoodHandler{
		service: service,
	}

	group.GET("/", h.GetAll)
}

func (h *FoodHandler) GetAll(c *gin.Context) {
	result, err := h.service.GetAllNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
