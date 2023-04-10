package service

import (
	"reciept/internal/models"
)

type RecipeRepository interface {
	Create(recipe models.Recipe) error
	GetByID(id int64) (models.Recipe, error)
	GetAll() ([]models.Recipe, error)
	Delete(id int64) error
	Update(id int64, inp models.RecipeUpdate) error
}

type Recipe struct {
	repo RecipeRepository
}

func NewRecipe(repo RecipeRepository) *Recipe {
	return &Recipe{
		repo: repo,
	}
}

func (r *Recipe) Create(recipe models.Recipe) error {

	return r.repo.Create(recipe)
}

func (r *Recipe) GetByID(id int64) (models.Recipe, error) {
	return r.repo.GetByID(id)
}

func (r *Recipe) GetAll() ([]models.Recipe, error) {
	return r.repo.GetAll()
}

func (r *Recipe) Delete(id int64) error {
	return r.repo.Delete(id)
}

func (r *Recipe) Update(id int64, inp models.RecipeUpdate) error {
	return r.repo.Update(id, inp)
}
