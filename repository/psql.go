package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reciept/models"
	"strings"
)

type Recipe struct {
	db *sql.DB
}

func NewRecipe(db *sql.DB) *Recipe {
	return &Recipe{db: db}
}

func (r *Recipe) Create(recipe models.Recipe) error {
	_, err := r.db.Exec("INSERT into recipe (name,description, ingredients, steps, total_time, rates, rates_quantity) VALUES($1,$2,$3,$4,$5,$6,$7)",
		recipe.Name, recipe.Description, recipe.Ingredients, recipe.Steps, recipe.TotalTime, recipe.Rates, recipe.RatesQuantity)

	return err
}
func (r *Recipe) GetByID(id int64) (models.Recipe, error) {
	var recipe models.Recipe
	err := r.db.QueryRow("SELECT id, name,description, ingredients, steps, total_time, rates, rates_quantity FROM recipe WHERE id=$1", id).
		Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.Steps, &recipe.TotalTime, &recipe.Rates, &recipe.RatesQuantity)
	if err == sql.ErrNoRows {
		return recipe, models.ErrRecipeNotFound
	}

	return recipe, err
}

func (r *Recipe) GetAll(ctx context.Context) ([]models.Recipe, error) {
	rows, err := r.db.Query("SELECT id, name,description, ingredients, steps, total_time, rates, rates_quantity FROM recipe")
	if err != nil {
		return nil, err
	}

	recipes := make([]models.Recipe, 0)
	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.Steps, &recipe.TotalTime, &recipe.Rates, &recipe.RatesQuantity); err != nil {
			return nil, err
		}

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

	if inp.Steps != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *inp.Steps)
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
