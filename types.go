package main

//Recipe is a baking recipe
type Recipe struct {
	Name        string
	Flower      Ingredient
	Ingredients []Ingredient
}

//Ingredient is a single element of Recipe
//Amount in gram
//Factor in % of flower, 1 for flower
type Ingredient struct {
	Name   string
	Amount int
	Factor float64
}
