package otphandler

import (
	"net/http"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

type OTPHandler struct {
	service *usecase.OtpService
}

func Register(group *gin.RouterGroup, service *usecase.OtpService) {
	o := OTPHandler{
		service: service,
	}

	group.POST("/create", o.RequestOTP)
	group.POST("/verify", o.VerifyOTP)
}

// RequestOTP godoc
// @Summary Request OTP
// @Description Send OTP code to user's phone number
// @Tags OTP
// @Accept json
// @Produce json
// @Param request body object{phone=string} true "Phone number"
// @Success 200 {object} object{message=string} "OTP sent successfully"
// @Failure 400 {object} object{error=string} "Invalid request"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /otp/create [post]
func (o *OTPHandler) RequestOTP(c *gin.Context) {
	var req struct {
		Phone string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := o.service.RequestOTP(req.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

// VerifyOTP godoc
// @Summary Verify OTP
// @Description Verify OTP code sent to user's phone
// @Tags OTP
// @Accept json
// @Produce json
// @Param request body object{phone=string,otp=string} true "Phone number and OTP code"
// @Success 200 {object} object{message=string} "OTP verified successfully"
// @Failure 400 {object} object{error=string} "Invalid request"
// @Failure 401 {object} object{error=string} "Invalid OTP"
// @Router /otp/verify [post]
func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req struct {
		Phone string `json:"phone" binding:"required"`
		OTP   string `json:"otp" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.VerifyOTP(req.Phone, req.OTP); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
