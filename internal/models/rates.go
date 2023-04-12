package models

type Rates struct {
	RecipeId     int64   `json:"recipe_id"`
	Rate         float64 `json:"rate"`
	RateQuantity int     `json:"rateQuantity"`
}
type RatesUpdate struct {
	Rate int `json:"rate"`
}
