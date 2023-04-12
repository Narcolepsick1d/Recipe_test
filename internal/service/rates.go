package service

import "reciept/internal/models"

type RateRepository interface {
	CreateRates(rate models.Rates) error
}
type Rates struct {
	repo RateRepository
}

func (r *Rates) CreateRates(rate models.Rates) error {
	return r.repo.CreateRates(rate)
}
