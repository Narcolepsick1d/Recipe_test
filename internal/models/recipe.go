package models

type Recipe struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Ingredients   string  `json:"ingredients"`
	Steps         []Steps `json:"steps"`
	TotalTime     int     `json:"totalTime"`
	Rates         float64 `json:"rates"`
	RatesQuantity int     `json:"rates_quantity"`
}
type RecipeUpdate struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Ingredients *string  `json:"ingredients"`
	Steps       *[]Steps `json:"steps"`
	TotalTime   *int     `json:"totalTime"`
}
