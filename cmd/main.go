package main

import (
	_ "pet-matching-service/docs" // Импорт для Swagger

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"pet-matching-service/internal/handler"
	"pet-matching-service/internal/repository"
	"pet-matching-service/internal/service"
)

// @title Pet Matching Service API
// @version 1.0
// @description This is a sample server for matching pets.
// @host localhost:8080
// @BasePath /

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	repo := repository.NewInMemoryPetRepository()
	svc := service.NewPetService(repo)
	h := handler.NewPetHandler(svc)

	e.POST("/pets", h.CreatePet) // @Summary Create a new pet
	// @Description Add a new pet to the system
	// @Tags pets
	// @Accept json
	// @Produce json
	// @Param pet body model.Pet true "New pet"
	// @Success 201 {object} model.Pet
	// @Failure 400 {object} echo.HTTPError
	// @Failure 500 {object} echo.HTTPError
	// @Router /pets [post]
	e.GET("/pets/:id", h.GetPetByID) // @Summary Get a pet by ID
	// @Description Get details of a pet by ID
	// @Tags pets
	// @Produce json
	// @Param id path int true "Pet ID"
	// @Success 200 {object} model.Pet
	// @Failure 400 {object} echo.HTTPError
	// @Failure 404 {object} echo.HTTPError
	// @Router /pets/{id} [get]
	e.GET("/pets", h.GetAllPets) // @Summary List all pets
	// @Description Get a list of all pets
	// @Tags pets
	// @Produce json
	// @Success 200 {array} model.Pet
	// @Failure 500 {object} echo.HTTPError
	// @Router /pets [get]
	e.DELETE("/pets/:id", h.DeletePet) // @Summary Delete a pet
	// @Description Delete a pet by ID
	// @Tags pets
	// @Param id path int true "Pet ID"
	// @Success 204
	// @Failure 400 {object} echo.HTTPError
	// @Failure 404 {object} echo.HTTPError
	// @Router /pets/{id} [delete]

	e.GET("/swagger/*", echoSwagger.WrapHandler) // Маршрут для Swagger

	e.Logger.Fatal(e.Start(":8080"))
}
