package main

import "time"

//Recipe is a baking recipe
type Recipe struct {
	Name            string
	Flower          Ingredient
	Ingredients     []Ingredient
	ExecFlower      Ingredient
	ExecIngredients []Ingredient
	Steps           []Process
	Locked          bool
}

//Ingredient is a single element of Recipe
//Amount in gram
//Factor in % of flower, 1 for flower
type Ingredient struct {
	Name   string
	Amount int
	Factor float64
}

//Process is...
type Process interface {
	execute(*Recipe) error
	print(*Recipe) error
}

//MixArray mixes all ingredients corresponding to ID array
type MixArray struct {
	IDs     []int
	Comment string
}

//MixSingle mixes a single ingredient according to factor
type MixSingle struct {
	ID      int
	Factor  float64
	Comment string
}

//Knead is a simple knead instruction
type Knead struct {
	Comment string
}

//Rest is a timed instruction
type Rest struct {
	Duration time.Duration
}

//Shape shapes into shape
type Shape struct {
	Shape string
}

//Bake bakes at a temp and duration
type Bake struct {
	Temperature int
	Duration    time.Duration
}
