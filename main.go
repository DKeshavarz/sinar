package main

import (
	"fmt"
	"time"

	"github.com/DKeshavarz/sinar/internal/config"
	pg "github.com/DKeshavarz/sinar/internal/interface/postgres"
	"github.com/DKeshavarz/sinar/internal/interface/redis"
	"github.com/DKeshavarz/sinar/internal/interface/server"
	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/DKeshavarz/sinar/pkg/sms"
)

func main() {
	conf := config.New()

	otpStorage := redis.New(*conf.Redis)
	otpSender := sms.New(*conf.SMS)

	optUsecase := usecase.NewOtpService(5, 3*time.Minute, otpStorage, otpSender)
	userRepo := pg.NewUserRepository()
	userUsecase := usecase.NewUser(userRepo)
	universityRepo := pg.NewUniversityRepository()
	universityUsecase := usecase.NewUnivercity(universityRepo)
	server := server.New(optUsecase, userUsecase, universityUsecase)

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
