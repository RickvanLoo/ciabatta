package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"text/tabwriter"
)

//AddFlower adds flower to the recipe
//Can also be used to edit
func (r *Recipe) AddFlower(amount int) {
	flower := Ingredient{"Flower", amount, 1}
	r.Flower = flower

	for i, ing := range r.Ingredients {
		newAmount := float64(amount) * ing.Factor
		ing.Amount = int(math.Round(newAmount))
		r.Ingredients[i] = ing
	}
}

//Rename changed the name of the Recipe
func (r *Recipe) Rename(name string) {
	r.Name = name
}

//AddIngredient adds an ingredient to the recipe
func (r *Recipe) AddIngredient(name string, amount int) {
	fweight := r.Flower.Amount

	ing := Ingredient{name, amount, float64(amount) / float64(fweight)}
	r.Ingredients = append(r.Ingredients, ing)
}

//DeleteIngredient removes and ingredient from a recipe
func (r *Recipe) DeleteIngredient(id int) {
	r.Ingredients = append(r.Ingredients[:id], r.Ingredients[id+1:]...)
}

//EditIngredient changes amount and factor of ingredient
func (r *Recipe) EditIngredient(id int, amount int) error {
	if id > len(r.Ingredients)-1 {
		return errors.New("Out of Bounds")
	}
	fweight := r.Flower.Amount

	r.Ingredients[id].Amount = amount
	r.Ingredients[id].Factor = float64(amount) / float64(fweight)
	return nil
}

func (r *Recipe) printIngredients(exec bool) {
	var flower Ingredient
	var ingredients []Ingredient
	if exec {
		flower = r.ExecFlower
		ingredients = r.ExecIngredients
		fmt.Println("Remaining Ingredients:")

	} else {
		flower = r.Flower
		ingredients = r.Ingredients
		fmt.Println("Recipe: " + r.Name)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "ID\tName\tAmount\tFactor")
	flowerString := "f\t" + flower.Name + "\t" + strconv.Itoa(flower.Amount) + "\t" + fmt.Sprintf("%f", flower.Factor)
	fmt.Fprintln(w, flowerString)

	for i, ing := range ingredients {
		printString := strconv.Itoa(i) + "\t" + ing.Name + "\t" + strconv.Itoa(ing.Amount) + "\t" + fmt.Sprintf("%f", ing.Factor)
		fmt.Fprintln(w, printString)
	}

	fmt.Fprintln(w)
	w.Flush()
}

func (r *Recipe) printSteps() {
	for _, step := range r.Steps {
		step.print(r)
	}
}

func (r *Recipe) execCopy() {
	r.ExecIngredients = make([]Ingredient, len(r.Ingredients))
	copy(r.ExecIngredients, r.Ingredients)
	r.ExecFlower = r.Flower
}
