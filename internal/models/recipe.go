package models

type Recipe struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Steps       []Steps `json:"steps"`
	TotalTime   int     `json:"totalTime"`
	Rates       Rates   `json:"rates"`
}
type RecipeUpdate struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Ingredients *string `json:"ingredients"`
	TotalTime   *int    `json:"totalTime"`
}
