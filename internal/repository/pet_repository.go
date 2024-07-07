package repository

import (
	"fmt"
	"pet-matching-service/internal/model"
)

type PetRepository interface {
	CreatePet(pet *model.Pet) error
	GetPetByID(id int) (*model.Pet, error)
	GetAllPets() ([]*model.Pet, error)
	DeletePet(id int) error
}

type InMemoryPetRepository struct {
	pets   map[int]*model.Pet
	nextID int
}

func NewInMemoryPetRepository() *InMemoryPetRepository {
	return &InMemoryPetRepository{
		pets:   make(map[int]*model.Pet),
		nextID: 1,
	}
}

func (r *InMemoryPetRepository) CreatePet(pet *model.Pet) error {
	pet.ID = r.nextID
	r.nextID++
	r.pets[pet.ID] = pet
	return nil
}

func (r *InMemoryPetRepository) GetPetByID(id int) (*model.Pet, error) {
	pet, exists := r.pets[id]
	if !exists {
		return nil, fmt.Errorf("pet not found")
	}
	return pet, nil
}

func (r *InMemoryPetRepository) GetAllPets() ([]*model.Pet, error) {
	pets := make([]*model.Pet, 0, len(r.pets))
	for _, pet := range r.pets {
		pets = append(pets, pet)
	}
	return pets, nil
}

func (r *InMemoryPetRepository) DeletePet(id int) error {
	if _, exists := r.pets[id]; !exists {
		return fmt.Errorf("pet not found")
	}
	delete(r.pets, id)
	return nil
}
