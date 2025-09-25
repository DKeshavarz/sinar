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
