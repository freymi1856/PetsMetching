package repository

import (
	"pet-matching-service/internal/model"

	"github.com/jinzhu/gorm"
)

type PetRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
	return &PetRepository{db: db}
}

func (r *PetRepository) CreatePet(pet *model.Pet) error {
	return r.db.Create(pet).Error
}

func (r *PetRepository) GetPetByID(id int) (*model.Pet, error) {
	var pet model.Pet
	err := r.db.First(&pet, id).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (r *PetRepository) GetAllPets() ([]model.Pet, error) {
	var pets []model.Pet
	err := r.db.Find(&pets).Error
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (r *PetRepository) DeletePet(id int) error {
	return r.db.Delete(&model.Pet{}, id).Error
}
