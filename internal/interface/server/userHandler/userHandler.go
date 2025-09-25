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

// GetByStudentNumber godoc
// @Summary Get user by student number
// @Description Get user information with university details by student number
// @Tags User
// @Accept json
// @Produce json
// @Param student_number path string true "Student Number"
// @Success 200 {object} dto.UserWithUniversity "User with university information"
// @Failure 400 {object} object{error=string} "Invalid student number"
// @Failure 404 {object} object{error=string} "User not found"
// @Router /user/{student_number} [get]
func (h *UserHandler) GetByStudentNumber(c *gin.Context) {
	number := c.Param("student_number")
	result, err := h.service.GetByStudentNumber(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
