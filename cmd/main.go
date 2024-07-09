package main

import (
	_ "pet-matching-service/docs"
	"pet-matching-service/internal/handler"
	"pet-matching-service/internal/model"
	"pet-matching-service/internal/repository"
	"pet-matching-service/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Pet Matching Service API
// @version 1.0
// @description This is a sample server for a pet matching service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	dsn := "host=db user=user password=password dbname=petdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Pet{}, &model.User{})

	petRepo := repository.NewPetRepository(db)
	petService := service.NewPetService(petRepo)
	petHandler := handler.NewPetHandler(petService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/pets", petHandler.CreatePet)
	e.GET("/pets/:id", petHandler.GetPetByID)
	e.GET("/pets", petHandler.GetAllPets)
	e.GET("/pets/type", petHandler.GetPetsByType)
	e.DELETE("/pets/:id", petHandler.DeletePet)

	e.POST("/users/register", userHandler.RegisterUser)
	e.POST("/users/login", userHandler.LoginUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.GET("/users", userHandler.GetAllUsers)

	e.Logger.Fatal(e.Start(":8080"))
}
