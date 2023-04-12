package psql

import (
	"database/sql"
	"reciept/internal/models"
)

type Rates struct {
	db *sql.DB
}

func NewRates(db *sql.DB) *Rates {
	return &Rates{db: db}
}
func (r *Rates) CreateRates(rate models.Rates) error {
	_, err := r.db.Exec("insert into rates (recipe_id, rate, rate_quantity) VALUES ($1,$2,$3,$4)",
		rate.RecipeId, rate.Rate, rate.RateQuantity)

	return err
}
