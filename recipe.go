package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"text/tabwriter"
)

//AddFlower adds flower to the recipe
func (r *Recipe) AddFlower(amount int) {
	flower := Ingredient{"Flower", amount, 1}
	r.Flower = flower

	for i, ing := range r.Ingredients {
		newAmount := float64(amount) * ing.Factor
		ing.Amount = int(math.Round(newAmount))
		r.Ingredients[i] = ing
	}
}

//AddIngredient adds an ingredient to the recipe
func (r *Recipe) AddIngredient(name string, amount int) {
	fweight := r.Flower.Amount

	ing := Ingredient{name, amount, float64(amount) / float64(fweight)}
	r.Ingredients = append(r.Ingredients, ing)
}

//AddMixM adds a Multiple Mixing Step
func (r *Recipe) AddMixM(indices []int) {
	var ings []Ingredient
	for _, i := range indices {
		ings = append(ings, r.Ingredients[i])
	}

	step := new(Step)
	step.Type = MixMT
	step.Time = 0
	step.Factor = 1
	step.Ingredients = ings
	step.Comment = ""

	r.Process = append(r.Process, *step)
}

func (r *Recipe) printIngredients() {
	w := new(tabwriter.Writer)
	fmt.Println("Recipe: " + r.Name)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Name\tAmount\tFactor")
	flowerString := r.Flower.Name + "\t" + strconv.Itoa(r.Flower.Amount) + "\t" + fmt.Sprintf("%f", r.Flower.Factor)
	fmt.Fprintln(w, flowerString)

	for _, ing := range r.Ingredients {
		printString := ing.Name + "\t" + strconv.Itoa(ing.Amount) + "\t" + fmt.Sprintf("%f", ing.Factor)
		fmt.Fprintln(w, printString)
	}

	fmt.Fprintln(w)
	w.Flush()
}
