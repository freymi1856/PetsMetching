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

func main() {
	dsn := "host=db user=user password=password dbname=petdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Pet{})

	petRepo := repository.NewPetRepository(db)
	petService := service.NewPetService(petRepo)
	petHandler := handler.NewPetHandler(petService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/pets", petHandler.CreatePet)
	e.GET("/pets/:id", petHandler.GetPetByID)
	e.GET("/pets", petHandler.GetAllPets)
	e.GET("/pets/type", petHandler.GetPetsByType)
	e.DELETE("/pets/:id", petHandler.DeletePet)

	e.Logger.Fatal(e.Start(":8080"))
}
