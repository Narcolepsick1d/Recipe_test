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
	var exist bool
	err := r.db.QueryRow("select exists(select rate,rate_quantity from rates where recipe_id=$1)", rate.RecipeId).
		Scan(&exist)
	if exist {
		ratesFromDb, _ := r.GetRatesFromDby(rate.RecipeId)

		rateAvg := (ratesFromDb.Rate + rate.Rate) / (ratesFromDb.RateQuantity + 1)
		_, err := r.db.Exec("insert into rates (recipe_id, rate, rate_quantity) VALUES ($1,$2,$3)",
			rate.RecipeId, rateAvg, rate.RateQuantity+1)
		if err != nil {
			return err
		}
	} else {
		_, err := r.db.Exec("insert into rates (recipe_id, rate, rate_quantity) VALUES ($1,$2,$3)",
			rate.RecipeId, rate.Rate, rate.RateQuantity)
		if err != nil {
			return err
		}
	}

	return err
}
func (r *Rates) GetRatesFromDby(id int64) (models.Rates, error) {
	var rate models.Rates
	err := r.db.QueryRow("select rate,rate_quantity from rates where recipe_id=$1", id).
		Scan(&rate.Rate, &rate.RateQuantity)
	if err == sql.ErrNoRows {
		return rate, models.ErrRecipeNotFound
	}
	return rate, err
}
