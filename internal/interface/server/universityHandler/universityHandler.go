package universityhandler

import (
	"net/http"
	"strconv"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UniversityHandler struct {
	service usecase.Univercity
}

func Register(group *gin.RouterGroup, service usecase.Univercity) {
	h := UniversityHandler{
		service: service,
	}

	group.GET(":id", h.GetByID)
}

// GetByID godoc
// @Summary Get university by ID
// @Description Get university information by university ID
// @Tags University
// @Accept json
// @Produce json
// @Param id path int true "University ID"
// @Success 200 {object} domain.University "University information"
// @Failure 400 {object} object{error=string} "Invalid university ID"
// @Failure 404 {object} object{error=string} "University not found"
// @Router /university/{id} [get]
func (h *UniversityHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid university ID"})
		return
	}

	result, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
