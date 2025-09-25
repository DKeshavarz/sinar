package server

import (
	otphandler "github.com/DKeshavarz/sinar/internal/interface/server/otpHandler"
	userhandler "github.com/DKeshavarz/sinar/internal/interface/server/userHandler"
	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

func New(otp *usecase.OtpService, user usecase.User) *gin.Engine {
	r := gin.Default()
	setup(r, otp, user)
	return r
}

func setup(r *gin.Engine, otp *usecase.OtpService, user usecase.User) {
	otphandler.Register(r.Group("/otp"), otp)
	userhandler.Register(r.Group("/user/"), user)
}
