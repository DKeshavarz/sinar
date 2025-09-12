package main

import (
	"fmt"
	"time"

	"github.com/DKeshavarz/sinar/internal/interface/redis"
	"github.com/DKeshavarz/sinar/internal/interface/server"
	"github.com/DKeshavarz/sinar/internal/interface/sms"
	"github.com/DKeshavarz/sinar/internal/usecase"
)

func main(){
	otpStorage := redis.New()
	otpSender := sms.New("ff", "fff")
	optUsecase := usecase.NewOtpService(5, 2 * time.Second, otpStorage, otpSender)
	server := server.New(optUsecase)

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}