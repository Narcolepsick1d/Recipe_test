package models

type Steps struct {
	RecipeId        int64  `json:"recipe_id"`
	StepNumber      int    `json:"step_number"`
	StepDescription string `json:"step_description"`
	TimePerStep     int    `json:"time_per_step"`
}
type StepsUpdate struct {
	StepNumber      int    `json:"step_number"`
	StepDescription string `json:"step_description"`
	TimePerStep     int    `json:"time_per_step"`
}
