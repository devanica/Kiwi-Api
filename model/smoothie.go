package model

type Smoothie struct {
	MSmoothie 	string `json:"smoothie"`
	Base      	string `json:"base"`
	Source    	string `json:"source"`
	Recipe    	string `json:"recipe"`
	Ingredient  string `json:"ingredients"`
	Image 		string `json:"image"`
	ID        	int    `json:"id"`
}