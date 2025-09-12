package main

import (
	"fmt"
	"time"

	"github.com/DKeshavarz/sinar/internal/config"
	"github.com/DKeshavarz/sinar/internal/interface/redis"
	"github.com/DKeshavarz/sinar/internal/interface/server"
	"github.com/DKeshavarz/sinar/internal/interface/sms"
	"github.com/DKeshavarz/sinar/internal/usecase"
)

func main(){
	conf := config.New()

	otpStorage := redis.New(*conf.Redis)
	otpSender := sms.New(*conf.SMS)

	optUsecase := usecase.NewOtpService(5, 2 * time.Second, otpStorage, otpSender)
	server := server.New(optUsecase)

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}