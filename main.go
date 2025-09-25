package main

import (
	"fmt"
	"time"

	"github.com/DKeshavarz/sinar/docs"
	"github.com/DKeshavarz/sinar/internal/config"
	pg "github.com/DKeshavarz/sinar/internal/interface/postgres"
	"github.com/DKeshavarz/sinar/internal/interface/redis"
	"github.com/DKeshavarz/sinar/internal/interface/server"
	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/DKeshavarz/sinar/pkg/sms"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sinar API
// @version 1.0
// @description A comprehensive food ordering and management system for universities
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	conf := config.New()

	otpStorage := redis.New(*conf.Redis)
	otpSender := sms.New(*conf.SMS)

	optUsecase := usecase.NewOtpService(5, 3*time.Minute, otpStorage, otpSender)
	userRepo := pg.NewUserRepository()
	userUsecase := usecase.NewUser(userRepo)
	universityRepo := pg.NewUniversityRepository()
	universityUsecase := usecase.NewUnivercity(universityRepo)
	userFoodRepo := pg.NewUserFoodRepository()
	userFoodUsecase := usecase.NewUserFood(userFoodRepo)
	restaurantRepo := pg.NewRestaurantRepository()
	restaurantUsecase := usecase.NewRestaurant(restaurantRepo)
	foodRepo := pg.NewFoodRepository()
	foodUsecase := usecase.NewFood(foodRepo)
	server := server.New(optUsecase, userUsecase, universityUsecase, userFoodUsecase, restaurantUsecase, foodUsecase)

	// Add Swagger documentation route
	docs.SwaggerInfo.BasePath = "/"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
