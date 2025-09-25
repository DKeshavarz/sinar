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

// GetAll godoc
// @Summary Get all foods
// @Description Get list of all available foods
// @Tags Food
// @Accept json
// @Produce json
// @Success 200 {array} domain.Food "List of all foods"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /food/ [get]
func (h *FoodHandler) GetAll(c *gin.Context) {
	result, err := h.service.GetAllNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
