package userfoodhandler

import (
	"net/http"
	"strconv"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserFoodHandler struct {
	service usecase.UserFood
}

func Register(group *gin.RouterGroup, service usecase.UserFood) {
	h := UserFoodHandler{
		service: service,
	}

	group.GET("/active", h.GetActiveFoods)
	group.GET(":id", h.GetByID)
	group.POST("/", h.Create)
	group.POST("/:id/use", h.UseFood)
}

func (h *UserFoodHandler) GetActiveFoods(c *gin.Context) {
	result, err := h.service.GetActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *UserFoodHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userfood ID"})
		return
	}

	result, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *UserFoodHandler) Create(c *gin.Context) {
	var req struct {
		UserID          int    `json:"user_id" binding:"required"`
		FoodID          int    `json:"food_id" binding:"required"`
		RestaurantID    int    `json:"restaurant_id" binding:"required"`
		Price           int    `json:"price" binding:"required"`
		SinarPrice      int    `json:"sinar_price" binding:"required"`
		Code            string `json:"code" binding:"required"`
		ExpirationHours int    `json:"expiration_hours" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Purchase(req.UserID, req.FoodID, req.RestaurantID, req.Price, req.SinarPrice, req.Code, req.ExpirationHours)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *UserFoodHandler) UseFood(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userfood ID"})
		return
	}

	err = h.service.MarkAsUsed(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "food marked as used"})
}
