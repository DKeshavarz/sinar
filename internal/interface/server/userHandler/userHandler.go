package userhandler

import (
	"net/http"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service usecase.User
}

func Register(group *gin.RouterGroup, service usecase.User) {
	h := UserHandler{service: service}

	group.GET(":student_number", h.GetByStudentNumber)
}

func (h *UserHandler) GetByStudentNumber(c *gin.Context) {
	number := c.Param("student_number")
	result, err := h.service.GetByStudentNumber(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
