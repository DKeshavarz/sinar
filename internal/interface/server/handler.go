package server

import (
	otphandler "github.com/DKeshavarz/sinar/internal/interface/server/otpHandler"
	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

func New(otp *usecase.OtpService)*gin.Engine {
	r := gin.Default()
	setup(r, otp)
	return r
}

func setup(r *gin.Engine, otp *usecase.OtpService) {
	otphandler.Register(r.Group("/otp"), otp)
}