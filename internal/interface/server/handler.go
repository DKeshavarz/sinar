package server

import (
	otphandler "github.com/DKeshavarz/sinar/internal/interface/server/otpHandler"
	restauranthandler "github.com/DKeshavarz/sinar/internal/interface/server/restaurantHandler"
	universityhandler "github.com/DKeshavarz/sinar/internal/interface/server/universityHandler"
	userhandler "github.com/DKeshavarz/sinar/internal/interface/server/userHandler"
	userfoodhandler "github.com/DKeshavarz/sinar/internal/interface/server/userfoodHandler"
	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/gin-gonic/gin"
)

func New(otp *usecase.OtpService, user usecase.User, university usecase.Univercity, userFood usecase.UserFood, restaurant usecase.Restaurant) *gin.Engine {
	r := gin.Default()
	setup(r, otp, user, university, userFood, restaurant)
	return r
}

func setup(r *gin.Engine, otp *usecase.OtpService, user usecase.User, university usecase.Univercity, userFood usecase.UserFood, restaurant usecase.Restaurant) {
	otphandler.Register(r.Group("/otp"), otp)
	userhandler.Register(r.Group("/user/"), user)
	universityhandler.Register(r.Group("/university/"), university)
	userfoodhandler.Register(r.Group("/userfood/"), userFood)
	restauranthandler.Register(r.Group("/restaurant/"), restaurant)
}
