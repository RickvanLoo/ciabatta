package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var currRecipe *Recipe

func main() {
	// recipe := newRecipe("Pizza")
	// recipe.AddFlower(1500)
	// recipe.AddIngredient("Water", 1000)
	// recipe.AddIngredient("Oil", 50)
	// recipe.AddIngredient("Salt", 25)
	// recipe.AddIngredient("Honey", 10)
	// recipe.AddIngredient("Dry Yeast", 5)
	// printIngredients(recipe)

	// recipe2 := convertRecipe(1000, recipe)
	// printIngredients(recipe2)

	fmt.Println("Ciabatta v0.1")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)
		inputSpaced := strings.Split(input, " ")

		if strings.HasPrefix(input, ":n") {
			fmt.Println("New Recipe: " + inputSpaced[1])
			currRecipe = newRecipe(inputSpaced[1])
		}

		if strings.HasPrefix(input, ":f") {
			i, _ := strconv.ParseInt(inputSpaced[1], 10, 64)

			currRecipe.AddFlower(int(i))
		}

		if strings.HasPrefix(input, ":a") {
			i, _ := strconv.ParseInt(inputSpaced[2], 10, 64)

			currRecipe.AddIngredient(inputSpaced[1], int(i))
		}

		if strings.HasPrefix(input, ":p") {
			printIngredients(currRecipe)
		}

		if strings.HasPrefix(input, ":c") {
			i, _ := strconv.ParseInt(inputSpaced[1], 10, 64)

			currRecipe = convertRecipe(int(i), currRecipe)
		}

		if strings.HasPrefix(input, ":q") {
			os.Exit(0)
		}

	}
}

func newRecipe(name string) *Recipe {
	recipe := new(Recipe)
	recipe.Name = name
	return recipe
}

func printIngredients(r *Recipe) {
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

func convertRecipe(amount int, r *Recipe) *Recipe {
	covR := newRecipe(r.Name + strconv.Itoa(amount))
	covR.AddFlower(amount)

	for _, ing := range r.Ingredients {
		newAmount := float64(amount) * ing.Factor
		covR.AddIngredient(ing.Name, int(math.Round(newAmount)))
	}

	return covR
}
