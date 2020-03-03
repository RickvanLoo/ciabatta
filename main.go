package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var currRecipe *Recipe
var execRecipe *Recipe
var defaultFolder string

func main() {
	fmt.Println("Ciabatta v0.2")
	defaultFolder = "/home/rick/Documents/Recipes/"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")
	currRecipe = NewRecipe("empty")
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
	recipe.Locked = false
	return recipe
}

//ProcessInput processes input
func ProcessInput(input string) {
	if currRecipe.Locked {
		processInputLocked(input)
	} else {
		processInputUnlocked(input)
	}
}

//ProcessInputUnlocked takes arguments when recipe is not locked
func processInputUnlocked(input string) {
	inputSpaced := strings.Split(input, " ")

	if strings.HasPrefix(input, ":n") {
		err := InputArgumentCheck(inputSpaced, 1)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("New Recipe: " + inputSpaced[1])
		currRecipe = NewRecipe(inputSpaced[1])
	}

	if strings.HasPrefix(input, ":r") {
		err := InputArgumentCheck(inputSpaced, 1)
		if err != nil {
			log.Println(err)
			return
		}

		currRecipe.Rename(inputSpaced[1])
	}

	if strings.HasPrefix(input, ":f") {
		err := InputArgumentCheck(inputSpaced, 1)
		if err != nil {
			log.Println(err)
			return
		}

		i, _ := strconv.ParseInt(inputSpaced[1], 10, 64)

		currRecipe.AddFlower(int(i))
	}

	if strings.HasPrefix(input, ":a") {
		err := InputArgumentCheck(inputSpaced, 2)
		if err != nil {
			log.Println(err)
			return
		}

		i, _ := strconv.ParseInt(inputSpaced[2], 10, 64)

		currRecipe.AddIngredient(inputSpaced[1], int(i))
	}

	if strings.HasPrefix(input, ":e") {
		err := InputArgumentCheck(inputSpaced, 2)
		if err != nil {
			log.Println(err)
			return
		}

		i, _ := strconv.Atoi(inputSpaced[1])
		j, _ := strconv.Atoi(inputSpaced[2])

		err = currRecipe.EditIngredient(i, j)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if strings.HasPrefix(input, ":d") {
		err := InputArgumentCheck(inputSpaced, 1)
		if err != nil {
			log.Println(err)
			return
		}

		i, err := strconv.Atoi(inputSpaced[1])
		if err != nil {
			log.Println(err)
			return
		}

		currRecipe.DeleteIngredient(i)
	}

	if strings.HasPrefix(input, "ls") {
		currRecipe.printIngredients(false)
		currRecipe.printSteps()
	}

	if strings.HasPrefix(input, ":s") {
		Save()
	}

	if strings.HasPrefix(input, ":lock") {
		fmt.Println("Locked!")
		currRecipe.execCopy()
		currRecipe.Steps = nil
		currRecipe.Locked = true
		return
	}

	if strings.HasPrefix(input, ":l") {
		err := InputArgumentCheck(inputSpaced, 1)
		if err != nil {
			log.Println(err)
			return
		}

		Open(inputSpaced[1])
	}

	if strings.HasPrefix(input, ":q") {
		os.Exit(0)
	}

}

func processInputLocked(input string) {
	inputSpaced := strings.Split(input, " ")

	if strings.HasPrefix(input, ":q") {
		os.Exit(0)
	}

	if strings.HasPrefix(input, ":lock") {
		fmt.Println("Unlocked!")
		currRecipe.Locked = false
	}

	if strings.HasPrefix(input, "mix") {
		ma := createMixArray(inputSpaced[1:], "")
		currRecipe.Steps = append(currRecipe.Steps, &ma)
		i := len(currRecipe.Steps) - 1

		err := currRecipe.Steps[i].execute(currRecipe)
		if err != nil {
			log.Println(err)
			delLastStep()
			return
		}
		currRecipe.Steps[i].print(currRecipe)
		currRecipe.printIngredients(true)
		currRecipe.printSteps()
	}

	if strings.HasPrefix(input, "smix") {
		ms := createMixSingle(inputSpaced[1:], currRecipe, "")
		currRecipe.Steps = append(currRecipe.Steps, &ms)
		i := len(currRecipe.Steps) - 1

		err := currRecipe.Steps[i].execute(currRecipe)
		if err != nil {
			log.Println(err)
			delLastStep()
			return
		}
		currRecipe.Steps[i].print(currRecipe)
		currRecipe.printIngredients(true)
		currRecipe.printSteps()
	}

	if strings.HasPrefix(input, "ls") {
		currRecipe.printIngredients(true)
		currRecipe.printSteps()
	}

	if strings.HasPrefix(input, "del") {
		delLastStep()
	}

}

func delLastStep() {
	currRecipe.Steps = currRecipe.Steps[:len(currRecipe.Steps)-1]
}

//InputArgumentCheck checks spaced input array for correct amount of arguments
func InputArgumentCheck(inputSpaced []string, amount int) error {
	if len(inputSpaced) != amount+1 {
		amountStr := strconv.Itoa(amount)
		err := errors.New("Command requires " + amountStr + " argument(s)!")
		return err
	} else {
		return nil
	}
}

//Save recipe to defaultFolder/RecipeName.json
func Save() {
	location := defaultFolder + currRecipe.Name + ".json"
	fmt.Println(location)

	file, err := json.MarshalIndent(currRecipe, "", " ")

	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(location, file, 0644)

	if err != nil {
		log.Println(err)
		return
	}
}

//Open recipe from defaultFolder/name.json
func Open(name string) {
	location := defaultFolder + name + ".json"

	file, err := os.Open(location)

	if err != nil {
		log.Println(err)
		return
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(byteValue, &currRecipe)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(location)
}
