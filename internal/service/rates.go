package service

import "reciept/internal/models"

type RateRepository interface {
	UpdateRates(rate models.Rates) error
	FilteredByRates(rate float64) ([]models.Recipe, error)
}
type Rates struct {
	repo RateRepository
}

func (r *Rates) UpdateRates(rate models.Rates) error {
	return r.repo.UpdateRates(rate)
}
func (r *Rates) FilteredByRates(rate float64) ([]models.Recipe, error) {
	return r.repo.FilteredByRates(rate)
}
