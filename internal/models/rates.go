package models

type Rates struct {
	UserId       int64 `json:"user_id"`
	RecipeId     int64 `json:"recipe_id"`
	Rate         int   `json:"rate"`
	RateQuantity int   `json:"rateQuantity"`
}
type RatesUpdate struct {
	Rate int `json:"rate"`
}
