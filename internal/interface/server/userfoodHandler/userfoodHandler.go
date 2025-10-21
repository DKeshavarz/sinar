package userfoodhandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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

// GetActiveFoods godoc
// @Summary Get active user foods
// @Description Get all user foods that have not expired yet
// @Tags UserFood
// @Accept json
// @Produce json
// @Success 200 {array} dto.UserFood "List of active user foods"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /userfood/active [get]
func (h *UserFoodHandler) GetActiveFoods(c *gin.Context) {
	result, err := h.service.GetActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetByID godoc
// @Summary Get user food by ID
// @Description Get specific user food details by ID
// @Tags UserFood
// @Accept json
// @Produce json
// @Param id path int true "UserFood ID"
// @Success 200 {object} dto.UserFood "User food details"
// @Failure 400 {object} object{error=string} "Invalid userfood ID"
// @Failure 404 {object} object{error=string} "User food not found"
// @Router /userfood/{id} [get]
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

// Create godoc
// @Summary Create user food purchase
// @Description Create a new user food purchase (supports both single object and array formats)
// @Tags UserFood
// @Accept json
// @Produce json
// @Param request body object{user_id=int,food_id=int,restaurant_id=int,price=int,sinar_price=int,code=string,expiration_hours=int} true "Purchase details (single object format)"
// @Param request body array true "Purchase details (array format)"
// @Success 201 {object} domain.UserFood "Created user food (single object)"
// @Success 201 {array} domain.UserFood "Created user foods (array)"
// @Failure 400 {object} object{error=string} "Invalid request"
// @Router /userfood/ [post]
func (h *UserFoodHandler) Create(c *gin.Context) {
	// Read the raw body once
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
		return
	}

	// Try to parse as single object with expiration_hours first (most common case)
	var req struct {
		UserID          int    `json:"user_id"`
		FoodID          int    `json:"food_id"`
		RestaurantID    int    `json:"restaurant_id"`
		Price           int    `json:"price"`
		SinarPrice      int    `json:"sinar_price"`
		Code            string `json:"code"`
		ExpirationHours int    `json:"expiration_hours"`
	}

	if err := json.Unmarshal(body, &req); err == nil && req.UserID > 0 && req.FoodID > 0 && req.RestaurantID > 0 && req.Code != "" && req.ExpirationHours > 0 {
		// Handle single object input
		result, err := h.service.Purchase(req.UserID, req.FoodID, req.RestaurantID, req.Price, req.SinarPrice, req.Code, req.ExpirationHours)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, result)
		return
	}

	// If single object parsing failed, try to parse as array
	var reqArray []struct {
		UserID       int    `json:"user_id"`
		FoodID       int    `json:"food_id"`
		RestaurantID int    `json:"restaurant_id"`
		Price        int    `json:"price"`
		SinarPrice   int    `json:"sinar_price"`
		Code         string `json:"code"`
		ExpiresAt    string `json:"expires_at"`
	}

	if err := json.Unmarshal(body, &reqArray); err == nil && len(reqArray) > 0 {
		// Handle array input
		var results []interface{}
		for _, req := range reqArray {
			// Calculate expiration hours from expires_at timestamp
			expirationHours := 24 // Default to 24 hours
			if expiresAt, err := time.Parse(time.RFC3339, req.ExpiresAt); err == nil {
				hours := int(time.Until(expiresAt).Hours())
				if hours > 0 {
					expirationHours = hours
				}
			}

			result, err := h.service.Purchase(req.UserID, req.FoodID, req.RestaurantID, req.Price, req.SinarPrice, req.Code, expirationHours)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			results = append(results, result)
		}
		c.JSON(http.StatusCreated, results)
		return
	}

	// If both parsing attempts failed, return error
	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format or missing required fields"})
}

// UseFood godoc
// @Summary Mark food as used
// @Description Mark a user food as used/expired
// @Tags UserFood
// @Accept json
// @Produce json
// @Param id path int true "UserFood ID"
// @Success 200 {object} object{message=string} "Food marked as used"
// @Failure 400 {object} object{error=string} "Invalid userfood ID"
// @Failure 409 {object} object{error=string} "Food already used/expired"
// @Router /userfood/{id}/use [post]
func (h *UserFoodHandler) UseFood(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userfood ID"})
		return
	}

	err = h.service.MarkAsUsed(id)
	if err != nil {

		if err.Error() == "food is already used/expired" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "food marked as used"})
}
