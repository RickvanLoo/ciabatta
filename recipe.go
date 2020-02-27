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

func (r *Recipe) printIngredients() {
	w := new(tabwriter.Writer)
	fmt.Println("Recipe: " + r.Name)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "ID\tName\tAmount\tFactor")
	flowerString := "f\t" + r.Flower.Name + "\t" + strconv.Itoa(r.Flower.Amount) + "\t" + fmt.Sprintf("%f", r.Flower.Factor)
	fmt.Fprintln(w, flowerString)

	for i, ing := range r.Ingredients {
		printString := strconv.Itoa(i) + "\t" + ing.Name + "\t" + strconv.Itoa(ing.Amount) + "\t" + fmt.Sprintf("%f", ing.Factor)
		fmt.Fprintln(w, printString)
	}

	fmt.Fprintln(w)
	w.Flush()
}
