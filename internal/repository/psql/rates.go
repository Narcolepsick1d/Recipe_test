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
func (r *Rates) UpdateRates(rate models.Rates) error {
	ratesFromDb, err := r.GetRatesFromDb(rate.RecipeId)
	rateAvg := (ratesFromDb.Rate + rate.Rate) / (float64(ratesFromDb.RateQuantity) + 1.0)

	_, err = r.db.Exec("Update rates set rate = $1, rate_quantity=$2 where recipe_id=$3",
		rateAvg, ratesFromDb.RateQuantity+1, rate.RecipeId)
	if err != nil {
		return err
	}
	return err
}
func (r *Rates) GetRatesFromDb(id int64) (models.Rates, error) {
	var rate models.Rates
	err := r.db.QueryRow("select rate,rate_quantity  from rates where recipe_id=$1", id).
		Scan(&rate.Rate, &rate.RateQuantity)
	if err == sql.ErrNoRows {
		return rate, models.ErrRecipeNotFound
	}

	return rate, err
}
func (r *Rates) FilteredByRates(rate float64) ([]models.Recipe, error) {
	rows, err := r.db.Query("select recipe.id, recipe.name,recipe.description,recipe.ingredients,rates.rate,rates.rate_quantity,recipe.total_time from recipe  join rates  on recipe.id = rates.recipe_id where rates.rate>=$1 order by rates.rate desc;", rate)
	if err != nil {
		return nil, err
	}
	recipes := make([]models.Recipe, 0)
	for rows.Next() {
		var recipe models.Recipe
		steps := make([]models.Steps, 0)
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.TotalTime, &recipe.Rates.Rate, &recipe.Rates.RateQuantity); err != nil {
			return nil, err
		}

		rowsForStep, err := r.db.Query("select recipe_id,step_number,step_description,time_per_step from steps where recipe_id=$1 ", recipe.ID)
		if err != nil {
			return nil, err
		}
		for rowsForStep.Next() {
			var step models.Steps
			if err := rowsForStep.Scan(&step.RecipeId, &step.StepNumber, &step.StepDescription, &step.TimePerStep); err != nil {
				return nil, err
			}
			steps = append(steps, step)
		}
		recipe.Steps = steps
		recipes = append(recipes, recipe)

	}

	return recipes, rows.Err()
}
