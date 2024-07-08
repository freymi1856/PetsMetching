package handler

import (
	"net/http"
	"strconv"

	"pet-matching-service/internal/model"
	"pet-matching-service/internal/service"

	"github.com/labstack/echo/v4"
)

type PetHandler struct {
	service *service.PetService
}

func NewPetHandler(service *service.PetService) *PetHandler {
	return &PetHandler{service: service}
}

// CreatePet godoc
// @Summary Create a new pet
// @Description Add a new pet to the system
// @Tags pets
// @Accept json
// @Produce json
// @Param pet body model.Pet true "New pet"
// @Success 201 {object} model.Pet
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /pets [post]
func (h *PetHandler) CreatePet(c echo.Context) error {
	pet := new(model.Pet)
	if err := c.Bind(pet); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	if err := h.service.CreatePet(pet); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, pet)
}

// GetPetByID godoc
// @Summary Get a pet by ID
// @Description Get details of a pet by ID
// @Tags pets
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} model.Pet
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /pets/{id} [get]
func (h *PetHandler) GetPetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid pet ID"})
	}

	pet, err := h.service.GetPetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "pet not found"})
	}

	return c.JSON(http.StatusOK, pet)
}

// GetAllPets godoc
// @Summary List all pets
// @Description Get a list of all pets
// @Tags pets
// @Produce json
// @Success 200 {array} model.Pet
// @Failure 500 {object} ErrorResponse
// @Router /pets [get]
func (h *PetHandler) GetAllPets(c echo.Context) error {
	pets, err := h.service.GetAllPets()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, pets)
}

// GetPetsByType godoc
// @Summary Get pets by type
// @Description Get a list of pets by type
// @Tags pets
// @Produce json
// @Param type query string true "Pet Type"
// @Success 200 {array} model.Pet
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /pets/type [get]
func (h *PetHandler) GetPetsByType(c echo.Context) error {
	petType := c.QueryParam("type")
	if petType == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "pet type is required"})
	}

	pets, err := h.service.GetPetsByType(petType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, pets)
}

// DeletePet godoc
// @Summary Delete a pet
// @Description Delete a pet by ID
// @Tags pets
// @Param id path int true "Pet ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /pets/{id} [delete]
func (h *PetHandler) DeletePet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid pet ID"})
	}

	if err := h.service.DeletePet(id); err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "pet not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
