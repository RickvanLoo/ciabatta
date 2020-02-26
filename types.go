package main

//Recipe is a baking recipe
type Recipe struct {
	Name        string
	Flower      Ingredient
	Ingredients []Ingredient
	Process     []Step
}

//Ingredient is a single element of Recipe
//Amount in gram
//Factor in % of flower, 1 for flower
type Ingredient struct {
	Name   string
	Amount int
	Factor float64
}

//Step is a global single step of Recipe
//Time (in minutes) is for Rest step, 0 for instant
//Factor is for mixing, 1 for all/remainder
//Ingredients, list of ingredients concerned
//Comment, added info
type Step struct {
	Type        StepType
	Time        int
	Factor      float64
	Ingredients []Ingredient
	Comment     string
}

//StepType is a type of step: mixing, kneading, shaping, etc.
type StepType int

//Mix is ...
const (
	MixMT StepType = iota
	MixST
	RestT
	KneedT
	ShapeST
	ShapePT
	CoverIT
	CoverT
	BakeT
)
