package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var currRecipe *Recipe
var defaultFolder string

func main() {
	fmt.Println("Ciabatta v0.1")
	defaultFolder = "/home/rick/Documents/Recipes/"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)
		ProcessInput(input)

	}
}

//NewRecipe creates a new recipe and replaces the recipe in memory
func NewRecipe(name string) *Recipe {
	recipe := new(Recipe)
	recipe.Name = name
	return recipe
}

//ProcessInput takes console input as string and executes commands
func ProcessInput(input string) {
	inputSpaced := strings.Split(input, " ")

	if strings.HasPrefix(input, ":new") {
		fmt.Println("New Recipe: " + inputSpaced[1])
		currRecipe = NewRecipe(inputSpaced[1])
	}

	if strings.HasPrefix(input, ":f") {
		i, _ := strconv.ParseInt(inputSpaced[1], 10, 64)

		currRecipe.AddFlower(int(i))
	}

	if strings.HasPrefix(input, ":a") {
		i, _ := strconv.ParseInt(inputSpaced[2], 10, 64)

		currRecipe.AddIngredient(inputSpaced[1], int(i))
	}

	if strings.HasPrefix(input, "ls") {
		currRecipe.printIngredients()
	}

	if strings.HasPrefix(input, ":s") {
		Save()
	}

	if strings.HasPrefix(input, ":o") {
		Open(inputSpaced[1])
	}

	if strings.HasPrefix(input, ":q") {
		os.Exit(0)
	}

}

//Save recipe to defaultFolder/RecipeName.json
func Save() {
	location := defaultFolder + currRecipe.Name + ".json"
	fmt.Println(location)

	file, _ := json.MarshalIndent(currRecipe, "", " ")
	_ = ioutil.WriteFile(location, file, 0644)
}

//Open recipe from defaultFolder/name.json
func Open(name string) {
	location := defaultFolder + name + ".json"

	file, _ := os.Open(location)
	byteValue, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(byteValue, &currRecipe)

	fmt.Println(location)
}
