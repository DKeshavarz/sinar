package server

import (
	otphandler "github.com/DKeshavarz/sinar/internal/interface/server/otpHandler"
	universityhandler "github.com/DKeshavarz/sinar/internal/interface/server/universityHandler"
	userhandler "github.com/DKeshavarz/sinar/internal/interface/server/userHandler"
	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

func New(otp *usecase.OtpService, user usecase.User, university usecase.Univercity) *gin.Engine {
	r := gin.Default()
	setup(r, otp, user, university)
	return r
}

func setup(r *gin.Engine, otp *usecase.OtpService, user usecase.User, university usecase.Univercity) {
	otphandler.Register(r.Group("/otp"), otp)
	userhandler.Register(r.Group("/user/"), user)
	universityhandler.Register(r.Group("/university/"), university)
}
