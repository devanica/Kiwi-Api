package model

type Ingredient struct {
	ID          int    `json:"id"`
	Smoothie    string `json:"smoothie"`
	MIngredient string `json:"ingredient"`
}
