package restauranthandler

import (
	"net/http"
	"strconv"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	service usecase.Restaurant
}

func Register(group *gin.RouterGroup, service usecase.Restaurant) {
	h := RestaurantHandler{
		service: service,
	}

	group.GET(":university_id", h.GetByUniversityID)
}

// GetByUniversityID godoc
// @Summary Get restaurants by university ID
// @Description Get all restaurants for a specific university
// @Tags Restaurant
// @Accept json
// @Produce json
// @Param university_id path int true "University ID"
// @Success 200 {array} domain.Restaurant "List of restaurants"
// @Failure 400 {object} object{error=string} "Invalid university ID"
// @Failure 404 {object} object{error=string} "No restaurants found"
// @Router /restaurant/{university_id} [get]
func (h *RestaurantHandler) GetByUniversityID(c *gin.Context) {
	universityIDStr := c.Param("university_id")
	universityID, err := strconv.Atoi(universityIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid university ID"})
		return
	}

	result, err := h.service.GetAll(universityID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
