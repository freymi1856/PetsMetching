package service

import (
	"pet-matching-service/internal/model"
	"pet-matching-service/internal/repository"
)

type PetService struct {
	repo *repository.PetRepository
}

func NewPetService(repo *repository.PetRepository) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) CreatePet(pet *model.Pet) error {
	return s.repo.CreatePet(pet)
}

func (s *PetService) GetPetByID(id int) (*model.Pet, error) {
	return s.repo.GetPetByID(id)
}

func (s *PetService) GetAllPets() ([]model.Pet, error) {
	return s.repo.GetAllPets()
}

func (s *PetService) DeletePet(id int) error {
	return s.repo.DeletePet(id)
}
