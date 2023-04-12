package psql

import (
	"database/sql"
	"fmt"
	"reciept/internal/models"
	"strings"
)

type Recipe struct {
	db *sql.DB
}

func NewRecipe(db *sql.DB) *Recipe {
	return &Recipe{db: db}
}

func (r *Recipe) Create(recipe models.Recipe) error {
	totalTime := r.getTotalTime(recipe)
	_, err := r.db.Exec("INSERT into recipe (name,description, ingredients,total_time) VALUES($1,$2,$3,$4)",
		recipe.Name, recipe.Description, recipe.Ingredients, totalTime)
	for _, step := range recipe.Steps {
		recipeId, _ := r.getId(recipe.Name)
		_, _ = r.db.Exec("INSERT into steps(recipe_id,step_number,step_description,time_per_step) VALUES ($1,$2,$3,$4)",
			recipeId, step.StepNumber, step.StepDescription, step.TimePerStep)
	}
	return err
}
func (r *Recipe) GetByID(id int64) (models.Recipe, error) {
	var recipe models.Recipe
	err := r.db.QueryRow("select recipe.name,recipe.description,recipe.ingredients,recipe.total_time,rates.rate,rates.rate_quantity from recipe join rates  on recipe.id = rates.recipe_id where recipe_id=1;\n", id).
		Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.TotalTime)
	if err == sql.ErrNoRows {
		return recipe, models.ErrRecipeNotFound
	}
	rowsForStep, err := r.db.Query("select step_number,step_description,time_per_step from steps where recipe_id=$1", id)
	if err != nil {
		return recipe, err
	}
	steps := make([]models.Steps, 0)
	for rowsForStep.Next() {
		var step models.Steps
		if err := rowsForStep.Scan(&step.StepNumber, &step.StepDescription, &step.TimePerStep); err != nil {
			return recipe, err
		}
		steps = append(steps, step)
	}
	recipe.Steps = steps
	return recipe, err
}

func (r *Recipe) GetAll() ([]models.Recipe, error) {
	rows, err := r.db.Query("select recipe.name,recipe.description,recipe.ingredients,recipe.total_time,rates.rate,rates.rate_quantity from recipe join rates  on recipe.id = rates.recipe_id ")
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

		rowsForStep, err := r.db.Query("select step_number,step_description,time_per_step from steps where recipe_id=$1", recipe.ID)
		if err != nil {
			return nil, err
		}
		for rowsForStep.Next() {
			var step models.Steps
			if err := rowsForStep.Scan(&step.StepNumber, &step.StepDescription, &step.TimePerStep); err != nil {
				return nil, err
			}
			steps = append(steps, step)
		}
		recipe.Steps = steps
		recipes = append(recipes, recipe)

	}

	return recipes, rows.Err()
}

func (r *Recipe) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM recipe WHERE id=$1", id)

	return err
}
func (r *Recipe) Update(id int64, inp models.RecipeUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *inp.Name)
		argId++
	}

	if inp.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *inp.Description)
		argId++
	}

	if inp.Ingredients != nil {
		setValues = append(setValues, fmt.Sprintf("publish_date=$%d", argId))
		args = append(args, *inp.Ingredients)
		argId++
	}

	if inp.TotalTime != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *inp.TotalTime)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE recipe SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
func (r *Recipe) GetByIngredient(ingredient string) ([]models.Recipe, error) {
	rows, err := r.db.Query("select recipe.name,recipe.description,recipe.ingredients,recipe.total_time,rates.rate,rates.rate_quantity from recipe join rates  on recipe.id = rates.recipe_id  WHERE recipe.ingredients like  '%' || $1 || '%'", ingredient)
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

		rowsForStep, err := r.db.Query("select step_number,step_description,time_per_step from steps where recipe_id=$1", recipe.ID)
		if err != nil {
			return nil, err
		}
		for rowsForStep.Next() {
			var step models.Steps
			if err := rowsForStep.Scan(&step.StepNumber, &step.StepDescription, &step.TimePerStep); err != nil {
				return nil, err
			}
			steps = append(steps, step)
		}
		recipe.Steps = steps
		recipes = append(recipes, recipe)

	}

	return recipes, rows.Err()
}

func (r *Recipe) FilteredByTime(totalTime int) ([]models.Recipe, error) {
	rows, err := r.db.Query("select recipe.name,recipe.description,recipe.ingredients,recipe.total_time,rates.rate,rates.rate_quantity from recipe join rates  on recipe.id = rates.recipe_id  WHERE  total_time<=$1 order by total_time", totalTime)
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

		rowsForStep, err := r.db.Query("select step_number,step_description,time_per_step from steps where recipe_id=$1", recipe.ID)
		if err != nil {
			return nil, err
		}
		for rowsForStep.Next() {
			var step models.Steps
			if err := rowsForStep.Scan(&step.StepNumber, &step.StepDescription, &step.TimePerStep); err != nil {
				return nil, err
			}
			steps = append(steps, step)
		}
		recipe.Steps = steps
		recipes = append(recipes, recipe)

	}

	return recipes, rows.Err()
}

func (r *Recipe) getTotalTime(recipe models.Recipe) int {
	totalTime := 0
	for _, stepTime := range recipe.Steps {
		totalTime += stepTime.TimePerStep
	}
	return totalTime
}
func (r *Recipe) getId(name string) (int64, error) {
	var recipe models.Recipe
	err := r.db.QueryRow("SELECT id FROM recipe WHERE name=$1", name).
		Scan(&recipe.ID)
	if err == sql.ErrNoRows {
		return 0, models.ErrRecipeNotFound
	}

	return recipe.ID, err
}
